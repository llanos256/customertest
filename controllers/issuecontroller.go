package controllers

import (
	"net/http"
	"customermod/service"
	"customermod/models"
	"encoding/json"
	"strconv"
)

func GetIssues(rw http.ResponseWriter, r *http.Request) {
	issues , err := service.AllIssues()
	BuildResponse(rw, issues, err)
}

func CreateIssue(rw http.ResponseWriter, r *http.Request) {
	issue := models.Issues{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&issue)
	saved , err := service.CreateIssue(issue)
	BuildResponse(rw, saved, err)
}

func FindIssue(rw http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	issueId, _ := strconv.Atoi(queryValues.Get("issueId"))
	found, err := service.FindIssue(issueId)
	BuildResponse(rw, found, err)
}

func DeleteIssue(rw http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	issueId, _ := strconv.Atoi(queryValues.Get("issueId"))
	message, err := service.DeleteIssue(issueId)
	BuildResponse(rw, message, err)
}

