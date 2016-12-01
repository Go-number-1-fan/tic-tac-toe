package ui

import "github.com/stretchr/testify/assert"
import "testing"

func TestValidator_ChecksWhetherTheValueIsInTheGivenAllowedValues_WhenItIsValid(t *testing.T) {
	validator := ConsoleInputValidator{}
	assert.True(t, validator.IsValid(0, []int{0, 1, 2, 3}))
}

func TestValidator_ChecksWhetherTheValueIsInTheGivenAllowedValues_WhenItIsInvalid(t *testing.T) {
	validator := ConsoleInputValidator{}
	assert.False(t, validator.IsValid(0, []int{1, 2, 3}))
}
