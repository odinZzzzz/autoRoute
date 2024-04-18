package handler

import "autoRoute"

func InitHandler() *autoRoute.AutoRoute {
	route := autoRoute.AutoRoute{
		Debug: false,
	}
	route.Register("game", gameHandler{AutoHandler: autoRoute.AutoHandler{}})
	return &route
}
