package req

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/HandsFree/teacherui-backend/entity"

	"github.com/HandsFree/teacherui-backend/api"
	"github.com/HandsFree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func GetResource() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid resource ID")
			return
		}

		res, err := api.GetResource(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		resJSON, err := jsoniter.Marshal(res)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(resJSON))
	}
}

func GetResourceContent() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid resource ID")
			return
		}

		resp, err := api.GetResourceContent(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(resp))
	}
}

func DeleteResource() gin.HandlerFunc {
	return func(s *gin.Context) {
		id, err := strconv.ParseUint(s.Param("id"), 10, 64)
		if err != nil || id < 0 {
			util.Error("error when sanitising resource id", err.Error())
			s.String(http.StatusBadRequest, "Client Error: Invalid resource ID")
			return
		}

		body, err := api.DeleteResource(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(body))
	}
}

func CreateResourceHandle() gin.HandlerFunc {
	return func(s *gin.Context) {
		var resource entity.Resource
		if err := s.ShouldBindJSON(&resource); err != nil {
			util.Error("CreateResourceHandle", err.Error())
			s.AbortWithStatus(http.StatusBadRequest)
			return
		}

		body, err := api.PostResource(s, resource)
		if err != nil {
			util.Error(err)
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(body))
	}
}

func UpdateResource() gin.HandlerFunc {
	return func(s *gin.Context) {
		id, err := strconv.ParseUint(s.Param("id"), 10, 64)
		if err != nil {
			log.Println("UpdateResource failed to bind id", err)
			s.AbortWithStatus(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(s.Request.Body)
		if err != nil {
			log.Println("UpdateResource reading body failed", err)
			s.AbortWithStatus(http.StatusBadRequest)
			return
		}

		body, err := api.SetResourceContent(s, id, data)
		if err != nil {
			util.Error(err)
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(body))
	}
}
