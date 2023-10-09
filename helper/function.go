package helper

import (
	"fmt"
	apps_config "go-simple/utils/config"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gosimple/slug"
)

func Extension(message string) string {
	parts := strings.Split(message, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

func ValidateFile(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")

	res := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			res = true
			break
		}
	}

	return res
}

func UploadFile(ctx *gin.Context, fileHeader *multipart.FileHeader, name string) bool {
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf(apps_config.STATIC_DIR+"/%s", slug.Make(name)+"."+Extension(fileHeader.Filename)))

	if errUpload != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error, can't upload file.",
		})

		return false
	} else {
		return true
	}

}

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(os.Getenv("JWT_USER_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return webtoken, nil

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenJwt, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)

		if !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_USER_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenJwt, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isValid := token.Claims.(jwt.MapClaims)
	if !isValid || !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	return claims, nil
}
