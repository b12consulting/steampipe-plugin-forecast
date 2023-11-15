package forecast

import (
	"context"
    "steampipe-plugin-forecast/sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableForecastTimeRegistrations(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "forecast_time_registrations",
		Description: "Forecast Time Registrations",
		List: &plugin.ListConfig{
			Hydrate: listTimeRegistrations,
		},
		HydrateConfig: []plugin.HydrateConfig{},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The time registration ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "person",
				Description: "The person ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "project",
				Description: "The project ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "billable_minutes_registered",
				Description: "The billable minutes registered.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "phase",
				Description: "The phase ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "task",
				Description: "The task ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "task_project",
				Description: "The task project ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "non_project_time",
				Description: "The non project time.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "time_registered",
				Description: "The time registered.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "date",
				Description: "The date.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "notes",
				Description: "The notes.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "approval_status",
				Description: "The approval status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "invoice_entry",
				Description: "The invoice entry ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "invoice",
				Description: "The invoice ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_by",
				Description: "The person ID of the user who created the time registration.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "updated_by",
				Description: "The person ID of the user who updated the time registration.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "The time registration created at.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The time registration updated at.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listTimeRegistrations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("forecast_time_registration.listTimeRegistrations")

	var timeRegistrations, _ = sdk.ListTimeRegistrations(ctx)

	for _, timeRegistration := range timeRegistrations {
		d.StreamListItem(ctx, timeRegistration)
	}

    return nil, nil
}

//// HYDRATE FUNCTIONS
