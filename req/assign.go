package req

import (
	"encoding/gob"
	"net/http"
	"strconv"
	"time"

	"github.com/HandsFree/teacherui-backend/api"
	"github.com/HandsFree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(map[int]bool{})
}

// Assign's a student to the given GLP
//
// inputs:
// - student id
// - glp id
func GetAssignRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		studentID, err := strconv.ParseUint(s.Param("student"), 10, 64)
		if err != nil || studentID < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid student ID")
			return
		}

		glpID, err := strconv.ParseUint(s.Param("glp"), 10, 64)
		if err != nil || glpID < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		// FIXME clean this up.

		fromParam := s.Param("from")

		var from time.Time
		if fromParam != "" {
			fromTime, err := time.Parse(time.RFC3339, fromParam)
			if err != nil {
				util.Error("assign from time is bad", err.Error())
				fromTime = time.Now()

				// we aren't going to error here since we can
				// just say it's been assigned from the current
				// time.
			}
			from = fromTime
		} else {
			from = time.Now()
		}

		toParam := s.Param("to")

		var to time.Time
		if toParam != "" {
			toTime, err := time.Parse(time.RFC3339, toParam)
			if err != nil {
				util.Error("assign to time is bad", err.Error())
				s.AbortWithError(http.StatusBadRequest, err)
				return
			}
			to = toTime
		}

		// do the post request to the beaconing API
		// saying we're assigning said student to glp.
		resp, err := api.AssignStudentToGLP(s, studentID, glpID, from, to)
		if err != nil {
			s.String(http.StatusBadRequest, "Failed to assign student to glp")
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, resp)
	}
}

func GetGroupAssignRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		groupID, err := strconv.ParseUint(s.Param("group"), 10, 64)
		if err != nil || groupID < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid group ID")
			return
		}

		glpID, err := strconv.ParseUint(s.Param("glp"), 10, 64)
		if err != nil || glpID < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		fromParam := s.Param("from")
		from, err := time.Parse(time.RFC3339, fromParam)
		if err != nil {
			util.Error("assign from time is bad", err.Error())
			from = time.Now()

			// we aren't going to error here since we can
			// just say it's been assigned from the current
			// time.
		}

		toParam := s.Param("to")
		var to time.Time
		if toParam != "" {
			var err error
			to, err = time.Parse(time.RFC3339, toParam)
			if err != nil {
				util.Error("assign to time is bad", err.Error())
				s.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}

		resp, err := api.AssignGroupToGLP(s, groupID, glpID, from, to)
		if err != nil {
			s.String(http.StatusBadRequest, "Failed to assign group to glp")
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, resp)
	}
}
