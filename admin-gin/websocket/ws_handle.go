package websocket

import (
	"admin-gin/global"
	"admin-gin/utils"
	"admin-gin/websocket/common"
	"admin-gin/websocket/internal/action"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(c *gin.Context) {
	// 创建连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	// 终止程序时关闭连接
	signalCh := make(chan os.Signal, 1)
	go func() {
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM) //前者通常由用户按 Ctrl + C 产生，后者通常用于请求程序正常终止 情况后往通道发送信号
		<-signalCh                                               // 接受到信号 取消阻塞
		err = conn.Close()                                       // 关闭ws连接
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return
		}
	}()
	// 获得用户信息， 将连接写入syncMap中
	userID := utils.GetUserID(c)
	if ok := global.SetWsConnByUserID(userID, conn); !ok {
		global.GVA_LOG.Error("加入连接wsMap失败")
		return
	}

	global.GVA_LOG.Info("websocket 连接成功")

	//listener  //每一个连接就监听一个连接， 这样做可以同时监听多个连接
	wsReader(userID)
}

type WsReceiveMessage struct {
	ActionName string      `json:"action_name"`
	Data       interface{} `json:"data"`
}

func wsReader(userID uint) {
	conn := global.GetWsConnByUserID(userID)
	if conn == nil {
		global.GVA_LOG.Error("连接不存在")
		return
	}
	for {
		// 读取消息
		_, massage, err := conn.ReadMessage()
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			conn.Close() // 出现错误时关闭连接
			return
		}
		var wsReceive WsReceiveMessage
		err = json.Unmarshal(massage, &wsReceive)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}

		// 分发action
		// 抓取数据
		if wsReceive.ActionName == "grab_action" {
			go action.GrabAction(wsReceive.Data, userID)
		}

		if wsReceive.ActionName == "simulate_coupon_action" {
			go action.SimulateCouponAction(wsReceive.Data, userID)
		}

		if wsReceive.ActionName == "goods_add_tags_task_action" {
			go action.GoodsAddTagsAction(userID)
		}

		if wsReceive.ActionName == "goods_bind_coupon_task_action" {
			go action.GoodsBandCouponAction(userID)
		}

		if wsReceive.ActionName == "simulate_loads_of_goods_task_action" {
			go action.SimulateLoadsOfGoodsAction(userID)
		}

		// 检查耗时任务的生命
		if wsReceive.ActionName == "task_live" {
			_, ok := global.LockMap.Load("task_lock")
			common.WsSendMsg(0, "", ok, "task_live", []uint{}, "")
		}
	}
}
