package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUp(t *testing.T) {

	testCases := []struct {
		Input         string
		ExpectedError bool
	}{
		{"", true},
		{"bbc.co.uk", true},
		{"https://google.co.uk", false},
		{"hltsps://google.co.uk", true},
		{"http://google.co.uk", false},
		{"https://jsonplaceholder.typicode.com/posts", false},
	}

	for _, test := range testCases {

		result := IsUp(test.Input)

		if test.ExpectedError {
			assert.Error(t, result)

		}

		if !test.ExpectedError {
			assert.NoError(t, result)
		}

		t.Fail()
	}
}
