package session

import (
	"github.com/insionng/makross/session"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/makross/session",

	"CONTEXT_FLASH_KEY":   session.CONTEXT_FLASH_KEY,
	"CONTEXT_SESSION_KEY": session.CONTEXT_SESSION_KEY,
	"COOKIE_FLASH_KEY":    session.COOKIE_FLASH_KEY,
	"SESSION_FLASH_KEY":   session.SESSION_FLASH_KEY,
	"SESSION_INPUT_KEY":   session.SESSION_INPUT_KEY,

	"GlobalManager": session.GlobalManager,

	"DecodeGob":         session.DecodeGob,
	"EncodeGob":         session.EncodeGob,
	"FlashValue":        session.FlashValue,
	"GetFlash":          session.GetFlash,
	"NewFlash":          session.NewFlash,
	"RandomCreateBytes": session.RandomCreateBytes,
	"Register":          session.Register,
	"Sessioner":         session.Sessioner,

	"GetStore": session.GetStore,

	"CookieProvider":     spec.StructOf((*session.CookieProvider)(nil)),
	"CookieSessionStore": spec.StructOf((*session.CookieSessionStore)(nil)),
	"FileProvider":       spec.StructOf((*session.FileProvider)(nil)),
	"FileSessionStore":   spec.StructOf((*session.FileSessionStore)(nil)),
	"Manager":            spec.StructOf((*session.Manager)(nil)),
	"NewManager":         session.NewManager,
	"MemProvider":        spec.StructOf((*session.MemProvider)(nil)),
	"MemSessionStore":    spec.StructOf((*session.MemSessionStore)(nil)),
	"Options":            spec.StructOf((*session.Options)(nil)),
}
