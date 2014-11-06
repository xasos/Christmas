package main

import "github.com/codegangsta/martini"

func main() {
	m := martini.Classic()
	
	m.get("/", func() string {
		return "Merry Christmas!"
	})

	m.run()
}
