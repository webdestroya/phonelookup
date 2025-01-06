package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPhoneNumber(t *testing.T) {
	tables := []struct {
		input string
		match bool
	}{
		{"1234567890", true},
		{"+1234567890", true},
		{"234-567-8900", true},
		{"+1000-000-0000", true},

		{"1+1", false},
		{"config", false},
		{"lookup", false},
		{"--help", false},
	}

	for _, table := range tables {
		t.Run(table.input, func(t *testing.T) {
			require.Equal(t, table.match, isPhoneNumber(table.input))
		})
	}
}
