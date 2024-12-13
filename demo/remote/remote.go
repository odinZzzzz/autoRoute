package remote

import (
	"github.com/odinZzzzz/autoRoute"
)

func InitRemote() *autoRoute.AutoRoute {
	route := autoRoute.AutoRoute{}
	route.Register("rpc", gameRemote{AutoHandler: autoRoute.AutoHandler{}})
	return &route
}
