package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOffsetAndLimit(c *gin.Context) (int, int, error) {
	offset := c.Query("offset")
	limit := c.Query("limit")
	offsetInt := 0
	limitInt := 20
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return 0, 0, err
		}
		offsetInt = o
	}
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return 0, 0, err
		}
		limitInt = l
	}
	return offsetInt, limitInt, nil
}
