package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PeterKWIlliams/my-to-do-go/internal/database"
	"github.com/PeterKWIlliams/my-to-do-go/utils"
)

type ProjectOperation int

const (
	CreateOperation ProjectOperation = iota
	UpdateOperation
)

func (s *Service) CreateProject(options map[string]string, args []string) error {
	project, err := formatCreateProjectData(options, args)
	context := context.Background()
	err = s.Config.DB.CreateProject(
		context,
		database.CreateProjectParams{
			Name:             project.Name,
			Deadline:         project.Deadline,
			TargetDuration:   project.TargetDuration,
			WorkingDirectory: project.WorkingDirectory,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// func (s *Service) formatUpdateProjectData(input map[string]string, args []string, existing *database.Project) (*database.Project, error) {
// 	project := &database.Project{}
// 	if existing == nil {
// 		return nil, fmt.Errorf("project not found")
// 	}
//
// 	if directory, ok := input["directory"]; ok {
// 		projectParams.WorkingDirectory = utils.ToNullString(&directory)
// 	}
//
// 	if timeStr, ok := input["targetDuration"]; ok {
//
// 		targetDuration, err := utils.ParseTimeToSeconds(timeStr)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid time format: %v", err)
// 		}
// 		projectParams.TargetDuration = utils.ToNullInt64(&targetDuration)
//
// 	} else if op == CreateOperation {
// 		projectParams.TargetDuration = sql.NullInt64{Valid: false}
// 	}
//
// 	return projectParams, nil
// }

func formatCreateProjectData(input map[string]string, args []string) (*database.CreateProjectParams, error) {
	projectParams := &database.CreateProjectParams{}

	if len(args) > 0 {
		projectParams.Name = args[0]
	} else {
		return nil, fmt.Errorf("project name required for creation")
	}
	if directory, ok := input["directory"]; ok {
		projectParams.WorkingDirectory = utils.ToNullString(&directory)
	}
	if timeStr, ok := input["targetDuration"]; ok {
		targetDuration, err := utils.ParseTimeToSeconds(timeStr)
		if err != nil {
			return nil, fmt.Errorf("invalid time format: %v", err)
		}
		projectParams.TargetDuration = utils.ToNullInt64(&targetDuration)

	} else if op == CreateOperation {
		projectParams.TargetDuration = sql.NullInt64{Valid: false}
	}

	return projectParams, nil
}
