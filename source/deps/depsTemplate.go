package deps

type depsTemplateData struct {
	ImportPrefix string
}

const (
	depsFileName = "deps.go"

	depsTemplate = `package deps

import (
	"context"
	"fmt"

	_paths_ "{{ .ImportPrefix }}/deps/paths"
)

// Start starts the deps.
func Start(ctx context.Context, ctxCancel context.CancelFunc) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("deps.Start: %w", err)
		}
	}()

	// App _paths_.
	if err = _paths_.Init(); err != nil {
		return
	}
	return
}
`
)
