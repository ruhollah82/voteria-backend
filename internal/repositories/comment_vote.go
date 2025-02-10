package repositories

import (
	"github.com/yaghoubi-mn/voter/internal/custom_errors"
	"github.com/yaghoubi-mn/voter/internal/models"
	"gorm.io/gorm"
)

type CommentVoteRepository interface {
	Create(commentVoteote models.CommentVote) error
	Delete(commentId uint64, userId uint64) (models.CommentVote, error)
}

type commentVoteRepository struct {
	db *gorm.DB
}

func NewCommentVoteRepository(db *gorm.DB) CommentVoteRepository {
	return &commentVoteRepository{
		db: db,
	}
}

func (r *commentVoteRepository) Create(commentVote models.CommentVote) error {

	err := r.db.Create(&commentVote).Error

	return err
}

func (r *commentVoteRepository) Delete(commentId uint64, userId uint64) (models.CommentVote, error) {

	var commentVote models.CommentVote
	err := r.db.Where("comment_id=? AND user_id=?", commentId, userId).First(&commentVote).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return commentVote, err
	}

	if err := r.db.Model(&models.CommentVote{}).Where("comment_id = ? AND user_id=?", commentId, userId).Delete(&models.CommentVote{}).Error; err != nil {
		return commentVote, err
	}

	if err == gorm.ErrRecordNotFound {
		err = custom_errors.RecordNotFound
	}

	return commentVote, err
}
