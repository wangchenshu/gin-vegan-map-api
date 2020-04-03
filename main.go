package main

import (
	"fmt"
	"gin-vegan-map-api/server/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	myEnv, err := godotenv.Read()

	os.Setenv("TZ", myEnv["TZ"])

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	log.Println("APP VERSION: " + myEnv["APP_VERSION"])
	log.Fatal(http.ListenAndServe(":"+myEnv["SERVER_PORT"], routes.Engine()))
}
