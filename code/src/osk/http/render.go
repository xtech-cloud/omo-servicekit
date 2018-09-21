package http

import (
	"net/http"

	"osk/core"

	"github.com/gin-gonic/gin"
)

func CatchRenderError() {
	if err := recover(); nil != err {
		core.Logger.Error(err)
	}
}

func TryRenderBindError(_context *gin.Context, _err error) {
	if nil == _err {
		return
	}

	_context.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": _err.Error(),
		"data":    gin.H{},
	})
	panic(_err)
}

func TryRenderDatabaseError(_context *gin.Context, _err error) {
	if nil == _err {
		return
	}

	_context.JSON(http.StatusOK, gin.H{
		"code":    2,
		"message": _err.Error(),
		"data":    gin.H{},
	})
	panic(_err)
}

func TryRenderInternalError(_context *gin.Context, _err error) {
	if nil == _err {
		return
	}

	_context.JSON(http.StatusOK, gin.H{
		"code":    3,
		"message": _err.Error(),
		"data":    gin.H{},
	})
	panic(_err)
}

func RenderError(_context *gin.Context, _code int, _message string) {
	_context.JSON(http.StatusOK, gin.H{
		"code":    _code,
		"message": _message,
		"data":    gin.H{},
	})
}
func RenderOK(_context *gin.Context, _data interface{}) {
	_context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    _data,
	})
}
