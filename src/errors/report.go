package errors

import (
	"os"

	"github.com/Spriithy/go-colors"
)

// Report is used to report a runtime error
// This process halts execution
func Report(msg string) {
	println(colors.Red(colors.None, "Error report")+":", msg)
	os.Exit(1)
}
