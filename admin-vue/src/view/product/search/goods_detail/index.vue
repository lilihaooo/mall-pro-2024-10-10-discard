<script setup>
import GoodsDetail from "@/view/product/search/goods_detail/components/goods-detail.vue";

</script>

<template>
  <div class="card-container">
  <el-card style="margin-top:10px; width: 80%; background-color: #ddecf6;" shadow="never">
    <goods-detail/>
  </el-card>
  </div>
</template>

<style scoped lang="scss">
.card-container {
  display: flex;
  justify-content: center;
  background-color: #ddecf6;
}
</style>
