package aws

import (
	"context"
    "fmt"
    "io"
    "net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableForecastProjects(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "forecast_projects",
		Description: "Forecast Projects",
		Get: &plugin.GetConfig{},
		List: &plugin.ListConfig{
			Hydrate: listForecastProjects,
		},
		HydrateConfig: []plugin.HydrateConfig{},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The project name.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION


//// HYDRATE FUNCTIONS

func listForecastProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("listForecastProjects")

    url := "https://api.forecast.it/api/v1/projects"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
		plugin.Logger(ctx).Trace("listForecastProjects", "ERROR", err)
        return nil, err
    }

    req.Header.Add("X-FORECAST-API-KEY", os.Getenv("FORECAST_API_KEY"))
    response, err := http.DefaultClient.Do(req)
    if err != nil {
		plugin.Logger(ctx).Trace("listForecastProjects", "ERROR", err)
        return nil, err
    }

    defer response.Body.Close()
    body, readErr := io.ReadAll(response.Body)
    if readErr != nil {
		plugin.Logger(ctx).Trace("listForecastProjects", "ERROR", err)
        return nil, err
    }
    fmt.Println(string(body))

    return nil, nil
}
