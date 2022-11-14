package marsrover

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestShouldMoveRoverAround(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Run("5 5\n"+"1 1 N\n"+"MM\n"), "1 3 N\n")

}

func TestShouldTurnRoverRight(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Run("5 5\n"+"1 1 N\n"+"RRR\n"), "1 1 W\n")

}

func TestShouldTurnRoverLeft(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Run("5 5\n"+"1 1 N\n"+"LLL\n"), "1 1 E\n")

}

func TestShouldMoveMultipleRoversCorrectly(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Run("5 5 \n 1 2 N \nLMLMLMLMM  \n 3 3 E \nMMRMMRMRRM"), "1 3 N\n5 1 E\n")

}
