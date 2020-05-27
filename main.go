package main

import "upper-tweet/http"

func main() {
	router := http.BuildRouter()
	http.BuildServer(router).Run()
}