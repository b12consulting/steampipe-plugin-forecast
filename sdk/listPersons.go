package sdk

import (
	"context"
	"net/http"
	"os"
	"encoding/json"
)

type Person struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	JobTitle string `json:"job_title"`
	Monday int `json:"monday"`
	Tuesday int `json:"tuesday"`
	Wednesday int `json:"wednesday"`
	Thursday int `json:"thursday"`
	Friday int `json:"friday"`
	Saturday int `json:"saturday"`
	Sunday int `json:"sunday"`
	Active bool `json:"active"`
	DefaultRole int `json:"default_role"`
	DepartmentID int `json:"department_id"`
	Cost int `json:"cost"`
	Language string `json:"language"`
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ClientID int `json:"client_id"`
	HolidayCalendarID int `json:"holiday_calendar_id"`
	StartDate string `json:"start_date"`
	Permissions []string `json:"permissions"`
	IsSystemUser bool `json:"is_system_user"`
}

func ListPersons(ctx context.Context) ([]Person, error) {
    url := "https://api.forecast.it/api/v2/persons"
	var persons []Person

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

	json.NewDecoder(response.Body).Decode(&persons)
	if err != nil {
        panic(err)
	}

    return persons, nil
}