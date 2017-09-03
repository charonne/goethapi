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
  "io/ioutil"
  "context"

  "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"

  "github.com/charonne/goethapi/config"
)

func main() {
	fmt.Printf("%s v%s\n", config.Config.App.Name, config.Config.App.Version)

  endPoint := config.Config.Blockchain.Rawurl
  client, err := ethclient.Dial(endPoint)
  if err != nil {
      log.Fatal(err)
  }
	log.Println("Connected to: ", config.Config.Blockchain.Rawurl)
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

  // Get context
  ctx := context.Background()

  // Figure out the gas price
  gasPrice, err := client.SuggestGasPrice(ctx)
  if err != nil {
    log.Fatal(err)
  }
	log.Println("Gas price: ", gasPrice)

  // Figure out the gas limit
  from := common.HexToAddress(config.Config.Account.Address)
  to := common.HexToAddress("0x56a4683218dcf19d5b62b667c4348e2ac06d174b")
  value := big.NewInt(1000000000000000000) // 1 ether

  input := []byte(config.Config.Account.Passphrase)
  msg := ethereum.CallMsg{From: from, To: &to, Value: value, Data: input}
  gasLimit, err := client.EstimateGas(ctx, msg)
  if err != nil {
    log.Fatal(err)
  }
	log.Println("Gas limit: ", gasLimit)


  // Get nonce
  nonce, err := client.PendingNonceAt(ctx, from)
  if err != nil {
    log.Fatal(err)
  }

  // Create the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, input)

  // Sign
  transact_opt := bind.NewKeyedTransactor(key.PrivateKey)
  //log.Printf("transact_opt: %T", transact_opt)
  tx, err = transact_opt.Signer(types.HomesteadSigner{}, common.HexToAddress(config.Config.Account.Address), tx)
  if err != nil {
      log.Fatal(err)
  }
  log.Printf("Tx: %v", tx)

  err = client.SendTransaction(ctx, tx)
  if err != nil {
      log.Fatalf("Transaction failed: %v", err)
  }
}
