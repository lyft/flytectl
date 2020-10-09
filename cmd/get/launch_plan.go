package get

import (
	"context"
	"encoding/json"
	"github.com/lyft/flytectl/cmd/config"
	cmdCore "github.com/lyft/flytectl/cmd/core"
	"github.com/lyft/flytectl/pkg/adminutils"
	"github.com/lyft/flytectl/pkg/printer"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/lyft/flytestdlib/logger"
)

var launchPlanStructure = map[string]string{
	"Version": "$.id.version",
	"Name":    "$.id.name",
	"Type":    "$.closure.compiledTask.template.type",
}

func transformLaunchPlan(jsonbody []byte) (interface{}, error) {
	results := PrintableLaunchPlan{}
	if err := json.Unmarshal(jsonbody, &results); err != nil {
		return results, err
	}
	return results, nil
}

func getLaunchPlanFunc(ctx context.Context, args []string, cmdCtx cmdCore.CommandContext) error {
	launchPlanPrinter := printer.Printer{}


	if len(args) == 1 {
		name := args[0]
		launchPlan, err := cmdCtx.AdminClient().ListLaunchPlans(ctx, &admin.ResourceListRequest{
			Limit: 10,
			Id: &admin.NamedEntityIdentifier{
				Project: config.GetConfig().Project,
				Domain:  config.GetConfig().Domain,
				Name : name,
			},
		})
		if err != nil {
			return err
		}
		logger.Debugf(ctx, "Retrieved %v excutions", len(launchPlan.LaunchPlans))
		err = launchPlanPrinter.Print(config.GetConfig().MustOutputFormat(), launchPlan, launchPlanStructure, transformLaunchPlan)
		if err != nil {
					return err
				}
		return nil
	}
	launchPlans, err := adminutils.GetAllNamedEntities(ctx, cmdCtx.AdminClient().ListLaunchPlanIds, adminutils.ListRequest{Project: config.GetConfig().Project, Domain: config.GetConfig().Domain})
	if err != nil {
		return err
	}
	logger.Debugf(ctx, "Retrieved %v launch plan", len(launchPlans))
	return launchPlanPrinter.Print(config.GetConfig().MustOutputFormat(), launchPlans, launchPlanStructure, transformLaunchPlan)
	return nil
}
