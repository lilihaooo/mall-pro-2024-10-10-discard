package common

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var mu sync.Mutex

const (
	MESSAGE_LEVEL_WARNING = 2
	MESSAGE_LEVEL_ERROR   = 1
	MESSAGE_LEVEL_INFO    = 0

	FORWHERE_GRABACTIONLOG = "grab_action_log"
	FORWHERE_ELMESSAGE     = "ElMessage"
	FORWHERE_GRABLIVE      = "grab_live"
)

type messageListener struct {
	Code           int         `json:"code"`
	Msg            string      `json:"msg"`
	Data           interface{} `json:"data"`
	GrabRecodeUuid string      ` json:"grab_recode_uuid"`
	ForWhere       string      `json:"for_where"`
	ToUserIDs      []uint      `json:"to_user_ids"`
}

func WsSendMsg(code int, msg string, data interface{}, forWhere string, toUserIDs []uint, grabRecordUuid string) {
	m := messageListener{
		Code:           code,
		Msg:            msg,
		Data:           data,
		ForWhere:       forWhere,
		ToUserIDs:      toUserIDs,
		GrabRecodeUuid: grabRecordUuid,
	}
	m.SeedMsg()
	m.saveMsg()
}

func (m *messageListener) SeedMsg() {
	respByte, _ := json.Marshal(m)
	// 没有指定用户, 则给全部连接发送消息
	if len(m.ToUserIDs) == 0 {
		for _, conn := range global.WsConnMap {
			if conn != nil {
				mu.Lock()
				_ = conn.WriteMessage(websocket.TextMessage, respByte)
				mu.Unlock()
			}

		}
	} else {
		// 指定连接发送
		for _, userID := range m.ToUserIDs {
			conn := global.GetWsConnByUserID(userID)
			if conn != nil {
				mu.Lock()
				_ = conn.WriteMessage(websocket.TextMessage, respByte)
				mu.Unlock()
			}
		}
	}

}

func (m *messageListener) saveMsg() {
	if m.GrabRecodeUuid == "" {
		return
	}
	var log product.GrabLog
	log.GrabRecordUuid = m.GrabRecodeUuid
	log.Level = m.Code
	log.Msg = m.Msg
	log.CreatedAt = time.Now()

	err := global.XTK_DB.Create(&log).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}
