package utils

import (
	"go-fiber-template/lib/constant"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetPaginationHeader(ctx *fiber.Ctx, page, limit, totalCount int64) {
	// Prevent division by zero
	if limit == 0 {
		limit = 10 // Default limit
	}
	if page == 0 {
		page = 1 // Default page
	}

	var nextPage *int64
	var prevPage *int64
	totalPages := (totalCount + limit - 1) / limit

	if page < totalPages {
		next := page + 1
		nextPage = &next
	}

	if page > 1 {
		prev := page - 1
		prevPage = &prev
	}

	ctx.Set(constant.HeaderXTotalCount, strconv.FormatInt(totalCount, 10))
	ctx.Set(constant.HeaderXTotalPages, strconv.FormatInt(totalPages, 10))
	ctx.Set(constant.HeaderXPage, strconv.FormatInt(page, 10))
	ctx.Set(constant.HeaderXLimit, strconv.FormatInt(limit, 10))

	if nextPage != nil {
		ctx.Set(constant.HeaderXNextPage, strconv.FormatInt(*nextPage, 10))
	}

	if prevPage != nil {
		ctx.Set(constant.HeaderXPrevPage, strconv.FormatInt(*prevPage, 10))
	}
}
