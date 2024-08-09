<script setup>
import {ref} from "vue";
import {getCategoryTree} from "@/api/product";
const tableData = ref([])

const getTableData = async () => {
  const table = await getCategoryTree()
  if (table.code === 0) {
    tableData.value = table.data
  }
}
getTableData()



</script>

<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
            type="primary"
            icon="plus"
            @click="addMenu(0)"
        >新增根菜单</el-button>
      </div>

      <el-table
          :data="tableData"
          style="width: 100%; margin-bottom: 20px"
          row-key="id"
      >
        <el-table-column prop="id" label="ID"  width="120" />
        <el-table-column prop="level" label="级别"  width="80" />
        <el-table-column prop="parent_id" label="父ID"  width="80" />
        <el-table-column prop="name" label="名称"  width="180" />
        <el-table-column prop="status" label="状态"  width="120" />
        <el-table-column label="图标"  width="180" >
          <template #default="scope">
            <img :src="scope.row.image" style="width: 30px"/>
          </template>
        </el-table-column>
        <el-table-column prop="num" label="数量"  width="100" />
        <el-table-column prop="sort" label="排序"  width="80" />
        <el-table-column align="center" label="操作" min-width="50">
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="plus"
                @click="deleteRow(scope.row)"
            >增加子分类
            </el-button>

            <el-button
                type="primary"
                link
                icon="edit"
                @click="deleteRow(scope.row)"
            >修改
            </el-button>

            <el-button
                type="primary"
                link
                icon="delete"
                @click="deleteRow(scope.row)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<style scoped lang="scss">





</style>