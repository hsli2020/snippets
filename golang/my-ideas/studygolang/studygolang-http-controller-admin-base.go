// Copyright 2016 The StudyGolang Authors. All rights reserved.

package admin

import (
	"encoding/json"
	"net/http"

	"github.com/studygolang/studygolang/context"
	. "github.com/studygolang/studygolang/http"
	"github.com/studygolang/studygolang/logic"

	echo "github.com/labstack/echo/v4"
	"github.com/polaris1119/goutils"
	"github.com/polaris1119/logger"
	"github.com/polaris1119/nosql"
)

func parsePage(ctx echo.Context) (curPage, limit int) {
	curPage = goutils.MustInt(ctx.FormValue("page"), 1)
	limit = goutils.MustInt(ctx.FormValue("limit"), 20)
	return
}

func parseConds(ctx echo.Context, fields []string) map[string]string {
	conds := make(map[string]string)

	for _, field := range fields {
		if value := ctx.FormValue(field); value != "" {
			conds[field] = value
		}
	}

	return conds
}

func getLogger(ctx echo.Context) *logger.Logger {
	return logic.GetLogger(context.EchoContext(ctx))
}

// render html 输出
func render(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	return RenderAdmin(ctx, contentTpl, data)
}

func renderQuery(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	return RenderQuery(ctx, contentTpl, data)
}

func success(ctx echo.Context, data interface{}) error {
	result := map[string]interface{}{
		"ok":   1,
		"msg":  "操作成功",
		"data": data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	go func(b []byte) {
		if cacheKey := ctx.Get(nosql.CacheKey); cacheKey != nil {
			nosql.DefaultLRUCache.CompressAndAdd(cacheKey, b, nosql.NewCacheData())
		}
	}(b)

	if ctx.Response().Committed {
		getLogger(ctx).Flush()
		return nil
	}

	return ctx.JSONBlob(http.StatusOK, b)
}

func fail(ctx echo.Context, code int, msg string) error {
	if ctx.Response().Committed {
		getLogger(ctx).Flush()
		return nil
	}

	result := map[string]interface{}{
		"ok":    0,
		"error": msg,
	}

	getLogger(ctx).Errorln("operate fail:", result)

	return ctx.JSON(http.StatusOK, result)
}

// ## http-controller-admin-routes.go

package admin

import echo "github.com/labstack/echo/v4"

func RegisterRoutes(g *echo.Group) {
	g.GET("", AdminIndex)
	new(AuthorityController).RegisterRoute(g)
	new(UserController).RegisterRoute(g)
	new(TopicController).RegisterRoute(g)
	new(NodeController).RegisterRoute(g)
	new(ArticleController).RegisterRoute(g)
	new(ProjectController).RegisterRoute(g)
	new(RuleController).RegisterRoute(g)
	new(ReadingController).RegisterRoute(g)
	new(ToolController).RegisterRoute(g)
	new(SettingController).RegisterRoute(g)
	new(MetricsController).RegisterRoute(g)
}
