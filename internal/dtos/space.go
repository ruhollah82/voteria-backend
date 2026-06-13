package dtos

import (
	"time"

	"github.com/yaghoubi-mn/voter/internal/models"
)

type SpaceCreateInput struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Username    string `validate:"required,username" json:"username"`
}

type SpaceEditInput struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
}

func (s SpaceCreateInput) GetSubModel(ownerID uint64) models.Space {
	return models.Space{
		Title:       s.Title,
		Description: s.Description,
		OwnerID:     ownerID,
		Username:    s.Username,
	}
}

func (s SpaceEditInput) UpdateSub(space *models.Space) {
	space.Title = s.Title
	space.Description = s.Description
}

type SpaceOutput struct {
	ID               uint64
	Username         string
	Title            string
	Description      string
	CreatedAt        time.Time
	Views            uint64
	SubscribersCount uint64
}

func GetSubOutputFromSub(space models.Space) SpaceOutput {
	return SpaceOutput{
		ID:               space.ID,
		Username:         space.Username,
		Title:            space.Title,
		Description:      space.Description,
		CreatedAt:        space.CreatedAt,
		Views:            space.Views,
		SubscribersCount: space.SubscribersCount,
	}
}
