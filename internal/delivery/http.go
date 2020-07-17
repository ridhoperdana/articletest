// GENERATED BY goruda
// This file was generated automatically at
// 2020-07-17 07:53:27.975243 +0700 WIB m=+0.022328012

package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ridhoperdana/articletest"
)

type httpHandler struct {
	service articletest.ArticleService
}

func (h httpHandler) createArticles(c echo.Context) error {

	fromRequest0 := articletest.Article{}
	if err := c.Bind(&fromRequest0); err != nil {
		return err
	}

	input0 := fromRequest0

	result0, err := h.service.Createarticles(input0)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result0)

}

func (h httpHandler) deleteArticleById(c echo.Context) error {
	fromRequest0 := c.Param("articleId")

	input0 := fromRequest0

	err := h.service.Deletearticlebyid(input0)
	if err != nil {
		switch v := err.(type) {
		case articletest.ErrorNotFound:
			return c.JSON(http.StatusNotFound, v)
		}
		return err
	}

	return c.NoContent(http.StatusNoContent)

}

func (h httpHandler) listArticles(c echo.Context) error {
	fromRequest0 := c.QueryParam("num")
	var (
		num int64
		err error
	)

	if fromRequest0 != "" {
		num, err = strconv.ParseInt(fromRequest0, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, articletest.ErrorBadRequest{
				Message: err.Error(),
			})
		}
	}

	input0 := num

	fromRequest1 := c.QueryParam("cursor")

	input1 := fromRequest1

	result0, result1, err := h.service.Listarticles(int32(input0), input1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, articletest.Error{
			Message: err.Error(),
		})
	}

	c.Response().Header().Set("X-Cursor", result1)

	return c.JSON(http.StatusOK, result0)

}

func (h httpHandler) showArticleById(c echo.Context) error {

	fromRequest0 := c.Param("articleId")

	input0 := fromRequest0

	result0, err := h.service.Showarticlebyid(input0)
	if err != nil {
		switch v := err.(type) {
		case articletest.ErrorNotFound:
			return c.JSON(http.StatusNotFound, v)
		}
		return c.JSON(http.StatusInternalServerError, articletest.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result0)

}

func (h httpHandler) updateArticle(c echo.Context) error {

	fromRequest0 := c.Param("articleId")

	input0 := fromRequest0

	fromRequest1 := articletest.Article{}
	if err := c.Bind(&fromRequest1); err != nil {
		return err
	}

	input1 := fromRequest1

	result0, err := h.service.Updatearticle(input0, input1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, articletest.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result0)

}

func RegisterHTTPPath(e *echo.Echo, service articletest.ArticleService) {
	handler := httpHandler{
		service: service,
	}
	e.POST("/article", handler.createArticles)
	e.DELETE("/article/:articleId", handler.deleteArticleById)
	e.GET("/article", handler.listArticles)
	e.GET("/article/:articleId", handler.showArticleById)
	e.PUT("/article/:articleId", handler.updateArticle)
}
