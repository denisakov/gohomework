package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	worldMssg := emoji.Sprint("Hello world!:smile:")
	fmt.Println(worldMssg)
}
