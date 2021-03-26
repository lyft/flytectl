package version

import (
	"context"
	"fmt"
	"os"

	adminclient "github.com/flyteorg/flyteidl/clients/go/admin"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flytestdlib/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Long descriptions are whitespace sensitive when generating docs using sphinx.
const (
	versionCmdShort = `Used for fetching flyte version`
	versionCmdLong  = `
Example version.
::

 bin/flytectl version
`
)

// VersionCommand will return version of flyte
func GetVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   versionCmdShort,
		Aliases: []string{"versions"},
		Long:    versionCmdLong,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			adminClient, err := adminclient.InitializeAdminClientFromConfig(ctx)
			if err != nil {
				fmt.Printf("err %v:", err)
				os.Exit(1)
			}
			v, err := adminClient.GetVersion(ctx, &admin.GetVersionRequest{})
			if err != nil {
				fmt.Printf("err %v:", err)
				os.Exit(1)
			}
			version.LogBuildInformation("flytectl")
			PrintVersion("flyteadmin", v)
		},
	}

	return versionCmd
}

func PrintVersion(appName string, res *admin.GetVersionResponse) {
	logrus.Info("------------------------------------------------------------------------")
	logrus.Infof("App [%s], Version [%s], BuildSHA [%s], BuildTS [%s]", appName, res.ControlPlaneVersion.Version, res.ControlPlaneVersion.Build, res.ControlPlaneVersion.BuildTime)
	logrus.Info("------------------------------------------------------------------------")
}
