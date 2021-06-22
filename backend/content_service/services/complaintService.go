package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"

	"gorm.io/gorm"
)

type ComplaintService struct {
	complaintRepository repositories.ComplaintRepository
}

func NewComplaintService(db *gorm.DB) (*ComplaintService, error) {
	complaintRepository, err := repositories.NewComplaintRepo(db)
	if err != nil {
		return nil, err
	}

	return &ComplaintService{
		complaintRepository,
	}, err
}

func (service *ComplaintService) CreateContentComplaint(ctx context.Context, contentComplaint domain.ContentComplaint) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateContentComplaint")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.complaintRepository.CreateContentComplaint(ctx, contentComplaint)
}

func (service *ComplaintService) GetAllContentComplaints(ctx context.Context) ([]domain.ContentComplaint, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateContentComplaint")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.complaintRepository.GetAllContentComplaints(ctx)
}
