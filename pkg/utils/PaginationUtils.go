package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page      int    `json:"page"`
	TotalPage int64  `json:"totalPage"`
	Limit     int    `json:"limit"`
	OrderBy   string `json:"order_by"`
	Sort      string `json:"sort"` //asc
	//Keyword      string            `json:"keyword"`
	Filters map[string]string `json:"filters"`
	Total   int64             `json:"total"`
	//ActiveColumn []string          `json:"activeColumn"`
}

func BindPagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "asc")
	orderBy := c.DefaultQuery("order_by", "id")
	//keyword := c.DefaultQuery("keyword", "")

	filters := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		if strings.HasPrefix(key, "filter[") && strings.HasSuffix(key, "]") {
			column := key[7 : len(key)-1] // ambil nama kolom dari key filter[col]
			filters[column] = values[0]
		}
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return Pagination{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		OrderBy: orderBy,
		//Keyword: keyword,
		Filters: filters,
	}
}

func (p *Pagination) Paginate(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	order := p.OrderBy + " " + p.Sort
	return db.Offset(offset).Limit(p.Limit).Order(order)
}

/* func BuildSearchCondition(term string, columns []string) string {
	like := "%" + term + "%"
	var conditions []string
	for _, col := range columns {
		conditions = append(conditions, col+" ILIKE '"+like+"'")
	}
	return "(" + strings.Join(conditions, " OR ") + ")"
} */

func BuildFilterCondition(filters map[string]string) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	for col, val := range filters {
		conditions = append(conditions, fmt.Sprintf("%s ILIKE ?", col))
		args = append(args, "%"+val+"%")
	}

	return strings.Join(conditions, " AND "), args
}
