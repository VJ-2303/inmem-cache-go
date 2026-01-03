package main

import "github.com/VJ-2303/cache/cache"

func main() {
	c := cache.New[int, string]()

	c.Set(1, "Hello")
	c.Set(2, "Guys")
	c.Set(3, "Friends")

	c.Remove(2)

	for i := 1; i < 4; i++ {
		value, _ := c.Get(i)
		println(value)
	}
}
