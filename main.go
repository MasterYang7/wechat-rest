package main

import (
	"embed"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/httpd"
)

//go:embed public
var efs embed.FS

func main() {

	args.Efs = &efs

	// dbase.Connect()

	// crond.Daemon()
	// plugin.CronjobPluginSetup()
	// plugin.KeywordPluginSetup()

	// robot.Start()

	httpd.Server()

}
