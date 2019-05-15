package tracker

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/watercompany/cx-tracker/src/api"
)

// Controller handles tracker requests
type Controller struct {
	service Service
}

// DefaultController creates new instance of controller
func DefaultController() Controller {
	return Controller{
		service: DefaultService(),
	}
}

// RegisterAPIs registration of controller routes
func (ctrl Controller) RegisterAPIs(public *gin.RouterGroup, closed *gin.RouterGroup) {
	public.PUT("/config", ctrl.updateConfig)
	public.GET("/config/:hash", ctrl.getConfig)
}

// @Summary Returns uptime info for previous month
// @Description Returns uptime info for nodes from the request
// @Tags config
// @Produce json
// @Success 201 string
// @Failure 500 {object} api.ErrorResponse
// @Router /config [put]
func (ctrl Controller) updateConfig(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	hash, err := ctrl.service.saveConfig(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, hash)
}

// @Summary Returns config file content
// @Description Returns config file content from memory
// @Tags config
// @Produce json
// @Param Hash query string true "Config hash"
// @Success 200 {array} string
// @Failure 404 {object} api.ErrorResponse
// @Router /config [get]
func (ctrl Controller) getConfig(c *gin.Context) {
	hash := c.Param("hash")
	response := ctrl.service.readConfig(hash)
	if len(response) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, api.ErrorResponse{Error: "No data matching hash: " + hash})
		return
	}

	if _, err := c.Writer.Write(response); err != nil {
		log.Error("Error writing response: ", err)
	}
}
