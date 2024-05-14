package main

import "os"

func main() {
	print(os.Getenv("TEST"))
}
