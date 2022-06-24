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

func Test_Add_With_Some_Numbers(t *testing.T) {
	result, err := internal.Add("1,2,3")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 6, result)
}
