package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	emoji.Println(":smile: Beer!!!")
	worldMssg := emoji.Sprint("Hello world!:smile:")
	fmt.Println(worldMssg)
}
