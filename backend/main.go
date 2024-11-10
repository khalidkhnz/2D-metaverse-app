package main

import (
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
)

func main() {
	lib.InitEnv()

	server := NewAPIServer(lib.GetPort(), lib.GetDBURI())
	server.Run(RunOptions{
		EnableProxyServer: true,
		EnableFileServer: false,
	})
}