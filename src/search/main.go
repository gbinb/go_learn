package main

import (
	"log"
	"os"
	"search/search"

	_ "search/matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
