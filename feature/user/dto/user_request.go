package dto

type ApplicantRequestCreatingDTO struct {
	UserID uint   `json:"user_id" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Status int    `json:"status" binding:"required"`
}
