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

type YardPlanHandler struct {
	//autogenerate :handlerStruct
	service *services.YardPlanService

	//endautogenerate :handlerStruct
}

// autogenerate :handlerNewStruct
func NewYardPlanHandler(service *services.YardPlanService) *YardPlanHandler {
	return &YardPlanHandler{service}
}

//endautogenerate :handlerNewStruct

// CRUD

// @Summary	Create data
// @Tags		Yard Plan
// @Param		request	formData	entities.YardPlanRequest	true	"Payload data"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yardPlan [post]
func (h *YardPlanHandler) Insert(c *gin.Context) {

	//autogenerate :handlerfunc
	var payload entities.YardPlanRequest
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var block entities.Block
	block.ID = uint(payload.Block)
	if err := h.service.GetBLockByID(&block); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := entities.YardPlan{
		YardID:          block.YardID,
		BlockID:         payload.Block,
		ContainerSize:   payload.ContainerSize,
		ContainerHeight: payload.ContainerHeight,
		ContainerType:   payload.ContainerType,
		SlotFrom:        payload.SlotFrom,
		SlotTo:          payload.SlotTo,
		RowFrom:         payload.RowFrom,
		RowTo:           payload.RowTo,
		TierFrom:        payload.TierFrom,
		TierTo:          payload.TierTo,
	}
	if err := h.service.Insert(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating yardPlan"})
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
// @Tags		Yard Plan
// @Param		request	formData	entities.YardPlanRequest	true	"Payload data"
// @Param		id		path		string						true	"ID"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yardPlan/{id} [put]
func (h *YardPlanHandler) Update(c *gin.Context) {

	//autogenerate :handlerfunc
	var payload entities.YardPlanRequest

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	//binding json
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var block entities.Block
	block.ID = uint(payload.Block)
	if err := h.service.GetBLockByID(&block); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//update
	data := entities.YardPlan{
		ID:              uint(id),
		YardID:          block.YardID,
		BlockID:         payload.Block,
		ContainerSize:   payload.ContainerSize,
		ContainerHeight: payload.ContainerHeight,
		ContainerType:   payload.ContainerType,
		SlotFrom:        payload.SlotFrom,
		SlotTo:          payload.SlotTo,
		RowFrom:         payload.RowFrom,
		RowTo:           payload.RowTo,
		TierFrom:        payload.TierFrom,
		TierTo:          payload.TierTo,
	}
	if err := h.service.Update(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "Updated",
		Data:    payload,
	})
	//endautogenerate :handlerfunc
}

func (h *YardPlanHandler) SetEnable(c *gin.Context) {

	//autogenerate :handlerfunc
	var yardPlan entities.YardPlan

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

	yardPlan.ID = uint(id)
	if err := h.service.SetEnable(&yardPlan, enable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    yardPlan,
	})
	//endautogenerate :handlerfunc
}

func (h *YardPlanHandler) Delete(c *gin.Context) {

	//autogenerate :handlerfunc
	var yardPlan entities.YardPlan

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	yardPlan.ID = uint(id)
	if err := h.service.Delete(&yardPlan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    yardPlan,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Get all data
// @Tags		Yard Plan
// @Param		page		query	integer	false	"Page"
// @Param		limit		query	integer	false	"Limit"
// @Param		sort		query	string	false	"default asc"	Enums(asc, desc)	default(asc)
// @Param		order_by	query	string	false	"Order By"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yardPlan [get]
func (h *YardPlanHandler) GetAll(c *gin.Context) {

	//autogenerate :handlerfunc
	var yardPlans []entities.YardPlan
	pagination := utils.BindPagination(c)
	if err := h.service.GetAll(&pagination, &yardPlans); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    yardPlans,
		Meta:    pagination,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Get spesifik data
// @Tags		Yard Plan
// @Param		id	path	string	true	"ID"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/yardPlan/{id} [get]
func (h *YardPlanHandler) GetById(c *gin.Context) {

	//autogenerate :handlerfunc
	var returnObject entities.ReturnObject
	var yardPlan entities.YardPlan

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	yardPlan.ID = uint(id)
	if err := h.service.GetById(&yardPlan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnObject.Data = yardPlan
	returnObject.Status = http.StatusOK
	returnObject.Message = "Success"
	c.JSON(http.StatusOK, returnObject)
	//endautogenerate :handlerfunc
}
