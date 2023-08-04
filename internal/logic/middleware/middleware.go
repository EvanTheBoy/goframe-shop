package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility/response"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/backend/login",
	}
}

const (
	CtxAccountId      = "account_id"       //token获取
	CtxAccountName    = "account_name"     //token获取
	CtxAccountAvatar  = "account_avatar"   //token获取
	CtxAccountSex     = "account_sex"      //token获取
	CtxAccountStatus  = "account_status"   //token获取
	CtxAccountSign    = "account_sign"     //token获取
	CtxAccountIsAdmin = "account_is_admin" //token获取
	CtxAccountRoleIds = "account_role_ids" //token获取
)

type TokenInfo struct {
	Id   int
	Name string
}

// ResponseHandler 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, code.Code(), "", res)
	}
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	service.BizCtx().Init(r, customCtx)
	if userEntity := service.Session().GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:      uint(userEntity.Id),
			Name:    userEntity.Name,
			IsAdmin: uint8(userEntity.IsAdmin),
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
