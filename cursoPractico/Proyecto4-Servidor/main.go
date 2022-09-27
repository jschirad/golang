package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Handle("/facu", HandleFacu)
	server.Listen()
}
