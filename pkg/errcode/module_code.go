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
)
