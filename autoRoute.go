package autoRoute

var autoHandlerMap map[string]map[string]interface{} = make(map[string]map[string]interface{})

type AutoRoute struct {
	Debug bool
}
type AutoHandler struct {
	HandlerName string
	WhiteList   []string
}

func (a *AutoRoute) Register(group string, handler interface{}) {
	a.dealHandler(group, handler)
}
