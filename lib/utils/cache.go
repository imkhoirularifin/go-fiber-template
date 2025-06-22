package utils

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var cacheHeaderKeys = []string{
	fiber.HeaderAcceptLanguage,
	fiber.HeaderAuthorization,
	fiber.HeaderContentType,
	fiber.HeaderUserAgent,
	fiber.HeaderXForwardedFor,
	fiber.HeaderXForwardedHost,
	fiber.HeaderXForwardedProto,
}

// CacheKeyWithQueryAndHeaders generates a cache key including both query parameters and headers
func CacheKeyWithQueryAndHeaders(c *fiber.Ctx) string {
	parts := []string{c.Path()}

	// Add query parameters
	queryArgs := c.Context().QueryArgs()
	if queryArgs.Len() > 0 {
		params := make([]string, 0)
		queryArgs.VisitAll(func(key, value []byte) {
			params = append(params, fmt.Sprintf("%s=%s", string(key), string(value)))
		})
		sort.Strings(params)
		parts = append(parts, fmt.Sprintf("q:%s", strings.Join(params, "&")))
	}

	// Add headers
	if len(cacheHeaderKeys) > 0 {
		headerValues := make([]string, 0)
		for _, key := range cacheHeaderKeys {
			value := c.Get(key)
			if value != "" {
				headerValues = append(headerValues, fmt.Sprintf("%s:%s", key, value))
			}
		}
		if len(headerValues) > 0 {
			sort.Strings(headerValues)
			parts = append(parts, fmt.Sprintf("h:%s", strings.Join(headerValues, "|")))
		}
	}

	return strings.Join(parts, "|")
}
