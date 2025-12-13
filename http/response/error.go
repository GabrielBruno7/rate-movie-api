package response

import (
	"crud/domain/errs"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(c *gin.Context, code string, httpStatus int) {
	domainErr := errs.NewWithCode(code, nil)
	c.JSON(httpStatus, domainErr.GetDetails())
}

func BadRequest(c *gin.Context, code string) {
	Error(c, code, 400)
}

func InternalServerError(c *gin.Context) {
	Error(c, errs.ErrInternalServer, 500)
}

func HandleError(c *gin.Context, err error) {
	if de, ok := err.(*errs.DomainError); ok {
		c.JSON(400, de.GetDetails())
		return
	}

	InternalServerError(c)
}
