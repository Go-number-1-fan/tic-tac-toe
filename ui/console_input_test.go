package ui

import "github.com/stretchr/testify/assert"
import "testing"

func TestConsoleInput_CanConvertAStringToANumAndTellIfError(t *testing.T) {
	num, anyError := convertStringToInt("123")
	assert.Equal(t, 123, num)
	assert.False(t, anyError)
}

func TestConsoleInput_TriesToConvertAStringToNumberAndTellsAboutAnyErrors(t *testing.T) {
	num, anyError := convertStringToInt("tom")
	assert.Equal(t, 0, num)
	assert.True(t, anyError)
}
