package main

import (
	"github.com/bensivo/mini-tcm/packages/webapp/pkg/server"
	"github.com/spf13/pflag"
)

func main() {
	var dir = pflag.StringP("dir", "d", ".", "The directory to load test cases from")
	pflag.Parse()

	svr := server.Server{
		Port:        8080,
		TestCaseDir: *dir,
	}
	svr.Serve()
}
