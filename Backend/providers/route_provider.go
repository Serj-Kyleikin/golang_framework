package providers

import (
	"github.com/gin-gonic/gin"

	v1routes "subscriptions/Backend/routes/api/v1"
)

type RouteRegistrar interface {
	Register(engine *gin.Engine)
}

type RouteProvider struct {
	registrars []RouteRegistrar
}

func BindRoutes() *RouteProvider {
	return &RouteProvider{
		registrars: []RouteRegistrar{
			v1routes.Bind(),
		},
	}
}

func (routeProvider *RouteProvider) RegisterRoutes(engine *gin.Engine) {
	for _, r := range routeProvider.registrars {
		r.Register(engine)
	}
}
