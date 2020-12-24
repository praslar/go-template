package handlers

import (
	"github.com/astaxie/beego"
	"net/http"
	"sync"
	"time"
)

// HealthCheckResponse a json encoded health check response
type HealthCheckResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

type healthControllerImpl struct {
	beego.Controller
	name      string    // service name
	version   string    // service version
	startedAt time.Time // started time of application
}

var (
	once = sync.Once{}
	// since beego doesn't keep instance data, we could use single to achieve same result
	singleton *healthControllerImpl
)

func (c *healthControllerImpl) Get() {
	c.Data["json"] = HealthCheckResponse{
		Name:    singleton.name,
		Version: singleton.version,
		Uptime:  time.Now().Sub(singleton.startedAt).String(),
	}
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
}

// NewHealthController ...
func NewHealthController(name, version string) beego.ControllerInterface {
	once.Do(func() {
		singleton = &healthControllerImpl{
			name:      name,
			version:   version,
			startedAt: time.Now(),
		}
	})
	return singleton
}
