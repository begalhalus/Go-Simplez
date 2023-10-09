package handler

import (
	"go-simple/entity"
	"go-simple/helper"
	"go-simple/roles"
	"go-simple/utils/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RolesHandler struct {
	rolesUsecase roles.RolesUsecase
}

func CreateRolesHandler(r *gin.Engine, rolesUsecase roles.RolesUsecase) {
	rolesHandler := RolesHandler{rolesUsecase}

	v1 := r.Group("/role/v1", middleware.AuthMiddleware)
	v1.POST("/create", rolesHandler.create)
	v1.GET("/detail/:id", rolesHandler.detail)
	v1.GET("/list", rolesHandler.list)
	v1.PUT("/update/:id", rolesHandler.update)
	v1.DELETE("/delete/:id", rolesHandler.delete)
}

// use credential
// credential := c.MustGet("credential").(jwt.MapClaims)
// fmt.Println("credential => user_id => ", credential["user_id"])

func (e *RolesHandler) create(c *gin.Context) {

	var roleReq entity.RolesRequest

	if err := c.ShouldBind(&roleReq); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// validate
	role := entity.Roles{
		RoleNm:   &roleReq.Name,
		RoleDesc: &roleReq.Description,
	}

	newRole, err := e.rolesUsecase.Create(&role)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, newRole, "success")

}

func (e *RolesHandler) list(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")
	sort := c.DefaultQuery("sort", "asc")
	sortField := c.DefaultQuery("sortField", "role_id")

	roles, totalData, err := e.rolesUsecase.List(page, limit, search, sort, sortField)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(*roles) == 0 {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.DataNotFound("Role"))
		return
	}

	totalPage := float64((totalData + int64(limit) - 1) / int64(limit))

	response := entity.Paginate{
		List:      roles,
		Limit:     int64(limit),
		Page:      int64(page),
		TotalData: totalData,
		TotalPage: totalPage,
	}

	helper.SuccessResponse(c, response, "success")

}

func (e *RolesHandler) detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.NotValid(idStr))
		return
	}
	roles, err := e.rolesUsecase.Detail(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.SuccessResponse(c, roles, "success")
}

func (e *RolesHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.NotValid(idStr))
		return
	}

	existingRoles, err := e.rolesUsecase.Detail(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// validate
	var roleReq entity.RolesRequest
	if err := c.ShouldBind(&roleReq); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if roleReq.Name != "" {
		existingRoles.RoleNm = &roleReq.Name
	}
	if roleReq.Description != "" {
		existingRoles.RoleDesc = &roleReq.Description
	}

	updateData := map[string]interface{}{
		"role_nm":   roleReq.Name,
		"role_desc": roleReq.Description,
	}

	_, err = e.rolesUsecase.Update(id, updateData)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, make(map[string]interface{}), "success")
}

func (e *RolesHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, helper.NotValid(idStr))
		return
	}
	err = e.rolesUsecase.Delete(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.SuccessResponse(c, make(map[string]interface{}), "success")
}
