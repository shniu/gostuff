package csp

import "testing"

func TestStopSubroutine(t *testing.T) {
	StopSubroutine()
}

func TestStopSubroutineWithContext(t *testing.T) {
	StopSubroutineWithContext()
}

func TestProducerToConsumer(t *testing.T) {
	producerToConsumer()
}

// Test: select
func TestSelectExample(t *testing.T) {
	selectExample1()
}
