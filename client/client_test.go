package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendWithoutDupes(t *testing.T) {
	cases := []struct {
		Base     []string
		Add      []string
		Expected []string
	}{
		{
			Base:     nil,
			Add:      []string{"a", "b", "a"},
			Expected: []string{"a", "b"},
		},
		{
			Base:     []string{"a", "b"},
			Add:      []string{"c", "b", "a"},
			Expected: []string{"a", "b", "c"},
		},
		{
			Base:     []string{"a", "b"},
			Add:      []string{"b"},
			Expected: []string{"a", "b"},
		},
	}
	for _, tc := range cases {
		appendWithoutDupes(&tc.Base, tc.Add)
		assert.Equal(t, tc.Expected, tc.Base)
	}
}
