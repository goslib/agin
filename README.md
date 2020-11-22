# Agin

<!-- > 2020-11-23T04:39:24+08:00 -->

[A](http://github.com/goslib/agin) RESTful API helper for [Gin](https://github.com/gin-gonic/gin).

> See documentations [here](https://github.com/goslib/rest).


## I. Installation

```bash
go get -u github.com/goslib/agin
```


## II. Quick Start

You may check out the demo below to build your own app on.

> The demo file is [here](https://github.com/goslib/demos/blob/main/agin-main.go).

```go
package main

import "github.com/goslib/agin"

func DemoPingHandler(ctx *agin.Context, res *agin.ResponseHelper) *agin.ResponseBundle {
	return res.Done("A hello world from [agin](https://github.com/goslib/agin)!")
}

func main() {
	module := agin.NewGroupedRouter("a-demo-router", "Demo Router", "/demo",
		"The description for the router.",
		agin.NewEndingRouter("sub-router", "Sub Router", "/router",
			"A ending router with corresponding grouped routes in implemented methods.",
			agin.NewGetRoute("API Ping", "Ping your API to test your network connecting.",
				"/ping", agin.NewHandlerWrapper(DemoPingHandler)),
		),
	)

	app := agin.New()
	router := app.Group("/api")

	module.Use(router)

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
```
