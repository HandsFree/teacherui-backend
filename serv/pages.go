package serv

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handsfree/teacherui-backend/api"
	"github.com/handsfree/teacherui-backend/req"
)

// Page represents a web page in the beaconing
// teacher ui. This wraps around a 'SimpleManagedRoute'
// this will inject template information into the
// index.html webpage with the relevant title, script
// to load and the host.
type Page struct {
	Title    string
	Script   string
	Template string
	Host     string
}

const pagesDirectory = "dist/beaconing/pages"

func newPage(title string, script string) *Page {
	host := api.GetLink()
	scriptLoc := fmt.Sprintf("%s/%s", pagesDirectory, script)
	return &Page{
		Title:    title,
		Script:   scriptLoc,
		Template: "index.html",
		Host:     fmt.Sprintf("%s/%s", host, scriptLoc),
	}
}

// wrapper for serving pages.
func getPage(p *Page) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		fmt.Println("---!", p.Host)
		c.HTML(http.StatusOK, p.Template, &gin.H{
			"pageTitle":  p.Title,
			"pageScript": p.Script,
			"host":       p.Host,
		})
	}
	return gin.HandlerFunc(handler)
}

// TODO make the frontend agnostic so
// that we dont need any of this stuff
func registerPages(router *gin.Engine) {
	router.GET("/", getPage(newPage("Home", "home/page.js")))

	classroomGroup := router.Group("/classroom")
	{
		classroomGroup.GET("/", getPage(newPage("Classroom", "classroom/page.js")))
		classroomGroup.GET("/groups", getPage(newPage("Classroom", "classroom/groups/page.js")))
		classroomGroup.GET("/student", getPage(newPage("Classroom", "classroom/student/page.js")))
		classroomGroup.GET("/group", getPage(newPage("Classroom", "classroom/group/page.js")))
	}

	router.GET("/authoring_tool", getPage(newPage("Authoring Tool", "authoring_tool/page.js")))

	lessonManager := router.Group("/lesson_manager")
	{
		lessonManager.GET("/", getPage(newPage("Lesson Manager", "lesson_manager/page.js")))
	}

	router.GET("/calendar", getPage(newPage("Calendar", "calendar/page.js")))
	router.GET("/search", getPage(newPage("Search", "search/page.js")))
	router.GET("/profile", getPage(newPage("Profile", "profile/page.js")))

	authPage := router.Group("auth")
	{
		authPage.GET("logout", req.GetLogOutRequest())
	}

	if gin.IsDebugging() {
		router.GET("/glpviewer", getPage(newPage("GLP Viewer", "glpviewer/page.js")))
	}
}
