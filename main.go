package main

import (
	"github.com/alecthomas/kong"
	"io/ioutil"
	"log"
	"os"
)

var App struct {
	List  LsCommand   `cmd help:"List indices"`
	Dump  DumpCommand `cmd help:"Dump indices"`
	Url   string      `required name:"url" help:"Base Kibana/ES url"`
	Quiet bool        `short:"q" help:"Quiet mode" default:"false"`
}

func main() {
	var err error
	ctx := kong.Parse(&App)
	queryDispatcher := &EsQueryDispatcher{
		BaseUrl:   App.Url,
		LogOutput: os.Stderr,
	}
	if App.Quiet {
		queryDispatcher.LogOutput = ioutil.Discard
	}
	log.SetOutput(queryDispatcher.LogOutput)
	err = queryDispatcher.DetectEsVersion()
	if err != nil {
		ctx.FatalIfErrorf(err)
	}
	err = ctx.Run(queryDispatcher)
	ctx.FatalIfErrorf(err)
}
