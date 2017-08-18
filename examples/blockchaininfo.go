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

import(
    "fmt"
    "log"
    "context"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/common"

    "github.com/charonne/goethapi/config"
    "github.com/charonne/goethapi/converter"
)

func getClient() (client *ethclient.Client, err error) {
    log.Println("Connect to: " + config.Config.Blockchain.Rawurl)
    endPoint := "/home/raph/studio/private/geth.ipc"
    client, err = ethclient.Dial(endPoint)
    return
}

func main() {
  fmt.Println(config.Config.App.Name)

  // Connect to node
  client, err := getClient()
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }
  fmt.Println()

  // Block info
  fmt.Println("Block info:")
  ctx := context.Background()
  blockInfo, err := client.BlockByNumber(ctx, nil)
  if err != nil {
    log.Fatalf("could not get block info: %v", err)
  }
  fmt.Println(blockInfo)
  fmt.Println()

  // Gas price
  fmt.Println("Gas price:")
  gasPrice, err := client.SuggestGasPrice(ctx)
  if err != nil {
    log.Print("ERR=>", err)
    return
  }
  fmt.Println(gasPrice)
  fmt.Println()

  // Account list
  fmt.Printf("Get account list from: %s\n", config.Config.Blockchain.Keystorepath)
  keystore_path := config.Config.Blockchain.Keystorepath
  keystoreList := keystore.NewKeyStore(keystore_path, keystore.LightScryptN, keystore.LightScryptP)
  for i, ks := range keystoreList.Accounts() {
    balance, err := client.BalanceAt(ctx, common.HexToAddress("0x7844cd3b9de632fef079c8dabc7fc946e48e23ef"), nil)
    if err != nil {
      log.Fatalf("could not get balance: %v", err)
    }
    balanceEther := converter.Convert(balance, "wei", "ether")
    fmt.Printf("%d: %s, balance: %v ether\n", i + 1, ks.Address.String(), balanceEther)
  }
  fmt.Println()


}
