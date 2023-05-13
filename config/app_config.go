package config

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var infuraConn *ethclient.Client

func connectInfura() *ethclient.Client {
	LoadEnv()
	INFURA_GOERLI_KEY := os.Getenv("INFURA_GOERLI_KEY")
	infuraConnection, err := ethclient.Dial(INFURA_GOERLI_KEY)

	if err != nil {
		log.Fatal("Unable to connect to the chosen ethereum network: ", err)
	}

	infuraConn = infuraConnection
	return infuraConn
}

func LoadEnv() {
	// err := godotenv.Load(".env")
	err := godotenv.Load() //
	if err != nil {
		log.Fatal("Error: unable to load the .env file")
	}
}

func GetInfuraConnection() *ethclient.Client {
	connectInfura()
	return infuraConn
}
