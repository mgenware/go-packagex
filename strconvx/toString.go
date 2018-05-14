package strconvx

import (
	"fmt"
)

// ToString fmt.Sprint(a...).
func ToString(a ...interface{}) string {
	return fmt.Sprint(a...)
}
