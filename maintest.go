package main

import (
    "cash.go/cash"
    "fmt"
)

func main () {
	caches := Cache.New()

	caches.Set("Loh", 123)
	caches.Item["three"] = 1222
	fmt.Println(caches.Get("three"))

	caches.Delete("three")
	fmt.Println(caches.Get("three"))

}