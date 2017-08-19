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
  fmt.Printf("Auth 1 address: %v\nAuth 1 key: %v\n", auth.From.String(), key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	key2, _ := crypto.GenerateKey()
	auth2 := bind.NewKeyedTransactor(key2)
	alloc[auth2.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	sim := backends.NewSimulatedBackend(alloc)

	// deploy contract
	addr, _, contract, err := contracts.DeployCoinFidToken(auth, sim, big.NewInt(0), "Toto", 0 , "TO", auth.From)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	// interact with contract
	fmt.Printf("Contract deployed to: %s\n", addr.String())

	fmt.Println("Mining...")
	sim.Commit()

  name, err := contract.Name(nil)
  fmt.Printf("Get Name: %s\n", name)

  tx, err := contract.SetName(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	}, "Coucou")
  if err != nil {
		log.Fatalf("could not execute contract: %v", err)
	}
  fmt.Printf("Contract executed: %s\n", tx.String())

	fmt.Println("Mining...")
	sim.Commit()

  name2, err := contract.Name(nil)
  fmt.Printf("Get Name 2: %s\n", name2)
}
