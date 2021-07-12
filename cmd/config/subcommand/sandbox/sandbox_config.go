package sandbox

//go:generate pflags Config --default-var DefaultConfig --bind-default-var
var (
	DefaultConfig = &Config{}
)

//Config
type Config struct {
	Source    string `json:"source" pflag:",Path of your source code"`
	Kustomize string `json:"kustomize" pflag:",kustomize file path"`
	Version   string `json:"version" pflag:",Version of flyte"`
}
