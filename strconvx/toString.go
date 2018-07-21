package strconvx

import (
	"fmt"
)

// ToString calls fmt.Sprint(a...).
func ToString(a ...interface{}) string {
	return fmt.Sprint(a...)
}
