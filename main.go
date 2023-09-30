package main

import (
	"fmt"
	"log"
	"todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	log.Println("test")
}
