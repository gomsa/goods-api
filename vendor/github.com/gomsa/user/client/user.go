package client

import (
	"context"
	"os"

	"github.com/gomsa/tools/config"
	"github.com/gomsa/tools/k8s/client"

	authPB "github.com/gomsa/user/proto/auth"
	casbinPB "github.com/gomsa/user/proto/casbin"
	frontPermitPB "github.com/gomsa/user/proto/frontPermit"
	permissionPB "github.com/gomsa/user/proto/permission"
	rolePB "github.com/gomsa/user/proto/role"
	userPB "github.com/gomsa/user/proto/user"
)

var (
	// User 用户客户端
	User userPB.UsersClient
	// Auth 认证客户端
	Auth authPB.AuthClient
	// FrontPermit 前端权限客户端
	FrontPermit frontPermitPB.FrontPermitsClient
	// Permission 权限客户端
	Permission permissionPB.PermissionsClient
	// Role 角色客户端
	Role rolePB.RolesClient
	// Casbin 权限管理客户端
	Casbin casbinPB.CasbinClient
)

func init() {
	srvName := os.Getenv("USER_NAME")

	User = userPB.NewUsersClient(srvName, client.DefaultClient)
	Auth = authPB.NewAuthClient(srvName, client.DefaultClient)
	FrontPermit = frontPermitPB.NewFrontPermitsClient(srvName, client.DefaultClient)
	Permission = permissionPB.NewPermissionsClient(srvName, client.DefaultClient)
	Role = rolePB.NewRolesClient(srvName, client.DefaultClient)
	Casbin = casbinPB.NewCasbinClient(srvName, client.DefaultClient)
}

// SyncPermission 同步权限
func SyncPermission(permission []config.Permission) error {
	for _, p := range permission {
		if p.Policy {
			req := permissionPB.Permission{}
			req.Service = p.Service
			req.Method = p.Method
			req.Name = p.Name
			req.Description = p.Description
			_, err := Permission.UpdateOrCreate(context.TODO(), &req)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
