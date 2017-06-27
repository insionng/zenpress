package logger

import (
	"github.com/insionng/makross/logger"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/makross/logger",

	"DefaultLoggerConfig": logger.DefaultLoggerConfig,

	"Logger":           logger.Logger,
	"LoggerWithConfig": logger.LoggerWithConfig,

	"LoggerConfig": spec.StructOf((*logger.LoggerConfig)(nil)),
}
