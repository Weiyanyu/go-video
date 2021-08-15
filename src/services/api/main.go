package main

import "go-video/src/services/api/router"

func main() {
	r := router.NewRouter()

	r.Run()
}
