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


    //"math/big"
)

func getClient() (client *ethclient.Client, err error) {
    log.Println("Connect to: " + config.Config.Blockchain.Rawurl)
    endPoint := "/home/raph/studio/private/geth.ipc"
    client, err = ethclient.Dial(endPoint)
    return
}

func main() {

  fmt.Println(config.Config.App.Name)

  /*
  value := big.NewInt(8046)
  converter.Convert(value, "", "")
  /*
  var value = big.NewInt(8046000000000)
  //  value = 8046000000000

  bigint := converter.Convert(value, "", "")
  fmt.Printf("Convert=== %v\n", bigint)
  log.Fatalf("Stop")
  */


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
