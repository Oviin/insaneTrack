package main

import (
	"github.com/FuJT/insaneTrack/web"
)

func main() {
	server := web.New()
	server.InitHandlers()
	server.Run("localhost:9000")

}
