package api

import (
	"github.com/Campus-Hub/server/internal/service"
	"github.com/Campus-Hub/server/pkg/errno"
	"github.com/Campus-Hub/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ListCourse(c *gin.Context) {
	var (
		listCourseSvr service.CourseSrv
		err           error
	)

	// Parameters Binding
	if err = c.ShouldBind(&listCourseSvr); err != nil {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}

	res := listCourseSvr.List()
	c.JSON(errno.Success, res)
}
