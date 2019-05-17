package tracker

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	public.PUT("/config", ctrl.saveConfig)
	public.GET("/configs", ctrl.getConfigs)
	public.GET("/config/:hash", ctrl.getConfig)
}

// @Summary Returns uptime info for previous month
// @Description Returns uptime info for nodes from the request
// @Tags config
// @Produce json
// @Success 201 string
// @Failure 500 {object} api.ErrorResponse
// @Router /config [put]
func (ctrl Controller) saveConfig(c *gin.Context) {
	var conf cxApplicationConfig //TODO consider this approach for validating the request
	if err := c.BindJSON(&conf); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, api.ErrorResponse{Error: errUnableToParseConfig.Error()})
		return
	}
	if len(conf.GenesisHash) == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, api.ErrorResponse{Error: errUnableToParseConfig.Error()})
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	ipAddress := getIPAddress(c.Request)

	hash, err := ctrl.service.createCxApplication(data, ipAddress)
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
	app, err := ctrl.service.getApplicationBy(hash)

	if err != nil {
		if err == errCannotFindUser {
			c.AbortWithStatusJSON(http.StatusNotFound, api.ErrorResponse{Error: err.Error() + hash})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// @Summary Returns list of all stored configs
// @Description Returns list of all stored configs
// @Tags config
// @Produce json
// @Success 200 {array} string
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
