// Copyright 2016 - Authors included on AUTHORS file.
//
// Use of this source code is governed by a Apache License
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"github.com/apuigsech/seekret"
	"github.com/apuigsech/seekret-source-dir"
	"github.com/urfave/cli"
)

const (
	DefaultCommitCount = 10
)

var s *seekret.Seekret

func main() {
	s = seekret.NewSeekret()

	app := cli.NewApp()

	app.Name = "seekret"
	app.Version = "0.0.1"
	app.Usage = "seek for secrets on various sources."

	app.Author = "Albert Puigsech Galicia"
	app.Email = "albert@puigsech.com"

	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "exception, x",
			Usage: "load exceptions from `FILE`.",
		},
		cli.StringFlag{
			Name:   "rules",
			Usage:  "`PATH` with rules.",
			EnvVar: "SEEKRET_RULES_PATH",
		},
		cli.StringFlag{
			Name:  "format, f",
			Usage: "specify the output format.",
			Value: "human",
		},
		// TODO: To be implemented.
		/*
			cli.StringFlag{
				Name: "groupby, g",
				Usage: "Group output by specific field",
			},
		*/
		cli.StringFlag{
			Name:  "known, k",
			Usage: "load known secrets from `FILE`.",
		},
		cli.IntFlag{
			Name:  "workers, w",
			Usage: "number of workers used for the inspection",
			Value: 4,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "dir",
			Usage:    "seek for seecrets on a directory",
			Category: "seek",
			Action:   seekretDir,

			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "recursive, r",
				},
				cli.BoolFlag{
					Name: "hidden",
				},
			},
		},
	}

	app.Before = seekretBefore
	app.After = seekretAfter

	app.Run(os.Args)
}

func seekretBefore(c *cli.Context) error {
	var err error

	rulesPath := c.String("rules")

	err = s.LoadRulesFromPath(rulesPath, true)
	if err != nil {
		return err
	}

	LoadKnownFromFile(s, c.String("known"))

	err = s.LoadExceptionsFromFile(c.String("exception"))
	if err != nil {
		return err
	}

	return nil
}

func seekretDir(c *cli.Context) error {
	source := c.Args().Get(0)
	if source == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	options := map[string]interface{}{
		"hidden":    c.Bool("hidden"),
		"recursive": c.Bool("recursive"),
	}

	err := s.LoadObjects(sourcedir.SourceTypeDir, source, options)
	if err != nil {
		return err
	}

	return nil
}

func seekretAfter(c *cli.Context) error {
	s.Inspect(c.Int("workers"))

	fmt.Println(FormatOutput(s.ListSecrets(), c.String("format")))

	return nil
}
