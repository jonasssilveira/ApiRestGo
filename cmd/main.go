package main

import (
	server2 "booksAPiProject/server"
)

func main() {
	server := server2.NewServer()
	server.Run()
}
