<script setup>
import {ref, reactive} from "vue";
import {sendWsMessage} from "@/utils/websocket";
import WarningBar from "@/components/warningBar/warningBar.vue";



const data = reactive({
  begin_date: null,
  count: null,
  amount_from: null,
  amount_to: null,
  effective_day: null,
})




const simulateCouponData = {
  action_name: "simulate_coupon_action",
  data: data
}


function sendSimulateCouponData() {
  sendWsMessage(JSON.stringify(simulateCouponData))
}


</script>

<template>
  <div>

    <el-form>
      <warning-bar title="注：最多生成100条记录"/>
      <el-form-item label="生成数量">
        <el-input v-model.number="data.count"  style="width:200px"/>
      </el-form-item>

      <el-form-item label="优惠券额" inline>
        <el-input  v-model.number="data.amount_from"  style="width:90px; margin-right: 5px"/>
        <span style="width: 10px">~</span>
        <el-input v-model.number="data.amount_to"  style="width:90px; margin-left: 5px"/>
      </el-form-item>


      <span class="demonstration">开始 结束时间为当天23:59:59</span>
      <el-form-item label="开始时间">
          <el-date-picker
              v-model="data.begin_date"
              type="date"
              style="width:200px"
          />
      </el-form-item>

      <el-form-item label="有效天数">
        <el-input v-model.number="data.effective_day" placeholder="天" style="width:200px"/>
      </el-form-item>


      <el-form-item>
        <el-button type="primary" @click="sendSimulateCouponData()">生成</el-button>
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