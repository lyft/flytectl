// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package workflow

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.StringVar(&(DefaultConfig.Version), fmt.Sprintf("%v%v", prefix, "version"), *new(string), "version of the workflow to be fetched.")
	cmdFlags.BoolVar(&(DefaultConfig.Latest), fmt.Sprintf("%v%v", prefix, "latest"), *new(bool), " flag to indicate to fetch the latest version,  version flag will be ignored in this case")
	cmdFlags.StringVar(&(DefaultConfig.Filter.FieldSelector), fmt.Sprintf("%v%v", prefix, "filter.field-selector"), *new(string), "Specifies the Field selector")
	cmdFlags.StringVar((&DefaultConfig.Filter.SortBy), fmt.Sprintf("%v%v", prefix, "filter.sort-by"), *new(string), "Specifies which field to sort result by ")
	cmdFlags.Int32Var((&DefaultConfig.Filter.Limit), fmt.Sprintf("%v%v", prefix, "filter.limit"), 100, "Specifies the limit")
	cmdFlags.BoolVar((&DefaultConfig.Filter.Asc), fmt.Sprintf("%v%v", prefix, "filter.asc"), false, "Specifies the sorting order. By default flytectl sort result in descending order")

	return cmdFlags
}
