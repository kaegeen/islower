package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	input := "AaslEma , Baba!"
	fmt.Println(toLowerCase(input))
}
