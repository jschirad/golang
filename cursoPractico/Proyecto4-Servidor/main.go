package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Handle("GET", "/facu", HandleFacu)
	server.Listen()
}
