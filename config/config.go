package config

import (
  "log"
  "github.com/spf13/viper"

  "runtime"
  "path"
)

// Get config path
func getConfigPath() string {

  _, filename, _, ok := runtime.Caller(0)
  if !ok {
      panic("No caller information")
  }
  path := path.Dir(filename)

  return path
}

// Get config
func Config() *viper.Viper {

  configPath := getConfigPath()

  config := viper.New()
  config.SetConfigName("app")     // no need to include file extension
  config.AddConfigPath(configPath)  // set the path of your config file
  config.AddConfigPath("config")  // set the path of your config file

  err := config.ReadInConfig()
  if err != nil {
    log.Fatalf("Config file error: %v", err)
  }

  return config
}
