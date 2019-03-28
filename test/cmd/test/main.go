package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/phogolabs/cli"
	"github.com/phogolabs/cli/provider/aws/s3"
	"github.com/phogolabs/log"
	"github.com/phogolabs/log/handler/json"
	"github.com/phogolabs/log/handler/rollbar"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Build number
	Build = "dev"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:   "log-level",
		Value:  "INFO",
		Usage:  "level of logging",
		EnvVar: "APP_API_LOG_LEVEL",
	},
	&cli.StringFlag{
		Name:   "rollbar-token",
		Value:  "",
		Usage:  "Rollbar Token",
		EnvVar: "APP_API_ROLLBAR_TOKEN",
	},
	&cli.StringFlag{
		Name:   "rollbar-environment",
		Value:  "development",
		Usage:  "Rollbar Environment",
		EnvVar: "APP_API_ROLLBAR_ENVIRONMENT",
	},
	&cli.StringFlag{
		Name:   "metrics-listen-addr",
		Value:  ":8282",
		Usage:  "listen address of metrics HTTP server",
		EnvVar: "APP_API_METRICS_LISTEN_ADDR",
	},
}

func main() {
	app := &cli.App{
    Name:       "pet-store",
    HelpName:   "pet-store",
    Usage:      "pet-store http service",
    UsageText:  "pet-store [global options]",
    Version:    Build,
		Writer:     os.Stdout,
		ErrWriter:  os.Stderr,
		Flags:      flags,
		OnSignal:   term,
		Action:     run,
		AfterInit:  afterInit,
		BeforeInit: beforeInit,
		Signals:    []os.Signal{syscall.SIGTERM},
		Providers:  []cli.Provider{&s3.Provider{}},
	}

	app.Run(os.Args)
}

func beforeInit(ctx *cli.Context) error {
	level, err := log.ParseLevel(ctx.String("log-level"))
	if err != nil {
		return err
	}

	log.SetLevel(level)

	fields := log.FieldMap{
		"app":     ctx.Command.Name,
		"version": ctx.Manifest.Version,
	}

	log.SetDefaultFields(fields)
	return nil
}

func afterInit(ctx *cli.Context) error {
	jhandler := json.New(os.Stdout)

	rhandler := &log.LevelHandler{
		Level: log.WarnLevel,
		Handler: rollbar.New(&rollbar.Config{
			CodeVersion: ctx.Manifest.Version,
			Token:       ctx.String("rollbar-token"),
			Environment: ctx.String("rollbar-environment"),
			ServerHost:  os.Getenv("HOSTNAME"),
      ServerRoot:  "https://github.com/phogolabs/pet-store",
		}),
	}

	handler := log.CompositeHandler{
		jhandler,
		rhandler,
	}

	log.SetHandler(handler)
	return nil
}

func run(ctx *cli.Context) error {
	go http.ListenAndServe(ctx.String("metrics-listen-addr"), promhttp.Handler())
	return nil
}

func term(ctx *cli.Context) error {
	return nil
}
