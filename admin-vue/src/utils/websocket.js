// src/utils/websocket.js
import {ElMessage} from "element-plus";
// import {useWebSocketStore} from "@/pinia";

let socket;
let messageCallback;



export function connectWebSocket(url) {
    socket = new WebSocket(url);

    socket.onopen = () => {
        console.log('WebSocket connection established');
    };

    socket.onmessage = (event) => {

        // console.log(event.data);
        // 解析消息并根据内容执行操作
        const parsedMessage = JSON.parse(event.data)

        // 处理消息
        // code = 1(error) parsedMessage.for_action_name === "" 弹层提示消息 直接处理
        if (parsedMessage.code === 1 && parsedMessage.for_where === "ElMessage") {
            ElMessage({
                message: parsedMessage.msg,
                type: "error",
            })
        }else if (parsedMessage.code === 0 && parsedMessage.for_where === "ElMessage") {
            ElMessage({
                message: parsedMessage.msg,
                type: "success",
            })
        }else  {
            if (messageCallback) {
                messageCallback(parsedMessage);
            }
        }





    };

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    socket.onerror = (error) => {
        console.error('WebSocket error:', error);
    };
}

export function sendWsMessage(message) {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(message);
    } else {
        console.error('WebSocket is not open');
        ElMessage.error('websocket未连接, 请刷新后再试')
    }
}

export function setMessageCallback(callback) {
    messageCallback = callback;
}

export function closeWebSocket() {
    if (socket) {
        socket.close();
    }
}
