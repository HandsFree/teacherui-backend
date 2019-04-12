package req

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/hands-free/teacherui-backend/entity"
	"github.com/hands-free/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func GetActiveLessonPlansWidget() gin.HandlerFunc {
	return func(s *gin.Context) {
		limitParam, err := strconv.Atoi(s.DefaultQuery("limit", "3"))
		if err != nil || limitParam <= 0 {
			limitParam = 3 // NaN
			util.Warn("ALP limit has illegal value, defaulting to 5")
		}

		lps := []entity.LessonPlanWidget{}

		//FIXME

		lpsCount := float64(len(lps))
		size := int(math.Min(float64(limitParam), lpsCount))

		json, err := jsoniter.Marshal(lps[0:size])
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(json))
	}
}

func NewLessonPlanWidget(name string, desc string, glpID uint64) entity.LessonPlanWidget {
	return entity.LessonPlanWidget{
		Name: name,
		Desc: desc,
		Link: fmt.Sprintf("/lesson_manager#view?id=%d", glpID),
	}
}
