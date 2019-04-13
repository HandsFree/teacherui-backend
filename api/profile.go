package api

import (
	"bytes"
	"net/http"

	"github.com/handsfree/teacherui-backend/entity"
	"github.com/handsfree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// PutProfile handles the put profile api request
func PutProfile(s *gin.Context) (string, error) {
	var json *entity.CurrentUser
	if err := s.ShouldBindJSON(&json); err != nil {
		util.Error("PutProfile", err.Error())
		return "", err
	}

	profilePut, err := jsoniter.Marshal(json)
	if err != nil {
		util.Error("PutProfile", err.Error())
		return "", err
	}

	resp, err, status := DoTimedRequestBody(s, "PUT",
		API.getPath(s, "currentuser"),
		bytes.NewBuffer(profilePut),
	)
	if err != nil {
		util.Error("PutProfile", err.Error())
		return "", err
	}
	if status != http.StatusOK {
		util.Info("[PutProfile] Status Returned: ", status)
		return "", nil
	}

	return string(resp), nil
}
