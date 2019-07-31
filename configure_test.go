package hystrix_test

import (
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	mhystrix "github.com/x-punch/micro-hystrix"
)

func TestConfigureCommand(t *testing.T) {
	command := "testing.cmd"
	timeout := 100
	mhystrix.ConfigureCommand(command, mhystrix.CommandConfig{Timeout: timeout})
	configures := hystrix.GetCircuitSettings()
	if c, ok := configures[command]; !ok || c.Timeout != time.Duration(timeout)*time.Millisecond {
		t.Fail()
	}
}

func TestConfigureDefault(t *testing.T) {
	timeout := 100
	mhystrix.ConfigureDefault(mhystrix.CommandConfig{Timeout: timeout})
	if hystrix.DefaultTimeout != timeout {
		t.Fail()
	}
}
