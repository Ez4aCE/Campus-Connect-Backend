package middleware


import (
	"net/http"


	"github.com/gin-gonic/gin"
)

func RequireRoles(roles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		roleVal, exists :=c.Get("role")

		if !exists{
			c.JSON(http.StatusForbidden, gin.H{"error":"role not found"})
			c.Abort()
			return
		}

		userRole :=roleVal.(string)

		for _,role := range roles{
			if userRole==role{
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error":"access denied"})
		c.Abort()
	}
}