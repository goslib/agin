package agin

import (
	"github.com/gin-gonic/gin"
	"github.com/goslib/rest"
)

type Context = gin.Context
type RouterGroup = gin.RouterGroup

var New = gin.New

var conf = rest.GetConfigure()
