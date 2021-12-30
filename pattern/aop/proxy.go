package aop

// UserProxy 代理，也是切面
type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) UserProxy {
	return UserProxy{user: user}
}

// Auth 校验，切入点
func (p UserProxy) Auth() {
	group := UserAopGroup{}
	group.Add(&ValidateName{})
	group.Add(&ValidatePassword{MaxLength: 10, MinLength: 6})

	// 前置通知
	if err := group.Before(p.user); err != nil {
		panic(err)
	}

	// 后置通知
	defer group.After(p.user)
	// 实际逻辑
	p.user.Auth()
}
