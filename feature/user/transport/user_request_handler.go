package transport

import (
	"net/http"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	RequestUsecase usecase.ApplicantRequestUsecaseInterface
}

func NewApplicantRequestHandler(requestUsecase usecase.ApplicantRequestUsecaseInterface) *RequestHandler {
	return &RequestHandler{RequestUsecase: requestUsecase}
}

func (h *RequestHandler) CreateApplicantRequest(c *gin.Context) {
	var request dto.ApplicantRequestCreatingDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.RequestUsecase.CreateApplicantRequest(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Request created successfully"})
}
