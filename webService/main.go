package main

import (
	"webService/router"
)

func main() {
	router := router.SetupRouter()
	router.Run()
}
