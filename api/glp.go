package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HandsFree/teacherui-backend/entity"
	"github.com/HandsFree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func containsGLP(glpID uint64, glpArr []*entity.GLP) bool {
	for _, glp := range glpArr {
		if glp.ID == glpID {
			return true
		}
	}

	return false
}

// To be used later on
// type glpPutJSON struct {
// 	ID                 uint64    `json:"id"`
// 	Name               string    `json:"name"`
// 	Desc               string    `json:"description"`
// 	Author             string    `json:"author"`
// 	Category           string    `json:"category"`
// 	Domain             string    `json:"domain"`
// 	Topic              string    `json:"topic"`
// 	AgeGroup           string    `json:"ageGroup"`
// 	Year               int       `json:"year"`
// 	LearningObjectives []string  `json:"learningObjectives"`
// 	Competences        []string  `json:"competences"`
// 	Content            string    `json:"content"`
// 	Public             bool      `json:"public"`
// 	GamePlotID         int       `json:"gamePlotId"`
// 	ExternConfig       string    `json:"externConfig"`
// 	CreatedAt          time.Time `json:"createdAt"`
// 	UpdatedAt          time.Time `json:"updatedAt"`
// 	Owner              string    `json:"owner"`
// 	OwnedByMe          bool      `json:"ownedByMe"`
// 	ReadOnly           bool      `json:"readOnly"`
// }

type glpPostJSON struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Author             string   `json:"author"`
	Category           string   `json:"category"`
	Domain             string   `json:"domain"`
	Topic              string   `json:"topic"`
	AgeGroup           string   `json:"ageGroup"`
	Year               int      `json:"year"`
	LearningObjectives []string `json:"learningObjectives"`
	Competences        []string `json:"competences"`
	Public             bool     `json:"public"`
	GamePlotID         int      `json:"gamePlotId"`
	ExternConfig       string   `json:"externConfig"`
}

// PutGLP ...
func PutGLP(s *gin.Context, id int) (string, error) {
	var json glpPostJSON
	if err := s.ShouldBindJSON(&json); err != nil {
		util.Error("PutGLP", err.Error())
		return "", err
	}

	glpPut, err := jsoniter.Marshal(json)
	if err != nil {
		util.Error("PutGLP", err.Error())
		return "", err
	}

	apiPath := API.getPath(s, "gamifiedlessonpaths/", fmt.Sprintf("%d", id))
	resp, err, status := DoTimedRequestBody(s, "PUT",
		apiPath,
		bytes.NewBuffer(glpPut),
	)
	if err != nil {
		util.Error("PutGLP", err.Error())
		return "", err
	}

	if status != http.StatusOK {
		util.Info("[PutGLP] Status Returned: ", status)
		return "", nil
	}

	return string(resp), nil
}

// CreateGLP handles the CreateGLP POST request.
func CreateGLP(s *gin.Context) (string, error) {
	var json glpPostJSON
	if err := s.ShouldBindJSON(&json); err != nil {
		util.Error("CreateGLP", err.Error())
		return "", err
	}

	glpPost, err := jsoniter.Marshal(json)
	if err != nil {
		util.Error("CreateGLP", err.Error())
		return "", err
	}

	resp, err, status := DoTimedRequestBody(s, "POST",
		API.getPath(s, "gamifiedlessonpaths"),
		bytes.NewBuffer(glpPost),
	)
	if err != nil {
		util.Error("CreateGLP", err.Error())
		return "", err
	}

	if status != http.StatusCreated {
		util.Info("[CreateGLP] Status Returned: ", status)
		return "", nil
	}

	return string(resp), nil
}

// GetGLPS requests all of the GLPs from the core
// API returned as a json string
func GetGLPS(s *gin.Context, minify bool) (string, error) {
	apiPath := API.getPath(s, "gamifiedlessonpaths/", fmt.Sprintf("?noContent=%s", strconv.FormatBool(minify)))

	resp, err, status := DoTimedRequest(s, "GET", apiPath)
	if err != nil {
		util.Error("GetGLPS", err.Error())
		return "", err
	}
	if status != http.StatusOK {
		util.Info("[GetGLPS] Status Returned: ", status)
		return "", err
	}

	return string(resp), nil
}

func GetAssignedGroups(s *gin.Context, id uint64) ([]*entity.Assignee, error) {
	apiPath := API.getPath(s, "gamifiedlessonpaths/",
		fmt.Sprintf("%d", id),
		fmt.Sprintf("/assignedstudentgroups?includeAnalytics=false"))

	resp, err, status := DoTimedRequest(s, "GET", apiPath)

	if err != nil {
		util.Error("GetAssignedGroups", err.Error())
		return nil, err
	}

	if status != http.StatusOK {
		util.Info("[GetAssignedGroups] Status Returned: ", status)
		return nil, err
	}

	// the entity.Assignee is the same as the
	// core api crams in both student group id and
	// student id into the same ntity.
	data := []*entity.Assignee{}
	if err := jsoniter.Unmarshal(resp, &data); err != nil {
		util.Error("GetAssignedGroups", err.Error())
		return nil, err
	}

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, resp); err != nil {
		util.Error("GetGLP", err.Error())
		return nil, err
	}

	return data, nil
}

