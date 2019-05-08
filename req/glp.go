package req

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"

	"github.com/HandsFree/teacherui-backend/api"
	"github.com/HandsFree/teacherui-backend/cfg"
	"github.com/HandsFree/teacherui-backend/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"

	"net/http"
	"strconv"
)

type GLPModel struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"description"`
	Author       string `json:"author"`
	Category     string `json:"category"`
	Content      string `json:"content"`
	GamePlotID   uint64 `json:"gamePlotId"`
	ExternConfig string `json:"externConfig"`
}

func PutGLPRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		glpID, err := strconv.Atoi(s.Param("id"))
		if err != nil {
			util.Error("PutGLPRequest", err.Error())
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		resp, err := api.PutGLP(s, glpID)
		if err != nil {
			util.Error("PutGLPRequest", err.Error())
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(resp))
	}
}

func PostGLPRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		resp, err := api.CreateGLP(s)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(resp))
	}
}

// deletes the given glp
//
// inputs:
// - glp id
func DeleteGLPRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		id, err := strconv.ParseUint(s.Param("id"), 10, 64)
		if err != nil || id < 0 {
			util.Error("error when sanitising glp id", err.Error())
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		body, err := api.DeleteGLP(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(body))
	}
}

type glpFile struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Mime string `json:"mime"`
	Link string `json:"link"`
}

func loadGLPFiles(folderName string) []glpFile {
	base, _ := filepath.Abs(filepath.Join(cfg.Beaconing.Server.FileRootPath, cfg.Beaconing.Server.GlpFilesPath))

	path := filepath.Join(base, folderName)

	util.Error("Loading glp files from ", path)

	fileList := []glpFile{}
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		util.Error("Failed to walk dir ", path, " because ", err.Error())
		return []glpFile{}
	}

	for _, fileInfo := range fileInfo {
		fullPath := filepath.Join(cfg.Beaconing.Server.FileRootPath, cfg.Beaconing.Server.GlpFilesPath, folderName, fileInfo.Name())

		file, err := os.Open(fullPath)
		if err != nil {
			util.Error(err.Error())
			continue
		}
		defer file.Close()

		fstPart := make([]byte, 1024)
		_, err = file.Read(fstPart)
		if err != nil {
			util.Error(err.Error())
			continue
		}
		file.Seek(0, os.SEEK_SET)

		fileType := ""

		absPath := filepath.Join(cfg.Beaconing.Server.GlpFilesPath, folderName, fileInfo.Name())

		kind, unknown := filetype.Match(fstPart)
		if unknown != nil {
			util.Error("Unknown file type for file ", fullPath)
			// mime std. says not to send
			// mime type for unknown files
			fileType = ""
		} else {
			fileType = kind.MIME.Value
		}

		glpFile := glpFile{
			Name: fileInfo.Name(),
			Size: fileInfo.Size(),
			Mime: fileType,
			Link: absPath,
		}

		fileList = append(fileList, glpFile)
	}

	return fileList
}

func GetGLPAssignedStudents() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		plan, err := api.GetAssignedStudents(s, id, true)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		planJSON, err := jsoniter.Marshal(plan)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(planJSON))
	}
}

func GetGLPUnAssignedGroups() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		groups, err := api.GetUnAssignedGroups(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		groupJSON, err := jsoniter.Marshal(groups)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(groupJSON))
	}
}

func GetGLPUnAssignedStudents() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		students, err := api.GetUnAssignedStudents(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		studentsJSON, err := jsoniter.Marshal(students)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(studentsJSON))
	}
}

func DeleteGLPResourceRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		ridParam := s.Param("resource")
		rid, err := strconv.ParseUint(ridParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid Resource ID")
			return
		}

		if err := api.UnlinkResourceFromGLP(s, id, rid); err != nil {
			util.Error(err)
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.JSON(http.StatusOK, map[string]string{
			"message": "OK",
		})
	}
}

func PostGLPResourceRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		ridParam := s.Param("resource")
		rid, err := strconv.ParseUint(ridParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid Resource ID")
			return
		}

		if err := api.LinkResourceToGLP(s, id, rid); err != nil {
			util.Error(err)
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		s.Header("Content-Type", "application/json")
		s.JSON(http.StatusOK, map[string]string{
			"message": "OK",
		})
	}
}

func GetGLPFilesRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		files, err := api.GetGLPFiles(s, id)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		filesJSON, err := jsoniter.Marshal(files)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(filesJSON))
	}
}

// retrieves the given glp
//
// inputs:
// - glp id
// - minify (bool, optional)
//   whether the "contents" of the GLP is omitted or not
func GetGLPRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		idParam := s.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil || id < 0 {
			s.String(http.StatusBadRequest, "Client Error: Invalid GLP ID")
			return
		}

		minify := s.Query("minify")

		// dont minify by default, however if
		// we have a minify parameter with the value
		// 1 then we will minify this glp request.
		// NOTE: that if the parameter fails to parse, etc.
		// then it is completely ignored in the request.
		shouldMinify := false
		if minify != "" {
			minifyParam, errConv := strconv.Atoi(minify)
			if errConv == nil {
				shouldMinify = minifyParam == 1
			} else {
				util.Error("Note: failed to atoi minify param in glp.go", err.Error())
			}
		}

		plan, err := api.GetGLP(s, id, shouldMinify)
		if err != nil {
			s.AbortWithError(http.StatusBadRequest, err)
			return
		}

		planJSON, err := jsoniter.Marshal(plan)
		if err != nil {
			util.Error(err.Error())
			return
		}

		s.Header("Content-Type", "application/json")
		s.String(http.StatusOK, string(planJSON))
	}
}
