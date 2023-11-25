package utils

import (
	cliv1 "github.com/codegangsta/cli"
	"github.com/urfave/cli/v2"
)

func AsV2IntFlag(v1Flag cliv1.IntFlag) *cli.IntFlag {
	var envVars []string
	if v1Flag.EnvVar != "" {
		envVars = []string{v1Flag.EnvVar}
	}
	return &cli.IntFlag{
		Name:        v1Flag.Name,
		Usage:       v1Flag.Usage,
		Value:       v1Flag.Value,
		EnvVars:     envVars,
		Destination: v1Flag.Destination,
	}
}

func AsV2StringFlag(v1Flag cliv1.StringFlag) *cli.StringFlag {
	var envVars []string
	if v1Flag.EnvVar != "" {
		envVars = []string{v1Flag.EnvVar}
	}
	return &cli.StringFlag{
		Name:        v1Flag.Name,
		Usage:       v1Flag.Usage,
		Value:       v1Flag.Value,
		EnvVars:     envVars,
		Destination: v1Flag.Destination,
	}
}

func AsV2StringSliceFlag(v1Flag cliv1.StringSliceFlag) *cli.StringSliceFlag {
	var envVars []string
	if v1Flag.EnvVar != "" {
		envVars = []string{v1Flag.EnvVar}
	}
	var values []string
	if v1Flag.Value != nil {
		values = *v1Flag.Value
	}
	return &cli.StringSliceFlag{
		Name:        v1Flag.Name,
		Usage:       v1Flag.Usage,
		Value:       cli.NewStringSlice(values...),
		EnvVars:     envVars,
	}
}

