<script setup>

import {useWebSocketStore} from "@/pinia";
import {computed, ref} from "vue";

const emit = defineEmits(['showLog']);

const props = defineProps({
  data: {
    type: Array,
    required: true
  }
});
const NowUuid = ref("");
const showLogs= (uuid)=>{
  NowUuid.value = uuid;
  emit('showLog',uuid)
}

// 抓取时禁用按钮
const webSocketStore = useWebSocketStore();
const grabLive = computed(() => webSocketStore.grab_live);


</script>

<template>
  <div>
    <el-scrollbar height="58vh">
    <el-table
        ref="multipleTable"
        style="width: 98%"
        tooltip-effect="dark"
        :data="data"
        row-key="ID"
    >
      <el-table-column align="center" label="用户名" min-width="120" prop="user_name"/>
      <el-table-column align="center" label="分类数量" width="100" prop="category_num"/>
      <el-table-column align="center" label="携程数" width="80" prop="go_num"/>
      <el-table-column align="center" label="分页深度" width="100" prop="page_num"/>
      <el-table-column align="center" label="运行时间" prop="run_time" width="160"/>
      <el-table-column align="center" label="抓取量" width="200" prop="grab_num"/>
      <el-table-column align="center" label="时间" min-width="200" prop="created_time"/>
      <el-table-column align="center" label="操作" fixed="right" min-width="120">
        <template #default="scope">
          <el-button
              type="primary"
              :class="{ 'active-button': scope.row.uuid === NowUuid }"
              @click="showLogs(scope.row.uuid)"
              :disabled ="grabLive"
              plain
          >
            <el-icon><DocumentCopy /></el-icon>日志
          </el-button>
        </template>

      </el-table-column>


    </el-table>
    </el-scrollbar>
  </div>

</template>

<style scoped lang="scss">
.active-button {
  background: var(--el-color-primary-light-1);
  color: white;
}
</style>