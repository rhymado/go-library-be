package handlers

import (
	"fmt"
	"lib/internal/helpers"
	"net/http"
	"path"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadHandler struct{}

func InitUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (u *UploadHandler) UploadToCloud(ctx *gin.Context) {
	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fieldName := "image"
	formFile, err := ctx.FormFile(fieldName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer file.Close()

	// id, err := uuid.NewRandom()
	id := "1"
	publicId := fmt.Sprintf("%s_%s-%s", "golang", fieldName, id)
	folder := ""
	res, err := cld.Uploader(ctx, file, publicId, folder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload sukses",
		"data": gin.H{
			"url": res.SecureURL,
		},
	})
}

func (u *UploadHandler) SavedToDirectory(ctx *gin.Context) {
	fieldName := "image"
	formFile, err := ctx.FormFile(fieldName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ext := path.Ext(formFile.Filename)
	folder := "public/docs"
	isImage, err := regexp.Match("png|jpg|jpeg", []byte(ext))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if isImage {
		folder = "public/images"
	}

	uid := uuid.NewString()
	fileName := fmt.Sprintf("%s-%s%s", fieldName, uid, ext)
	dst := path.Join(folder, fileName)
	if err := ctx.SaveUploadedFile(formFile, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload sukses",
		"data": gin.H{
			"url": dst,
		},
	})
}
