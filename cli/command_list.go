package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var list = cli.Command{
	Name:   "list",
	Usage:  "Lists all the packages in the index.",
	Action: listAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "filter",
			Usage: "Filter packages.",
		},
	},
}

var (
	//TODO: Trim the end slashes spaces.
	packagesListTemplate = ` Listing Packages From: {{.Location}}{{if .Packages}}
{{range .Packages}}
   {{ . }}{{ end }}{{else}} 
No Packages Found.{{end}}
`
)

func listAction(c *cli.Context) {
	//TODO: Prettify this.

	arg := c.Args().First()

	if arg == "" {
		arg = "*"
	}

	repo.SetFilter(arg)

	err := repo.Update()
	if err != nil {
		Log(err, true, LOG_ERR)
	}

	templates.New("packageslist").Parse(packagesListTemplate)

	err = templates.ExecuteTemplate(os.Stdout, "packageslist", repo)
	if err != nil {
		Log(err, true, LOG_ERR)
	}
}