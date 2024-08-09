<script setup>

import {ref, reactive, watch} from "vue";
// 父子通信
const emit = defineEmits(['change']);

const props = defineProps({
  data: {
    type: Object,
    required: true
  }
});

const page = ref(0)
const total = ref(0)
const pageSize = ref(0)

watch(props.data,(newValue)=>{
  page.value = newValue.page
  total.value = newValue.total
  pageSize.value = newValue.pageSize
})


const handleSizeChange = (val) => {
 pageSize.value = val
  emit('change', {page: page.value, pageSize: pageSize.value});

}
const handleCurrentChange = (val) => {
  page.value = val
  emit('change', {page: page.value, pageSize: pageSize.value});
}
</script>

<template>
  <div class="gva-pagination">
    <el-pagination
        :currentPage="page"
        :page-size="pageSize"
        :page-sizes="[10,20, 30, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
    />
  </div>
</template>

<style scoped lang="scss">

</style>