package router

import (
	"pandax/apps/system/api"

	"github.com/emicklei/go-restful/v3"
)

func InitIpamRouter(container *restful.Container) {
	s := &api.System{}
	ws := new(restful.WebService)
	ws.Path("/ipam").Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(s.ConnectWs))
	ws.Route(ws.GET("/server").To(s.ServerInfo))
	container.Add(ws)
}
