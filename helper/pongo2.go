package helper

import (
	"encoding/base64"
	"github.com/flosch/pongo2"
	"strings"
	"time"
)

func ConvertToBase64(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func ConvertToBase64ByPongo2(in *pongo2.Value) *pongo2.Value {
	return pongo2.AsValue(base64.StdEncoding.EncodeToString([]byte(in.String())))
}

func SplitByPongo2(in *pongo2.Value, splitor *pongo2.Value) *pongo2.Value {
	return pongo2.AsValue(Split(in.String(), splitor.String()))
}

func MarkdownByPongo2(in *pongo2.Value) *pongo2.Value {
	return pongo2.AsValue(Markdown(in.String()))
}

func CropwordByPongo2(in *pongo2.Value, start *pongo2.Value, length *pongo2.Value, symbol *pongo2.Value) *pongo2.Value {
	return pongo2.AsValue(Substr(in.String(), start.Integer(), length.Integer(), symbol.String()))
}

func Cropword(in string, start int, length int, symbol string) string {
	return Substr(in, start, length, symbol)
}

func File(s string) string {
	if len(s) > 0 {
		if strings.HasPrefix(s, "http") || strings.HasPrefix(s, "/identicon") {
			return s
		} else {
			return "/file" + s
		}
	}
	return s
}

func Unix2TimeByPongo2(in *pongo2.Value, timeLayout *pongo2.Value) *pongo2.Value {
	return pongo2.AsValue(time.Unix(int64(in.Integer()), 0).Format(timeLayout.String()))
}
