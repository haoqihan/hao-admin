package errcode

var (
	ErrorGetTagListFail = NewError(2000, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(2001, "创建标签失败")
	ErrorUpdateTagFail  = NewError(2002, "更新标签失败")
	ErrorDeleteTagFail  = NewError(2003, "删除标签失败")
	ErrorCountTagFail   = NewError(2004, "统计标签失败")

	ErrorGetArticleFail    = NewError(2005, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(2006, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(2007, "创建文章失败")
	ErrorUpdateArticleFail = NewError(2008, "更新文章失败")
	ErrorDeleteArticleFail = NewError(2009, "删除文章失败")

	ErrorUploadFileFail = NewError(2010, "上传文件失败")

	ErrorRegisterUserFail    = NewError(2011, "注册用户失败,用户已存在")
	ErrorCaptchaFail         = NewError(2012, "生成验证码失败")
	ErrorCaptchaErr          = NewError(2013, "验证码错误")
	ErrorLoginFail           = NewError(2014, "用户名或密码错误")
	ErrorDeleteUserFail      = NewError(2015, "删除用户失败")
	ErrorChangePasswordFail  = NewError(2016, "修改失败，请检查用户名和密码")
	ErrorCountUserFail       = NewError(2017, "统计用户失败")
	ErrorGetUserListFail     = NewError(2018, "获取用户列表失败")
	ErrorUpdateUserFail      = NewError(2019, "更新用户信息失败")
	ErrorUpdateUserAuthFaile = NewError(2020, "更新用户权限失败")
)
