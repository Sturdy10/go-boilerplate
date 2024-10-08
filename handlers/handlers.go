package handlers

import (
	"go-boilerplate/services"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Test(c *gin.Context)
}

type handler struct {
	s services.IServices
}

func NewHandler(s services.IServices) IHandler {
	return &handler{s: s}
}

func (h *handler) Test(c *gin.Context) {
	
}
