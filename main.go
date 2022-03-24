package main

import "go-mux/app"

func main() {
	a := app.App{}
	a.Initialize(
		"postgres",
		"1q2w3e4r",
		"postgres")

	a.Run(":8010")
}
