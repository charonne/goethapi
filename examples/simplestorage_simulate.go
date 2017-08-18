package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/charonne/goethapi/test/contracts"
)

func main() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	fmt.Printf("The key: %s\n\n", key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc)

	// deploy contract
	addr, _, contract, err := contracts.DeploySimpleStorage(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	// interact with contract
	fmt.Printf("Contract deployed to: %s\n", addr.String())

	fmt.Println("Mining...")
	sim.Commit()

  fmt.Println("Adding new value...")
	tx, err := contract.Set(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(10),
	}, big.NewInt(42))
  if err != nil {
		log.Fatalf("could not execute contract: %v", err)
	}
  fmt.Printf("Contract executed: %s\n", tx.String())

	fmt.Println("Mining...")
	sim.Commit()

	info, _ := contract.Get(nil)
	fmt.Printf("Get data: %d\n", info)

	fmt.Println("Mining...")
	sim.Commit()

  // Get contract
  contract2, err := contracts.NewSimpleStorage(addr, sim)

  data2, _ := contract2.Get(nil)
  fmt.Printf("Get data 2: %d\n", data2)
}
