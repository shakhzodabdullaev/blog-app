package handlers

import (
	"blog/app/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
// @tags article
// @ID create-article-handler
// @Summary Create Article
// @Description Create Article By Given Info and Author ID
// @Param data body models.ArticleCreateModel true "Article Body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SuccessResponse{data=string}
// @Failure default {object} models.DefaultError
// @Router /articles [POST]
func (h *Handler) CreateArticle(c *gin.Context) {
	var entity models.ArticleCreatedModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.Article().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "article has been created",
		Data:    "ok",
	})
}

// GetArticleList godoc
// @tags article
// @ID get-all-handler
// @Summary List articles
// @Description get all articles
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search string"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ArticleListItem
// @Failure default {object} models.DefaultError
// @Router /articles [get]
func (h *Handler) GetArticleList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	resp, err := h.strg.Article().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}
