<script setup>
import {reactive, ref} from "vue";
import {getShopList} from "@/api/product";
import ShopTable from "@/view/product/shop/list/components/shop-table.vue";
import Pagination from "@/view/product/search/goods_search/components/pagination.vue";

const tableData = ref([])
const condition = ref({})

const getTableData = async (query) => {
  const res = await getShopList(query)
    tableData.value = res.data.list
    // 分页数据
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


</script>

<template>
  <div>
    <div class="gva-table-box">
      <shop-table :data="tableData"></shop-table>
      <pagination :data="paginationData" @change="handleChange"/>
    </div>
  </div>
</template>

<style scoped lang="scss">


</style>