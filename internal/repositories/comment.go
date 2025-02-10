package repositories

import (
	"fmt"

	"github.com/yaghoubi-mn/voter/internal/config"
	"github.com/yaghoubi-mn/voter/internal/custom_errors"
	"github.com/yaghoubi-mn/voter/internal/enums"
	"github.com/yaghoubi-mn/voter/internal/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment models.Comment) error
	Update(comment models.Comment) error
	Delete(commentId uint64) error
	GetByID(commentId uint64) (models.Comment, error)
	GetAll(postId uint64, sortBy enums.SortBy, page int) ([]models.Comment, error)
	GetAllReplies(commentId uint64) ([]models.Comment, error)
	AddCommentScore(commentId uint64, number int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment models.Comment) error {
	return r.db.Create(&comment).Error
}

func (r *commentRepository) Update(comment models.Comment) error {
	return r.db.Updates(&comment).Error

}

func (r *commentRepository) Delete(commentId uint64) error {
	return r.db.Delete(models.Comment{ID: commentId}).Error
}

func (r *commentRepository) GetByID(commentId uint64) (models.Comment, error) {
	var comment models.Comment
	if err := r.db.Preload("Author").First(&comment, &models.Comment{ID: commentId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return comment, custom_errors.RecordNotFound
		}

		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) GetAll(postId uint64, sortBy enums.SortBy, page int) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Preload("Author").Where("post_id=?", postId).Order(sortBy).Offset((page - 1) * config.PageLimit).Limit(config.PageLimit).Find(&comments).Error

	return comments, err
}

func (r *commentRepository) GetAllReplies(commentId uint64) ([]models.Comment, error) {

	var comments []models.Comment
	if err := r.db.Where("comment_id=?", commentId).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) AddCommentScore(commentId uint64, number int) error {
	var expr string
	if number >= 0 {
		expr = fmt.Sprintf("score + %v", number)
	} else {
		expr = fmt.Sprintf("score %v", number)
	}

	if err := r.db.Model(&models.Comment{}).Where("id=?", commentId).Update("score", gorm.Expr(expr)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return custom_errors.RecordNotFound
		}

		return err
	}

	return nil
}
