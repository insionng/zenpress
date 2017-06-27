package static

import (
	"github.com/insionng/makross/static"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/makross/static",

	"DefaultStaticConfig": static.DefaultStaticConfig,

	"Static":           static.Static,
	"StaticWithConfig": static.StaticWithConfig,

	"StaticConfig": spec.StructOf((*static.StaticConfig)(nil)),
}