func GetAssignedStudents(s *gin.Context, id uint64, includeGroups bool) ([]*entity.Assignee, error) {
	apiPath := API.getPath(s, "gamifiedlessonpaths/",
		fmt.Sprintf("%d", id),
		fmt.Sprintf("/assignedstudents?includeGroups=%s&includeAnalytics=false", strconv.FormatBool(includeGroups)))

	resp, err, status := DoTimedRequest(s, "GET", apiPath)

	if err != nil {
		util.Error("GetAssignedStudents", err.Error())
		return nil, err
	}

	if status != http.StatusOK {
		util.Info("[GetAssignedStudents] Status Returned: ", status)
		return nil, err
	}

	data := []*entity.Assignee{}
	if err := jsoniter.Unmarshal(resp, &data); err != nil {
		util.Error("GetAssignedStudents", err.Error())
		return nil, err
	}

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, resp); err != nil {
		util.Error("GetGLP", err.Error())
		return nil, err
	}

	return data, nil
}

func GetUnAssignedGroups(s *gin.Context, id uint64) ([]*entity.StudentGroup, error) {
	assigned, err := GetAssignedGroups(s, id)
	if err != nil {
		util.Error("GetUnAssignedGroups", err.Error())
		return nil, err
	}

	fmt.Println("The assigned students are ", assigned)

	assignedMap := map[uint64]bool{}
	for _, assignedGroup := range assigned {
		assignedMap[assignedGroup.StudentGroupID] = true
	}

	groupsJSON, err := GetStudentGroups(s)
	if err != nil {
		util.Error("GetUnAssignedGroups", err.Error())
		return nil, err
	}

	// conv json -> objects
	var groups []*entity.StudentGroup
	if err := jsoniter.Unmarshal([]byte(groupsJSON), &groups); err != nil {
		util.Error("GetUnAssignedGroups", err)
		return nil, err
	}

	unassigned := []*entity.StudentGroup{}
	for _, group := range groups {
		// the student is not assigned to this GLP,
		// add it to the list.
		if _, ok := assignedMap[group.ID]; !ok {
			unassigned = append(unassigned, group)
		}
	}

	return unassigned, nil
}

// Note this returns _student_ objects only. It does not show
// unassigned student groups.
func GetUnAssignedStudents(s *gin.Context, id uint64) ([]*entity.Student, error) {
	assigned, err := GetAssignedStudents(s, id, true)
	if err != nil {
		util.Error("GetUnAssignedStudents", err.Error())
		return nil, err
	}

	fmt.Println("The assigned students are ", assigned)

	assignedMap := map[uint64]bool{}
	for _, assignedStudent := range assigned {
		assignedMap[assignedStudent.StudentID] = true
	}

	studentsJSON, err := GetStudents(s)
	if err != nil {
		util.Error("GetUnAssignedStudents", err.Error())
		return nil, err
	}

	// conv json -> objects
	var students []*entity.Student
	if err := jsoniter.Unmarshal([]byte(studentsJSON), &students); err != nil {
		util.Error("GetUnAssignedStudents", err)
		return nil, err
	}

	unassigned := []*entity.Student{}
	for _, student := range students {
		// the student is not assigned to this GLP,
		// add it to the list.
		if _, ok := assignedMap[student.ID]; !ok {
			unassigned = append(unassigned, student)
		}
	}

	return unassigned, nil
}

func GetGLPFiles(s *gin.Context, id uint64) ([]*entity.ResourceLinkResponse, error) {
	apiPath := API.getPath(s, "gamifiedlessonpaths/", fmt.Sprintf("%d", id), "/resources")
	resp, err, status := DoTimedRequest(s, "GET", apiPath)

	if err != nil {
		util.Error("GetGLPFiles", err.Error())
		return nil, err
	}

	if status != http.StatusOK {
		util.Info("[GetGLPFiles] Status Returned: ", status)
		return nil, err
	}

	data := []*entity.ResourceLinkResponse{}
	if err := jsoniter.Unmarshal(resp, &data); err != nil {
		util.Error("GetGLPFiles", err.Error())
		return nil, err
	}

	return data, nil
}

// GetGLP requests the GLP with the given id, this function returns
// the string of json retrieved _as well as_ the parsed json object
// see entity.GLP
func GetGLP(s *gin.Context, id uint64, minify bool) (*entity.GLP, error) {
	apiPath := API.getPath(s, "gamifiedlessonpaths/",
		fmt.Sprintf("%d", id),
		fmt.Sprintf("?noContent=%s", strconv.FormatBool(minify)))

	resp, err, status := DoTimedRequest(s, "GET", apiPath)

	if err != nil {
		util.Error("GetGLP", err.Error())
		return nil, err
	}

	if status != http.StatusOK {
		util.Info("[GetGLP] Status Returned: ", status)
		return nil, err
	}

	data := &entity.GLP{}
	if err := jsoniter.Unmarshal(resp, data); err != nil {
		util.Error("GetGLP", err.Error())
		return nil, err
	}

	// should we compact everything?
	// we do here because the json for glps request is stupidly long
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, resp); err != nil {
		util.Error("GetGLP", err.Error())
		return nil, err
	}

	return data, nil
}

// DeleteGLP deletes the given GLP of {id} from the
// core database.
func DeleteGLP(s *gin.Context, id uint64) (string, error) {
	resp, err, status := DoTimedRequest(s, "DELETE",
		API.getPath(s, "gamifiedlessonpaths/", fmt.Sprintf("%d", id)),
	)
	if err != nil {
		util.Error(err)
		return "", err
	}

	if status != http.StatusOK {
		util.Info("[DeleteGLP] Status Returned: ", status)
		return "", nil
	}

	return string(resp), nil
}
