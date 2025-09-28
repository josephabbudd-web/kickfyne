package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_framework_ "github.com/josephabbudd-web/kickfyne/commands/framework"
	_frontend_ "github.com/josephabbudd-web/kickfyne/commands/frontend"
	_help_ "github.com/josephabbudd-web/kickfyne/commands/help"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

func main() {

	// Build the args to pass on to the handlers.
	lArgs := len(os.Args)
	if lArgs < 2 {
		fmt.Println(_help_.Usage())
		return
	}

	var err error
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer func() {
		ctxCancel()
		if err != nil {
			log.Println("Error: ", err.Error())
			os.Exit(1)
		}
	}()

	var pathWD string
	if pathWD, err = os.Getwd(); err != nil {
		return
	}

	var folderPaths *_utils_.FolderPaths
	if folderPaths, err = _utils_.NewFolderPaths(pathWD); err != nil {
		return
	}
	go notify(ctx, ctxCancel)

	// IF the current folder does not contain the go.mod file then display help.
	var importPrefix string
	if importPrefix, err = _utils_.ImportPrefix(pathWD); err != nil {
		// Display help and return.
		err = _help_.Handler(nil)
		return
	}

	// The user is in the correct folder so proceed.
	switch os.Args[1] {
	case _framework_.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = _framework_.Handler(handlerArgs, importPrefix, folderPaths)
	case _frontend_.CmdScreen:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[1:]
		}
		err = _frontend_.Handler(handlerArgs, importPrefix, folderPaths)
	case _help_.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = _help_.Handler(handlerArgs)
	default:
		fmt.Println(_help_.Usage())
	}
}

func notify(ctx context.Context, ctxCancel context.CancelFunc) {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-signalChan:
			ctxCancel()
			// terminate after second signal before callback is done
			go func() {
				<-signalChan
				os.Exit(1)
			}()
			return
		}
	}
}
