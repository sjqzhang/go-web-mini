package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-mini/common"
	"go-web-mini/model"
	"go-web-mini/repository"
	"go-web-mini/response"
	"go-web-mini/vo"
	"strconv"
)

type IMenuController interface {
	GetMenus(c *gin.Context)             // 获取菜单列表
	GetMenuTree(c *gin.Context)          // 获取菜单树
	CreateMenu(c *gin.Context)           // 创建菜单
	UpdateMenuById(c *gin.Context)       // 更新菜单
	BatchDeleteMenuByIds(c *gin.Context) // 批量删除菜单

	GetUserMenusByUserId(c *gin.Context)    // 获取用户的可访问菜单列表
	GetUserMenuTreeByUserId(c *gin.Context) // 获取用户的可访问菜单树
}

//@middleware auth
//@router /api
type MenuController struct {
	MenuRepository repository.IMenuRepository
	UserRepository repository.IUserRepository
}

//func NewMenuController() IMenuController {
//	menuRepository := repository.NewMenuRepository()
//	menuController := MenuController{MenuRepository: menuRepository}
//	return menuController
//}
//@tags menu
// 获取菜单列表
//@router /menu/list [get]
func (mc MenuController) GetMenus(c *gin.Context) {
	menus, err := mc.MenuRepository.GetMenus(nil)
	if err != nil {
		response.Fail(c, nil, "获取菜单列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"menus": menus}, "获取菜单列表成功")
}

//@tags menu
// 获取菜单树
//@router /menu/tree [get]
func (mc MenuController) GetMenuTree(c *gin.Context) {
	menuTree, err := mc.MenuRepository.GetMenuTree(nil)
	if err != nil {
		response.Fail(c, nil, "获取菜单树失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"menuTree": menuTree}, "获取菜单树成功")
}

//@tags menu
// 创建菜单
//@router /menu/create [post]
func (mc MenuController) CreateMenu(c *gin.Context) {
	var req vo.CreateMenuRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取当前用户

	ctxUser, err := mc.UserRepository.GetCurrentUser(nil, c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	menu := model.Menu{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       &req.Icon,
		Path:       req.Path,
		Redirect:   &req.Redirect,
		Component:  req.Component,
		Sort:       req.Sort,
		Status:     req.Status,
		Hidden:     req.Hidden,
		NoCache:    req.NoCache,
		AlwaysShow: req.AlwaysShow,
		Breadcrumb: req.Breadcrumb,
		ActiveMenu: &req.ActiveMenu,
		ParentId:   &req.ParentId,
		Creator:    ctxUser.Username,
	}

	err = mc.MenuRepository.CreateMenu(nil, &menu)
	if err != nil {
		response.Fail(c, nil, "创建菜单失败: "+err.Error())
		return
	}
	response.Success(c, nil, "创建菜单成功")
}

//@tags menu
// 更新菜单
//@router /menu/update/:menuId [patch]
func (mc MenuController) UpdateMenuById(c *gin.Context) {
	var req vo.UpdateMenuRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取路径中的menuId
	menuId, _ := strconv.Atoi(c.Param("menuId"))
	if menuId <= 0 {
		response.Fail(c, nil, "菜单ID不正确")
		return
	}

	// 获取当前用户
	ur := mc.UserRepository //repository.NewUserRepository()

	fmt.Println(ur == nil)
	ctxUser, err := ur.GetCurrentUser(nil, c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	menu := model.Menu{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       &req.Icon,
		Path:       req.Path,
		Redirect:   &req.Redirect,
		Component:  req.Component,
		Sort:       req.Sort,
		Status:     req.Status,
		Hidden:     req.Hidden,
		NoCache:    req.NoCache,
		AlwaysShow: req.AlwaysShow,
		Breadcrumb: req.Breadcrumb,
		ActiveMenu: &req.ActiveMenu,
		ParentId:   &req.ParentId,
		Creator:    ctxUser.Username,
	}

	err = mc.MenuRepository.UpdateMenuById(nil, uint(menuId), &menu)
	if err != nil {
		response.Fail(c, nil, "更新菜单失败: "+err.Error())
		return
	}

	response.Success(c, nil, "更新菜单成功")

}

//@tags menu
// 批量删除菜单
//@router /menu/delete/batch [delete]
func (mc MenuController) BatchDeleteMenuByIds(c *gin.Context) {
	var req vo.DeleteMenuRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	err := mc.MenuRepository.BatchDeleteMenuByIds(nil, req.MenuIds)
	if err != nil {
		response.Fail(c, nil, "删除菜单失败: "+err.Error())
		return
	}

	response.Success(c, nil, "删除菜单成功")
}

//@tags menu
// 根据用户ID获取用户的可访问菜单列表
//@router /access/list/user/:userId [get]
func (mc MenuController) GetUserMenusByUserId(c *gin.Context) {
	// 获取路径中的userId
	userId, _ := strconv.Atoi(c.Param("userId"))
	if userId <= 0 {
		response.Fail(c, nil, "用户ID不正确")
		return
	}

	menus, err := mc.MenuRepository.GetUserMenusByUserId(nil, uint(userId))
	if err != nil {
		response.Fail(c, nil, "获取用户的可访问菜单列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"menus": menus}, "获取用户的可访问菜单列表成功")
}

//@tags menu
// 根据用户ID获取用户的可访问菜单树
//@router /menu/access/tree/:userId [get]
func (mc MenuController) GetUserMenuTreeByUserId(c *gin.Context) {
	// 获取路径中的userId
	userId, _ := strconv.Atoi(c.Param("userId"))
	if userId <= 0 {
		response.Fail(c, nil, "用户ID不正确")
		return
	}

	menuTree, err := mc.MenuRepository.GetUserMenuTreeByUserId(nil, uint(userId))
	if err != nil {
		response.Fail(c, nil, "获取用户的可访问菜单树失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"menuTree": menuTree}, "获取用户的可访问菜单树成功")
}
