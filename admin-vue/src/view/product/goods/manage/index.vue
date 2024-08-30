<script setup>
import GoodsDetail from "@/view/product/goods/manage/components/goods-detail.vue";

</script>

<template>
  <div class="card-container">


  <el-card  style="margin-top:10px; width: 90%;">
    <el-scrollbar >
    <goods-detail/>
      </el-scrollbar >
  </el-card>
  </div>
</template>

<style scoped lang="scss">
.card-container {
  display: flex;
  justify-content: center;
}
</style>