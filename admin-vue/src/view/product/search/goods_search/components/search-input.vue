<script setup>
import { ref,watch } from 'vue';
import {getGoodsListV2, getSuggestionList} from "@/api/product";

const condition = ref({});
const kw = ref('');
const keyword = ref('');
const downListShow = ref(false);

const emit = defineEmits(['search']);
const onSearch = () => {
  emit('search', condition.value);
};
const submitKG = () => {
  if (typeof kw.value === 'string' && kw.value.trim() !== '') {
    condition.value.keyword = kw.value.trim();
  }
  onSearch()
};

watch(kw, (newV) => {
  if (newV.trim() !== ""){
    keyword.value = newV;
    console.log(keyword.value);
    getSuggestionData({keyword: newV});
  }else {
    downListShow.value = false
  }
})

const suggestionList = ref([])
watch(suggestionList, (newV) => {
  console.log(newV)
  if (newV != null){
    if (newV.length > 0) {
      downListShow.value= true
    }else{
      downListShow.value = false
    }
  }else {
    downListShow.value = false
  }
})

const getSuggestionData = async (query) => {
  const res = await getSuggestionList(query)
  // 填充推荐列表
  suggestionList.value = res.data
}


const searchItem = (item)=>{
  kw.value = item
  submitKG()
}





</script>



<template>
  <div>
    <div class="search-container-box">
      <div class="search-box">
        <div class="search-container">
          <input
              type="text"
              v-model="kw"
              class="search-input"
              placeholder="商品/店铺/分类"
              @keyup.enter="submitKG"
          >
          <el-button
              type="primary"
              :round="true"
              icon="search"
              :plain="false"
              :shadow="true"
              @click="submitKG"
              style="width: 32px;"
          >
          </el-button>
        </div>
        <div
            class="search-container-down"
            v-show="downListShow"
        >
          <div>
            <el-scrollbar>
              <p v-for="item in suggestionList" :key="item" class="scrollbar-demo-item">
                <el-button type="primary" link @click="searchItem(item)">
                  {{ item }}
                </el-button>
              </p>
            </el-scrollbar>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>



<style scoped lang="scss">


.scrollbar-demo-item {
  display: flex;
  height: 20px;
  margin: 6px 16px;
  text-align: center;
  border-radius: 4px;
  color: var(--el-color-primary);
}



.search-input {
  height: 18px;
  margin-left: 10px;
}

.search-container-down {
  display: flex;
  position: absolute;

  width: 568px;
  margin-left: 25px;
  border-radius: 5px;
  z-index: 10; /* Higher z-index to float above other elements */
  background-color: white;
}

.search-container-box {
  display: flex;
  justify-content: center;
  width: 100%;
  margin-top: -10px;
  padding-top: 15px;
}

.search-box {
  width: 620px;
  position: relative;
}

.search-container {
  display: flex;
  align-items: center;
  margin: 0 auto;
  border: 2px solid var(--el-color-primary);
  border-radius: 50px;
  padding: 0 4px;
  height: 48px;
  color: #999;
}

.search-container input[type="text"] {
  border: none;
  outline: none;
  width: 100%;
  padding: 10px;
  border-radius: 50px;
  font-size: 18px;
  color: #888;
}

.search-container input[type="text"]::placeholder {
  color: #ccc;
}
</style>

