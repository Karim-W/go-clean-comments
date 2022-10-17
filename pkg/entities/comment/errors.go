package comment

import "fmt"

var (
	Err_INVALID_ID   = fmt.Errorf("invalid id")
	Err_NOT_FOUND    = fmt.Errorf("not found")
	Err_NO_USER_ID   = fmt.Errorf("no user id")
	Err_EMPTY_BODY   = fmt.Errorf("empty body")
	Err_NO_TIMESTAMP = fmt.Errorf("no timestamp provided")
)
