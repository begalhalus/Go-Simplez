package handler

import (
	"go-simple/entity"
	"go-simple/helper"
	"go-simple/users"
	apps_config "go-simple/utils/config"
	"go-simple/utils/middleware"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

type UsersHandler struct {
	usersUsecase users.UsersUsecase
}

func CreateUsersHandler(r *gin.Engine, usersUsecase users.UsersUsecase) {
	usersHandler := UsersHandler{usersUsecase}

	v1 := r.Group("/user/v1")
	v1.POST("/login", usersHandler.login)
	v1.POST("/create", middleware.AuthMiddleware, usersHandler.create)
	v1.GET("/detail/:id", middleware.AuthMiddleware, usersHandler.detail)
	v1.GET("/list", middleware.AuthMiddleware, usersHandler.list)
	v1.PUT("/update/:id", middleware.AuthMiddleware, usersHandler.update)
	v1.DELETE("/delete/:id", middleware.AuthMiddleware, usersHandler.delete)
}

// use credential
// credential := c.MustGet("credential").(jwt.MapClaims)
// fmt.Println("credential => user_id => ", credential["user_id"])

func (e *UsersHandler) login(c *gin.Context) {
	var userReq entity.LoginRequest
	if err := c.ShouldBind(&userReq); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	existingUsers, err := e.usersUsecase.Login(userReq.Email, userReq.Password, nil)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims := jwt.MapClaims{
		"name":  *existingUsers.Name,
		"email": *existingUsers.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := helper.GenerateToken(&claims)
	if errToken != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "failed generate token !")
		return

	}

	updateData := map[string]interface{}{
		"token": token,
	}
	_, err = e.usersUsecase.Update(*existingUsers.UserID, updateData)

	*existingUsers.Token = token

	helper.SuccessResponse(c, existingUsers, "success")
}

func (e *UsersHandler) create(c *gin.Context) {
	var userReq entity.UsersRequest
	var user entity.Users

	if err := c.ShouldBind(&userReq); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if userReq.Password == "" {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.Required("password"))
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// validate
	user.RoleID = &userReq.RoleID
	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Password = hashedPassword

	fileHeader, err := c.FormFile("file")

	if err != nil && err != http.ErrMissingFile {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if fileHeader != nil {
		// Checked validation
		fileType := []string{"image/jpeg", "image/jpg", "image/png"}
		isValid := helper.ValidateFile(fileHeader, fileType)

		if !isValid {
			helper.ErrorResponse(c, http.StatusBadRequest, "File not allowed !")
			return
		}

		isUpload := helper.UploadFile(c, fileHeader, userReq.Name)

		if !isUpload {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		user.File = c.Request.Host + apps_config.STATIC_ROUTE + "/" + slug.Make(userReq.Name) + "." + helper.Extension(fileHeader.Filename)
	}

	newUser, err := e.usersUsecase.Create(&user)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, newUser, "success")
}

func (e *UsersHandler) list(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")
	sort := c.DefaultQuery("sort", "asc")
	sortField := c.DefaultQuery("sortField", "user_id")

	users, totalData, err := e.usersUsecase.List(page, limit, search, sort, sortField)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(*users) == 0 {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.DataNotFound("Users"))
		return
	}

	totalPage := float64((totalData + int64(limit) - 1) / int64(limit))

	response := entity.Paginate{
		List:      users,
		Limit:     int64(limit),
		Page:      int64(page),
		TotalData: totalData,
		TotalPage: totalPage,
	}

	helper.SuccessResponse(c, response, "success")

}

func (e *UsersHandler) detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.NotValid(idStr))
		return
	}
	users, err := e.usersUsecase.Detail(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.SuccessResponse(c, users, "success")
}

func (e *UsersHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.NotValid(idStr))
		return
	}

	existingUsers, err := e.usersUsecase.Detail(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// validate
	var userReq entity.UsersRequest
	if err := c.ShouldBind(&userReq); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if userReq.Name != "" {
		existingUsers.Name = &userReq.Name
	}
	if userReq.Email != "" {
		existingUsers.Email = &userReq.Email
	}

	// if password not nil
	if userReq.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
		if err != nil {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		existingUsers.Password = hashedPassword
	}

	fileHeader, err := c.FormFile("file")

	if err != nil && err != http.ErrMissingFile {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if fileHeader != nil {
		// Checked validation
		fileType := []string{"image/jpeg", "image/jpg", "image/png"}
		isValid := helper.ValidateFile(fileHeader, fileType)

		if !isValid {
			helper.ErrorResponse(c, http.StatusBadRequest, "File not allowed !")
			return
		}

		isUpload := helper.UploadFile(c, fileHeader, userReq.Name)

		if !isUpload {
			helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		existingUsers.File = c.Request.Host + apps_config.STATIC_ROUTE + "/" + slug.Make(userReq.Name) + "." + helper.Extension(fileHeader.Filename)
	}

	updateData := map[string]interface{}{
		"name":     userReq.Name,
		"email":    userReq.Email,
		"password": existingUsers.Password,
		"file":     existingUsers.File,
	}

	_, err = e.usersUsecase.Update(id, updateData)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, make(map[string]interface{}), "success")
}

func (e *UsersHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.NotValid(idStr))
		return
	}
	err = e.usersUsecase.Delete(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.SuccessResponse(c, make(map[string]interface{}), "success")
}
