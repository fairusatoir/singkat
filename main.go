package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
		os.Exit(0)
	}
}

func run() error {
	Sh, err := NewShorten("https://github.com/jackc/pgx")
	if err != nil {
		return err
	}

	fmt.Println(Sh)

	return nil
}
