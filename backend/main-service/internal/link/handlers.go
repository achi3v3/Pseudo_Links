package link

import (
	"main-service/backend/main-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service *LinkService
}

func NewLinkHandler(service *LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

func (h *LinkHandler) Create(ctx *gin.Context) {
	var jsonBody struct {
		OriginLink string `json:"origin_link" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&jsonBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format: " + err.Error(),
		})
		return
	}

	pseudoLink, err := h.service.Create(jsonBody.OriginLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create link: " + err.Error(),
		})
		return
	}
	link := models.Link{
		OriginLink: jsonBody.OriginLink,
		PseudoLink: pseudoLink,
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Link created successfully",
		"data":    link,
	})
}

func (h *LinkHandler) Get(ctx *gin.Context) {
	originLink := ctx.Query("origin_link")
	if originLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "origin_link parameter is required",
		})
		return
	}
	link, err := h.service.GetLink(originLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get link: " + err.Error(),
		})
		return
	}

	if link == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Link not found",
		})
		return
	}

	// Возвращаем найденную ссылку
	ctx.JSON(http.StatusOK, gin.H{
		"data": link,
	})
}

func (h *LinkHandler) Delete(ctx *gin.Context) {
	originLink := ctx.Query("origin_link")
	if originLink == "" {

	}
	if err := h.service.DeleteLink(originLink); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete link: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Link deleted successfully",
	})

}
