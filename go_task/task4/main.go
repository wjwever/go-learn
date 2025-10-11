package main

import (
	"fmt"
	"log"
	"task4/app"
)

func main() {
	app := app.App{SecretKey: "secrect_key"}
	if err := app.Init(); err != nil {
		log.Fatal("服务启动失败") // TODO
	}

	fmt.Println("Connect DataBase Successfully!")
	app.Start()
}
