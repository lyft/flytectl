package sandbox

import (
	"context"
	"fmt"
	"strings"

	"github.com/enescakir/emoji"

	"github.com/docker/docker/client"
	cmdCore "github.com/flyteorg/flytectl/cmd/core"
)

const (
	teardownShort = "Teardown will cleanup the sandbox environment"
	teardownLong  = `
Teardown will remove docker container and all the flyte config 
::

 bin/flytectl sandbox teardown 
	

Usage
`
)

var defaultInput = strings.NewReader("y")

func teardownSandboxCluster(ctx context.Context, args []string, cmdCtx cmdCore.CommandContext) error {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = removeIfSandboxExist(cli, defaultInput)
	if err != nil {
		return err
	}

	if err := configCleanup(); err != nil {
		fmt.Printf("Config cleanup failed. Which Failed due to %v \n ", err)
	}
	fmt.Printf("Sandbox cluster is removed successfully %v \n", emoji.Rocket)
	return nil
}
