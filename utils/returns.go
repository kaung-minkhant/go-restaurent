package utils

import "fmt"

func ReturnAccessDenied() error {
	return fmt.Errorf("access denied")
}
func ReturnSomethingWentWrong() error {
	return fmt.Errorf("something went wrong")
}
