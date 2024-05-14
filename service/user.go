package service

//type UserService struct {
//	UserName        string `json:"user_name" form:"user_name"`
//	Password        string `json:"password" form:"password"`
//	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
//	Email           string `json:"email" form:"email"`
//	Phone           string `json:"phone" form:"phone"`
//}
//
//func (service UserService) Register(ctx context.Context) serializer.Response {
//	var user model.User_basic
//
//	userDao := dao.NewUserBasicDao()
//
//	_, exist, err := userDao.ExistOrNotByUserName(user.UserName)
//	if err != nil {
//		return serializer.Response{
//			Code:    1,
//			Message: "error" + err.Error(),
//		}
//	}
//	if exist {
//		return serializer.Response{
//			Code:    1,
//			Message: "error" + "用户名重复注册!!!",
//		}
//	}
//	user = model.User_basic{
//		UserName:     service.UserName,
//		Email:        service.Email,
//		Status:       "regular",
//		Authority:    "Ordinary employees",
//		UserIdentity: "tourist",
//		Phone:        service.Phone,
//	}
//	// 密码加密
//	if err = user.SetPassword(service.Password); err != nil {
//		return serializer.Response{
//			Code:    1,
//			Message: "error" + err.Error(),
//		}
//	}
//	// 创建用户
//	err = userDao.CreateUser(&user)
//	if err != nil {
//		return serializer.Response{
//			Code:    1,
//			Message: "error" + err.Error(),
//		}
//	}
//	return serializer.Response{
//		Code:    0,
//		Data:    user,
//		ReqCode: utils.GenReqCode(),
//		Message: "success!",
//	}
//}
//
//func (service *UserService) Login(ctx context.Context) serializer.Response {
//	var user *model.User
//	userDao := dao.NewUserBasicDao()
//	// 判断该用户是否存在
//	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
//	if !exist || err != nil {
//		return serializer.Response{
//			Code:    1,
//			Message: "用户不存在,请先注册",
//		}
//	}
//	// 检验密码
//	if user.CheckPassword(service.Password) == false {
//		return serializer.Response{
//			Code:    1,
//			Message: "密码错误，请重新登录",
//		}
//	}
//
//	return serializer.Response{
//		Code:    0,
//		Data:    user,
//		ReqCode: utils.GenReqCode(),
//		Message: "success",
//	}
//
//}
