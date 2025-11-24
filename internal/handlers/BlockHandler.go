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

type BlockHandler struct {
	//autogenerate :handlerStruct
	service *services.BlockService
	//endautogenerate :handlerStruct
}

// autogenerate :handlerNewStruct
func NewBlockHandler(service *services.BlockService) *BlockHandler {
	return &BlockHandler{service}
}

//endautogenerate :handlerNewStruct

// CRUD

// @Summary	Create data
// @Tags		Block
// @Param		request	formData	entities.BlockRequest	true	"Payload data"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/block [post]
func (h *BlockHandler) Insert(c *gin.Context) {
	//autogenerate :handlerfunc
	var payload entities.BlockRequest
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := entities.Block{
		Name:   payload.Name,
		Code:   payload.Code,
		YardID: payload.Yard,
		Slot:   payload.Slot,
		Row:    payload.Row,
		Tier:   payload.Tier,
	}

	if err := h.service.Insert(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating block"})
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
// @Tags		Block
// @Param		request	formData	entities.BlockRequest	true	"Payload data"
// @Param		id		path		string					true	"ID"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/block/{id} [put]
func (h *BlockHandler) Update(c *gin.Context) {

	//autogenerate :handlerfunc
	var payload entities.BlockRequest

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
	data := entities.Block{
		ID:     uint(id),
		Name:   payload.Name,
		Code:   payload.Code,
		YardID: payload.Yard,
		Slot:   payload.Slot,
		Row:    payload.Row,
		Tier:   payload.Tier,
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

func (h *BlockHandler) SetEnable(c *gin.Context) {

	//autogenerate :handlerfunc
	var block entities.Block

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

	block.ID = uint(id)
	if err := h.service.SetEnable(&block, enable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    block,
	})
	//endautogenerate :handlerfunc
}

func (h *BlockHandler) Delete(c *gin.Context) {

	//autogenerate :handlerfunc
	var block entities.Block

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	block.ID = uint(id)
	if err := h.service.Delete(&block); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    block,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Get all data
// @Tags		Block
// @Param		page			query	integer	false	"Page"
// @Param		limit			query	integer	false	"Limit"
// @Param		sort			query	string	false	"default asc"	Enums(asc, desc)	default(asc)
// @Param		order_by		query	string	false	"Order By"
// @Param		filter[name]	query	string	false	"keyword nama"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/block [get]
func (h *BlockHandler) GetAll(c *gin.Context) {

	//autogenerate :handlerfunc
	var blocks []entities.Block
	pagination := utils.BindPagination(c)
	if err := h.service.GetAll(&pagination, &blocks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	c.JSON(http.StatusOK, entities.ReturnObject{
		Status:  http.StatusOK,
		Message: "success",
		Data:    blocks,
		Meta:    pagination,
	})
	//endautogenerate :handlerfunc
}

// @Summary	Get spesifik data
// @Tags		Block
// @Param		id	path	string	true	"ID"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/block/{id} [get]
func (h *BlockHandler) GetById(c *gin.Context) {

	//autogenerate :handlerfunc
	var returnObject entities.ReturnObject
	var block entities.Block

	//cek paramter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter(s)"})
		return
	}

	block.ID = uint(id)
	if err := h.service.GetById(&block); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnObject.Data = block
	returnObject.Status = http.StatusOK
	returnObject.Message = "Success"
	c.JSON(http.StatusOK, returnObject)
	//endautogenerate :handlerfunc
}
