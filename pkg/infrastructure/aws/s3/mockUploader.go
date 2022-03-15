package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/col3name/images-search/pkg/domain"
	"io"
)

type MockUploader struct {
}

func NewMockUploader() *MockUploader {
	s := new(MockUploader)
	return s
}

func (s *MockUploader) Upload(_ string, _ io.Reader, _ types.ObjectCannedACL) (*domain.UploadOutput, error) {

	var id = "uploadOutput.VersionID"
	return &domain.UploadOutput{
		Location:  " uploadOutput.Location",
		VersionID: &id,
		UploadID:  "uploadOutput.UploadID",
	}, nil
}