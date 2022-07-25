package pretty

import (
	"bytes"
	"fmt"
	"strconv"
)

const supportedFlags = "0-+# "

func constructOrigFormat(f fmt.State, verb rune) (format string) {
	buf := bytes.NewBuffer([]byte("%"))

	for _, flag := range supportedFlags {
		if f.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}

	if width, ok := f.Width(); ok {
		buf.WriteString(strconv.Itoa(width))
	}

	if precision, ok := f.Precision(); ok {
		buf.WriteString(".")
		buf.WriteString(strconv.Itoa(precision))
	}

	buf.WriteRune(verb)

	format = buf.String()
	return format
}
