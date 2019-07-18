package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
)

// CommandConfig is used to tune circuit settings at runtime
type CommandConfig struct {
	Timeout                int
	MaxConcurrentRequests  int
	RequestVolumeThreshold int
	SleepWindow            int
	ErrorPercentThreshold  int
}

// Configure applies settings for a set of circuits
func Configure(cmds map[string]CommandConfig) {
	for k, v := range cmds {
		ConfigureCommand(k, v)
	}
}

// ConfigureCommand applies settings for a circuit
func ConfigureCommand(name string, config CommandConfig) {
	hystrix.ConfigureCommand(name, hystrix.CommandConfig{
		Timeout:                config.Timeout,
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.SleepWindow,
		ErrorPercentThreshold:  config.ErrorPercentThreshold,
	})
}
