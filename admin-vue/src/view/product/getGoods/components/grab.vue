<script setup lang="ts">
import {reactive} from "vue";
import {sendWsMessage} from "@/utils/websocket";
import {ref} from 'vue'
import {ElTree} from 'element-plus'
import {getCategoryTree} from "@/api/product";
import WarningBar from "@/components/warningBar/warningBar.vue";
import { useWebSocketStore } from "@/pinia";
// 使用 computed 包装 grab_live.value 以保持响应式
const webSocketStore = useWebSocketStore();



interface Data {
  option: number
  page_num: number
  go_num: number
  category_ids: number[]
}

interface SendWsMassage {
  action_name: string,
  data: Data,
}

const data = reactive<Data>({
  option: 0,
  page_num: 200,
  go_num: 5,
  category_ids: [],
})


const stopWsMsg: SendWsMassage = {
  action_name: "grab_action",
  data: data,
}
const beginWsMsg: SendWsMassage = {
  action_name: "grab_action",
  data: data,
}


function sendBegin() {
  webSocketStore.clearMessages()
  data.option = 1
  data.category_ids = treeRef.value!.getCheckedKeys(false) as number[]
  sendWsMessage(JSON.stringify(beginWsMsg))
}

function sendStop() {
  data.option = 0
  sendWsMessage(JSON.stringify(stopWsMsg))

}


// tree
const defaultProps = {
  children: 'children',
  label: 'name',
}

interface Tree {
  id: number
  label: string
  children?: Tree[]
}


const treeData = ref([] as Tree[])

const treeRef = ref<InstanceType<typeof ElTree>>()

const resetChecked = () => {
  treeRef.value!.setCheckedKeys([], false)
  data.option = 0
  data.category_ids = []
  data.go_num = 5
  data.page_num = 200
}


interface Res {
  code: number
  data: any
}

// declare关键字用于告诉编译器关于已经存在的变量、函数、类或模块的类型信息，而不需要实际的实现代码
declare function getCategoryTree(): Promise<Res>;

// 获取分类信息
const getCategoryData = async () => {
  const res = await getCategoryTree()
  if (res.code === 0) {
    treeData.value = res.data

  }
}
getCategoryData()


</script>

<template>
  <div>

    <el-form

        @keyup.enter="sendBegin"
    >
      <el-form-item>
        <el-button type="primary" @click="sendBegin()">抓取</el-button>
        <el-button type="warning" @click="sendStop()">停止</el-button>
      </el-form-item>

      <el-form-item label="并发数量">
        <el-input v-model.number="data.go_num" placeholder="并发量" style="width:200px"/>
      </el-form-item>

      <el-form-item label="分页深度">
        <el-input v-model.number="data.page_num" placeholder="每个分类抓取页数" style="width:200px"/>
      </el-form-item>



      <warning-bar title="注：不选则抓取全部分类"/>
      <el-form-item label="选择分类">
        <el-scrollbar height="280px">
          <el-tree
              ref="treeRef"
              :data="treeData"
              show-checkbox
              :default-expand-all="false"
              node-key="id"
              highlight-current
              :props="defaultProps"
          />
        </el-scrollbar>
      </el-form-item>

      <el-form-item>
        <div class="buttons">
          <el-button @click="resetChecked">重置</el-button>
        </div>
      </el-form-item>


    </el-form>


  </div>
</template>

<style scoped lang="scss">

</style>