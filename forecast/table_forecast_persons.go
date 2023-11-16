package forecast

import (
	"context"
    "steampipe-plugin-forecast/sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableForecastPersons(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "forecast_persons",
		Description: "Forecast Persons",
		List: &plugin.ListConfig{
			Hydrate: listPersons,
		},
		HydrateConfig: []plugin.HydrateConfig{},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The person ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "first_name",
				Description: "The person first name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_name",
				Description: "The person last name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The person email.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "job_title",
				Description: "The person job title.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "monday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "tuesday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "wednesday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "thursday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "friday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "saturday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "sunday",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "active",
				Description: "Whether the person is archived.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "default_role",
				Description: "The person default role.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "department_id",
				Description: "The person department ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cost",
				Description: "The person cost.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "language",
				Description: "The person language.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by",
				Description: "The person created by.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "updated_by",
				Description: "The person updated by.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "The person created at.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_at",
				Description: "The person updated at.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "client_id",
				Description: "The person client ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "holiday_calendar_id",
				Description: "The person holiday calendar ID.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "start_date",
				Description: "The person start date.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "permissions",
				Description: "The person permissions.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "is_system_user",
				Description: "The person is system user.",
				Type:        proto.ColumnType_BOOL,
			},
		},
	}
}

//// LIST FUNCTION

func listPersons(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("forecast_persons.listPersons")

	var persons, _ = sdk.ListPersons(ctx)

	for _, person := range persons {
		d.StreamListItem(ctx, person)
	}

    return nil, nil
}

//// HYDRATE FUNCTIONS
