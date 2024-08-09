<script setup>
import { useWebSocketStore } from "@/pinia";
import { computed } from "vue";

const webSocketStore = useWebSocketStore();
const log_msg = webSocketStore.grab_log_messages;

// 抓取时禁用按钮
const grabLive = computed(() => webSocketStore.grab_live);

// 父组件传值
const props = defineProps({
  data: {
    type: Array,
    required: true
  }
});
</script>

<template>
  <div>
    <transition name="slide-fade" mode="out-in">
      <el-scrollbar v-if="grabLive" key="grab-live" class="scrollbar-container now" style="height: 66vh;">
        <p v-for="item in log_msg" :key="item.id" class="scrollbar-demo-item">{{ item.msg }}</p>
      </el-scrollbar>

      <el-scrollbar v-else key="data" class="scrollbar-container one" style="height: 66vh;">
        <p v-for="item in data" :key="item.id" class="scrollbar-demo-item">{{ item.msg }}</p>
      </el-scrollbar>
    </transition>
  </div>
</template>

<style scoped lang="scss">
.slide-fade-enter-active.now {
  transition: all 0s ease-out;
}
.slide-fade-leave-active.now{
  transition: all 1s ease-out;
}



.slide-fade-enter-active.one {
  transition: all 1s ease-out;
}
.slide-fade-leave-active.one{
  transition: all 0s ease-out;
}






.slide-fade-enter-from, .slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

.scrollbar-container {
  position: relative;
  width: 100%;
}

.scrollbar-demo-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 42px;
  margin: 10px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
}
</style>
