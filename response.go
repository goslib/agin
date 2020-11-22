package agin

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// ---------- ---------- Use a Wrapper to Transform your own Handler to the the Standard One ---------- ---------- //

// You may customize an extra response helper or wrapper, if needs.
func NewHandlerWrapper(
// handler func(res *agin.ResponseHelper, env *agin.Route, ctx *gin.Context) *agin.ResponseBundle,
	handler func(ctx *Context, res *ResponseHelper) *ResponseBundle,
) func(env *Route) func(ctx *Context) {
	// @see: [utils.go@github/gin#nameOfFunction()](https://github.com/gin-gonic/gin/blob/7742ff50e0a05d079a0c468ccfbf7c6ecfe2414b/utils.go#L123)
	name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return func(env *Route) func(ctx *Context) {
		env.HandlerName = name
		return func(ctx *Context) {
			res := NewResponseHelper(env, ctx)
			//handler(res, env, ctx)
			handler(ctx, res)
		}
	}
}

// ---------- ---------- Response Helper for Each Request ---------- ---------- //

func NewResponseHelper(env *Route, ctx *Context) *ResponseHelper {
	return &ResponseHelper{env, ctx}
}

// A response helper with the extra route context.
type ResponseHelper struct {
	Route *Route

	Context *Context
}

// ---------- ---------- Primary Request Handling ---------- ---------- //

func (m *ResponseHelper) Done(value interface{}) *ResponseBundle {
	m.Context.Status(http.StatusOK)
	m.Context.JSON(http.StatusOK, value)
	return nil
}

// Note something important.
func (m *ResponseHelper) Info(label string, values ...interface{}) {
	fmt.Println("[AGIN/INFO] ["+m.Route.GetTag()+"] [", label, "]:", values)
}

// ---------- ---------- Error Logging ---------- ---------- //

// A custom error consumer may be used following your own willingness.
// The params perfectly would be [ optional err, optional label, and optional interface{} values ].
func (m *ResponseHelper) errorout(tag string, err error, label string, values interface{}) {
	if label == "" {
		fmt.Println(tag, ">:", err, values)
	} else {
		fmt.Println(tag, "["+label+"]: <", err, ">", values)
	}
}

func (m *ResponseHelper) Error(err error, label string, values ...interface{}) {
	m.errorout("[AGIN/ERROR] ["+m.Route.GetTag()+"]", err, label, values)
}

// ---------- ---------- 4xx Client Errors ---------- ---------- //

func (m *ResponseHelper) endBadRequest(err error, label string, values []interface{}) *ResponseBundle {
	m.Context.Status(http.StatusBadRequest)
	m.errorout("[AGIN/404] ["+m.Route.GetTag()+"] [BAD_REQUEST]", err, label, values)
	return nil
}

func (m *ResponseHelper) EndBadRequest(values ...interface{}) *ResponseBundle {
	return m.endBadRequest(nil, "", values)
}

func (m *ResponseHelper) ErrorBadRequest(err error, values ...interface{}) *ResponseBundle {
	return m.endBadRequest(err, "", values)
}

func (m *ResponseHelper) BadRequest(err error, label string, values ...interface{}) *ResponseBundle {
	return m.endBadRequest(err, label, values)
}

// ---------- ---------- 5xx Server Errors ---------- ---------- //

func (m *ResponseHelper) endInternalServerError(err error, label string, values []interface{}) *ResponseBundle {
	m.Context.Status(http.StatusInternalServerError)
	m.errorout("[AGIN/500] ["+m.Route.GetTag()+"] [INTERNAL_SERVER_ERROR]", err, label, values)
	return nil
}

func (m *ResponseHelper) EndInternalServerError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "", values)
}

func (m *ResponseHelper) InternalServerError(err error, label string, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, label, values)
}

// ---------- Other Alias ---------- //

func (m *ResponseHelper) InternalDatabaseError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "Inner Database Error", values)
}

func (m *ResponseHelper) InternalServicesError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "Inner Services Error", values)
}

func (m *ResponseHelper) Internal3rdServicesError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "3rd-party Services Error", values)
}

// ---------- ---------- Response Hook Bundle ---------- ---------- //

//type IResponseBundle interface {
//}

// 按需自定义扩展
type ResponseBundle struct {
	Code int

	Error error

	Response interface{}

	Values []interface{}
}

//func NewHandlerWrapper(
//	fn func(ctx *gin.Context, res *agin.ResponseHelper) *agin.ResponseBundle,
//) func(env *agin.Route, ctx *gin.Context) interface{} {
//	return func(env *agin.Route, ctx *gin.Context) interface{} {
//		res := agin.NewResponseHelper(env, ctx)
//		return fn(ctx, res)
//	}
//}
