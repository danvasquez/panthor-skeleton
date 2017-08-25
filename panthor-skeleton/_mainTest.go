package main

import (
	"testing"
	"github.com/danvasquez/panthor-skeleton/php"
)

func TestTesting(t *testing.T) {
	output := php.SampleForTesting("widget")

	if output !="say widget" {
		t.Fail()
	}
}
