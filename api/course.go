package api

import (
	"github.com/Campus-Hub/server/internal/service"
	"github.com/Campus-Hub/server/pkg/errno"
	"github.com/Campus-Hub/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Summary Show a course
// @Description Get information about a specific course.
// @ID showCourse
// @Accept  json
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} pack.Course
// @Failure 401 {object} service.Response
// @Router /courses/{id} [get]
func ShowCourse(c *gin.Context) {
	var (
		svr service.CourseSrv
		err error
	)

	// Parameters Binding
	if err = c.ShouldBind(&svr); err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}

	// Call Show Course Service Return Response.
	res, err := svr.Show()
	if err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}
	c.JSON(errno.Success, res)
}

func ListCourse(c *gin.Context) {
	var (
		svr service.CourseSrv
		err error
	)

	// Parameters Binding
	if err = c.ShouldBind(&svr); err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}

	// Call Show Course Service Return Response.
	res, err := svr.List()
	if err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}
	c.JSON(errno.Success, res)
}

func CreateCourse(c *gin.Context) {
	var (
		createCourseSvr service.CourseSrv
		err             error
	)

	if err = c.ShouldBind(&createCourseSvr); err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}

	res, err := createCourseSvr.Create()
	if err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}
	c.JSON(errno.Success, res)
}

func UpdateCourse(c *gin.Context) {
	var (
		svr service.CourseSrv
		err error
	)

	// Parameters Binding.
	if err = c.ShouldBind(&svr); err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}

	res, err := svr.Update()
	if err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}
	c.JSON(errno.Success, res)
}

func RemoveCourse(c *gin.Context) {
	return
}
