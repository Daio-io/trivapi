package db

import "os"

const localConnection = "mongodb://localhost/trivapi"

func getConnectionString() string {
  connectString := os.Getenv("DB_CONNECTION")
  if connectString == "" {
    return localConnection
  }
  return connectString
}
