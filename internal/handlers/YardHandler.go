package handlers

import (
	//autogenerate :import
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"YardPlaning/internal/entities"
	"YardPlaning/internal/services"
	"YardPlaning/pkg/utils"
)

type YardHandler struct {
	//autogenerate :handlerStruct
	service *services.YardService
	//endautogenerate :handlerStruct
}

// autogenerate :handlerNewStruct
func NewYardHandler(service *services.YardService) *YardHandler {
	return &YardHandler{service}
}

//endautogenerate :handlerNewStruct

// CRUD

// @Summary	Create data
// @Tags		Yard
// @Param		request	formData	entities.YardRequest	true	"Payload data"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yard [post]
func (h *YardHandler) Insert(c *gin.Context) {

	//autogenerate :handlerfunc
	var payload entities.YardRequest
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := entities.Yard{
		Name: payload.Name,
	}

	if err := h.service.Insert(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating yard"})
		return
	}

	c.JSON(http.StatusCreated, entities.ReturnObject{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    payload,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Update data
// @Tags		Yard
// @Param		request	formData	entities.YardRequest	true	"Payload data"
// @Param		id		path		string					true	"ID"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yard/{id} [put]
func (h *YardHandler) Update(c *gin.Context) {

	//autogenerate :handlerfunc
	var payload entities.YardRequest

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	//binding json
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update
	var yard entities.Yard
	yard.ID = uint(id)
	yard.Name = payload.Name
	if err := h.service.Update(&yard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "Updated",
		Data:    yard,
	})
	//endautogenerate :handlerfunc
}

func (h *YardHandler) SetEnable(c *gin.Context) {

	//autogenerate :handlerfunc
	var yard entities.Yard

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	//cek paramter
	enable, err := strconv.ParseBool(c.Param("enable"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	yard.ID = uint(id)
	if err := h.service.SetEnable(&yard, enable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    yard,
	})
	//endautogenerate :handlerfunc
}

func (h *YardHandler) Delete(c *gin.Context) {

	//autogenerate :handlerfunc
	var yard entities.Yard

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	yard.ID = uint(id)
	if err := h.service.Delete(&yard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    yard,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Get all data
// @Tags		Yard
// @Param		page			query	integer	false	"Page"
// @Param		limit			query	integer	false	"Limit"
// @Param		sort			query	string	false	"default asc"	Enums(asc, desc)	default(asc)
// @Param		order_by		query	string	false	"Order By"
// @Param		filter[name]	query	string	false	"keyword nama"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yard [get]
func (h *YardHandler) GetAll(c *gin.Context) {

	//autogenerate :handlerfunc
	var yards []entities.Yard
	pagination := utils.BindPagination(c)
	if err := h.service.GetAll(&pagination, &yards); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    yards,
		Meta:    pagination,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Get spesifik data
// @Tags		Yard
// @Param		id	path	string	true	"ID"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yard/{id} [get]
func (h *YardHandler) GetById(c *gin.Context) {

	//autogenerate :handlerfunc
	var returnObject entities.ReturnObject
	var yard entities.Yard

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	yard.ID = uint(id)
	if err := h.service.GetById(&yard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnObject.Data = yard
	returnObject.Status = http.StatusOK
	returnObject.Message = "Success"
	c.JSON(http.StatusOK, returnObject)
	//endautogenerate :handlerfunc
}
