package moduletest

import (
	"github.com/opentofu/opentofu/external/configs"
	"github.com/opentofu/opentofu/external/tfdiags"
)

type File struct {
	Config *configs.TestFile

	Name   string
	Status Status

	Runs []*Run

	Diagnostics tfdiags.Diagnostics
}
