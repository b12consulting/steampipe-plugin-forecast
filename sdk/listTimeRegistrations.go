package sdk

import (
	"context"
	"net/http"
	"os"
	"encoding/json"
	"fmt"
)

type TimeRegistration struct {
	ID int `json:"id"`
	Person int `json:"person"`
	Project int `json:"project"`
	BillableMinutesRegistered int `json:"billable_minutes_registered"`
	Phase int `json:"phase"`
	Task int `json:"task"`
	TaskProject int `json:"task_project"`
	NonProjectTime int `json:"non_project_time"`
	TimeRegistered int `json:"time_registered"`
	Date string `json:"date"`
	Notes string `json:"notes"`
	ApprovalStatus string `json:"approval_status"`
	InvoiceEntry int `json:"invoice_entry"`
	Invoice int `json:"invoice"`
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ListTimeRegistrations(ctx context.Context) ([]TimeRegistration, error) {
	var timeRegistrations []TimeRegistration
	pageNumber := 1

	for {
		url := fmt.Sprintf("https://api.forecast.it/api/v4/time_registrations?pageNumber=%d&pageSize=500", pageNumber)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("X-FORECAST-API-KEY", os.Getenv("FORECAST_API_KEY"))
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		var body struct {
			PageContents []TimeRegistration `json:"pageContents"`
			PageNumber int `json:"pageNumber"`
			PageSize int `json:"pageSize"`
			TotalObjectCount int `json:"totalObjectCount"`
		}
		err = json.NewDecoder(response.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		timeRegistrations = append(timeRegistrations, body.PageContents...)

		if body.PageNumber == body.TotalObjectCount / body.PageSize + 1 {
			break
		}

		pageNumber++
	}

	return timeRegistrations, nil
}