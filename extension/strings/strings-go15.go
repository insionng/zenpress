// +build go1.5

package strings

import "strings"

func init() {
	Exports["Compare"] = strings.Compare
	Exports["LastIndexByte"] = strings.LastIndexByte
}
