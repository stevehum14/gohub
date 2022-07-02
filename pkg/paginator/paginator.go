// Package paginator 处理分页逻辑
package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// paging 分页数据
type Paging struct {
	CurrentPage int 	//当前页
	PerPage int			//每页条数
	TotalPage int		//总页数
	TotalCount int		//总条数
	NextPageUrl string	//下一页的链接
	PrevPageUrl string	//上一页的链接
}

// Paginator 分页操作类
type Paginator struct {
	BaseURL string 		// 用以拼接 URL
	PerPage int 		// 每页条数
	Page int			// 当前页
	Offset int			// 数据库读取数据时 Offset 的值
	TotalCount int64	// 总条数
	TotalPage int		// 总页数 = TotalCount/PerPage
	Sort string			// 排序规则
	Order string		// 排序顺序

	query *gorm.DB		// db query 句柄
	ctx *gin.Context	// gin context，方便调用
}
// Paginate 分页
// c —— gin.context 用来获取分页的 URL 参数
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// baseURL —— 用以分页链接
// data —— 模型数组，传址获取数据
// PerPage —— 每页条数，优先从 url 参数里取，否则使用 perPage 的值
// 用法:
//         query := database.DB.Model(Topic{}).Where("category_id = ?", cid)
//      var topics []Topic
//         paging := paginator.Paginate(
//             c,
//             query,
//             &topics,
//             app.APIURL(database.TableName(&Topic{})),
//             perPage,
//         )

//func Paginate(c *gin.Context,db *gorm.DB,data interface{},baseURL string,perPage int) Paging  {
//	// 初始化 Paginator 实例
//	p := &Paginator{
//		query: db,
//		ctx: c,
//	}
//}
// 拼接分页链接
func (p Paginator) getPageLink(page int) string {

}
// getPrevPageURL 返回下一页的链接
func (p Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.get
}