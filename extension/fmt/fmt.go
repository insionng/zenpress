package fmt

import (
	"fmt"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fmt",

	"Errorf":   fmt.Errorf,
	"Fprint":   fmt.Fprint,
	"Fprintf":  fmt.Fprintf,
	"Fprintln": fmt.Fprintln,
	"Fscan":    fmt.Fscan,
	"Fscanf":   fmt.Fscanf,
	"Fscanln":  fmt.Fscanln,
	"Print":    fmt.Print,
	"Printf":   fmt.Printf,
	"Println":  fmt.Println,
	"Scan":     fmt.Scan,
	"Scanf":    fmt.Scanf,
	"Scanln":   fmt.Scanln,
	"Sprint":   fmt.Sprint,
	"Sprintf":  fmt.Sprintf,
	"Sprintln": fmt.Sprintln,
	"Sscan":    fmt.Sscan,
	"Sscanf":   fmt.Sscanf,
	"Sscanln":  fmt.Sscanln,
}
