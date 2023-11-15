package sdk

import (
	"context"
	"net/http"
	"os"
	"encoding/json"
)

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


func ListProjects(ctx context.Context) ([]Project, error) {
    url := "https://api.forecast.it/api/v1/projects"
	var projects []Project

    req, err := http.NewRequest("GET", url, nil)
	if err != nil {
        panic(err)
    }

    req.Header.Add("X-FORECAST-API-KEY", os.Getenv("FORECAST_API_KEY"))
    response, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&projects)
	if err != nil {
        panic(err)
	}

    return projects, nil
}