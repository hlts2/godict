# godict

[![Go Report Card](https://goreportcard.com/badge/github.com/hlts2/godict)](https://goreportcard.com/report/github.com/hlts2/godict)
[![GoDoc](http://godoc.org/github.com/hlts2/godict?status.svg)](http://godoc.org/github.com/hlts2/godict)

godict is a simple library that generates string patterns based on dictionary.


## Requirement

Go 1.15

## Installing

```
go get github.com/hlts2/godict
```

## Example

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


## Contribution
1. Fork it ( https://github.com/hlts2/godict/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## Author
[hlts2](https://github.com/hlts2)
