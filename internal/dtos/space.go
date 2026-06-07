package dtos

import (
	"time"

	"github.com/yaghoubi-mn/voter/internal/models"
)

type SpaceInput struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
}

func (s SpaceInput) GetSubModel(ownerID uint64) models.Space {
	return models.Space{
		Title:       s.Title,
		Description: s.Description,
		OwnerID:     ownerID,
	}
}

func (s SpaceInput) UpdateSub(sub *models.Space) {
	sub.Title = s.Title
	sub.Description = s.Description
}

type SpaceOutput struct {
	ID               uint64
	Title            string
	Description      string
	CreatedAt        time.Time
	Views            uint64
	SubscribersCount uint64
}

func GetSubOutputFromSub(sub models.Space) SpaceOutput {
	return SpaceOutput{
		ID:               sub.ID,
		Title:            sub.Title,
		Description:      sub.Description,
		CreatedAt:        sub.CreatedAt,
		Views:            sub.Views,
		SubscribersCount: sub.SubscribersCount,
	}
}
