package forecast

import (
	"context"
    "steampipe-plugin-forecast/sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableForecastPersonTimeRegistrations(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "forecast_person_time_registrations",
		Description: "Forecast Time Registrations by Person",
		List: &plugin.ListConfig{
			Hydrate: listPersonTimeRegistrations,
			KeyColumns: plugin.SingleColumn("person"),
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
				Name:        "card",
				Description: "The card ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "task",
				Description: "The task ID.",
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
				Name:        "created_by",
				Description: "The created by ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "updated_by",
				Description: "The updated by ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "The created at timestamp.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_at",
				Description: "The updated at timestamp.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
		},
	}
}

//// LIST FUNCTION

func listPersonTimeRegistrations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("forecast_time_registrations.listTimeRegistrations")

	var timeRegistrations, _ = sdk.ListTimeRegistrationsByPerson(ctx, d.EqualsQuals["person"].GetInt64Value())

	for _, timeRegistration := range timeRegistrations {
		d.StreamListItem(ctx, timeRegistration)
	}

    return nil, nil
}

//// HYDRATE FUNCTIONS
