package editor

import (
	"io"
	"os"

	"github.com/cli/cli/v2/pkg/surveyext"
	"github.com/cli/go-gh/pkg/config"
)

func Edit(initialValue string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (string, error) {
	editorCmd := os.Getenv("GH_EDITOR")
	if editorCmd == "" {
		cfg, err := config.Read()
		if err != nil {
			return "", err
		}
		// ignore KeyNotFoundError
		editorCmd, _ = cfg.Get([]string{"editor"})
	}
	return surveyext.Edit(editorCmd, "*.md", initialValue, stdin, stdout, stderr)
}
