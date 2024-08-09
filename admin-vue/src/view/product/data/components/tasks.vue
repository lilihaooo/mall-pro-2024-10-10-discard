<script setup>
import {ref, reactive} from "vue";
import {sendWsMessage} from "@/utils/websocket";


// 商品绑定优惠券
const data1 = reactive({})
const goodsBandCouponTaskData = {
  action_name: "goods_bind_coupon_task_action",
  data: data1
}

// 商品同步ES
const data2 = reactive({})
const goods2ESTaskData = {
  action_name: "goods_2_es_task_action",
  data: data2
}

// 商品添加标签
const data3 = reactive({})
const goodsAddTagsTaskData = {
  action_name: "goods_add_tags_task_action",
  data: data3
}


// 商品添加标签
const data4 = reactive({})
const simulateLoadsOfGoodsTaskData = {
  action_name: "simulate_loads_of_goods_task_action",
  data: data4
}


function sendGoodsBandCoupon() {
  sendWsMessage(JSON.stringify(goodsBandCouponTaskData))
}

function sendGoods2ES() {
  sendWsMessage(JSON.stringify(goods2ESTaskData))
}

function sendGoodsAddTags() {
  sendWsMessage(JSON.stringify(goodsAddTagsTaskData))
}


function sendSimulateLoadsOfGoods() {
  sendWsMessage(JSON.stringify(simulateLoadsOfGoodsTaskData))
}



</script>

<template>
  <div>

    <el-form>
      <el-form-item label="商品绑定优惠券" label-width="130px">
        <el-button type="primary" @click="sendGoodsBandCoupon()">开始</el-button>
      </el-form-item>

      <el-form-item label="商品添加标签" label-width="130px">
        <el-button type="primary" @click="sendGoodsAddTags()">开始</el-button>
      </el-form-item>

      <el-form-item label="千万商品数据模拟" label-width="130px">
        <el-button type="primary" @click="sendSimulateLoadsOfGoods()">开始</el-button>
      </el-form-item>

      <el-form-item label="商品同步ES" label-width="130px">
        <el-button type="primary" @click="sendGoods2ES()">开始</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<style scoped lang="scss">
.demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 12px;
  margin-left: 88px;
}

</style>