package router

import "ARTIST_PROJECT_NAME/service/api/control"

var postHandleList = map[string]*routerMethod{
	"/": {
		Handle: control.Index,
		Filter: false,
	},
}
