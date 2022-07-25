package pretty

import (
	"encoding/json"
	"fmt"
)

func JSON(v interface{}) fmt.Formatter {
	return jsonFormatter{v: v}
}

type jsonFormatter struct {
	v interface{}
}

var _ fmt.Formatter = jsonFormatter{}

func (p jsonFormatter) Format(f fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		var (
			content []byte
			err     error
		)
		if f.Flag('#') {
			content, err = json.MarshalIndent(p.v, "", "  ")
		} else {
			content, err = json.Marshal(p.v)
		}
		if err == nil {
			fmt.Fprintf(f, "%s", content)
			break
		}
		fallthrough
	default:
		// fall back to default formatter
		fmt.Fprintf(f, constructOrigFormat(f, verb), p.v)
	}
}
