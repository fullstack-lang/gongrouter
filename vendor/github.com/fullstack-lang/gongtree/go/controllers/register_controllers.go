package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongtree/go")
	{ // insertion point for registrations
		v1.GET("/v1/buttons", GetController().GetButtons)
		v1.GET("/v1/buttons/:id", GetController().GetButton)
		v1.POST("/v1/buttons", GetController().PostButton)
		v1.PATCH("/v1/buttons/:id", GetController().UpdateButton)
		v1.PUT("/v1/buttons/:id", GetController().UpdateButton)
		v1.DELETE("/v1/buttons/:id", GetController().DeleteButton)

		v1.GET("/v1/nodes", GetController().GetNodes)
		v1.GET("/v1/nodes/:id", GetController().GetNode)
		v1.POST("/v1/nodes", GetController().PostNode)
		v1.PATCH("/v1/nodes/:id", GetController().UpdateNode)
		v1.PUT("/v1/nodes/:id", GetController().UpdateNode)
		v1.DELETE("/v1/nodes/:id", GetController().DeleteNode)

		v1.GET("/v1/trees", GetController().GetTrees)
		v1.GET("/v1/trees/:id", GetController().GetTree)
		v1.POST("/v1/trees", GetController().PostTree)
		v1.PATCH("/v1/trees/:id", GetController().UpdateTree)
		v1.PUT("/v1/trees/:id", GetController().UpdateTree)
		v1.DELETE("/v1/trees/:id", GetController().DeleteTree)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
