package main

import (
	"log"
	"os"

	"github.com/starkandwayne/pivy/commands"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var Version string

func main() {
	logger := log.New(os.Stdout, "", 0)
	app := kingpin.New("pivy", "PivNets little helper")
	app.Version(Version)
	commands.Configure(logger, app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
