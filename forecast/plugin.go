package forecast

import (
    "context"

    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name:             "steampipe-plugin-forecast",
        DefaultTransform: transform.FromGo().NullIfZero(),
        TableMap: map[string]*plugin.Table{
            "forecast_projects": tableForecastProjects(),
        },
    }
    return p
}