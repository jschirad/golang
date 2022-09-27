package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleAPI)
	server.Handle("/facu", HandleFacu)
	server.Listen()
}
