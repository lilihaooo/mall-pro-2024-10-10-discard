<script setup>

import Grab from "@/view/product/getGoods/components/grab.vue";
import GrabRecordTable from "@/view/product/getGoods/components/grab-record-table.vue";
import Pagination from "@/components/pagination/pagination.vue";
import { getGrabLog, getGrabRecord} from "@/api/product";
import {reactive, ref, computed, watch} from "vue";
import GrabLog from "@/view/product/getGoods/components/grab-log.vue";
import {sendWsMessage} from "@/utils/websocket";
import {useWebSocketStore} from "@/pinia";


const condition = ref({
  page: 1,
  pageSize: 10,
})

const grabRecord = ref([])
const getTableData = async (query) => {
  const res = await getGrabRecord(condition.value)
  grabRecord.value = res.data.list

  // 赋值分页信息
  paginationData.pageSize = res.data.pageSize
  paginationData.total = res.data.total
  paginationData.page = res.data.page
}
getTableData()

const paginationData = reactive({
  pageSize: 10,
  page: 1,
  total: null,
})

// 分页改变触发的
const handleChange = (data) => {
  condition.value.page = data.page
  condition.value.pageSize = data.pageSize
  getTableData(condition.value)
}


// 发消息询问后端抓取任务是否正在进行
function sendWsMessageForGrabLive() {
  const wsMsg = {
    action_name: "grab_action",
    data: {
      option: 2
    }
  }
  sendWsMessage(JSON.stringify(wsMsg))
}

sendWsMessageForGrabLive()


// 加载动画
// 使用 computed 包装 grab_live.value 以保持响应式
const webSocketStore = useWebSocketStore();
const grabLive = computed(() => webSocketStore.grab_live);


watch(grabLive, (v) => {
  if (!v){
    getTableData(condition.value)
  }
})


const grabLog = ref([])
const showLog = async (uuid) => {
  const res = await getGrabLog({uuid: uuid})
  grabLog.value = res.data
}

</script>

<template>


  <el-row :gutter="10">
    <!--        抓取-->
    <el-col :span="6">
      <el-card class="el-card">
        <template #header>
          <div class="card-header">
            <span style="margin-left: 10px; font-size: 18px">抓取</span>
          </div>
        </template>

        <div>
          <Grab/>
        </div>
      </el-card>
    </el-col>


    <!--        记录-->
    <el-col :span="12">
      <el-card class="el-card">
        <template #header>
          <div class="card-header">
            <span style="margin-left: 10px; font-size: 18px">记录</span>
          </div>
        </template>

        <div>
          <grab-record-table :data="grabRecord" @show-log="showLog"/>
          <pagination :data="paginationData" @change="handleChange"/>
        </div>


      </el-card>
    </el-col>


    <!--        日志-->
    <el-col :span="6">
      <el-card class="el-card">
        <template #header>
          <div class="card-header">
            <span style="margin-left: 10px; font-size: 18px">日志</span>
            <div style="margin-right: 40px" v-loading="grabLive"></div>
          </div>
        </template>

        <div>
          <grab-log :data="grabLog"/>
        </div>

      </el-card>
    </el-col>
  </el-row>
  <!--  </div>-->
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