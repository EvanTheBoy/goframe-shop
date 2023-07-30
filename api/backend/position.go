package backend

import "github.com/gogf/gf/v2/frame/g"

type PositionReq struct {
	g.Meta    `path:"/position/add" tags:"Position" method:"post" summary:"You first position api"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
	Sort      int    `json:"sort"    dc:"排序"`
}

type PositionRes struct {
	PositionId int `json:"position_id"`
}
