package main

import "smallServer/router"

func main() {
	r := router.SetupRouter()
	_ = r.Run(":80")
}
