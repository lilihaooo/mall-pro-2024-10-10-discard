<template>
  <div id="app" class="bg-gray-50 text-slate-700 dark:text-slate-500  dark:bg-slate-800">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import {useRoute} from "vue-router";
import {useAppStore} from "@/pinia";
import {watch} from "vue";
const route = useRoute();
const appStore = useAppStore();



watch(route, (newRoute) => {
  if (newRoute.name === 'search') {
    appStore.config.showTabs = false
  } else {
    appStore.config.showTabs = true
  }
});


// element 2.3.8之前使用下面的语句
// import zhCn from 'element-plus/lib/locale/lang/zh-cn'
  defineOptions({
  name: 'App'
})


</script>
<style lang="scss">

// 引入初始化样式
#app {
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button{
  font-weight: 400 !important;
}

.gva-body-h{
  min-height: calc(100% - 3rem);
}

.gva-container{
  height: calc(100% - 2.5rem);
}
.gva-container2{
  height: calc(100% - 4.5rem);
}

</style>
