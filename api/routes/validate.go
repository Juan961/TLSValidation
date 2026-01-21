package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type Endpoint struct {
	IpAddress         string `json:"ipAddress"`         // 69.67.183.100
	StatusMessage     string `json:"statusMessage"`     // Ready
	Grade             string `json:"grade"`             // A+
	GradeTrustIgnored string `json:"gradeTrustIgnored"` // A+
	HasWarnings       bool   `json:"hasWarnings"`       // false
	IsExceptional     bool   `json:"isExceptional"`     // true
	Progress          int    `json:"progress"`          // 100
	Duration          int    `json:"duration"`          // 47359
	Eta               int    `json:"eta"`               // 3684
	Delegation        int    `json:"delegation"`        // 2
}

type Response struct {
	Host            string     `json:"host"`            // www.ssllabs.com
	Port            int        `json:"port"`            // 443
	Protocol        string     `json:"protocol"`        // http
	IsPublic        bool       `json:"isPublic"`        // false
	Status          string     `json:"status"`          // READY
	StartTime       int        `json:"startTime"`       // 1768886648445
	TestTime        int        `json:"testTime"`        // 1768886695867
	EngineVersion   string     `json:"engineVersion"`   // 2.4.1
	CriteriaVersion string     `json:"criteriaVersion"` // 2009q
	Endpoints       []Endpoint `json:"endpoints"`
}

const baseURL string = "https://api.ssllabs.com/api/v2/analyze?host=%s"

func ValidateRoute(w http.ResponseWriter, r *http.Request) {
	user := map[string]bool{"error": true}

	host := r.URL.Query().Get("host")

	if host == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, user)
		return
	}

	url := fmt.Sprintf(baseURL, host)

	resp, err := http.Get(url)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, user)
		return
	}

	if resp.StatusCode != 200 {
		render.Status(r, http.StatusPaymentRequired)
		render.JSON(w, r, user)
		return
	}

	defer resp.Body.Close()

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Println("failed: %w", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, user)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
