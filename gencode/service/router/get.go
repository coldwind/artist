package router

import (
	"ARTIST_PROJECT_NAME/service/control"
)

var getHandleList = map[string]*routerMethod{
	"/": {
		Handle: control.Index,
		Filter: false,
	},
}
