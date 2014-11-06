package main

import 'github.com/codegangsta/martini"

func main() {
	m := martini.Classic()
	
	m.get("/", func() {
		return "Merry Christmas!"
	})

	m.run()
}
