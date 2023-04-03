package config

import (
	"log"

	"github.com/joho/godotenv"
)

// databaseConnection := 'null'

func connectDB() {

	// dbConnection, err := gorm.Open("mysql", dsn(dbname))

	// // handle error
	// if err != nil {
	// 	panic(err)
	// }

	// // err = dbConnection.Ping()
	// databaseConnection = dbConnection

	// // if err != nil {
	// // 	fmt.Print("Unable to migrate DB tables")
	// // }
	// // defer dbConnection.Close()
	// if err != nil {
	// 	panic("Unable to open database")
	// }
	// return databaseConnection
}

func LoadEnv() {
	// err := godotenv.Load(".env")
	err := godotenv.Load() //
	if err != nil {
		log.Fatal("Error: unable to load the .env file")
	}
}

func GetDBConnection() {
	// connectDB()
	// return databaseConnection
}
