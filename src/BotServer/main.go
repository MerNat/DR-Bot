package main

// Init ServerSide Implemenation
func main() {
	go InitInput()
	InitConn("tcp4", "localhost", "8080")
}