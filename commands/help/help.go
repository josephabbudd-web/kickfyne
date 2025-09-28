package help

import (
	"fmt"

	_frontend_ "github.com/josephabbudd-web/kickfyne/commands/frontend"
)

const (
	Cmd = "help"
)

// Handler displays the requested help.
func Handler(args []string) (err error) {

	if len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("help.Handler: %w", err)
		}
	}()

	switch args[0] {
	case _frontend_.CmdScreen:
		fmt.Println(_frontend_.UsageScreen())
	default:
		Usage()
	}
	return
}
