package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, 20, 30, "perhitungan volume salah")
}

func main() {
	TestHitungVolume(&testing.T{})
}
