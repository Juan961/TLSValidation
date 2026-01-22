package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

type ParamValidation struct {
	BaseURL     string
	NewAnalysis string
}

type ValidateResponse struct {
	Finished bool
	Response Response
}

func validationPulling(baseURL string) (ValidateResponse, error) {
	var response Response
	var valRes ValidateResponse

	finished := false

	for !finished {
		time.Sleep(1 * time.Second)

		resp, err := http.Get(baseURL)

		if err != nil {
			return valRes, errors.New("Request failed")
		}

		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&response)

		if response.Status == "READY" {
			valRes.Response = response
			valRes.Finished = true

			finished = true
		} else if response.Status == "ERROR" {
			return valRes, errors.New("Error while pulling from new")
		} else {
			fmt.Println("Pending " + strconv.Itoa(response.Endpoints[0].Progress) + "%")
		}
	}

	return valRes, nil
}

func validateCache(baseURL string) (ValidateResponse, error) {
	var response Response
	var valRes ValidateResponse

	resp, err := http.Get(baseURL)

	if err != nil {
		return valRes, errors.New("Request failed")
	}

	if resp.StatusCode != 200 {
		return valRes, errors.New("Request not succeded")
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&response)

	valRes.Finished = response.Status == "READY"

	if valRes.Finished {
		valRes.Response = response
	}

	return valRes, nil
}

func validateParams(args []string) (ParamValidation, error) {
	res := ParamValidation{}

	args_len := len(args)

	if args_len == 1 {
		return res, errors.New("'host' param not found")
	}

	host := os.Args[1]

	if len(host) > 255 || len(host) < 1 {
		return res, errors.New("'host' param not valid")
	}

	res.BaseURL = "https://api.ssllabs.com/api/v2/analyze?host=" + host

	if args_len == 3 && os.Args[2] == "new" {
		res.NewAnalysis = "&startNew=on"
	} else {
		fmt.Println("No 'new' param found, getting cached value")
	}

	return res, nil
}

func main() {
	fmt.Println("Script started...")

	obj, obj_err := validateParams(os.Args)

	if obj_err != nil {
		log.Fatal(obj_err)
	}

	validation, err := validateCache(obj.BaseURL + obj.NewAnalysis)

	if obj_err != nil {
		log.Fatal(err)
	}

	if !validation.Finished {
		validation, err = validationPulling(obj.BaseURL)
	}

	if obj_err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished, result:")
	fmt.Println(validation.Response)
}
