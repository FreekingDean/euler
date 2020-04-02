package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	tests := []struct {
		x, y int
	}{
		{10, 23},
		{16, 60},
		{1000, 233168},
	}
	for _, test := range tests {
		assert.Equal(t, test.y, P1(test.x))
	}
}

func TestP2(t *testing.T) {
	tests := []struct {
		x, y int
	}{
		{3, 2},
		{10, 10},
		{55, 44},
		{4000000, 4613732},
	}
	for _, test := range tests {
		assert.Equal(t, test.y, P2(test.x))
	}
}

func TestP3(t *testing.T) {
	tests := []struct {
		x int
		y int
	}{
		{3, 3},
		{13195, 29},
		{600851475143, 6857},
	}
	for _, test := range tests {
		assert.Equal(t, test.y, P3(test.x))
	}
}
