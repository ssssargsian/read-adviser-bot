package e

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("can't do request: %w", err)
}
func WrapIfError(msg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}
