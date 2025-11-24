package handlers

/*
Copyright © 2025 Dinas Komunikasi dan Informatika DIY <diskominfo@jogjaprov.go.id>
Pusat Layanan Transformasi Digital

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	//autogenerate :import
	//endautogenerate :import

	"github.com/gin-gonic/gin"

	"YardPlaning/internal/services"
)

type StackingHandler struct {
	//autogenerate :handlerStruct
	service *services.YardPlanService
	//endautogenerate :handlerStruct
}

// autogenerate :handlerNewStruct
func NewStackingHandler(service *services.YardPlanService) *StackingHandler {
	return &StackingHandler{service}
}

//endautogenerate :handlerNewStruct

// CRUD

// @Summary	Suggest Placement
// @Tags		Stacking
// @Param		request	formData	entities.SuggestRequest	true	"Payload data"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/suggestion [post]
func (h *StackingHandler) Suggest(c *gin.Context) {

	//autogenerate :handlerfunc
	//endautogenerate :handlerfunc
}

// @Summary	Suggest Placement
// @Tags		Stacking
// @Param		request	formData	entities.SuggestRequest	true	"Payload data"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/placement [post]
func (h *StackingHandler) Placement(c *gin.Context) {

	//autogenerate :handlerfunc
	//endautogenerate :handlerfunc
}

// @Summary	Suggest Placement
// @Tags		Stacking
// @Param		request	formData	entities.SuggestRequest	true	"Payload data"
// @Produce	json
// @Success	201	{object}	entities.ReturnObject
// @Router		/pickup [post]
func (h *StackingHandler) Pickup(c *gin.Context) {

	//autogenerate :handlerfunc
	//endautogenerate :handlerfunc
}
