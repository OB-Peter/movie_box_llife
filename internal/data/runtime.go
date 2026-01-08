package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntime = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	js := fmt.Sprintf(`"%d mins"`, r)

	return []byte(js), nil
}

func (r *Runtime) UnmarshalJSON(js []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(js))
	if err != nil {
		return ErrInvalidRuntime
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntime
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntime
	}

	*r = Runtime(i)
	return nil
}
