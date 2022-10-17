package main

import "fmt"
import "os"

func main() {
	cwd, err := os.Getwd()
	fmt.Println("cwd:", cwd, "err", err)
}
