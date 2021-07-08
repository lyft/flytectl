// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package sandbox

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.StringVar(&DefaultConfig.Source, fmt.Sprintf("%v%v", prefix, "source"), DefaultConfig.Source, " Path of your source code")
	cmdFlags.StringVar(&DefaultConfig.Version, fmt.Sprintf("%v%v", prefix, "version"), DefaultConfig.Version, "Version of flyte")
	return cmdFlags
}
