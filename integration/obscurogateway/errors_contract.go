package faucet

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func init() { //nolint: gochecknoinits
	contractABI, err := abi.JSON(strings.NewReader(errorsContractABIString))
	if err != nil {
		panic(err)
	}
	errorsContractABI = contractABI
}

var (
	errorsContractABI       abi.ABI
	errorsContractBytecode  = "0x60806040526040518060400160405280600381526020017f666f6f0000000000000000000000000000000000000000000000000000000000815250600090816200004a9190620002d9565b503480156200005857600080fd5b50620003c0565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620000e157607f821691505b602082108103620000f757620000f662000099565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620001617fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000122565b6200016d868362000122565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001ba620001b4620001ae8462000185565b6200018f565b62000185565b9050919050565b6000819050919050565b620001d68362000199565b620001ee620001e582620001c1565b8484546200012f565b825550505050565b600090565b62000205620001f6565b62000212818484620001cb565b505050565b5b818110156200023a576200022e600082620001fb565b60018101905062000218565b5050565b601f82111562000289576200025381620000fd565b6200025e8462000112565b810160208510156200026e578190505b620002866200027d8562000112565b83018262000217565b50505b505050565b600082821c905092915050565b6000620002ae600019846008026200028e565b1980831691505092915050565b6000620002c983836200029b565b9150826002028217905092915050565b620002e4826200005f565b67ffffffffffffffff8111156200030057620002ff6200006a565b5b6200030c8254620000c8565b620003198282856200023e565b600060209050601f8311600181146200035157600084156200033c578287015190505b620003488582620002bb565b865550620003b8565b601f1984166200036186620000fd565b60005b828110156200038b5784890151825560018201915060208501945060208101905062000364565b86831015620003ab5784890151620003a7601f8916826200029b565b8355505b6001600288020188555050505b505050505050565b610a0280620003d06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063812002241161005b57806381200224146100c6578063852055c1146100e257806398849947146100ec578063e072b5f6146100f65761007d565b80633943380c146100825780634ed3885e146100a05780637e479d4a146100bc575b600080fd5b61008a610114565b60405161009791906103bd565b60405180910390f35b6100ba60048036038101906100b59190610528565b6101a2565b005b6100c46101b5565b005b6100e060048036038101906100db9190610528565b6101f8565b005b6100ea61024f565b005b6100f4610260565b005b6100fe61029b565b60405161010b91906103bd565b60405180910390f35b60008054610121906105a0565b80601f016020809104026020016040519081016040528092919081815260200182805461014d906105a0565b801561019a5780601f1061016f5761010080835404028352916020019161019a565b820191906000526020600020905b81548152906001019060200180831161017d57829003601f168201915b505050505081565b80600090816101b19190610787565b5050565b60006101f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed906108a5565b60405180910390fd5b565b600081510361023c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023390610911565b60405180910390fd5b806000908161024b9190610787565b5050565b600061025e5761025d610931565b5b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610292906109ac565b60405180910390fd5b6060600080546102aa906105a0565b80601f01602080910402602001604051908101604052809291908181526020018280546102d6906105a0565b80156103235780601f106102f857610100808354040283529160200191610323565b820191906000526020600020905b81548152906001019060200180831161030657829003601f168201915b5050505050905090565b600081519050919050565b600082825260208201905092915050565b60005b8381101561036757808201518184015260208101905061034c565b60008484015250505050565b6000601f19601f8301169050919050565b600061038f8261032d565b6103998185610338565b93506103a9818560208601610349565b6103b281610373565b840191505092915050565b600060208201905081810360008301526103d78184610384565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61043582610373565b810181811067ffffffffffffffff82111715610454576104536103fd565b5b80604052505050565b60006104676103df565b9050610473828261042c565b919050565b600067ffffffffffffffff821115610493576104926103fd565b5b61049c82610373565b9050602081019050919050565b82818337600083830152505050565b60006104cb6104c684610478565b61045d565b9050828152602081018484840111156104e7576104e66103f8565b5b6104f28482856104a9565b509392505050565b600082601f83011261050f5761050e6103f3565b5b813561051f8482602086016104b8565b91505092915050565b60006020828403121561053e5761053d6103e9565b5b600082013567ffffffffffffffff81111561055c5761055b6103ee565b5b610568848285016104fa565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806105b857607f821691505b6020821081036105cb576105ca610571565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026106337fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826105f6565b61063d86836105f6565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061068461067f61067a84610655565b61065f565b610655565b9050919050565b6000819050919050565b61069e83610669565b6106b26106aa8261068b565b848454610603565b825550505050565b600090565b6106c76106ba565b6106d2818484610695565b505050565b5b818110156106f6576106eb6000826106bf565b6001810190506106d8565b5050565b601f82111561073b5761070c816105d1565b610715846105e6565b81016020851015610724578190505b610738610730856105e6565b8301826106d7565b50505b505050565b600082821c905092915050565b600061075e60001984600802610740565b1980831691505092915050565b6000610777838361074d565b9150826002028217905092915050565b6107908261032d565b67ffffffffffffffff8111156107a9576107a86103fd565b5b6107b382546105a0565b6107be8282856106fa565b600060209050601f8311600181146107f157600084156107df578287015190505b6107e9858261076b565b865550610851565b601f1984166107ff866105d1565b60005b8281101561082757848901518255600182019150602085019450602081019050610802565b868310156108445784890151610840601f89168261074d565b8355505b6001600288020188555050505b505050505050565b7f466f726365642072657175697265000000000000000000000000000000000000600082015250565b600061088f600e83610338565b915061089a82610859565b602082019050919050565b600060208201905081810360008301526108be81610882565b9050919050565b7f4e6577206b65792063616e6e6f7420626520656d707479000000000000000000600082015250565b60006108fb601783610338565b9150610906826108c5565b602082019050919050565b6000602082019050818103600083015261092a816108ee565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f466f726365642072657665727400000000000000000000000000000000000000600082015250565b6000610996600d83610338565b91506109a182610960565b602082019050919050565b600060208201905081810360008301526109c581610989565b905091905056fea264697066735822122096d8546ec9908bc71f158582f76aba34c9d82114e3afbfc84a3722ce20ae975d64736f6c63430008120033"
	errorsContractABIString = `[
	{
		"inputs": [],
		"name": "force_assert",
		"outputs": [],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "force_require",
		"outputs": [],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "force_revert",
		"outputs": [],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "get_key",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "key",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_key",
				"type": "string"
			}
		],
		"name": "set",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "newKey",
				"type": "string"
			}
		],
		"name": "set_key",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
)

// // SPDX-License-Identifier: GPL-3.0
//pragma solidity ^0.8.0;
//
//contract Error {
//    string public key = "foo";
//
//    function set(string memory _key) public {
//        key = _key;
//    }
//
//    function force_require() public view {
//        require(false, "Forced require");
//    }
//
//    function force_revert() public view {
//        revert("Forced revert");
//    }
//
//    function force_assert() public view{
//        assert(false);
//    }
//
//    function set_key(string memory newKey) public {
//        require(bytes(newKey).length != 0, "New key cannot be empty");
//        key = newKey;
//    }
//
//    function get_key() public view returns (string memory) {
//        return key;
//    }
//}
