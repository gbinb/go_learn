package main

import (
	"log"
	"os"
	"search/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
