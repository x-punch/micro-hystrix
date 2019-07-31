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

# Default Configure in hystrix
```
var (
	// DefaultTimeout is how long to wait for command to complete, in milliseconds
	DefaultTimeout = 1000
	// DefaultMaxConcurrent is how many commands of the same type can run at the same time
	DefaultMaxConcurrent = 10
	// DefaultVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
	DefaultVolumeThreshold = 20
	// DefaultSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
	DefaultSleepWindow = 5000
	// DefaultErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests
	DefaultErrorPercentThreshold = 50
)
```

# Update default config in hystrix
```
package main

import (
	"github.com/micro/go-micro"
	"github.com/x-punch/micro-hystrix"
)

func main() {
	hystrix.ConfigureDefault(hystrix.CommandConfig{Timeout: 1000})
	service := micro.NewService(micro.WrapClient(hystrix.NewClientWrapper()))
	service.Init(micro.Name("test.srv"), micro.Address(":80"))
	if err := service.Run(); err != nil {
		panic("Failed to run service: " + err.Error())
	}
}
```
