package common

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockfunction() {}

func TestPKGCloseFunc(t *testing.T) {
	function := PKGCloseFunc(mockfunction)
	assert.NoError(t, function())
}

func TestStringInSlice(t *testing.T) {
	result := StringInSlice("a", []string{"a", "b"})
	assert.True(t, result)
	result = StringInSlice("c", []string{"a", "b"})
	assert.False(t, result)
}

func TestCheckExpires(t *testing.T) {
	result, err := CheckExpires(time.Now().Add(time.Hour).Unix())
	assert.False(t, result)
	assert.NoError(t, err)
	result, err = CheckExpires(time.Now().Add(-time.Hour).Unix())
	assert.True(t, result)
	assert.NoError(t, err)
}

func TestParameterError(t *testing.T) {
	result_status, result_message := GetShortError()
	assert.Equal(t, http.StatusNotFound, result_status)
	assert.Equal(t,
		gin.H{
			"description": "Parameter Error.",
		},
		result_message,
	)
}
