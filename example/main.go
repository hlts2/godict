package main

import (
	"context"
	"log"

	"github.com/hlts2/godict"
)

func main() {
	dict, err := godict.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start to search pattern: 3dmax")

	err = dict.Do(context.Background(), func(pattern string) {
		if pattern == "3dmax" {
			log.Printf("find: %s\n", pattern)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
