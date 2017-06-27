package pongor

import (
	"github.com/insionng/makross/pongor"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/makross/pongor",

	"Option":   spec.StructOf((*pongor.Option)(nil)),
	"Renderer": spec.StructOf((*pongor.Renderer)(nil)),
	"Renderor": pongor.Renderor,
}
