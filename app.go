package main

import (

	// Dependencies of Turbine
	"github.com/meroxa/turbine-go"
	"github.com/meroxa/turbine-go/runner"
)

func main() {
	runner.Start(App{})
}

var _ turbine.App = (*App)(nil)

type App struct{}

func (a App) Run(v turbine.Turbine) error {
	source, err := v.Resources("cck")
	if err != nil {
		return err
	}

	rr, err := source.Records("inbound", nil)
	if err != nil {
		return err
	}

	res := v.Process(rr, Format{})

	dest, err := v.Resources("demopg")
	if err != nil {
		return err
	}

	err = dest.Write(res, "events")
	if err != nil {
		return err
	}

	return nil
}

type Format struct{}

func (f Format) Process(stream []turbine.Record) []turbine.Record {
	for i, record := range stream {
	}
	return stream
}
