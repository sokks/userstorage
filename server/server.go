package main

import (
	"flag"

	"github.com/sokks/userstorage"
)

func main() {
	portPtr := flag.String("port", "9080", "server listening port")
	inMemPtr := flag.Bool("inmemory", false, "uses in-memory storage(not full and only for testing)")
	dbFilePtr := flag.String("dbFile", "", "server sqlite database filepath")
	flag.Parse()
	userstorage.StartServer(*inMemPtr, *portPtr, *dbFilePtr)
}