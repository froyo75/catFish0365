package main

import (
	"catFish0365/libs"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Context struct {
	Debug bool
}
type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

var cli struct {
	Clientid string      `short:"c" name:"clientid" optional:"" default:"d3590ed6-52b3-4102-aeff-aad2292ab01c" help:"The Application (client) ID of an Azure App to target (default: d3590ed6-52b3-4102-aeff-aad2292ab01c)."`
	Resource string      `short:"r" name:"resource" optional:"" default:"https://graph.windows.net" help:"The resource to request access tokens."`
	OutPath  string      `optional:"" short:"o" help:"Logging output to a specific file."`
	Version  VersionFlag `name:"version" short:"v" help:"Print version information and quit."`

	Exploitdcauth struct {
	} `cmd:"" help:"Launch device code phishing attack."`

	Getaccesstoken struct {
		RefreshToken string `arg:"" name:"refreshtoken" help:"The refresh token to provide." type:"refreshtoken"`
		Scope        string `short:"s" name:"scope" optional:"" default:"openid" help:"The scope parameter to explicitly request a refresh token (default: openid)."`
	} `cmd:"" help:"Get a new access token for the given clientid and resource using the given refresh token."`
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "02 Jan 2006 15:04:05"})
	ctx := kong.Parse(&cli, kong.Name("catFish0365"), kong.Description("A tool to automate device code phishing attacks by @froyo75"), kong.UsageOnError(), kong.Vars{"version": "1.0"})
	if cli.OutPath != "" {
		libs.Logging(cli.OutPath)
	}
	switch ctx.Command() {
	case "exploitdcauth":
		libs.ExploitDeviceCodeAuth(cli.Clientid, cli.Resource)
	case "getaccesstoken <refreshtoken>":
		libs.GetNewAccessToken(cli.Clientid, cli.Resource, cli.Getaccesstoken.Scope, cli.Getaccesstoken.RefreshToken)
	default:
		panic(ctx.Command())
	}
}
