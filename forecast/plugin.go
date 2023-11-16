package forecast

import (
    "context"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-forecast"

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name: pluginName,
        DefaultTransform: transform.FromGo().NullIfZero(),
        TableMap: map[string]*plugin.Table{
            "forecast_projects": tableForecastProjects(ctx),
            "forecast_time_registrations": tableForecastTimeRegistrations(ctx),
            "forecast_project_time_registrations": tableForecastProjectTimeRegistrations(ctx),
            "forecast_person_time_registrations": tableForecastPersonTimeRegistrations(ctx),
            "forecast_persons": tableForecastPersons(ctx),
        },
    }
    return p
}