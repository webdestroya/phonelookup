package twilio

import (
	"fmt"
)

func GetErrorString(codeThing any) string {

	code := 0

	switch v := codeThing.(type) {
	case float32:
		code = int(v)
	case float64:
		code = int(v)

	case int:
		// nothing
		code = v

	case int64:
		code = int(v)
	case int32:
		code = int(v)
	case uint:
		code = int(v)
	case uint64:
		code = int(v)
	case uint32:
		code = int(v)

	default:
		return fmt.Sprintf("%v - Unhandled (%T)", v, v)
	}

	return fmt.Sprintf("%d - %s", code, GetErrorMessage(code))
}
