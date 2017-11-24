package node

import(
  "fmt"
  "log"
  "context"
  "github.com/fatih/color"

  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/accounts/keystore"
  "github.com/ethereum/go-ethereum/crypto"

  "github.com/charonne/goethapi/config"
  "github.com/charonne/goethapi/converter"
)

func getClient() (client *ethclient.Client, err error) {
  log.Println("Connect to: " + config.Config.Blockchain.Rawurl)
  client, err = ethclient.Dial(config.Config.Blockchain.Rawurl)
  return
}

// Get gas price
func getGasPrice() {
  color.Blue("Get gas price:")
  // Connect to node
  client, err := getClient()
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }
  // Get context
  ctx := context.Background()
  // Gas price
  gasPrice, err := client.SuggestGasPrice(ctx)
  if err != nil {
    log.Print("ERR=>", err)
    return
  }
  fmt.Println(gasPrice)
}

// Get block info
func getBlockInfo() {
  color.Blue("Get block info:")
  // Connect to node
  client, err := getClient()
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }
  // Get context
  ctx := context.Background()
  // Block info
  blockInfo, err := client.BlockByNumber(ctx, nil)
  if err != nil {
    log.Fatalf("could not get block info: %v", err)
  }
  fmt.Println(blockInfo)
}

// Get account list
func getAccounts() {
  color.Blue("Get accounts:")
  // Connect to node
  client, err := getClient()
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }
  // Get context
  ctx := context.Background()
  // Account list
  fmt.Printf("Get account list from: %s\n", config.Config.Blockchain.Keystorepath)
  keystore_path := config.Config.Blockchain.Keystorepath
  keystoreList := keystore.NewKeyStore(keystore_path, keystore.LightScryptN, keystore.LightScryptP)
  for i, ks := range keystoreList.Accounts() {
    balance, err := client.BalanceAt(ctx, ks.Address, nil)
    if err != nil {
      log.Fatalf("could not get balance: %v", err)
    }
    balanceEther := converter.Convert(balance, "wei", "ether")
    fmt.Printf("%d: %s, balance: %v ether,  %v wei\n", i + 1, ks.Address.String(), balanceEther, balance)
  }
}

// Get private key
func getPrivateKey() {
  color.Blue("Get private key:")
  key, err := keystore.DecryptKey([]byte(config.Config.Account.Json), config.Config.Account.Passphrase)
  if err != nil {
      log.Fatal("Json key decrypted with bad password")
  }
  fmt.Printf("Private key: %x\n", key.PrivateKey.D.Bytes())

  publicKey := key.PrivateKey.PublicKey
  fmt.Printf("Public key: %x\n", crypto.PubkeyToAddress(publicKey))
  /*
  // Alternative with pointers
  publicKey := &key.PrivateKey.PublicKey
  fmt.Printf("Public key %x", crypto.PubkeyToAddress(*publicKey))
  */
}
