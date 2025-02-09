package custom_errors

import "errors"

var (
	RecordNotFound error = errors.New("record not found")
)
