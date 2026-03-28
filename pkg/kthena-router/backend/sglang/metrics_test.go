package sglang

import "testing"

func TestNewSglangEngine_UsesDefaultMetricPort(t *testing.T) {
	engine := NewSglangEngine()

	if engine.MetricPort != 30000 {
		t.Fatalf("expected default metric port 30000, got %d", engine.MetricPort)
	}
}

func TestNewSglangEngine_UsesConfiguredMetricPort(t *testing.T) {
	engine := NewSglangEngine(31000)

	if engine.MetricPort != 31000 {
		t.Fatalf("expected metric port 31000, got %d", engine.MetricPort)
	}
}
