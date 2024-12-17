package main

import "testing"

func TestConfig(t *testing.T) {
	name := readConfigNoViper("config.toml")
	if name != "yo" {
		t.Fail()
	}
}
func BenchmarkConfig(b *testing.B) {
	_ = readConfigNoViper("config.toml")
}
