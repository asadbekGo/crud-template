package handler

import (
	"app/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory
	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		h.handlerResponse(c, "error category should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Category().Create(&createCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create category resposne", http.StatusOK, resp)
}

func (h *handler) GetByIdCategory(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id category resposne", http.StatusOK, resp)
}

func (h *handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list category offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list category limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Category().GetList(&models.CategoryGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.category.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list category resposne", http.StatusOK, resp)
}
