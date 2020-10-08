package main

import "errors"

type NumberRepository struct {
	number int
}

type NumberService interface {
	SumAllPrimesNumber() (int, error)
}

func (_this NumberRepository) SumAllPrimesNumber() (int, error) {
	sum := 0
	if _this.number < 1 {
		return 0, errors.New("The input number has to be greater 0")
	}
	for i := 1; i <= _this.number; i++ {
		if isPrime(i) {
			sum = sum + i
		}
	}
	return sum, nil
}

func isPrime(numberToCheck int) bool {
	if numberToCheck == 1 {
		return false
	}
	for i := 2; i*i <= numberToCheck; i++ {
		if numberToCheck%i == 0 {
			return false
		}
	}
	return true
}
