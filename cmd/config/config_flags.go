// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "project"), defaultConfig.Project, "Specifies the project to work on.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "domain"), defaultConfig.Domain, "Specified the domain to work on.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "output"), defaultConfig.Domain, "Specified the output type.")
	return cmdFlags
}
