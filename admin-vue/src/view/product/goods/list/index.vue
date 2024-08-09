<template>
  <div>
    <warning-bar title="注：MySQL中查询"/>
    <el-card>
      <el-scrollbar>
        <search @search="handleSearch"/>
      </el-scrollbar>
    </el-card>


    <el-card style="margin-top:10px">
      <goodsTable :data="tableData"/>
      <pagination :data="paginationData" @change="handleChange"/>
    </el-card>

  </div>
</template>

<script setup>
import search from "@/view/product/goods/list/components/search.vue";
import goodsTable from "@/view/product/goods/list/components/goods-table.vue";
import {ref, reactive} from "vue";
import {getGoodsListV1} from "@/api/product";
import Pagination from "@/view/product/search/goods_search/components/pagination.vue";
import WarningBar from "@/components/warningBar/warningBar.vue";

const handleSearch = async (query) => {
  // 将参数保存到外部
  condition.value = query
  // 发送请求到后端获取数据
  await getTableData(query);
};
const condition = ref({})

const tableData = ref([])

const getTableData = async (query) => {
  const res = await getGoodsListV1(query)
  if (res.code === 0) {
    tableData.value = res.data.list || []
    // 分页数据
    paginationData.pageSize = res.data.pageSize
    paginationData.total = res.data.total
    paginationData.page = res.data.page
  }
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
</script>

<style>

</style>
