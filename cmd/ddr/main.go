package main

import (
	"os"

	"git.home.jm0.eu/josua/dd-remote/client"
	"github.com/jessevdk/go-flags"
)

// Define CLI options
type Options struct {
	URI   string `short:"u" long:"uri" default:"http://localhost/" description:"dd-remote server address"`
	Input string `short:"i" long:"input" default:"/dev/stdin" description:"input file to send" `
}

func main() {
	var err error
	var options Options
	var status bool

	// parse cli options
	_, err = flags.Parse(&options)
	if err != nil {
		os.Exit(1)
	}

	// invoke client
	status = client.Write(options.URI, options.Input)

	// exit with appropriate error status
	if !status {
		os.Exit(1)
	}
}
