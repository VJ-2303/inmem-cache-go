package main

import (
	"time"

	"github.com/VJ-2303/cache/cache"
)

func main() {
	c := cache.New[int, string]()

	c.Set(1, "Hello", 3*time.Second)
	c.Set(2, "Guys", 3*time.Second)
	c.Set(3, "Friends", 3*time.Second)

	time.Sleep(5 * time.Second)

	for i := 1; i < 4; i++ {
		value, _ := c.Get(i)
		println(value)
	}
}
