package serv

import (
	"github.com/HandsFree/teacherui-backend/req"
	"github.com/gin-gonic/gin"
)

func registerAPI(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	{
		// FIXME move this under a category. do we still need this?
		v1.GET("active_lesson_plans", req.GetActiveLessonPlansWidget())
	}

	lang := v1.Group("lang")
	{
		lang.GET("/:code/phrase/:key", req.GetTranslation())
		lang.POST("/phrases", req.GetTranslationPhrases())
	}

	// FIXME: move somewhere, e.g. /students/
	v1.GET("student_overview", req.GetStudentOverview())

	authAPI := v1.Group("auth")
	{
		authAPI.GET("gettoken", req.GetCheckAuthRequest())
	}

	tokens := v1.Group("token")
	{
		tokens.GET("/", req.GetTokenRequest())
	}

	assign := v1.Group("assign")
	{
		assign.GET("/:student/to/:glp", req.GetAssignRequest())
	}

	assignGroup := v1.Group("assigngroup")
	{
		assignGroup.GET("/:group/to/:glp", req.GetGroupAssignRequest())
	}

	student := v1.Group("student")
	{
		student.GET("/:id", req.GetStudentRequest())
		student.PUT("/:id", req.PutStudentRequest())
		student.DELETE("/:id", req.DeleteStudentRequest())
		student.POST("/", req.PostStudentRequest())

		// kind of like assignedglps, but a HARD search i.e.
		// it will go through each GLP and get the GLPs. this is predominently
		// for the calendar since this is what happens on the client side, but i want
		// to squish it into one request to make it a bit faster.
		student.GET("/:id/assignedglps_hard", req.GetAssignedGLPsHardRequest())

		student.GET("/:id/assignedglps", req.GetAssignedGLPsRequest())
		student.DELETE("/:id/assignedglps/:glp", req.DeleteAssignedGLPsRequest())
	}

	students := v1.Group("students")
	{
		students.GET("/", req.GetStudentsRequest())

		// FIXME(Felix): _hard request here. this is mirrored?
		students.GET("/:id/assignedglps", req.GetAssignedGLPsRequest())

		students.DELETE("/:id/assignedglps/:glp", req.DeleteAssignedGLPsRequest())
	}

	profile := v1.Group("profile")
	{
		profile.GET("/", req.GetProfileRequest())
		profile.PUT("/", req.PutProfileRequest())
		profile.GET("recent_activities", req.GetRecentActivities())
		profile.GET("recently_assigned", req.GetRecentlyAssigned())
	}

	glps := v1.Group("glps")
	{
		glps.GET("/", req.GetGLPSRequest())
		glps.GET("/exists_with_name", req.GetGLPSWithName())
	}

	res := v1.Group("resource")
	{
		res.GET("/:id", req.GetResource())
		res.GET("/:id/content", req.GetResourceContent())
		res.DELETE("/:id", req.DeleteResource())

		// when creating a resource, we create a 'handle'
		// name, etc. the content is then PUT into the resource.
		res.POST("/", req.CreateResourceHandle())

		// this is whree we set the content of the resource
		res.PUT("/:id/content", req.UpdateResource())
	}

	glp := v1.Group("glp")
	{
		glp.GET("/:id", req.GetGLPRequest())
		glp.GET("/:id/files/", req.GetGLPFilesRequest())

		glp.POST("/:id/link_resource/:resource", req.PostGLPResourceRequest())
		glp.DELETE("/:id/unlink_resource/:resource", req.DeleteGLPResourceRequest())

		// returns all of the students assigned to this glp
		glp.GET("/:id/assigned_students", req.GetGLPAssignedStudents())

		// returns all of the students that are not, or _can be_
		// assigned to this glp.
		glp.GET("/:id/unassigned_students", req.GetGLPUnAssignedStudents())
		glp.GET("/:id/unassigned_groups", req.GetGLPUnAssignedGroups())

		glp.DELETE("/:id", req.DeleteGLPRequest())
		glp.POST("/", req.PostGLPRequest())
		glp.PUT("/:id", req.PutGLPRequest())
	}

	studentGroups := v1.Group("studentgroups")
	{
		studentGroups.GET("/", req.GetStudentGroupsRequest())
	}

	studentsFromStudentGroup := v1.Group("students_from_studentgroup")
	{
		studentsFromStudentGroup.GET("/:id", req.GetStudentsFromStudentGroupRequest())
	}

	studentGroup := v1.Group("studentgroup")
	{
		studentGroup.GET("/:id", req.GetStudentGroupRequest())
		studentGroup.PUT("/:id", req.PutStudentGroupRequest())

		studentGroup.GET("/:id/assignedglps_hard", req.GetStudentGroupAssignedHardRequest())
		studentGroup.GET("/:id/assignedglps", req.GetStudentGroupAssignedRequest())

		studentGroup.DELETE("/:id/assignedglps/:glp", req.DeleteGroupAssignedRequest())
		studentGroup.POST("/", req.PostStudentGroupRequest())
		studentGroup.DELETE("/:id", req.DeleteStudentGroupRequest())
	}

	gameplots := v1.Group("gameplots")
	{
		gameplots.GET("/", req.GetGameplots())
	}

	v1.POST("search", req.PostSearchRequest())
}
