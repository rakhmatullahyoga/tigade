package tigade

// Define custom errors here

type BaseError struct {
	Message string
	Field   string
	Code    int
}

func (e BaseError) Error() string {
	return e.Message
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
