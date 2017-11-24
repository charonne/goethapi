package tools

import(
  "fmt"
  "errors"

  "context"
  "github.com/charonne/goethapi/internal/database"
	"github.com/ethereum/go-ethereum/ethclient"

  "github.com/ethereum/go-ethereum/common"
  "github.com/charonne/goethapi/config"
)

var apiconfig = config.Config()

// Get nonce
func GetNodeNonce(address common.Address) uint64 {
  // Call node
  endPoint := apiconfig.GetString("blockchain.rawurl")
  client, err := ethclient.Dial(endPoint)
  if err != nil {
    fmt.Printf("Dial error: %s", err)
  }
  // Get nonce
  ctx := context.Background()
  nonce, err := client.PendingNonceAt(ctx, address)
  if err != nil {
    fmt.Println(err)
  }
  return nonce
}

// Get current nonce number
func GetLastNonce(address common.Address) (database.Nonce, error) {
  // Get database
  db := database.DatabaseConnection()
  nonceNb := GetNodeNonce(address)

  // Create nonce if not exists
  var existNonce database.Nonce
  if err := db.Where("address = ?", address.String()).First(&existNonce).Error; err != nil {
    // Remove one to get the real last number
    nonceNb -= 1
    fmt.Printf("last nonce: %v\n", nonceNb)
    newNonce := database.Nonce{Address: address.String(), Number:nonceNb}
    db.Save(&newNonce)
    return newNonce, nil
  } else {
    // Node as last nonce
    if (existNonce.Number < nonceNb) {
      // Remove one to get the real last number
      nonceNb -= 1
      db.Model(&existNonce).Update("Number", nonceNb)
    }
    fmt.Printf("nonce: %v\n", existNonce)
    return existNonce, nil
  }
  return existNonce, errors.New("nonce error")
}

// Get new nonce number
func GetNewNonce(address common.Address) database.Nonce {
  // Get current nonce
  nonce, err := GetLastNonce(address)
  if err != nil {
    fmt.Printf("Nonce error: %s", err)
    return nonce
  }
  nonce.Number += 1
  // Get database
  db := database.DatabaseConnection()
  db.Model(&nonce).Update("Number", nonce.Number)

  return nonce
}
