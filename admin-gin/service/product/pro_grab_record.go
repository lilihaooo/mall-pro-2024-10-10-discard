package product

//
//import (
//	"admin-gin/global"
//	"admin-gin/model/common/request"
//	"admin-gin/model/product"
//	"admin-gin/model/product/response"
//	"admin-gin/model/system"
//)
//
//func (*products) GrabRecord(req request.PageInfo) (list []response.ProGrabRecord, total int64, err error) {
//	query := global.XTK_DB
//	err = query.Model(product.GrabRecord{}).Count(&total).Error
//	if err != nil {
//		return
//	}
//
//	var gr []product.GrabRecord
//	offset := req.PageSize * (req.Page - 1)
//	err = query.Order("created_at desc").Limit(req.PageSize).Offset(offset).Find(&gr).Error
//	if err != nil {
//		return
//	}
//
//	// 获取用户数据
//	var Users []system.SysUser
//	err = global.GVA_DB.Select("id, username").Find(&Users).Error
//	if err != nil {
//		return
//	}
//	userIDNameMap := make(map[uint]string)
//	for _, user := range Users {
//		userIDNameMap[user.ID] = user.Username
//	}
//	// 序列化数据
//	for _, record := range gr {
//		list = append(list, response.ProGrabRecord{
//			CreatedTime: record.GVA_MODEL.CreatedAt.Format("2006-01-02 15:04:05"),
//			UserName:    userIDNameMap[record.UserID],
//			CategoryNum: record.CategoryNum,
//			GoNum:       record.GoNum,
//			PageNum:     record.PageNum,
//			RunTime:     record.RunTime,
//			GrabNum:     record.GrabNum,
//			Uuid:        record.Uuid,
//		})
//	}
//	return
//}
