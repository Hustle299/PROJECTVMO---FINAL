package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase usecase.AdminUsecaseInterface
}

func NewAuthenticationHandler(usecase usecase.AdminUsecaseInterface) *AdminHandler {
	return &AdminHandler{usecase: usecase}
}

func (h *AdminHandler) GetListPendingRequest(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, msg := h.usecase.GetListPendingRequest()
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AdminHandler) GetPendingRequestById(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	resp, msg := h.usecase.GetPendingRequestById(id)
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AdminHandler) GetListRequest(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, msg := h.usecase.GetListRequest()
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AdminHandler) GetRequestById(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	resp, msg := h.usecase.GetRequestById(id)
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AdminHandler) ApproveRequest(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	msg := h.usecase.ApproveRequest(id, userId.(int))
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) RejectRequest(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	msg := h.usecase.RejectRequest(id, userId.(int))
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) AddRejectNotes(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	var req dto.AddRejectNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg := h.usecase.AddRejectNotes(id, req.Notes)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) DeleteRequest(c *gin.Context) {
	if err := h.checkAdminRole(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	msg := h.usecase.DeleteRequest(id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) checkAdminRole(c *gin.Context) error {
	roleId, exists := c.Get("roleId")
	if !exists || roleId.(int) != 3 {
		return errors.New("forbidden: only admins can perform this action")
	}
	return nil
}
