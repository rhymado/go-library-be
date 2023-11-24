package handlers

import (
	"fmt"
	"lib/internal/models"
	"lib/internal/repositories"
	"lib/pkg"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type AuthHandler struct {
	*repositories.AuthRepository
}

func InitAuthHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar}
}

func (a *AuthHandler) RegisterHandler(ctx *gin.Context) {
	body := &models.UserModel{}
	if err := ctx.ShouldBind(body); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	h := pkg.HashConfig{
		Time:    3,
		Memory:  64 * 1024,
		Threads: 2,
		KeyLen:  32,
		SaltLen: 16,
	}
	hashedPassword, err := h.GenHashedPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := a.Register(body, hashedPassword); err != nil {
		pgErr, _ := err.(*pq.Error)
		if pgErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "username atau email sudah terdaftar",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User berhasil register",
		"data": gin.H{
			"username": body.Username,
			"email":    body.Email,
		},
	})
}

func (a *AuthHandler) LoginHandler(ctx *gin.Context) {
	body := &models.GetUserInfoModel{}
	if err := ctx.ShouldBind(body); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := a.GetPassword(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Email or Password is wrong",
		})
		return
	}

	hc := pkg.HashConfig{}
	isValid, err := hc.ComparePasswordAndHash(body.Password, result[0].Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email or Password is wrong",
		})
		return
	}

	// generate JWT
	payload := pkg.NewPayload(result[0].Id, result[0].Role)
	token, err := payload.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("Selamat Datang %s", result[0].Username)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data": gin.H{
			"token": token,
			"userInfo": gin.H{
				"email":    body.Email,
				"username": result[0].Username,
			},
		},
	})
}
