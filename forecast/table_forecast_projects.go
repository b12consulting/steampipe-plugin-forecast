package forecast

import (
	"context"
    "net/http"
	"os"
	"encoding/json"

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
				Name:        "companyProjectId",
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
				Name:        "statusDescription",
				Description: "The project status description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The project description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "priorityLevelId",
				Description: "The project priority level ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "color",
				Description: "The project color.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "estimationUnits",
				Description: "The project estimation units.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "minutesPerEstimationPoint",
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
				Name:        "budgetType",
				Description: "The project budget type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "useSprints",
				Description: "The project use sprints.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "sprintLength",
				Description: "The project sprint length.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "startDate",
				Description: "The project start date.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "endDate",
				Description: "The project end date.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "useTaskOwner",
				Description: "The project use task owner.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "useTaskFollowers",
				Description: "The project use task followers.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "useSubTasks",
				Description: "The project use sub tasks.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "client",
				Description: "The project client.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "rateCard",
				Description: "The project rate card.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "remainingAutoCalculated",
				Description: "The project remaining auto calculated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "useProjectAllocations",
				Description: "The project use project allocations.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "useBaseline",
				Description: "The project use baseline.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "baselineWinChance",
				Description: "The project baseline win chance.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "labels",
				Description: "The project labels.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "externalRefs",
				Description: "The project external refs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "createdBy",
				Description: "The project created by.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "updatedBy",
				Description: "The project updated by.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "createdAt",
				Description: "The project created at.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updatedAt",
				Description: "The project updated at.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("forecast_project.listProjects")

    url := "https://api.forecast.it/api/v1/projects"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
		plugin.Logger(ctx).Error("forecast_project.listProjects", "ERROR", err)
        return nil, err
    }

    req.Header.Add("X-FORECAST-API-KEY", os.Getenv("FORECAST_API_KEY"))
    response, err := http.DefaultClient.Do(req)
    if err != nil {
		plugin.Logger(ctx).Error("forecast_project.listProjects", "ERROR", err)
        return nil, err
    }

	type Project struct {
		ID int `json:"id"`
		CompanyProjectID int `json:"company_project_id"`
		Name string `json:"name"`
		Stage string `json:"stage"`
		Status string `json:"status"`
		StatusDescription string `json:"status_description"`
		Description string `json:"description"`
		PriorityLevelId string `json:"priority_level_id"`
		Color string `json:"color"`
		EstimationUnits string `json:"estimation_units"`
		MinutesPerEstimationPoint int `json:"minutes_per_estimation_point"`
		Budget string `json:"budget"`
		Billable bool `json:"billable"`
		BudgetType string `json:"budget_type"`
		UseSprints bool `json:"use_sprints"`
		SprintLength int `json:"sprint_length"`
		StartDate string `json:"start_date"`
		EndDate string `json:"end_date"`
		UseTaskOwner bool `json:"use_task_owner"`
		UseTaskFollowers bool `json:"use_task_followers"`
		UseSubTasks bool `json:"use_sub_tasks"`
		Client int `json:"client"`
		RateCard int `json:"rate_card"`
		RemainingAutoCalculated bool `json:"remaining_auto_calculated"`
		UseProjectAllocations bool `json:"use_project_allocations"`
		UseBaseline bool `json:"use_baseline"`
		BaselineWinChance int `json:"baseline_win_chance"`
		Labels []int `json:"labels"`
		ExternalRefs []int `json:"external_refs"`
		CreatedBy int `json:"created_by"`
		UpdatedBy int `json:"updated_by"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	var projects []Project
	json.NewDecoder(response.Body).Decode(&projects)
	if err != nil {
		plugin.Logger(ctx).Error("forecast_project.listProjects", "ERROR", err)
		return nil, err
	}

	for _, project := range projects {
		d.StreamListItem(ctx, project)
	}

    return projects, nil
}

//// HYDRATE FUNCTIONS
