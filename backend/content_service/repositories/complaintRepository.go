package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComplaintRepository interface {
	CreateContentComplaint(context.Context, domain.ContentComplaint) error
	GetAllContentComplaints(context.Context) ([]domain.ContentComplaint, error)
	DeleteByPostId (context.Context, string) error
	RejectById (context.Context, string) error
	DeleteComplaintByUserId( context.Context,  string) error
}

type complaintRepository struct {
	DB *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) (*complaintRepository, error) {
	if db == nil {
		panic("ComplaintRepository not created, gorm.DB is nil")
	}

	return &complaintRepository{
		DB: db,
	}, nil
}

func (repo *complaintRepository) DeleteComplaintByUserId(ctx context.Context,id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteComplaintByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.DB.Where("user_id = ?", id).Delete(&persistence.ContentComplaint{})
	if result.Error != nil {
		return errors.New("Could not delete all user complaints!")
	}else if result.RowsAffected == 0 {
		return errors.New("Could not delete all user complaints!")

	}
	return nil

}


func (repository *complaintRepository) CreateContentComplaint(ctx context.Context, contentComplaint domain.ContentComplaint) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateContentComplaint")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var contentComplaintPersistence = persistence.ContentComplaint{
		Id:       uuid.New().String(),
		Category: contentComplaint.Category,
		PostId:   contentComplaint.PostId,
		Status:   "Pending",
		IsPost:   contentComplaint.IsPost,
		UserId:   contentComplaint.UserId,
	}

	result := repository.DB.Create(&contentComplaintPersistence)
	return result.Error
}

func (repository *complaintRepository) GetAllContentComplaints(ctx context.Context) ([]domain.ContentComplaint, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllContentComplaints")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var complaintsPersistence []persistence.ContentComplaint
	var complaintsDomain []domain.ContentComplaint

	result := repository.DB.Find(&complaintsPersistence)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, complaint := range complaintsPersistence {
		complaintsDomain = append(complaintsDomain, domain.ContentComplaint{
			Id:       complaint.Id,
			Category: complaint.Category,
			PostId:   complaint.PostId,
			Status:   complaint.Status,
			IsPost:   complaint.IsPost,
			UserId:   complaint.UserId,
		})
	}

	return complaintsDomain, nil

}

func (repository *complaintRepository) DeleteByPostId (ctx context.Context, postId string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteByPostId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Where("post_id = ?", postId).Delete(&persistence.ContentComplaint{})
	if result.Error != nil {
		return result.Error
	}else if result.RowsAffected == 0 {
		return errors.New("Did not delete complaint")
	}

	return nil
}

func (repository *complaintRepository) RejectById (ctx context.Context,id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RejectById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Where("id = ?", id).Updates(&persistence.ContentComplaint{Status: "Rejected"})

	if result.Error != nil{
		return result.Error
	}else if result.RowsAffected == 0 {
		return errors.New("Could not update complaint!")
	}
	return nil
}