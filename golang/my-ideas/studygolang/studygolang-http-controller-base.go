// Copyright 2016 The StudyGolang Authors. All rights reserved.

package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/studygolang/studygolang/context"
	. "github.com/studygolang/studygolang/http"
	"github.com/studygolang/studygolang/logic"

	echo "github.com/labstack/echo/v4"
	"github.com/polaris1119/goutils"
	"github.com/polaris1119/logger"
	"github.com/polaris1119/nosql"
)

func getLogger(ctx echo.Context) *logger.Logger {
	return logic.GetLogger(context.EchoContext(ctx))
}

// render html 输出
func render(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	return Render(ctx, contentTpl, data)
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

	oldETag := ctx.Request().Header.Get("If-None-Match")
	if strings.HasPrefix(oldETag, "W/") {
		oldETag = oldETag[2:]
	}
	newETag := goutils.Md5Buf(b)
	if oldETag == newETag {
		return ctx.NoContent(http.StatusNotModified)
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

	ctx.Response().Header().Add("ETag", newETag)

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

// http-controller-routes.go

package controller

import echo "github.com/labstack/echo/v4"

func RegisterRoutes(g *echo.Group) {
	new(IndexController).RegisterRoute(g)
	new(AccountController).RegisterRoute(g)
	new(TopicController).RegisterRoute(g)
	new(ArticleController).RegisterRoute(g)
	new(ProjectController).RegisterRoute(g)
	new(ResourceController).RegisterRoute(g)
	new(ReadingController).RegisterRoute(g)
	new(WikiController).RegisterRoute(g)
	new(UserController).RegisterRoute(g)
	new(LikeController).RegisterRoute(g)
	new(FavoriteController).RegisterRoute(g)
	new(MessageController).RegisterRoute(g)
	new(SidebarController).RegisterRoute(g)
	new(CommentController).RegisterRoute(g)
	new(SearchController).RegisterRoute(g)
	new(WideController).RegisterRoute(g)
	new(ImageController).RegisterRoute(g)
	new(CaptchaController).RegisterRoute(g)
	new(BookController).RegisterRoute(g)
	new(MissionController).RegisterRoute(g)
	new(UserRichController).RegisterRoute(g)
	new(TopController).RegisterRoute(g)
	new(GiftController).RegisterRoute(g)
	new(OAuthController).RegisterRoute(g)
	new(WebsocketController).RegisterRoute(g)
	new(DownloadController).RegisterRoute(g)
	new(LinkController).RegisterRoute(g)
	new(SubjectController).RegisterRoute(g)
	new(GCTTController).RegisterRoute(g)

	new(FeedController).RegisterRoute(g)
	new(WechatController).RegisterRoute(g)
	new(InstallController).RegisterRoute(g)
}
