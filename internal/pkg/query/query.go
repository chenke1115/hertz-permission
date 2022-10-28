/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-26 10:33:53
 * @LastEditTime: 2022-09-26 14:54:09
 * @Description: Do not edit
 */
package query

// Herte handler query binding struct
type PaginationQuery struct {
	Offset uint   `form:"offset" json:"offset" query:"offset"`
	Limit  uint   `form:"limit" json:"limit" query:"limit"`
	Stime  string `json:"stime" form:"stime" query:"stime"`
	Etime  string `json:"etime" form:"etime" query:"etime"`
}
