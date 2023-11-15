package main

import (
    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
    "steampipe-plugin-forecast/forecast"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: forecast.Plugin})
}