<script setup>
import Coupon from "@/view/product/data/components/coupon.vue";
import {useWebSocketStore} from "@/pinia";
import {computed} from "vue";
import Tasks from "@/view/product/data/components/tasks.vue";
import {sendWsMessage} from "@/utils/websocket";
// 加载动画
// 使用 computed 包装 grab_live.value 以保持响应式
const webSocketStore = useWebSocketStore();
const simulateCouponLive = computed(() => webSocketStore.simulate_coupon_live);
const taskLive = computed(() => webSocketStore.task_live);



// 发消息询问后端耗时任务是否正在进行
function sendWsMessageForTaskLive() {
  const wsMsg = {
    action_name: "task_live",
  }
  sendWsMessage(JSON.stringify(wsMsg))
}
sendWsMessageForTaskLive()



</script>

<template>
  <el-row :gutter="10">
    <!--        抓取-->
    <el-col :span="6" style="height: 80vh">
      <el-card class="el-card">
        <template #header>
          <div class="card-header">
            <span style="margin-left: 10px; font-size: 18px">模拟优惠券数据</span>

            <div style="margin-right: 30px" v-loading="simulateCouponLive"></div>
          </div>
        </template>
        <div>
          <Coupon/>
        </div>
      </el-card>
    </el-col>

    <el-col :span="6" style="height: 80vh">
      <el-card class="el-card">
        <template #header>
          <div class="card-header">
            <span style="margin-left: 10px; font-size: 18px">耗时任务</span>
            <div style="margin-right: 30px" v-loading="taskLive"></div>
          </div>
        </template>
        <div>
          <Tasks/>
        </div>
      </el-card>
    </el-col>





    </el-row>


</template>

<style scoped lang="scss">
.el-card {
  height: 80vh;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}


</style>