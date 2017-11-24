package database

import (
	"log"
	"time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/charonne/fidever-api/config"
)

var apiconfig = config.Config()

// Smart contract structure
type Application struct {
  ID          uint `gorm:"primary_key"`
  Idkey       string `gorm:"size:50"`

  Name        string `gorm:"size:255"`
  Callback    string `gorm:"size:255"`

  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   *time.Time
}

// Smart contract structure
type Nonce struct {
  ID          uint `gorm:"primary_key"`

  Address     string `gorm:"size:50"`
  Number       uint64
}

// Connect to database
func DatabaseConnection() *gorm.DB {

  db, err := gorm.Open(apiconfig.GetString("database.connection"),
    "host=" + apiconfig.GetString("Database.host") +
    " port=" + apiconfig.GetString("database.port") +
    " dbname=" + apiconfig.GetString("database.dbname") +
    " user=" + apiconfig.GetString("Database.user") +
    " password=" + apiconfig.GetString("database.password") +
    " sslmode=disable")

  //db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
  if err != nil {
    log.Fatalf("failed to connect database: %v", err)
  }

  return db
}

// Initialise database
func InitialiseDb () {
  log.Println("Initialise database")

  db := DatabaseConnection()

  // Drop the schema
  db.DropTable(&Nonce{})
  // Migrate the schema
  db.AutoMigrate(&Nonce{})

  // Close db
  defer db.Close()
}
