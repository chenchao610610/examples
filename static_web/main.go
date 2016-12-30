package main

import "gopkg.in/kataras/iris.v5"

func main() {
	iris.StaticWeb("/static", "./static", 1)
	// or
	//iris.StaticWeb("/", "./static", 0)
	iris.Listen(":8080")
}
