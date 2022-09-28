package main

import (
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	fmt.Println("start server")
	server.Logger.Fatal(server.Start(":8000"))
}
