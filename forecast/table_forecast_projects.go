package forecast

import (
	"context"
    "steampipe-plugin-forecast/sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableForecastProjects(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "forecast_projects",
		Description: "Forecast Projects",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		HydrateConfig: []plugin.HydrateConfig{},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The project ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "company_project_id",
				Description: "The company project ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "name",
				Description: "The project name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "stage",
				Description: "The project stage.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The project status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_description",
				Description: "The project status description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The project description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "priority_level_id",
				Description: "The project priority level ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "color",
				Description: "The project color.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "estimation_units",
				Description: "The project estimation units.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "minutes_per_estimation_point",
				Description: "The project minutes per estimation point.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "budget",
				Description: "The project budget.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "billable",
				Description: "The project billable.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "budget_type",
				Description: "The project budget type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "use_sprints",
				Description: "The project use sprints.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "sprint_length",
				Description: "The project sprint length.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "start_date",
				Description: "The project start date.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "end_date",
				Description: "The project end date.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "use_task_owner",
				Description: "The project use task owner.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "use_task_followers",
				Description: "The project use task followers.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "use_sub_tasks",
				Description: "The project use sub tasks.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "client",
				Description: "The project client.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "rate_card",
				Description: "The project rate card.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "remaining_auto_calculated",
				Description: "The project remaining auto calculated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "use_project_allocations",
				Description: "The project use project allocations.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "use_baseline",
				Description: "The project use baseline.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "baseline_win_chance",
				Description: "The project baseline win chance.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "labels",
				Description: "The project labels.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "external_refs",
				Description: "The project external refs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "created_by",
				Description: "The project created by.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "updated_by",
				Description: "The project updated by.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "The project created at.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The project updated at.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("forecast_projects.listProjects")

	var projects, _ = sdk.ListProjects(ctx)

	for _, project := range projects {
		d.StreamListItem(ctx, project)
	}

    return nil, nil
}

//// HYDRATE FUNCTIONS
