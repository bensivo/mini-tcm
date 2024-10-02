package main

import (
	"github.com/alexflint/go-arg"
	"github.com/bensivo/mini-tcm/packages/webapp/pkg/service"
)

func main() {
	var args struct {
		Dir string `arg:"--dir,-d,required" help:"The directory to load test cases from"`
	}
	arg.MustParse(&args)

	service.LoadTestCasesFromDir(args.Dir)
}
