package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/webdav"
	"log"
	"net"
	"net/http"
	"strings"
)

func getIp() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}
	for _, addr := range address {
		fmt.Println("监听地址：", strings.Split(addr.String(), "/")[0])
	}
}

func main() {
	getIp()
	err := http.ListenAndServe(":81", &webdav.Handler{FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS()})
	if err != nil {
		log.Println(err)
	}

	app := fiber.New()

	// Custom config
	app.Static("/", "./", fiber.Static{
		Compress: true,
		Browse:   true,
	})

	err = app.Listen(":80")
	if err != nil {
		log.Fatal(err)
	}
}
