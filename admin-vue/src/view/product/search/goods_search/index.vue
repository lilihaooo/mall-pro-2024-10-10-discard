<style>


.search-left {
  margin-right: 10px;
}

.search-right {
  background-color: #ddecbf;
  width: 48px;
  position: absolute;
  left:96.8%;
  top: 63px;
  height: calc(100% - 66px);
  z-index: 100;



}

</style>

<template>
  <div>
    <div class="search-left">
      <div style="background-color: #ddecf6">
        <search-input @search="handleSearch"/>
      </div>
      <div style="background-color: #ddecf6">
        <el-scrollbar>
          <search @search="handleSearch"/>
        </el-scrollbar>
      </div>
      <div style=" background-color: #c0deef;" shadow="never">
        <desktop :data="tableData"/>
        <pagination :data="paginationData" @change="handleChange" style="margin-right: 250px"/>
      </div>
    </div>
    <div class="search-right">
      <search-right />
    </div>
  </div>
</template>

<script setup>
import search from "@/view/product/search/goods_search/components/search.vue";
import {ref, reactive, watch} from "vue";
import {getGoodsListV2} from "@/api/product";
import Pagination from "@/view/product/search/goods_search/components/pagination.vue";
import Desktop from "@/view/product/search/goods_search/components/desktop.vue";
import SearchInput from "@/view/product/search/goods_search/components/search-input.vue";
import SearchRight from "@/view/product/search/goods_search/components/search-right.vue";


const handleSearch = async (query) => {
  // 将参数保存到外部
  condition.value = query
  // 发送请求到后端获取数据
  await getTableData(query);
};


const condition = ref({})
const tableData = ref([])


const getTableData = async (query) => {
  const res = await getGoodsListV2(query)
  if (res.code === 0) {
    tableData.value = res.data.list
    // 分页数据
    paginationData.pageSize = res.data.pageSize
    paginationData.total = res.data.total
    paginationData.page = res.data.page
  }
}

getTableData(condition.value)

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

// 我的收藏发生变
const handleCollectChange = (ids) => {
  getTableData(condition.value)
}
</script>


