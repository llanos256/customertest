package service

import (
	"customermod/models"
	"customermod/database_config"
)

func AllIssues() ([]models.Issues, error) {
	issues := []models.Issues{}
	orm := database_config.GetConnection()
	err := orm.Find(&issues).Error
	return issues, err
}

func CreateIssue(issue models.Issues) (models.Issues, error) {
    orm := database_config.GetConnection()
	err := orm.Save(&issue).Error
	return issue, err
}

func FindIssue(issueId int) (models.Issues, error) {
    orm := database_config.GetConnection()
	issue := models.Issues{}
	err := orm.First(&issue, issueId).Error
	return issue, err
}

func DeleteIssue(issueId int) (string , error) {
    orm := database_config.GetConnection()
    // Find the issue first to ensure it exists
    existingIssue := models.Issues{}
    err := orm.First(&existingIssue, issueId).Error
    if err != nil {
        return "not deleted", err // Issue not found or other error occurred
    }

    // Issue exists, proceed with deletion
    err = orm.Delete(&existingIssue).Error
	if err != nil {
		return "not deleted", err
	}
    return "deleted", err
}