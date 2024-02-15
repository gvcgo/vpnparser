package cmd

import (
	"fmt"
	"os"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gvcgo/vpnparser/pkgs/outbound"
	cli "github.com/urfave/cli/v2"
)

type App struct {
	cmd *cli.App
}

func New() *App {
	return &App{
		cmd: &cli.App{
			Usage:       "vpnparser <Command> <SubCommand>...",
			Description: "vpnparser, download files from github for gvc.",
			Commands:    []*cli.Command{},
		},
	}
}

func (that *App) Add(command *cli.Command) {
	that.cmd.Commands = append(that.cmd.Commands, command)
}

func (that *App) Run() {
	that.cmd.Run(os.Args)
}

var app *App

func ShowOutboundStr(oStr string) {
	j := gjson.New(oStr)
	fmt.Println(j.MustToJsonIndentString())
}

func init() {
	app = New()
	app.Add(&cli.Command{
		Name:    "sing",
		Aliases: []string{"s"},
		Usage:   "Generate sing-box outbound from vpn url.",
		Action: func(ctx *cli.Context) error {
			rawUri := ctx.Args().First()
			if rawUri == "" {
				return nil
			}
			ob := outbound.GetOutbound(outbound.SingBox, rawUri)
			ob.Parse(rawUri)
			fmt.Println(rawUri)
			ShowOutboundStr(ob.GetOutboundStr())
			return nil
		},
	})

	app.Add(&cli.Command{
		Name:    "xray",
		Aliases: []string{"x"},
		Usage:   "Generate xray-core outbound from vpn url.",
		Action: func(ctx *cli.Context) error {
			rawUri := ctx.Args().First()
			if rawUri == "" {
				return nil
			}
			ob := outbound.GetOutbound(outbound.XrayCore, rawUri)
			ob.Parse(rawUri)
			fmt.Println(rawUri)
			ShowOutboundStr(ob.GetOutboundStr())
			return nil
		},
	})
}

func StartApp() {
	app.Run()
}
