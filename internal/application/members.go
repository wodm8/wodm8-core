package application

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/wodm8/wodm8-core/internal/domain"
	"github.com/wodm8/wodm8-core/internal/members"
	"github.com/wodm8/wodm8-core/internal/platform/aws"
	"path/filepath"
	"sync"
)

var bucketName = "wodm8-dev"

var folderBaseName = "members"

type MemberService struct {
	memberRepository members.MemberRepository
}

func NewMemberService(memberRepository members.MemberRepository) MemberService {
	return MemberService{memberRepository: memberRepository}
}

func (m MemberService) CreateMember(memberReq domain.CreateUserRequest) error {
	memberId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	member, err := members.NewMember(memberId.String(), memberReq.FirstName, memberReq.LastName, memberReq.Email, "", "", "", "", "", 0)
	if err != nil {
		fmt.Printf("error creating member: %v\n", err.Error())
		return err
	}

	if err := m.memberRepository.Save(member); err != nil {
		fmt.Printf("error saving member: %v\n", err.Error())
		return err
	}
	return nil
}

func (m MemberService) GetMember(email string) (members.Members, error) {
	member, err := m.memberRepository.FindByEmail(email)
	if err != nil {
		fmt.Printf("member missing: %v\n", err.Error())
		return member, err
	}
	return member, nil
}

func (m MemberService) UpdateMember(memberReq domain.MembersRequest) error {

	awsSrv := aws.NewAwsServiceImpl(context.Background())
	cfg, err := awsSrv.Load()
	if err != nil {
		return err
	}
	s3Client := awsSrv.S3(cfg)

	var profilePictureUrl string
	var frontPagePictureUrl string

	wg := sync.WaitGroup{}

	wg.Add(2)

	if memberReq.ProfilePicture != "" {
		go func() error {
			imageData, err := decodeBase64Image(memberReq.ProfilePicture)
			if err != nil {
				fmt.Printf("error decoding image data: %v\n", err)
				return err
			}

			key := filepath.Join(folderBaseName, memberReq.ID, fmt.Sprintf("%s.jpg", memberReq.ID))

			objUrl, err := awsSrv.SaveObjectInBucket(s3Client, bucketName, key, imageData)
			if err != nil {
				return err
			}
			profilePictureUrl = objUrl
			wg.Done()
			return nil
		}()
	}

	if memberReq.FrontPageImg != "" {

		go func() error {
			imageData, err := decodeBase64Image(memberReq.FrontPageImg)
			if err != nil {
				fmt.Printf("error decoding image data: %v\n", err)
				return err
			}

			key := filepath.Join(folderBaseName, memberReq.ID, "front_page.jpg")

			objUrl, err := awsSrv.SaveObjectInBucket(s3Client, bucketName, key, imageData)
			if err != nil {
				return err
			}
			fmt.Printf("object frontimg: %v\n", objUrl)
			frontPagePictureUrl = objUrl
			wg.Done()
			return nil
		}()
	}

	wg.Wait()

	member, err := members.NewMember(memberReq.ID, memberReq.FirstName, memberReq.LastName, memberReq.Email, memberReq.Gender, memberReq.City, memberReq.Country, profilePictureUrl, frontPagePictureUrl, memberReq.Age)
	fmt.Printf("update member: %v\n", member)
	if err != nil {
		fmt.Printf("error creating member: %v\n", err.Error())
		return err
	}

	if err := m.memberRepository.Update(member); err != nil {
		fmt.Printf("error saving member: %v\n", err.Error())
		return err
	}

	return nil
}

func decodeBase64Image(base64Image string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 image: %v", err)
	}
	return data, nil
}
