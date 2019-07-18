# Micro Hystrix
A go-micro plugin for hystrix-go.

# Usage
```
package main

import (
	"github.com/micro/go-micro"
	"github.com/x-punch/micro-hystrix"
)

func main() {
	service := micro.NewService(micro.WrapClient(hystrix.NewClientWrapper()))
	service.Init(micro.Name("test.srv"), micro.Address(":80"))
	if err := service.Run(); err != nil {
		panic("Failed to run service: " + err.Error())
	}
}
```
