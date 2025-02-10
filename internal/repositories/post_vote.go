package repositories

import (
	"github.com/yaghoubi-mn/voter/internal/custom_errors"
	"github.com/yaghoubi-mn/voter/internal/models"
	"gorm.io/gorm"
)

type PostVoteRepository interface {
	Create(postVote models.PostVote) error
	Delete(postId uint64, userId uint64) (models.PostVote, error)
}

type postVoteRepository struct {
	db *gorm.DB
}

func NewPostVoteRepository(db *gorm.DB) PostVoteRepository {
	return &postVoteRepository{
		db: db,
	}
}

func (r *postVoteRepository) Create(postVote models.PostVote) error {

	err := r.db.Create(&postVote).Error

	return err
}

func (r *postVoteRepository) Delete(postId uint64, userId uint64) (models.PostVote, error) {

	var postVote models.PostVote
	err := r.db.Where("post_id=? AND user_id=?", postId, userId).First(&postVote).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return postVote, err
	}

	if err := r.db.Model(&models.PostVote{}).Where("post_id = ? AND user_id=?", postId, userId).Delete(&models.PostVote{}).Error; err != nil {
		return postVote, err
	}

	if err == gorm.ErrRecordNotFound {
		err = custom_errors.RecordNotFound
	}

	return postVote, err
}
