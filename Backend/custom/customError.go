package custom

type MyError string

// Implement the Error() method for the custom error type
func (e MyError) Error() string {
	return string(e)
}