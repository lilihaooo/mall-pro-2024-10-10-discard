<script setup>
import router  from "@/router";
const props = defineProps({
  data: {
    type: Array,
    required: true
  }
});


const manageShop =(id)=>{
  router.push({
    name:'shopManage',
    query:{id:id},
  })

}






</script>

<template>
  <div>
    <el-table
        :data="data"
        style="width: 100%; margin-bottom: 20px"
        row-key="id"
        class="shop-list-table"
    >
      <el-table-column prop="id" label="ID" width="120"/>
      <el-table-column prop="name" label="店铺名称" width="240"/>


      <el-table-column prop="status" label="状态" width="120"/>

      <el-table-column label="Logo" width="200">
        <template #default="scope">
          <img :src="scope.row.logo" style="width: 80px"/>
        </template>
      </el-table-column>

      <el-table-column prop="logistics_score" label="物流分" width="120"/>
      <el-table-column prop="product_score" label="商品分" width="120"/>
      <el-table-column prop="service_score" label="服务分" width="120"/>
      <el-table-column align="center" label="操作" min-width="50">
        <template #default="scope">
          <el-button
              type="primary"
              link
              icon="plus"
              @click="manageShop(scope.row.id)"
          >管理店铺
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

</template>

<style scoped lang="scss">
.shop-list-table {
  height: 66vh;
}

</style>