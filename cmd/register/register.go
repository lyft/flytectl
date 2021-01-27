package register

import (
	cmdcore "github.com/lyft/flytectl/cmd/core"
	"github.com/spf13/cobra"
)

// RegisterCommand will return register command
func RegisterCommand() *cobra.Command {
	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Registers tasks/workflows from a given location of generated serialized files.",
	}

	registerResourcesFuncs := map[string]cmdcore.CommandEntry{
		"files":    {CmdFunc: registerFromFilesFunc, Aliases: []string{"file"}},
	}

	cmdcore.AddCommands(registerCmd, registerResourcesFuncs)
	registerCmd.PersistentFlags().StringVarP(&(GetConfig().version), "version", "v", "1", "Specifies the Version to use for registering the serialized files.")
	return registerCmd
}
