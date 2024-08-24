package application

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wodm8/wodm8-core/internal/box"
	"github.com/wodm8/wodm8-core/internal/domain"
	"github.com/wodm8/wodm8-core/internal/members"
)

// var bucketName = "wodm8-dev"
// var folderBaseName = "boxes"

type BoxService struct {
	boxRepository     box.BoxRepository
	membersRepository members.MemberRepository
}

func NewBoxService(boxRepository box.BoxRepository, membersRepository members.MemberRepository) BoxService {
	return BoxService{
		boxRepository:     boxRepository,
		membersRepository: membersRepository,
	}
}

func (b BoxService) CreateBox(ctx *gin.Context, boxReq domain.CreateBoxRequest, userEmail string) error {
	boxId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	member, err := b.membersRepository.FindByEmail(userEmail)
	if err != nil {
		return err
	}

	box_, err := box.NewBox(boxId.String(), member.ID, boxReq.Name, boxReq.Address, boxReq.Phone, boxReq.Logo)
	if err != nil {
		return err
	}

	if err := b.boxRepository.Save(box_); err != nil {
		return err
	}

	return nil
}
