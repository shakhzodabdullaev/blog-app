package handlers

import (
	"blog/app/config"
	"blog/app/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Author struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewAuthor(strg storage.StorageI, cfg config.Config) Handler {
	return Handler{
		strg: strg,
		cfg:  cfg,
	}
}

func (h *Author) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Author) getLimitParam(c *gin.Context) (limit int, err error) {
	limitStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(limitStr)
}
