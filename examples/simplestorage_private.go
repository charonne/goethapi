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
	"time"
	"math/big"
  "io/ioutil"
  "context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"

  "github.com/charonne/goethapi/config"
	"github.com/charonne/goethapi/examples/contracts"
)


func main() {

	fmt.Printf("%s v%s\n", config.Config.App.Name, config.Config.App.Version)

	endPoint := config.Config.Blockchain.Rawurl
	client, err := ethclient.Dial(endPoint)
	if err != nil {
			log.Fatal(err)
	}

	// Get context
	ctx := context.Background()

	// Figure out the gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Gas price: ", gasPrice)

	// 	Get transact opt
	keystoreFile := config.Config.Account.Keystore
  password := config.Config.Account.Passphrase
	keyjson, err := ioutil.ReadFile(keystoreFile)
  if err != nil {
      log.Fatal(err)
  }
	key, err := keystore.DecryptKey(keyjson, password)
  if err != nil {
      log.Fatal(err)
  }
	auth := bind.NewKeyedTransactor(key.PrivateKey)


	// Deploy contract
	addr, transaction, contract, err := contracts.DeploySimpleStorage(auth, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	// interact with contract
	fmt.Println("Addr: ", addr.String())
	fmt.Println("Txhash: ", transaction.Hash().String())
	fmt.Println("Transaction: ", transaction)

	// Wait for mining
	fmt.Println("Mining...")
	time.Sleep(60 * time.Second)

	// Get value
	info, _ := contract.Get(nil)
	fmt.Printf("Get data: %d\n", info)

	// Wait for mining
	fmt.Println("Mining...")
	time.Sleep(60 * time.Second)




	// Change value
	fmt.Println("Adding new value...")
	tx, err := contract.Set(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	}, big.NewInt(42))
  if err != nil {
		log.Fatalf("could not execute contract: %v", err)
	}
  fmt.Printf("Contract executed: %s\n", tx.String())

	// Wait for mining
	fmt.Println("Mining...")
	time.Sleep(60 * time.Second)

	// Get value
	info, _ = contract.Get(nil)
	fmt.Printf("Get data: %d\n", info)




	// Change value
	fmt.Println("Adding new value...")
	addr = common.HexToAddress(addr.String())
  contract2, err := contracts.NewSimpleStorage(addr, client)
  if err != nil {
    log.Fatalf("could not call contract: %v", err)
  }
	tx, err = contract2.Set(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	}, big.NewInt(13))
  if err != nil {
		log.Fatalf("could not execute contract: %v", err)
	}
  fmt.Printf("Contract executed: %s\n", tx.String())

	// Wait for mining
	fmt.Println("Mining...")
	time.Sleep(60 * time.Second)

	// Get value
	info, _ = contract2.Get(nil)
	fmt.Printf("Get data: %d\n", info)
}
