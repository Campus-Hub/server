package api

import (
	"github.com/Campus-Hub/server/internal/service"
	"github.com/Campus-Hub/server/pkg/errno"
	"github.com/Campus-Hub/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func UploadResource(c *gin.Context) {
	var (
		uploadResourceService service.ResourceService
	)

	if err := c.ShouldBind(&uploadResourceService); err != nil {
		res := uploadResourceService.Upload
		c.JSON(errno.Success, res)
	} else {
		c.JSON(errno.Error, errorResponse(err))
		logger.Logger.Error(err)
	}
}
