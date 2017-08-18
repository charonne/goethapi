// Copyright 2017 Charonne https://charonne.com
// This file is part of the goethapi library.
//
// The goethapi library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gethitihteg library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gethitihteg library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/charonne/goethapi/examples/contracts"
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
