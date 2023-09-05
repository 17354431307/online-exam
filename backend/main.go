package main

import "backend/core"

func main() {
	core.InitViper("etc/config.yaml")
	core.RunServer()
}
