package library


type Controller struct {
	BaseUrl string
	ActionName string
	ControllerName string
}

func (c Controller) setBashUrl(url string){
	c.BaseUrl = url
}

func (c Controller) setActionName(actionName string){
	c.ActionName = actionName
}

func (c Controller) setControllerName(cName string) {
	c.ControllerName = cName
}

