package config


import(
  "github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		Name    string
		Version string
		Url     string
		Port    string
	}

	Database struct {
		Connection  string
		Host        string
		Dbname      string
		User        string
		Password    string
	}

	Blockchain struct {
		Rawurl        string
		Keystorepath  string
	}

	Account struct {
    Address     string
    Keystore    string
		Json        string
		Passphrase  string
	}

	Ticker struct {
    Period      int
	}

  Callback      string
}{}

func init() {
  // Config
  configor.Load(&Config, "src/github.com/charonne/goethapi/config/app.yml")
}
