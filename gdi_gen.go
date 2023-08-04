package main

import (
	"embed"
	"github.com/sjqzhang/gdi"
	p2 "go-web-mini/config"
	p3 "go-web-mini/controller"
	p4 "go-web-mini/dto"
	p5 "go-web-mini/generator"
	p7 "go-web-mini/model"
	p8 "go-web-mini/repository"
	p11 "go-web-mini/service"
	p13 "go-web-mini/vo"
)

//go:embed common dto generator middleware routes service util config controller model repository response vo
var gdiEmbedFiles embed.FS

func init() {
	gdi.SetEmbedFs(&gdiEmbedFiles)
	_ = gdi.GDIPool{}
	gdi.PlaceHolder((*p2.SystemConfig)(nil))
	gdi.PlaceHolder((*p2.LogsConfig)(nil))
	gdi.PlaceHolder((*p2.MysqlConfig)(nil))
	gdi.PlaceHolder((*p2.CasbinConfig)(nil))
	gdi.PlaceHolder((*p2.JwtConfig)(nil))
	gdi.PlaceHolder((*p2.RateLimitConfig)(nil))
	gdi.PlaceHolder((*p3.ApiController)(nil))
	gdi.PlaceHolder((*p3.MenuController)(nil))
	gdi.PlaceHolder((*p3.NewsController)(nil))
	gdi.PlaceHolder((*p3.OperationLogController)(nil))
	gdi.PlaceHolder((*p3.RoleController)(nil))
	gdi.PlaceHolder((*p3.UserController)(nil))
	gdi.PlaceHolder((*p4.ApiTreeDto)(nil))
	gdi.PlaceHolder((*p4.UserInfoDto)(nil))
	gdi.PlaceHolder((*p4.UsersDto)(nil))
	gdi.PlaceHolder((*p5.FieldResult)(nil))
	gdi.PlaceHolder((*p5.TableResult)(nil))
	gdi.PlaceHolder((*p5.CommonObject)(nil))
	gdi.PlaceHolder((*p7.Api)(nil))
	gdi.PlaceHolder((*p7.RoleCasbin)(nil))
	gdi.PlaceHolder((*p7.Model)(nil))
	gdi.PlaceHolder((*p7.Menu)(nil))
	gdi.PlaceHolder((*p7.News)(nil))
	gdi.PlaceHolder((*p7.NewsQuery)(nil))
	gdi.PlaceHolder((*p7.News2)(nil))
	gdi.PlaceHolder((*p7.OperationLog)(nil))
	gdi.PlaceHolder((*p7.Role)(nil))
	gdi.PlaceHolder((*p7.User)(nil))
	gdi.PlaceHolder((*p8.ApiRepository)(nil))
	gdi.PlaceHolder((*p8.MenuRepository)(nil))
	gdi.PlaceHolder((*p8.News2Repository)(nil))
	gdi.PlaceHolder((*p8.NewsRepository)(nil))
	gdi.PlaceHolder((*p8.OperationLogRepository)(nil))
	gdi.PlaceHolder((*p8.RoleRepository)(nil))
	gdi.PlaceHolder((*p8.UserRepository)(nil))
	gdi.PlaceHolder((*p11.News2Service)(nil))
	gdi.PlaceHolder((*p11.NewsService)(nil))
	gdi.PlaceHolder((*p13.ApiListRequest)(nil))
	gdi.PlaceHolder((*p13.CreateApiRequest)(nil))
	gdi.PlaceHolder((*p13.UpdateApiRequest)(nil))
	gdi.PlaceHolder((*p13.DeleteApiRequest)(nil))
	gdi.PlaceHolder((*p13.CreateMenuRequest)(nil))
	gdi.PlaceHolder((*p13.UpdateMenuRequest)(nil))
	gdi.PlaceHolder((*p13.DeleteMenuRequest)(nil))
	gdi.PlaceHolder((*p13.News)(nil))
	gdi.PlaceHolder((*p13.ListNewsRequest)(nil))
	gdi.PlaceHolder((*p13.CreateNewsRequest)(nil))
	gdi.PlaceHolder((*p13.UpdateNewsRequest)(nil))
	gdi.PlaceHolder((*p13.DeleteNewsRequest)(nil))
	gdi.PlaceHolder((*p13.OperationLogListRequest)(nil))
	gdi.PlaceHolder((*p13.DeleteOperationLogRequest)(nil))
	gdi.PlaceHolder((*p13.CreateRoleRequest)(nil))
	gdi.PlaceHolder((*p13.RoleListRequest)(nil))
	gdi.PlaceHolder((*p13.DeleteRoleRequest)(nil))
	gdi.PlaceHolder((*p13.UpdateRoleMenusRequest)(nil))
	gdi.PlaceHolder((*p13.UpdateRoleApisRequest)(nil))
	gdi.PlaceHolder((*p13.RegisterAndLoginRequest)(nil))
	gdi.PlaceHolder((*p13.CreateUserRequest)(nil))
	gdi.PlaceHolder((*p13.UserListRequest)(nil))
	gdi.PlaceHolder((*p13.DeleteUserRequest)(nil))
	gdi.PlaceHolder((*p13.ChangePwdRequest)(nil))
}
