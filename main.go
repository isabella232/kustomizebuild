package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "kustomize"
	app.Version = VERSION
	app.Usage = "Helm chart customization"
	app.Action = func(c *cli.Context) error {
		logrus.Info("Issues with app")
		return nil
	}

	app.Run(os.Args)
}
