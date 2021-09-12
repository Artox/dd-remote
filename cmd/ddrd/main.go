package main

import (
	"os"

	"git.home.jm0.eu/josua/dd-remote/server"
	"github.com/jessevdk/go-flags"
)

// Define CLI options
type Options struct {
	ListenHost string `short:"l" long:"listen-host" default:"localhost" description:"address to listen on for incoming connections"`
	ListenPort uint   `short:"p" long:"listen-port" default:"80" description:"port to listen on for incoming connections"`
	Output     string `short:"o" long:"output" default:"/dev/stdout" description:"destination file for writing to"`
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

	// start server
	status = server.Start(options.ListenHost, options.ListenPort, options.Output)

	// exit with appropriate error status
	if !status {
		os.Exit(1)
	}
}
