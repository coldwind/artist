package router

import (
	"ARTIST_PROJECT_NAME/service/api/control"
)

var getHandleList = map[string]*routerMethod{
	"/": {
		Handle: control.Index,
		Filter: false,
	},
}
