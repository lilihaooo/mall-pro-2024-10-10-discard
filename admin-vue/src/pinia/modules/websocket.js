// src/pinia/modules/websocket.js
import {defineStore} from 'pinia';
import {connectWebSocket, sendWsMessage, setMessageCallback, closeWebSocket} from '@/utils/websocket';
import {ref} from "vue";

export const useWebSocketStore = defineStore('websocket', () => {
    const isConnected = ref(false);
    const grab_log_messages = ref([]);
    const grab_live = ref(false);
    const simulate_coupon_live = ref(false);
    const task_live = ref(false);

    function connect(url) {
        connectWebSocket(url);
        isConnected.value = true;

        setMessageCallback((message) => {
            if (message.for_where === "grab_action_log") {
                grab_log_messages.value.unshift(message);
            }

            if (message.for_where === "grab_live") {
                grab_live.value = !!message.data;
            }

            if (message.for_where === "simulate_coupon_live") {
                simulate_coupon_live.value = !!message.data;
            }


            if (message.for_where === "task_live") {
                task_live.value = !!message.data;
            }

        });
    }

    function clearMessages() {
        grab_log_messages.value.splice(0, grab_log_messages.value.length);
    }


    function close() {
        closeWebSocket();
        isConnected.value = false;
    }

    return {
        isConnected,
        grab_log_messages,
        grab_live,
        simulate_coupon_live,
        task_live,
        connect,
        clearMessages,
        close,
    };
});
