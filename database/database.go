package database

import (
	"log"
	"time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/charonne/goethapi/config"
)


// Smart contract structure
type Contract struct {
  ID          uint `gorm:"primary_key"`
  Idkey       string `gorm:"size:50"`

  Name        string `gorm:"size:255"`
  Source      string
  Bytecode    string
  Interface   string
	Deployed		bool

  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   *time.Time
}

// Smart contract structure
type ContractDeployed struct {
  ID          uint `gorm:"primary_key"`
  Address     string `gorm:"size:50"`
	Txhash      string `gorm:"size:50"`

  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   *time.Time
}

// Transaction structure
type Transaction struct {
  ID          uint `gorm:"primary_key"`

	Txhash      string `gorm:"size:50"`
	Type      	string `gorm:"size:50"` // contract, exec
  Contract    string `gorm:"size:50"`
	Status			int

  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   *time.Time
}

// Connect to database
func DatabaseConnection() *gorm.DB {
  db, err := gorm.Open(config.Config.Database.Connection,
    "host=" + config.Config.Database.Host +
    " dbname=" + config.Config.Database.Dbname +
    " user=" + config.Config.Database.User +
    " password=" + config.Config.Database.Password +
    " sslmode=disable")
  if err != nil {
    log.Fatalf("failed to connect database: %v", err)
  }

  return db
}

// Initialise database
func InitialiseDb () {
  log.Println("Initialise database")

  db := DatabaseConnection()

  // Migrate the schema
  db.AutoMigrate(&Contract{})
	db.AutoMigrate(&ContractDeployed{})
  db.AutoMigrate(&Transaction{})

  // Close db
  defer db.Close()
}
