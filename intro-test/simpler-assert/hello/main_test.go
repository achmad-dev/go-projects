package main
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(test *testing.T) {
	result := hello("human")
	assert.Equal(test, "Hello, human", result)
}
