package dtos

import (
	"time"

	"github.com/yaghoubi-mn/voter/internal/models"
)

type SubInput struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
}

func (s SubInput) GetSubModel(ownerID uint64) models.Sub {
	return models.Sub{
		Title:       s.Title,
		Description: s.Description,
		OwnerID:     ownerID,
	}
}

func (s SubInput) UpdateSub(sub *models.Sub) {
	sub.Title = s.Title
	sub.Description = s.Description
}

type SubOutput struct {
	ID               uint64
	Title            string
	Description      string
	CreatedAt        time.Time
	Views            uint64
	SubscribersCount uint64
}

func GetSubOutputFromSub(sub models.Sub) SubOutput {
	return SubOutput{
		ID:               sub.ID,
		Title:            sub.Title,
		Description:      sub.Description,
		CreatedAt:        sub.CreatedAt,
		Views:            sub.Views,
		SubscribersCount: sub.SubscribersCount,
	}
}
