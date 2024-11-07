package main

import "github.com/khalidkhnz/2D-metaverse-app/backend/lib"

func main() {
	server := NewAPIServer(lib.Port, lib.DbUrl)
	server.Run(RunOptions{
		EnableProxyServer: false,
		EnableFileServer: false,
	})
}