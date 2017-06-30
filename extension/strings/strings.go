package strings

import (
	"strings"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "strings",

	"Contains":       strings.Contains,
	"ContainsAny":    strings.ContainsAny,
	"ContainsRune":   strings.ContainsRune,
	"Count":          strings.Count,
	"EqualFold":      strings.EqualFold,
	"Fields":         strings.Fields,
	"FieldsFunc":     strings.FieldsFunc,
	"HasPrefix":      strings.HasPrefix,
	"HasSuffix":      strings.HasSuffix,
	"Index":          strings.Index,
	"IndexAny":       strings.IndexAny,
	"IndexByte":      strings.IndexByte,
	"IndexFunc":      strings.IndexFunc,
	"IndexRune":      strings.IndexRune,
	"Join":           strings.Join,
	"LastIndex":      strings.LastIndex,
	"LastIndexAny":   strings.LastIndexAny,
	"LastIndexFunc":  strings.LastIndexFunc,
	"Map":            strings.Map,
	"Repeat":         strings.Repeat,
	"Replace":        strings.Replace,
	"Split":          strings.Split,
	"SplitAfter":     strings.SplitAfter,
	"SplitAfterN":    strings.SplitAfterN,
	"SplitN":         strings.SplitN,
	"Title":          strings.Title,
	"ToLower":        strings.ToLower,
	"ToLowerSpecial": strings.ToLowerSpecial,
	"ToTitle":        strings.ToTitle,
	"ToTitleSpecial": strings.ToTitleSpecial,
	"ToUpper":        strings.ToUpper,
	"ToUpperSpecial": strings.ToUpperSpecial,
	"Trim":           strings.Trim,
	"TrimFunc":       strings.TrimFunc,
	"TrimLeft":       strings.TrimLeft,
	"TrimLeftFunc":   strings.TrimLeftFunc,
	"TrimPrefix":     strings.TrimPrefix,
	"TrimRight":      strings.TrimRight,
	"TrimRightFunc":  strings.TrimRightFunc,
	"TrimSpace":      strings.TrimSpace,
	"TrimSuffix":     strings.TrimSuffix,

	"Reader":      spec.StructOf((*strings.Reader)(nil)),
	"NewReader":   strings.NewReader,
	"Replacer":    spec.StructOf((*strings.Replacer)(nil)),
	"NewReplacer": strings.NewReplacer,
}
