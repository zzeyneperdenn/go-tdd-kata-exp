package test

import (
	"github.com/stretchr/testify/assert"
	"go-tdd-kata-exp/internal"
	"testing"
)

func Test_Add_With_Empty_String(t *testing.T) {
	result, err := internal.Add("")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 0, result)
}
