package main

import (
	"github.com/hampusad/book-analyser/bookanalyser"
)

func main() {
	b := bookanalyser.Book{}
	b.Read()
	b.Print()
}
