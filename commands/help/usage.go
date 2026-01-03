package help

import (
	_framework_ "github.com/josephabbudd-web/kickfyne/commands/framework"
	_screen_ "github.com/josephabbudd-web/kickfyne/commands/screen"
	_version_ "github.com/josephabbudd-web/kickfyne/commands/version"
)

const (
	newParagraph   = "\n\n"
	gettingStarted = `🍻 INTRODUCING kickfyne!
kickfyne is a tool to help build an application using the fyne toolkit which has among other things a very nice GUI. The fyne toolkit web site is located at https://fyne.io/. kickfyne is not in any way associated with the fyne projects.
`
)

func Usage() (usage string) {
	usage =
		_version_.V() + newParagraph +
			gettingStarted + newParagraph +
			_framework_.Usage() + newParagraph +
			_screen_.Usage() + newParagraph
	return
}
