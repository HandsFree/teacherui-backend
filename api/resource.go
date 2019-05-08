package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/HandsFree/teacherui-backend/entity"
	"github.com/HandsFree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func GetResourceContent(s *gin.Context, id uint64) (string, error) {
	apiPath := API.getPath(s, "resources/", fmt.Sprintf("%d", id), "/content")

	resp, err, status := DoTimedRequest(s, "GET", apiPath)

	if err != nil {
		util.Error("GetResourceContent", err.Error())
		return "", err
	}

	if status != http.StatusOK {
		util.Info("[GetResourceContent] Status Returned: ", status)
		return "", err
	}

	// response from the core api
	// is crammed into this json request
	// perhaps store the filetype here too?
	type response struct {
		Data string `json:"data"`
	}
	data := &response{string(resp)}

	respData, err := jsoniter.Marshal(&data)
	if err != nil {
		util.Info("[GetResourceContent]", err)
		return "", nil
	}

	return string(respData), nil
}

func GetResource(s *gin.Context, id uint64) (*entity.Resource, error) {
	apiPath := API.getPath(s, "resources/", fmt.Sprintf("%d", id))

	resp, err, status := DoTimedRequest(s, "GET", apiPath)

	if err != nil {
		util.Error("GetResource", err.Error())
		return nil, err
	}

	if status != http.StatusOK {
		util.Info("[GetResource] Status Returned: ", status)
		return nil, err
	}

	data := &entity.Resource{}
	if err := jsoniter.Unmarshal(resp, data); err != nil {
		util.Error("GetResource", err.Error())
		return nil, err
	}

	// should we compact everything?
	// we do here because the json for glps request is stupidly long
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, resp); err != nil {
		util.Error("GetResource", err.Error())
		return nil, err
	}

	return data, nil
}

func DeleteResource(s *gin.Context, id uint64) (string, error) {
	resp, err, status := DoTimedRequest(s, "DELETE", API.getPath(s, "resources/", fmt.Sprintf("%d", id)))
	if err != nil {
		util.Error(err)
		return "", err
	}

	if status != http.StatusOK {
		util.Info("[DeleteResource] Status Returned: ", status)
		return "", nil
	}

	return string(resp), nil
}

func UnlinkResourceFromGLP(s *gin.Context, glp uint64, to uint64) error {
	link := entity.ResourceLink{
		GlpID:      fmt.Sprintf("%d", glp),
		ResourceID: fmt.Sprintf("%d", to),
	}

	json, err := jsoniter.Marshal(link)
	if err != nil {
		return err
	}

	apiPath := API.getPath(s, "gamifiedlessonpaths/", fmt.Sprintf("%d", glp), "/resources/", fmt.Sprintf("%d", to))
	resp, err, status := DoTimedRequestBody(s, "DELETE", apiPath, bytes.NewBuffer(json))
	if err != nil {
		util.Error(err)
		return err
	}

	if status != http.StatusOK {
		type response struct {
			Err struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			} `json:"error"`
		}

		var jsonResp response
		jsoniter.Unmarshal(resp, &jsonResp)

		util.Info("[UnlinkResourceFromGLP] Status Returned: ", status, jsonResp.Err.Message)
		return errors.New(jsonResp.Err.Message)
	}

	// TODO(Felix): Parse the response and see if we error.
	// handle accordingly.

	return nil
}

func LinkResourceToGLP(s *gin.Context, glp uint64, to uint64) error {
	link := entity.ResourceLink{
		GlpID:      fmt.Sprintf("%d", glp),
		ResourceID: fmt.Sprintf("%d", to),
	}

	json, err := jsoniter.Marshal(link)
	if err != nil {
		return err
	}

	apiPath := API.getPath(s, "gamifiedlessonpaths/", fmt.Sprintf("%d", glp), "/resources")
	resp, err, status := DoTimedRequestBody(s, "POST", apiPath, bytes.NewBuffer(json))
	if err != nil {
		util.Error(err)
		return err
	}

	if status != http.StatusCreated {
		type response struct {
			Err struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			} `json:"error"`
		}

		var jsonResp response
		jsoniter.Unmarshal(resp, &jsonResp)

		util.Info("[LinkResourceToGLP] Status Returned: ", status, jsonResp.Err.Message)
		return errors.New(jsonResp.Err.Message)
	}

	// TODO(Felix): Parse the response and see if we error.
	// handle accordingly.

	return nil
}

func PostResource(s *gin.Context, res entity.Resource) (string, error) {
	json, err := jsoniter.Marshal(res)
	if err != nil {
		return "", err
	}

	resp, err, status := DoTimedRequestBody(s, "POST", API.getPath(s, "resources/"), bytes.NewBuffer(json))
	if err != nil {
		util.Error(err)
		return "", err
	}

	if status != http.StatusCreated {
		util.Info("[PostResource] Status Returned: ", status, string(resp))
		return "", errors.New("Failed to create resource")
	}

	return string(resp), nil
}

func SetResourceContent(s *gin.Context, id uint64, data []byte) (string, error) {
	apiPath := API.getPath(s, "resources/", fmt.Sprintf("%d", id), "/content")

	buff := bytes.NewBuffer(data)

	resp, err, status := DoTimedRequestAcceptBody(s, "PUT", "application/octet-stream", apiPath, buff)
	if err != nil {
		util.Error("[SetResourceContent]", err)
		return "", err
	}

	if status != http.StatusOK {
		util.Info("[SetResourceContent] Status Returned: ", status, string(resp))
		return "", nil
	}

	fmt.Println("SetResourceContent resp:", string(resp))

	return string(resp), nil
}
