package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumAllPrimesNumber(t *testing.T) {
	Testcases := [][]int{
		{10, 17},
		{20, 77},
		{30, 129},
	}
	for i := 0; i < len(Testcases); i++ {
		numberRepo := NumberRepository{Testcases[i][0]}
		sum, err := numberRepo.SumAllPrimesNumber()
		assert.Equal(t, Testcases[i][1], sum)
		assert.Nil(t, err)
	}
}

func Test_SumAllPrimesNumber_WithNegativeNumber_ShouldBeError(t *testing.T) {
	numberRepo := NumberRepository{-1}
	sum, err := numberRepo.SumAllPrimesNumber()
	assert.Equal(t, 0, sum)
	assert.Equal(t, err.Error(), "The input number has to be greater 0")
}
