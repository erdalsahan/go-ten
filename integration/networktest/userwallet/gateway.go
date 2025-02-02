package userwallet

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/ten-protocol/go-ten/go/common/retry"
	"github.com/ten-protocol/go-ten/go/rpc"
	"github.com/ten-protocol/go-ten/go/wallet"
	"github.com/ten-protocol/go-ten/tools/walletextension/lib"
)

type GatewayUser struct {
	wal wallet.Wallet

	gwLib  *lib.TGLib // TenGateway utility
	client *ethclient.Client

	// state managed by the wallet
	nonce uint64

	logger gethlog.Logger
}

func NewGatewayUser(wal wallet.Wallet, gatewayURL string, logger gethlog.Logger) (*GatewayUser, error) {
	gwLib := lib.NewTenGatewayLibrary(gatewayURL, "") // not providing wsURL for now, add if we need it

	err := gwLib.Join()
	if err != nil {
		return nil, fmt.Errorf("failed to join TenGateway: %w", err)
	}
	err = gwLib.RegisterAccount(wal.PrivateKey(), wal.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to register account with TenGateway: %w", err)
	}

	client, err := ethclient.Dial(gwLib.HTTP())
	if err != nil {
		return nil, fmt.Errorf("failed to dial TenGateway HTTP: %w", err)
	}

	fmt.Printf("Registered acc with TenGateway: %s (%s)\n", wal.Address(), gwLib.HTTP())

	return &GatewayUser{
		wal:    wal,
		gwLib:  gwLib,
		client: client,
		logger: logger,
	}, nil
}

func (g *GatewayUser) SendFunds(ctx context.Context, addr gethcommon.Address, value *big.Int) (*gethcommon.Hash, error) {
	txData := &types.LegacyTx{
		Nonce: g.nonce,
		Value: value,
		To:    &addr,
	}
	gasPrice, err := g.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to suggest gas price - %w", err)
	}
	txData.GasPrice = gasPrice
	gasLimit, err := g.client.EstimateGas(ctx, ethereum.CallMsg{
		From: g.wal.Address(),
		To:   &addr,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to estimate gas - %w", err)
	}
	txData.Gas = gasLimit
	signedTx, err := g.wal.SignTransaction(txData)
	if err != nil {
		return nil, fmt.Errorf("unable to sign transaction - %w", err)
	}
	err = g.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("unable to send transaction - %w", err)
	}
	txHash := signedTx.Hash()
	// transaction has been sent, we increment the nonce
	g.nonce++
	return &txHash, nil
}

func (g *GatewayUser) AwaitReceipt(ctx context.Context, txHash *gethcommon.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error
	err = retry.Do(func() error {
		receipt, err = g.client.TransactionReceipt(ctx, *txHash)
		if !errors.Is(err, rpc.ErrNilResponse) {
			return retry.FailFast(err)
		}
		return err
	}, retry.NewTimeoutStrategy(20*time.Second, 1*time.Second))
	if err != nil {
		return nil, fmt.Errorf("unable to get receipt - %w", err)
	}
	return receipt, nil
}

func (g *GatewayUser) NativeBalance(ctx context.Context) (*big.Int, error) {
	return g.client.BalanceAt(ctx, g.wal.Address(), nil)
}

func (g *GatewayUser) Wallet() wallet.Wallet {
	return g.wal
}
