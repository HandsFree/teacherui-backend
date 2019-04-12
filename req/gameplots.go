package req

import (
	"net/http"

	"github.com/hands-free/teacherui-backend/api"
	"github.com/hands-free/teacherui-backend/util"
	"github.com/gin-gonic/gin"
)

func GetGameplots() gin.HandlerFunc {
	return func(s *gin.Context) {
		resp, err := api.GetGameplots(s)
		if err != nil {
			util.Error("GetGameplots", err.Error())
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, resp)
	}
}
