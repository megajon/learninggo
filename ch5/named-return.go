package main

import "errors"

func main() {

}

func namedDivAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err := errors.New("cannot divide by 0")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator&denominator
	return result, remainder, err
}
