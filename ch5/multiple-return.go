package main

import "errors"

func main() {

}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
