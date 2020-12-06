package api

import "github.com/gin-gonic/gin"

func getDefaultQueries(c *gin.Context, queries map[string]string) map[string]string {
	for key, val := range queries {
		queries[key] = c.DefaultQuery(key, val)
	}
	return queries
}
