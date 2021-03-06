package api

import (
	"net/http"
)

func JsonResponse(ctx Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func ErrResponse(ctx Context, data interface{}) {
	ctx.JSON(http.StatusNotFound, data)
}
