package parse

import (
	"github.com/hands-free/teacherui-backend/api"
	"github.com/hands-free/teacherui-backend/entity"
	"github.com/hands-free/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func StudentGroups(s *gin.Context) ([]*entity.StudentGroup, error) {
	studentGroupsData, err := api.GetStudentGroups(s)
	if err != nil {
		util.Error("parse.StudentGroups", err.Error())
		return []*entity.StudentGroup{}, err
	}

	// conv json -> objects
	var studentGroups []*entity.StudentGroup
	if err := jsoniter.Unmarshal([]byte(studentGroupsData), &studentGroups); err != nil {
		util.Error("parse.StudentGroups", err)
		return nil, err
	}

	return studentGroups, nil
}
