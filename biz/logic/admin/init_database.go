/*
 * Copyright 2023 CoolFunAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package admin

import (
	"context"
	"sync"

	"coolfun-admin/pkg/encrypt"

	"github.com/casbin/casbin/v2"

	"coolfun-admin/data/ent"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cockroachdb/errors"
)

type InitDatabase struct {
	DB  *ent.Client
	Csb *casbin.Enforcer
	Mu  *sync.Mutex
}

var (
	mutex              = new(sync.Mutex)
	InitDatabaseStatus bool
)

func NewInitDatabase(db *ent.Client, csb *casbin.Enforcer) *InitDatabase {
	return &InitDatabase{
		DB:  db,
		Csb: csb,
		Mu:  mutex,
	}
}

func (I *InitDatabase) InitDatabase(ctx context.Context) error {
	// add lock to avoid duplicate initialization
	I.Mu.Lock()
	defer I.Mu.Unlock()

	// judge if the initialization had been done
	check, _ := I.DB.API.Query().Count(ctx)
	if InitDatabaseStatus || check > 0 {
		return errors.New("Database had been initialized")
	}

	// insert init data
	err := I.insertUserData(ctx)
	if err != nil {
		hlog.Error("insert user data failed", err)
		err = errors.Wrap(err, "insert user data failed")
		return err
	}

	err = I.insertRoleData(ctx)
	if err != nil {
		hlog.Error("insert role data failed", err)
		err = errors.Wrap(err, "insert role data failed")
		return err
	}

	err = I.insertMenuData(ctx)
	if err != nil {
		hlog.Error("insert menu data failed", err)
		err = errors.Wrap(err, "insert menu data failed")
		return err
	}

	err = I.insertApiData(ctx)
	if err != nil {
		hlog.Error("insert api data failed", err)
		err = errors.Wrap(err, "insert api data failed")
		return err
	}
	err = I.insertRoleMenuAuthorityData(ctx)
	if err != nil {
		hlog.Error("insert role menu authority data failed", err)
		err = errors.Wrap(err, "insert role menu authority data failed")
		return err
	}
	err = I.insertCasbinPoliciesData(ctx)
	if err != nil {
		hlog.Error("insert casbin policies data failed", err)
		err = errors.Wrap(err, "insert casbin policies data failed")
		return err
	}

	// set init status
	InitDatabaseStatus = true
	return nil
}

// insert init user data
func (I *InitDatabase) insertUserData(ctx context.Context) error {
	var users []*ent.UserCreate
	password, _ := encrypt.BcryptEncrypt("admin")
	users = append(users, I.DB.User.Create().
		SetUsername("admin").
		SetNickname("admin").
		SetPassword(password).
		SetEmail("admin@gmail.com").
		SetMobile("12345678901").
		SetRoleID(1),
	)

	err := I.DB.User.CreateBulk(users...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert init apis data
func (I *InitDatabase) insertRoleData(ctx context.Context) error {
	roles := make([]*ent.RoleCreate, 3)
	roles[0] = I.DB.Role.Create().
		SetName("role.admin").
		SetValue("admin").
		SetRemark("超级管理员").
		SetOrderNo(1)

	roles[1] = I.DB.Role.Create().
		SetName("role.stuff").
		SetValue("stuff").
		SetRemark("普通员工").
		SetOrderNo(2)

	roles[2] = I.DB.Role.Create().
		SetName("role.member").
		SetValue("member").
		SetRemark("注册会员").
		SetOrderNo(3)

	err := I.DB.Role.CreateBulk(roles...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert init API data
// insert init API data
func (I *InitDatabase) insertApiData(ctx context.Context) error {
	// 使用动态切片来存储已初始化的 APICreate 实例
	var apis []*ent.APICreate

	// USER
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/login").
		SetDescription("apiDesc.userLogin").
		SetAPIGroup("user").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/register").
		SetDescription("apiDesc.userRegister").
		SetAPIGroup("user").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/create").
		SetDescription("apiDesc.createUser").
		SetAPIGroup("user").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/update").
		SetDescription("apiDesc.updateUser").
		SetAPIGroup("user").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/change-password").
		SetDescription("apiDesc.userChangePassword").
		SetAPIGroup("user").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/info").
		SetDescription("apiDesc.OauthUserInfo").
		SetAPIGroup("user").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/list").
		SetDescription("apiDesc.userList").
		SetAPIGroup("user").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user").
		SetDescription("apiDesc.deleteUser").
		SetAPIGroup("user").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/perm").
		SetDescription("apiDesc.userPermissions").
		SetAPIGroup("user").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/profile").
		SetDescription("apiDesc.userProfile").
		SetAPIGroup("user").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/profile").
		SetDescription("apiDesc.updateProfile").
		SetAPIGroup("user").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/logout").
		SetDescription("apiDesc.logout").
		SetAPIGroup("user").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/user/status").
		SetDescription("apiDesc.updateUserStatus").
		SetAPIGroup("user").
		SetMethod("POST"))

	// ROLE
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/role/create").
		SetDescription("apiDesc.createRole").
		SetAPIGroup("role").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/role/update").
		SetDescription("apiDesc.updateRole").
		SetAPIGroup("role").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/role").
		SetDescription("apiDesc.deleteRole").
		SetAPIGroup("role").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/role/list").
		SetDescription("apiDesc.roleList").
		SetAPIGroup("role").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/role/status").
		SetDescription("apiDesc.setRoleStatus").
		SetAPIGroup("role").
		SetMethod("POST"))

	// MENU
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/create").
		SetDescription("apiDesc.createMenu").
		SetAPIGroup("menu").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/update").
		SetDescription("apiDesc.updateMenu").
		SetAPIGroup("menu").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu").
		SetDescription("apiDesc.deleteMenu").
		SetAPIGroup("menu").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/list").
		SetDescription("apiDesc.menuList").
		SetAPIGroup("menu").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/role").
		SetDescription("apiDesc.roleMenu").
		SetAPIGroup("menu").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/param/create").
		SetDescription("apiDesc.createMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/param/update").
		SetDescription("apiDesc.updateMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/param/list").
		SetDescription("apiDesc.menuParamListByMenuID").
		SetAPIGroup("menu").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/menu/param").
		SetDescription("apiDesc.deleteMenuParam").
		SetAPIGroup("menu").
		SetMethod("DELETE"))

	// CAPTCHA
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/captcha").
		SetDescription("apiDesc.captcha").
		SetAPIGroup("captcha").
		SetMethod("GET"))

	// AUTHORIZATION
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/authority/api/create").
		SetDescription("apiDesc.createApiAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/authority/api/update").
		SetDescription("apiDesc.updateApiAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/authority/api/role").
		SetDescription("apiDesc.APIAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/authority/menu/create").
		SetDescription("apiDesc.createMenuAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/authority/menu/update").
		SetDescription("apiDesc.updateMenuAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/authority/menu/role").
		SetDescription("apiDesc.menuAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST"))

	// API
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/api/create").
		SetDescription("apiDesc.createApi").
		SetAPIGroup("api").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/api/update").
		SetDescription("apiDesc.updateApi").
		SetAPIGroup("api").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/api").
		SetDescription("apiDesc.deleteAPI").
		SetAPIGroup("api").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/api/list").
		SetDescription("apiDesc.APIList").
		SetAPIGroup("api").
		SetMethod("GET"))

	// DICTIONARY
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/create").
		SetDescription("apiDesc.createDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/update").
		SetDescription("apiDesc.updateDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict").
		SetDescription("apiDesc.deleteDictionary").
		SetAPIGroup("dictionary").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/detail").
		SetDescription("apiDesc.deleteDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/detail/create").
		SetDescription("apiDesc.createDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/detail/update").
		SetDescription("apiDesc.updateDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/detail/list").
		SetDescription("apiDesc.getDictionaryListDetail").
		SetAPIGroup("dictionary").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/dict/list").
		SetDescription("apiDesc.getDictionaryList").
		SetAPIGroup("dictionary").
		SetMethod("GET"))

	// TOKEN
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/token/create").
		SetDescription("apiDesc.createToken").
		SetAPIGroup("token").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/token/update").
		SetDescription("apiDesc.updateToken").
		SetAPIGroup("token").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/token").
		SetDescription("apiDesc.deleteToken").
		SetAPIGroup("token").
		SetMethod("DELETE"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/token/list").
		SetDescription("apiDesc.getTokenList").
		SetAPIGroup("token").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/token/status").
		SetDescription("apiDesc.setTokenStatus").
		SetAPIGroup("token").
		SetMethod("POST"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/token/logout").
		SetDescription("user.forceLoggingOut").
		SetAPIGroup("token").
		SetMethod("POST"))

	// LOGS
	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/logs/list").
		SetDescription("apiDesc.getLogsList").
		SetAPIGroup("logs").
		SetMethod("GET"))

	apis = append(apis, I.DB.API.Create().
		SetPath("/api/admin/logs/deleteAll").
		SetDescription("apiDesc.deleteLogs").
		SetAPIGroup("logs").
		SetMethod("DELETE"))

	// 批量插入 API 数据
	err := I.DB.API.CreateBulk(apis...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init menu data
func (I *InitDatabase) insertMenuData(ctx context.Context) error {
	// 使用动态切片来存储已初始化的 MenuCreate 实例
	var menus []*ent.MenuCreate

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(0).
		SetMenuType(0).
		SetParentID(1).
		SetPath("").
		SetName("root").
		SetComponent("").
		SetOrderNo(0).
		SetTitle("").
		SetIcon("").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(1).
		SetPath("/dashboard/workbench/index").
		SetName("Dashboard").
		SetComponent("/dashboard/workbench/index").
		SetOrderNo(0).
		SetTitle("控制台").
		SetIcon("ant-design:home-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("").
		SetName("System Management").
		SetComponent("LAYOUT").
		SetOrderNo(1).
		SetTitle("系统管理").
		SetIcon("ant-design:tool-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/menu/index").
		SetName("MenuManagement").
		SetComponent("/sys/menu/index").
		SetOrderNo(1).
		SetTitle("菜单管理").
		SetIcon("ant-design:bars-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/role/index").
		SetName("Role Management").
		SetComponent("/sys/role/index").
		SetOrderNo(2).
		SetTitle("角色管理").
		SetIcon("ant-design:user-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/api/index").
		SetName("API Management").
		SetComponent("/sys/api/index").
		SetOrderNo(3).
		SetTitle("API管理").
		SetIcon("ant-design:api-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/user/index").
		SetName("User Management").
		SetComponent("/sys/user/index").
		SetOrderNo(4).
		SetTitle("用户管理").
		SetIcon("ant-design:user-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/dictionary/index").
		SetName("Dictionary Management").
		SetComponent("/sys/dictionary/index").
		SetOrderNo(5).
		SetTitle("字典管理").
		SetIcon("ant-design:book-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(2).
		SetParentID(3).
		SetPath("/sys/dictionary/detail").
		SetName("Dictionary Detail").
		SetComponent("/sys/dictionary/detail").
		SetOrderNo(6).
		SetTitle("字典明细").
		SetIcon("ant-design:align-left-outlined").
		SetHideMenu(true))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/token/index").
		SetName("Token Management").
		SetComponent("/sys/token/index").
		SetOrderNo(8).
		SetTitle("Token管理").
		SetIcon("ant-design:lock-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/logs/index").
		SetName("Logs Management").
		SetComponent("/sys/logs/index").
		SetOrderNo(9).
		SetTitle("日志管理").
		SetIcon("ant-design:profile-twotone").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("").
		SetName("Other Pages").
		SetComponent("LAYOUT").
		SetOrderNo(2).
		SetTitle("其他页面").
		SetIcon("ant-design:question-circle-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(13).
		SetPath("/sys/profile/index").
		SetName("Profile").
		SetComponent("/sys/profile/index").
		SetOrderNo(2).
		SetTitle("用户信息").
		SetIcon("ant-design:profile-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("").
		SetName("Dev Tool").
		SetComponent("LAYOUT").
		SetOrderNo(3).
		SetTitle("开发工具").
		SetIcon("ant-design:api-filled").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(16).
		SetPath("/devtool/structToProto").
		SetName("StructToProto").
		SetComponent("/devtool/structToProto").
		SetOrderNo(1).
		SetTitle("StructToProto").
		SetIcon("ant-design:disconnect-outlined").
		SetHideMenu(false))

	menus = append(menus, I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(16).
		SetPath("/devtool/structTag").
		SetName("DeleteStructTag").
		SetComponent("/devtool/structTag").
		SetOrderNo(2).
		SetTitle("DeleteStructTag").
		SetIcon("ant-design:disconnect-outlined").
		SetHideMenu(false))

	// 批量插入菜单数据
	err := I.DB.Menu.CreateBulk(menus...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert admin menu authority

func (I *InitDatabase) insertRoleMenuAuthorityData(ctx context.Context) error {
	count, err := I.DB.Menu.Query().Count(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	menuIDs := make([]uint64, count)

	for i := range menuIDs {
		menuIDs[i] = uint64(i + 1)
	}

	err = I.DB.Role.Update().AddMenuIDs(menuIDs...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init casbin policies

func (I *InitDatabase) insertCasbinPoliciesData(ctx context.Context) error {
	apis, err := I.DB.API.Query().All(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	var policies [][]string
	for _, v := range apis {
		policies = append(policies, []string{"1", v.Path, v.Method})
	}

	addResult, err := I.Csb.AddPolicies(policies)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	if !addResult {
		return errors.Wrap(err, "db failed")
	}
	return nil
}
