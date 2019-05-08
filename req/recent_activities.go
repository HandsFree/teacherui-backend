package req

import (
	"net/http"

	"github.com/HandsFree/teacherui-backend/api"
	"github.com/HandsFree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func GetRecentlyAssigned() gin.HandlerFunc {
	return func(c *gin.Context) {
		assigned, err := api.GetRecentlyAssigned(c)
		if err != nil {
			util.Error(err.Error())
			return
		}

		// convert the plans into json and display
		jsonResult, err := jsoniter.Marshal(&assigned)
		if err != nil {
			util.Error(err.Error())
			return
		}

		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(jsonResult))
	}
}

func GetRecentActivities() gin.HandlerFunc {
	// FIXME error handling!
	return func(s *gin.Context) {
		activities, err := api.GetRecentActivities(s)
		if err != nil {
			util.Error(err.Error())
			return
		}

		// convert the plans into json and display
		jsonResult, err := jsoniter.Marshal(&activities)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(jsonResult))
	}
}
