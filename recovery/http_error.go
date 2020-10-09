package recovery

import "fmt"

type HTTPError struct {
	Code   int
	Mesage string
	Err    error
}

func (err HTTPError) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("http got error: %s, status: %d, and message: %s", err.Err, err.Code, err.Mesage)
	}
	return fmt.Sprintf("http status: %d, and message: %s", err.Code, err.Mesage)
}
