package helpers

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CreateImageToCloudinary(file *multipart.FileHeader, ctx context.Context) *uploader.UploadResult {
	imageFile, err := file.Open()
	PanicIfError(err, "error at fileHeader.Open in cloudinary.go")
	cloudy, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_APY_KEY"))
	PanicIfError(err, "error at cloudinary.NewFromURL in cloudinary.go create image")
	cloudyRes, err := cloudy.Upload.Upload(ctx, imageFile, uploader.UploadParams{})
	PanicIfError(err, "error at cloudy.Upload.Upload in cloudinary.go")

	return cloudyRes
}
func DeleteImageToCloudinary(imageId string, ctx context.Context) error {
	cloudy, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_APY_KEY"))
	PanicIfError(err, "error at cloudinary.NewFromURL in cloudinary.go delete image")
	_, err = cloudy.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: imageId})

	return err
}
