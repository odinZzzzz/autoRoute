package autoRoute

var autoHandlerMap map[string]map[string]handlerData = make(map[string]map[string]handlerData)
var RouteOpt = RouteOption{
	Debug:    false,
	UseProto: false,
}

type AutoRoute struct {
}
type RouteOption struct {
	Debug    bool `是否开启Debug调试详细日志`
	UseProto bool `接口接参使用protoBuf`
}
type AutoHandler struct {
	HandlerName string
	WhiteList   []string
}

func (a *AutoRoute) Register(group string, handler interface{}) {
	a.registerHandler(group, handler)
}
