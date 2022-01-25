package pictures

type Repository interface {
	IsExists(pictureId string) error
	Search(dto SearchPictureDto) (SearchPictureResultDto, error)
	SaveInitialPicture(image *InitialImage) (int, error)
	SaveInitialPictures(image InitialDropboxImage) (*InitialDropboxImageResult, error)
	UpdateImageHandle(picture *Picture) error
	Store(imageTextDetectionDto *TextDetectionOnImageDto) error
	StoreAll(pictures []*TextDetectionOnImageDto) error
	Delete(imageId string) error
}
