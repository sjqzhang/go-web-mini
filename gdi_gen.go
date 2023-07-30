package main

import (
	"embed"
	"github.com/sjqzhang/gdi"
	p2 "go-web-mini/config"
	p3 "go-web-mini/controller"
	p4 "go-web-mini/dto"
	p6 "go-web-mini/model"
	p7 "go-web-mini/repository"
	p11 "go-web-mini/vo"
)

//go:embed repository config controller dto response routes util vo common middleware model
var gdiEmbededFiles embed.FS

func init() {
	gdi.SetEmbedFs(&gdiEmbededFiles)
	_ = gdi.GDIPool{}
	gdi.PlaceHolder((*p2.SystemConfig)(nil))
	gdi.PlaceHolder((*p2.LogsConfig)(nil))
	gdi.PlaceHolder((*p2.MysqlConfig)(nil))
	gdi.PlaceHolder((*p2.CasbinConfig)(nil))
	gdi.PlaceHolder((*p2.JwtConfig)(nil))
	gdi.PlaceHolder((*p2.RateLimitConfig)(nil))
	gdi.PlaceHolder((*p3.ApiController)(nil))
	gdi.PlaceHolder((*p3.MenuController)(nil))
	gdi.PlaceHolder((*p3.OperationLogController)(nil))
	gdi.PlaceHolder((*p3.RoleController)(nil))
	gdi.PlaceHolder((*p3.UserController)(nil))
	gdi.PlaceHolder((*p4.ApiTreeDto)(nil))
	gdi.PlaceHolder((*p4.UserInfoDto)(nil))
	gdi.PlaceHolder((*p4.UsersDto)(nil))
	gdi.PlaceHolder((*p6.Api)(nil))
	gdi.PlaceHolder((*p6.RoleCasbin)(nil))
	gdi.PlaceHolder((*p6.Menu)(nil))
	gdi.PlaceHolder((*p6.OperationLog)(nil))
	gdi.PlaceHolder((*p6.Role)(nil))
	gdi.PlaceHolder((*p6.User)(nil))
	gdi.PlaceHolder((*p7.ApiRepository)(nil))
	gdi.PlaceHolder((*p7.MenuRepository)(nil))
	gdi.PlaceHolder((*p7.OperationLogRepository)(nil))
	gdi.PlaceHolder((*p7.RoleRepository)(nil))
	gdi.PlaceHolder((*p7.UserRepository)(nil))
	gdi.PlaceHolder((*p11.ApiListRequest)(nil))
	gdi.PlaceHolder((*p11.CreateApiRequest)(nil))
	gdi.PlaceHolder((*p11.UpdateApiRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteApiRequest)(nil))
	gdi.PlaceHolder((*p11.CreateMenuRequest)(nil))
	gdi.PlaceHolder((*p11.UpdateMenuRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteMenuRequest)(nil))
	gdi.PlaceHolder((*p11.OperationLogListRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteOperationLogRequest)(nil))
	gdi.PlaceHolder((*p11.CreateRoleRequest)(nil))
	gdi.PlaceHolder((*p11.RoleListRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteRoleRequest)(nil))
	gdi.PlaceHolder((*p11.UpdateRoleMenusRequest)(nil))
	gdi.PlaceHolder((*p11.UpdateRoleApisRequest)(nil))
	gdi.PlaceHolder((*p11.RegisterAndLoginRequest)(nil))
	gdi.PlaceHolder((*p11.CreateUserRequest)(nil))
	gdi.PlaceHolder((*p11.UserListRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteUserRequest)(nil))
	gdi.PlaceHolder((*p11.ChangePwdRequest)(nil))
}
