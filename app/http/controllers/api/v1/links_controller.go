package v1

import (
    "github.com/gin-gonic/gin"
    "gohub/app/models/link"
    "gohub/pkg/response"
)

type LinksController struct {
    BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    links := link.All()
    response.Data(c, links)
}