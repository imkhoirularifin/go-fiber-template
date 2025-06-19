package utils

import (
	"go-fiber-template/lib/constant"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetPaginationHeader(ctx *fiber.Ctx, page, limit, totalCount int) {
	var nextPage *int
	var prevPage *int
	totalPages := (totalCount + limit - 1) / limit

	if page < totalPages {
		next := page + 1
		nextPage = &next
	}

	if page > 1 {
		prev := page - 1
		prevPage = &prev
	}

	ctx.Set(constant.HeaderXTotalCount, strconv.Itoa(totalCount))
	ctx.Set(constant.HeaderXTotalPages, strconv.Itoa(totalPages))
	ctx.Set(constant.HeaderXPage, strconv.Itoa(page))
	ctx.Set(constant.HeaderXLimit, strconv.Itoa(limit))

	if nextPage != nil {
		ctx.Set(constant.HeaderXNextPage, strconv.Itoa(*nextPage))
	}

	if prevPage != nil {
		ctx.Set(constant.HeaderXPrevPage, strconv.Itoa(*prevPage))
	}
}
