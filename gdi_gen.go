package main

import (
	"embed"
	"github.com/sjqzhang/gdi"
	p1 "go-web-mini/config"
	p2 "go-web-mini/controller"
	p4 "go-web-mini/dto"
	p5 "go-web-mini/generator"
	p8 "go-web-mini/model"
	p9 "go-web-mini/repository"
	p12 "go-web-mini/service"
	p14 "go-web-mini/vo"
)

//go:embed docs generator repository response service util config global middleware model routes controller dto vo
var gdiEmbedFiles embed.FS

func init() {
	gdi.SetEmbedFs(&gdiEmbedFiles)
	_ = gdi.GDIPool{}
	gdi.PlaceHolder((*p1.SystemConfig)(nil))
	gdi.PlaceHolder((*p1.LogsConfig)(nil))
	gdi.PlaceHolder((*p1.MysqlConfig)(nil))
	gdi.PlaceHolder((*p1.CasbinConfig)(nil))
	gdi.PlaceHolder((*p1.JwtConfig)(nil))
	gdi.PlaceHolder((*p1.RateLimitConfig)(nil))
	gdi.PlaceHolder((*p1.RedisConfig)(nil))
	gdi.PlaceHolder((*p2.ApiController)(nil))
	gdi.PlaceHolder((*p2.BranchController)(nil))
	gdi.PlaceHolder((*p2.MenuController)(nil))
	gdi.PlaceHolder((*p2.ModuleController)(nil))
	gdi.PlaceHolder((*p2.NewsController)(nil))
	gdi.PlaceHolder((*p2.OperationLogController)(nil))
	gdi.PlaceHolder((*p2.RoleController)(nil))
	gdi.PlaceHolder((*p2.UserController)(nil))
	gdi.PlaceHolder((*p4.ApiTreeDto)(nil))
	gdi.PlaceHolder((*p4.UserInfoDto)(nil))
	gdi.PlaceHolder((*p4.UsersDto)(nil))
	gdi.PlaceHolder((*p5.FieldResult)(nil))
	gdi.PlaceHolder((*p5.TableResult)(nil))
	gdi.PlaceHolder((*p5.CommonObject)(nil))
	gdi.PlaceHolder((*p5.Config)(nil))
	gdi.PlaceHolder((*p8.Api)(nil))
	gdi.PlaceHolder((*p8.Branch)(nil))
	gdi.PlaceHolder((*p8.BranchQuery)(nil))
	gdi.PlaceHolder((*p8.RoleCasbin)(nil))
	gdi.PlaceHolder((*p8.Model)(nil))
	gdi.PlaceHolder((*p8.PagerModel)(nil))
	gdi.PlaceHolder((*p8.Menu)(nil))
	gdi.PlaceHolder((*p8.Module)(nil))
	gdi.PlaceHolder((*p8.ModuleQuery)(nil))
	gdi.PlaceHolder((*p8.News)(nil))
	gdi.PlaceHolder((*p8.NewsQuery)(nil))
	gdi.PlaceHolder((*p8.OperationLog)(nil))
	gdi.PlaceHolder((*p8.Role)(nil))
	gdi.PlaceHolder((*p8.User)(nil))
	gdi.PlaceHolder((*p9.ApiRepository)(nil))
	gdi.PlaceHolder((*p9.BranchRepository)(nil))
	gdi.PlaceHolder((*p9.MenuRepository)(nil))
	gdi.PlaceHolder((*p9.ModuleRepository)(nil))
	gdi.PlaceHolder((*p9.NewsRepository)(nil))
	gdi.PlaceHolder((*p9.OperationLogRepository)(nil))
	gdi.PlaceHolder((*p9.RoleRepository)(nil))
	gdi.PlaceHolder((*p9.UserRepository)(nil))
	gdi.PlaceHolder((*p12.BranchService)(nil))
	gdi.PlaceHolder((*p12.ModuleService)(nil))
	gdi.PlaceHolder((*p12.NewsService)(nil))
	gdi.PlaceHolder((*p14.ApiListRequest)(nil))
	gdi.PlaceHolder((*p14.CreateApiRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateApiRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteApiRequest)(nil))
	gdi.PlaceHolder((*p14.BranchResponse)(nil))
	gdi.PlaceHolder((*p14.Branch)(nil))
	gdi.PlaceHolder((*p14.ListBranchResponse)(nil))
	gdi.PlaceHolder((*p14.ListBranchRequest)(nil))
	gdi.PlaceHolder((*p14.GetBranchResponse)(nil))
	gdi.PlaceHolder((*p14.CreateBranchRequest)(nil))
	gdi.PlaceHolder((*p14.CreateBranchResponse)(nil))
	gdi.PlaceHolder((*p14.UpdateBranchRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateBranchResponse)(nil))
	gdi.PlaceHolder((*p14.DeleteBranchRequest)(nil))
	gdi.PlaceHolder((*p14.GetBranchRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteBranchResponse)(nil))
	gdi.PlaceHolder((*p14.Response)(nil))
	gdi.PlaceHolder((*p14.CreateMenuRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateMenuRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteMenuRequest)(nil))
	gdi.PlaceHolder((*p14.ModuleResponse)(nil))
	gdi.PlaceHolder((*p14.Module)(nil))
	gdi.PlaceHolder((*p14.ListModuleResponse)(nil))
	gdi.PlaceHolder((*p14.ListModuleRequest)(nil))
	gdi.PlaceHolder((*p14.GetModuleResponse)(nil))
	gdi.PlaceHolder((*p14.CreateModuleRequest)(nil))
	gdi.PlaceHolder((*p14.CreateModuleResponse)(nil))
	gdi.PlaceHolder((*p14.UpdateModuleRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateModuleResponse)(nil))
	gdi.PlaceHolder((*p14.DeleteModuleRequest)(nil))
	gdi.PlaceHolder((*p14.GetModuleRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteModuleResponse)(nil))
	gdi.PlaceHolder((*p14.NewsResponse)(nil))
	gdi.PlaceHolder((*p14.News)(nil))
	gdi.PlaceHolder((*p14.ListNewsResponse)(nil))
	gdi.PlaceHolder((*p14.ListNewsRequest)(nil))
	gdi.PlaceHolder((*p14.GetNewsResponse)(nil))
	gdi.PlaceHolder((*p14.CreateNewsRequest)(nil))
	gdi.PlaceHolder((*p14.CreateNewsResponse)(nil))
	gdi.PlaceHolder((*p14.UpdateNewsRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateNewsResponse)(nil))
	gdi.PlaceHolder((*p14.DeleteNewsRequest)(nil))
	gdi.PlaceHolder((*p14.GetNewsRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteNewsResponse)(nil))
	gdi.PlaceHolder((*p14.OperationLogListRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteOperationLogRequest)(nil))
	gdi.PlaceHolder((*p14.CreateRoleRequest)(nil))
	gdi.PlaceHolder((*p14.RoleListRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteRoleRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateRoleMenusRequest)(nil))
	gdi.PlaceHolder((*p14.UpdateRoleApisRequest)(nil))
	gdi.PlaceHolder((*p14.RegisterAndLoginRequest)(nil))
	gdi.PlaceHolder((*p14.CreateUserRequest)(nil))
	gdi.PlaceHolder((*p14.UserListRequest)(nil))
	gdi.PlaceHolder((*p14.DeleteUserRequest)(nil))
	gdi.PlaceHolder((*p14.ChangePwdRequest)(nil))
}
