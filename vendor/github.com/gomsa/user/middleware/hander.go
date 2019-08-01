package middleware

import (
	"context"
	"errors"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	"github.com/gomsa/tools/config"
	"github.com/gomsa/user/client"
	authPb "github.com/gomsa/user/proto/auth"
	casbinPb "github.com/gomsa/user/proto/casbin"
)

// Handler 处理器
// 包含一些高阶函数比如中间件常用的 token 验证等
type Handler struct {
	Permissions []config.Permission
}

// Wrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func (h *Handler) Wrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		if h.IsAuth(req) {
			meta, ok := metadata.FromContext(ctx)
			if !ok {
				return errors.New("no auth meta-data found in request")
			}
			if token, ok := meta["x-token"]; ok {
				// Note this is now uppercase (not entirely sure why this is...)
				// token := strings.Split(meta["authorization"], "Bearer ")[1]
				// Auth here
				authResp, err := client.Auth.ValidateToken(ctx, &authPb.Request{
					Token: token,
				})
				if err != nil || authResp.Valid == false {
					return err
				}
				// 设置用户 id
				meta["user_id"] = authResp.User.Id
				meta["service"] = req.Service()
				meta["method"] = req.Method()
				ctx = metadata.NewContext(ctx, meta)
				if h.IsPolicy(req) {
					casbinResp, err := client.Casbin.Validate(ctx, &casbinPb.Request{})
					if err != nil || casbinResp.Valid == false {
						return err
					}	
				}
			} else {
				return errors.New("Empty Authorization")
			}
		}
		err = fn(ctx, req, resp)
		return err
	}
}

// IsAuth 检测用户授权
func (h *Handler) IsAuth(req server.Request) bool {
	for _, p := range h.Permissions {
		if p.Service == req.Service() && p.Method == req.Method() && p.Auth == true {
			return true
		}
	}
	return false
}
// IsPolicy 检查用户权限
func (h *Handler) IsPolicy(req server.Request) bool {
	for _, p := range h.Permissions {
		if p.Service == req.Service() && p.Method == req.Method() && p.Policy == true {
			return true
		}
	}
	return false
}
