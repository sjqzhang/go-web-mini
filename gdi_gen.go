package main

import (
	"embed"
	"github.com/sjqzhang/gdi"
	p1 "go-web-mini/apps/cms/controller"
	p2 "go-web-mini/apps/cms/model"
	p3 "go-web-mini/apps/cms/repository"
	p4 "go-web-mini/apps/cms/service"
	p5 "go-web-mini/apps/cms/vo"
	p6 "go-web-mini/apps/system/controller"
	p7 "go-web-mini/apps/system/dto"
	p8 "go-web-mini/apps/system/model"
	p9 "go-web-mini/apps/system/repository"
	p10 "go-web-mini/apps/system/service"
	p11 "go-web-mini/apps/system/vo"
	p12 "go-web-mini/config"
	p14 "go-web-mini/generator"
	p15 "go-web-mini/global"
)

//go:embed generator middleware apps/cms/model apps/cms/repository apps/system/dto apps/system/model apps/system/repository apps/system/service apps/cms/vo global apps/cms/controller docs response util apps/cms/service apps/system/controller apps/system/vo config routes
var gdiEmbedFiles embed.FS

func init() {
	gdi.SetEmbedFs(&gdiEmbedFiles)
	_ = gdi.GDIPool{}
	gdi.PlaceHolder((*p1.ArticleController)(nil))
	gdi.PlaceHolder((*p2.Article)(nil))
	gdi.PlaceHolder((*p2.ArticleQuery)(nil))
	gdi.PlaceHolder((*p2.Model)(nil))
	gdi.PlaceHolder((*p2.PagerModel)(nil))
	gdi.PlaceHolder((*p3.ArticleRepository)(nil))
	gdi.PlaceHolder((*p4.ArticleService)(nil))
	gdi.PlaceHolder((*p5.ArticleResponse)(nil))
	gdi.PlaceHolder((*p5.Article)(nil))
	gdi.PlaceHolder((*p5.ListArticleResponse)(nil))
	gdi.PlaceHolder((*p5.ListArticleRequest)(nil))
	gdi.PlaceHolder((*p5.GetArticleResponse)(nil))
	gdi.PlaceHolder((*p5.CreateArticleRequest)(nil))
	gdi.PlaceHolder((*p5.CreateArticleResponse)(nil))
	gdi.PlaceHolder((*p5.UpdateArticleRequest)(nil))
	gdi.PlaceHolder((*p5.UpdateArticleResponse)(nil))
	gdi.PlaceHolder((*p5.DeleteArticleRequest)(nil))
	gdi.PlaceHolder((*p5.GetArticleRequest)(nil))
	gdi.PlaceHolder((*p5.DeleteArticleResponse)(nil))
	gdi.PlaceHolder((*p5.Response)(nil))
	gdi.PlaceHolder((*p6.ApiController)(nil))
	gdi.PlaceHolder((*p6.MenuController)(nil))
	gdi.PlaceHolder((*p6.OperationLogController)(nil))
	gdi.PlaceHolder((*p6.RoleController)(nil))
	gdi.PlaceHolder((*p6.TableMetadataController)(nil))
	gdi.PlaceHolder((*p6.UserController)(nil))
	gdi.PlaceHolder((*p7.ApiTreeDto)(nil))
	gdi.PlaceHolder((*p7.UserInfoDto)(nil))
	gdi.PlaceHolder((*p7.UsersDto)(nil))
	gdi.PlaceHolder((*p8.Api)(nil))
	gdi.PlaceHolder((*p8.RoleCasbin)(nil))
	gdi.PlaceHolder((*p8.Model)(nil))
	gdi.PlaceHolder((*p8.PagerModel)(nil))
	gdi.PlaceHolder((*p8.Menu)(nil))
	gdi.PlaceHolder((*p8.OperationLog)(nil))
	gdi.PlaceHolder((*p8.Role)(nil))
	gdi.PlaceHolder((*p8.TableMetadata)(nil))
	gdi.PlaceHolder((*p8.TableMetadataQuery)(nil))
	gdi.PlaceHolder((*p8.User)(nil))
	gdi.PlaceHolder((*p9.ApiRepository)(nil))
	gdi.PlaceHolder((*p9.MenuRepository)(nil))
	gdi.PlaceHolder((*p9.OperationLogRepository)(nil))
	gdi.PlaceHolder((*p9.RoleRepository)(nil))
	gdi.PlaceHolder((*p9.TableMetadataRepository)(nil))
	gdi.PlaceHolder((*p9.UserRepository)(nil))
	gdi.PlaceHolder((*p10.TableMetadataService)(nil))
	gdi.PlaceHolder((*p11.ApiListRequest)(nil))
	gdi.PlaceHolder((*p11.CreateApiRequest)(nil))
	gdi.PlaceHolder((*p11.UpdateApiRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteApiRequest)(nil))
	gdi.PlaceHolder((*p11.Response)(nil))
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
	gdi.PlaceHolder((*p11.TableMetadataResponse)(nil))
	gdi.PlaceHolder((*p11.TableMetadata)(nil))
	gdi.PlaceHolder((*p11.ListTableMetadataResponse)(nil))
	gdi.PlaceHolder((*p11.ListTableMetadataRequest)(nil))
	gdi.PlaceHolder((*p11.GetTableMetadataResponse)(nil))
	gdi.PlaceHolder((*p11.CreateTableMetadataRequest)(nil))
	gdi.PlaceHolder((*p11.CreateTableMetadataResponse)(nil))
	gdi.PlaceHolder((*p11.UpdateTableMetadataRequest)(nil))
	gdi.PlaceHolder((*p11.UpdateTableMetadataResponse)(nil))
	gdi.PlaceHolder((*p11.DeleteTableMetadataRequest)(nil))
	gdi.PlaceHolder((*p11.GetTableMetadataRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteTableMetadataResponse)(nil))
	gdi.PlaceHolder((*p11.RegisterAndLoginRequest)(nil))
	gdi.PlaceHolder((*p11.CreateUserRequest)(nil))
	gdi.PlaceHolder((*p11.UserListRequest)(nil))
	gdi.PlaceHolder((*p11.DeleteUserRequest)(nil))
	gdi.PlaceHolder((*p11.ChangePwdRequest)(nil))
	gdi.PlaceHolder((*p12.SystemConfig)(nil))
	gdi.PlaceHolder((*p12.LogsConfig)(nil))
	gdi.PlaceHolder((*p12.MysqlConfig)(nil))
	gdi.PlaceHolder((*p12.CasbinConfig)(nil))
	gdi.PlaceHolder((*p12.JwtConfig)(nil))
	gdi.PlaceHolder((*p12.RateLimitConfig)(nil))
	gdi.PlaceHolder((*p12.RedisConfig)(nil))
	gdi.PlaceHolder((*p14.FieldResult)(nil))
	gdi.PlaceHolder((*p14.TableResult)(nil))
	gdi.PlaceHolder((*p14.CommonObject)(nil))
	gdi.PlaceHolder((*p14.Config)(nil))
	gdi.PlaceHolder((*p14.IndexResult)(nil))
	gdi.PlaceHolder((*p15.CustomLogger)(nil))
}
