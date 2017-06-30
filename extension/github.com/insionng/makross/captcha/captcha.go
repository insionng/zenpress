package captcha

import (
	"github.com/insionng/makross/captcha"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/makross/captcha",

	"Captchaer": captcha.Captchaer,
	"Version":   captcha.Version,

	"Captcha":    spec.StructOf((*captcha.Captcha)(nil)),
	"NewCaptcha": captcha.NewCaptcha,
	"Store":      captcha.Store,
	"Image":      spec.StructOf((*captcha.Image)(nil)),
	"NewImage":   captcha.NewImage,
	"Options":    spec.StructOf((*captcha.Options)(nil)),
}
