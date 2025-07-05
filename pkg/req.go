package pkg

import "github.com/gin-gonic/gin"

func ParseURL(c *gin.Context,req any) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	return nil
}
func ParseBody(c *gin.Context, req any) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}
	return nil
}