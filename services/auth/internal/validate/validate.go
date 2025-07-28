package validate

import (
	"unicode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StrIsLetterAndDigit(str string) bool {
	for _, symbol := range str {
		if !unicode.IsLetter(symbol) && !unicode.IsDigit(symbol) {
			return false
		}
	}
	return true
}

func NameIsValid(name string) error {
	if len(name) > 16 || len(name) < 2 {
		return status.Error(codes.InvalidArgument, "length of name is wrong")
	}
	if !StrIsLetterAndDigit(name) {
		return status.Error(codes.InvalidArgument, "symbols is wrong")
	}
	return nil
}

func PasswordIsValid(name string) error {
	if len(name) > 32 || len(name) < 8 {
		return status.Error(codes.InvalidArgument, "length of password is wrong")
	}
	if !StrIsLetterAndDigit(name) {
		return status.Error(codes.InvalidArgument, "symbols is wrong")
	}
	return nil
}
