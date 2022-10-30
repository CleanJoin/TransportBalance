package main

import (
	"os"
	"strconv"

	"transportBalance/balance"
)

// @title           Swagger transportBalance
// @version         1.0
// @description     This is a sample server transportBalance
// @termsOfService  https://github.com/CleanJoin/transportBalance/
// @contact.name   Github.com
// @contact.url    https://github.com/CleanJoin/transportBalance/
// @host      localhost:8000
// @BasePath  /api/v1
func main() {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))

	serverGin := balance.NewServerGin("localhost", serverPort)

	usersStorage := balance.NewUserStorageDB(new(balance.PasswordHasherSha1), balance.NewConnectDB(5432))
	balanceStorage := balance.NewBalanceStorageDB(balance.NewConnectDB(5432))
	serverGin.Use(usersStorage, balanceStorage)
	serverGin.Run()

}
