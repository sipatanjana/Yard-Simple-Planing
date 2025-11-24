package entities

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

type SuggestRequest struct {
	Yard            uint    `json:"yard" binding:"required"`
	ContainerNumber string  `json:"container_number" binding:"required"`
	ContainerSize   int8    `json:"container_size" binding:"required"`
	ContainerHeight float32 `json:"container_height" binding:"required"`
	ContainerType   string  `json:"container_type" binding:"required"`
}

type PlacementRequest struct {
	Yard            uint   `json:"yard" binding:"required"`
	ContainerNumber string `json:"container_number" binding:"required"`
	SLot            uint   `json:"slot" binding:"required"`
	Row             uint   `json:"row" binding:"required"`
	Tier            uint   `json:"tier" binding:"required"`
}

type PickupRequest struct {
	Yard            uint   `json:"yard" binding:"required"`
	ContainerNumber string `json:"container_number" binding:"required"`
}
