package tracker

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/SkycoinPro/cx-services-cx-tracker/src/api"
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
	public.PUT("/config", ctrl.saveConfig)
	public.GET("/configs", ctrl.getConfigs)
	public.GET("/config/:genesisHash", ctrl.getConfig)
	public.GET("/config/:genesisHash/file", ctrl.getConfigFile)
}

// @Summary Save/update configuration
// @Description Save/update configuration
// @Tags configuration
// @Param tracker.cxApplicationConfig body cxApplicationConfig true "Request for creating/updating configuration"
// @Success 201
// @Failure 500 {object} api.ErrorResponse
// @Router /config [put]
func (ctrl Controller) saveConfig(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	ipAddress := getIPAddress(c.Request)

	if err := ctrl.service.createCxApplication(data, ipAddress); err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// @Summary Returns configuration for genesisHash
// @Description Returns configuration for genesisHash
// @Tags configuration
// @Produce json
// @Param genesisHash query string true "Config genesisHash"
// @Success 200 {object} tracker.CxApplication
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /config/:genesisHash [get]
func (ctrl Controller) getConfig(c *gin.Context) {
	hash := c.Param("genesisHash")
	app, err := ctrl.service.getApplicationByGenesisHash(hash)

	if err != nil {
		if err == errCannotFindApplication {
			c.AbortWithStatusJSON(http.StatusNotFound, api.ErrorResponse{Error: fmt.Errorf("%v %v", err, hash).Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// @Summary Returns configuration file for genesisHash
// @Description Returns configuration file for genesisHash
// @Tags configuration
// @Produce json
// @Param genesisHash query string true "Config genesisHash"
// @Success 200 {object} tracker.cxApplicationConfig
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /config/:genesisHash/file [get]
func (ctrl Controller) getConfigFile(c *gin.Context) {
	hash := c.Param("genesisHash")
	app, err := ctrl.service.getApplicationByGenesisHash(hash)

	if err != nil {
		if err == errCannotFindApplication {
			c.AbortWithStatusJSON(http.StatusNotFound, api.ErrorResponse{Error: fmt.Errorf("%v %v", err, hash).Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, app.Config)
}

// @Summary Returns list of all stored configurations
// @Description Returns list of all stored configurations
// @Tags configuration
// @Produce json
// @Success 200 {array} tracker.CxApplication
// @Failure 500 {object} api.ErrorResponse
// @Router /configs [get]
func (ctrl Controller) getConfigs(c *gin.Context) {
	apps, err := ctrl.service.findAllApplications()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, apps)
}

func getIPAddress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		if len(addresses[0]) == 0 {
			return r.RemoteAddr
		}
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() {
				// bad address, go to next
				continue
			}
			return ip
		}
	}
	return ""
}
