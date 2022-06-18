package main

import (
	"fmt"
	"strings"
)

func main() {
	server := "127.0.0.1/"
	fmt.Println(strings.HasSuffix(server, "/"))
}
