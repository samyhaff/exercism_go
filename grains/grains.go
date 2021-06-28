package grains

import "errors"

func Square(square int) (uint64, error) {
	if square > 0 && square <= 64 {
		return 1 << uint(square-1), nil
	}
	return 0, errors.New("input out of board range")
}

func Total() uint64 {
	return 1<<64 - 1
}
