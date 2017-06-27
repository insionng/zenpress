package switchr

import (
	"github.com/insionng/zenpress/module/switchr"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/zenpress/module/switchr",

	"DefaultSwitchrConfig": switchr.DefaultSwitchrConfig,

	"Switchr":           switchr.Switchr,
	"SwitchrWithConfig": switchr.SwitchrWithConfig,

	"SwitchrConfig": spec.StructOf((*switchr.SwitchrConfig)(nil)),
}
