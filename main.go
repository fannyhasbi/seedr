package main

import (
	"log"

	_ "github.com/lib/pq"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Hello world")
}
