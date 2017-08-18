
package main




import(
    "fmt"
    "log"
    "context"

    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/charonne/goethapi/config"

    //"github.com/jinzhu/configor"
)
/*
var Config = struct {
	Name string

	Blockchain struct {
		Rawurl  string
	}

	Account struct {
		Key         string
		Passphrase  string
	}
}{}
*/
func getClient() (client *ethclient.Client, err error) {
    log.Println("Connect to: " + config.Config.Blockchain.Rawurl)
    endPoint := "/home/raph/studio/private/geth.ipc"
    client, err = ethclient.Dial(endPoint)
    return
}

func main() {

  fmt.Println(config.Config.App.Name)

  // Config
  //configor.Load(&Config, "src/github.com/charonne/goethapi/config/app.yml")

 // Connect to node
 client, err := getClient()
 if err != nil {
   log.Fatalf("could not create ipc client: %v", err)
 }

  // Block info
  fmt.Println("Block info")
  ctx := context.Background()
  blockInfo, err := client.BlockByNumber(ctx, nil)
  if err != nil {
    log.Fatalf("could not get block info: %v", err)
  }
  fmt.Println(blockInfo)

  // Gas price
  fmt.Println("Gas price")
  gasPrice, err := client.SuggestGasPrice(ctx)
  if err != nil {
    log.Print("ERR=>", err)
    return
  }
  fmt.Println(gasPrice)


}
