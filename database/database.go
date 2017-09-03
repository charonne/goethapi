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
  Idkey       string

  Name        string
  Source      string
  Bytecode    string
  Interface   string

  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   *time.Time
}

// Transaction structure
type Transaction struct {
  ID          uint `gorm:"primary_key"`

	Txhash      string
	Type      	string // contract, exec
  Contract    string
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
  db.AutoMigrate(&Transaction{})

  // Close db
  defer db.Close()
}
