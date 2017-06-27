package switchr

import (
	"fmt"
	"strings"

	"github.com/insionng/makross"
	"github.com/insionng/makross/pongor"
	"github.com/insionng/makross/skipper"
)

type (
	// SwitchrConfig defines the config for Switchr middleware.
	SwitchrConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper skipper.Skipper

		Theme          string
		Filter, Reload bool
	}
)

var (
	// DefaultSwitchrConfig is the default Switchr middleware config.
	DefaultSwitchrConfig = SwitchrConfig{
		Skipper: skipper.DefaultSkipper,
	}
)

// Switchr returns a Switchr middleware to serves Switchr content from the provided
// theme directory.
func Switchr(theme string) makross.Handler {
	c := DefaultSwitchrConfig
	c.Theme = theme
	return SwitchrWithConfig(c)
}

// SwitchrWithConfig returns a Switchr middleware with config.
// See `Switchr()`.
func SwitchrWithConfig(config SwitchrConfig) makross.Handler {
	// Defaults
	if config.Theme == "" {
		config.Theme = "default" // For security we want to restrict to default theme.
	}
	if config.Skipper == nil {
		config.Skipper = DefaultSwitchrConfig.Skipper
	}

	return func(c *makross.Context) error {
		if config.Skipper(c) {
			return c.Next()
		}

		if strings.HasPrefix(c.Request.URL.Path, "/root/") {
			c.Makross().SetRenderer(pongor.Renderor(pongor.Option{Directory: "template", Reload: config.Reload, Filter: config.Filter}))
		} else {
			c.Makross().SetRenderer(pongor.Renderor(pongor.Option{Directory: fmt.Sprintf("content/theme/%s/template", config.Theme), Reload: config.Reload, Filter: config.Filter}))
		}

		return nil
	}

}
