<script setup>
import {ref, watch, onMounted} from 'vue';
import {getSuggestionList} from "@/api/product";
import {Close, Edit} from "@element-plus/icons-vue";

const condition = ref({});
const kw = ref('');
const keyword = ref('');
const kwOld = ref('');
// 显示联想框
const downListShow = ref(false);
const recommendShow = ref(false);

const emit = defineEmits(['search']);
const onSearch = () => {
  emit('search', condition.value);
};
const submitKG = () => {
  if (typeof kw.value === 'string' && kw.value.trim() !== '') {
    condition.value.keyword = kw.value.trim();
    onSearch()
    // 向历史kw的数组中添加元素
    if (kw.value !== kwOld.value) {
      if (historyKW.value.length >= 20) {
        historyKW.value.pop();
      }
      historyKW.value = historyKW.value.filter(item => item !== kw.value); // 先删除已存在的元素
      historyKW.value.unshift(kw.value); // 再添加到头部
     // 将数组转换为字符串并保存到 localStorage
      localStorage.setItem('historyKW', JSON.stringify(historyKW.value));
      kwOld.value = kw.value;
    }
  }

};

const suggestionList = ref([])
watch(kw, (newV) => {
  if (newV.trim() !== "") {
    keyword.value = newV;
    console.log(keyword.value);
    recommendShow.value =false
    getSuggestionData({keyword: newV});
  } else {
    suggestionList.value = []
    downListShow.value = false

  }
})


watch(suggestionList, (newV) => {
  console.log(newV)
  if (newV != null) {
    if (newV.length > 0) {
      downListShow.value = true
      recommendShow.value = false
    } else {
      recommendShow.value = true
      downListShow.value = false
    }
  } else {
    downListShow.value = false
  }
})

const getSuggestionData = async (query) => {
  const res = await getSuggestionList(query)
  // 填充推荐列表
  suggestionList.value = res.data
}

const searchItem = (item) => {
  kw.value = String(item);
  submitKG()
}
const isFocus = ref(false)
const onFocus = () => {
  isFocus.value = true
  if (suggestionList.value == null) {
      recommendShow.value = true
  } else {
    if (suggestionList.value.length > 0) {
      downListShow.value = true
    } else {
        recommendShow.value = true
    }
  }
}
const isMouseInside = ref(false);
// 输入框失去焦点时
const onBlur = () => {
  isFocus.value = false
  setTimeout(() => {
    if (!isMouseInside.value && !isFocus.value) {
      downListShow.value = false;
      recommendShow.value = false;
    }
  }, 100); // 延迟关闭，确保鼠标进入联想框时不立即消失
};

// 鼠标进入联想框或推荐框
const onMouseEnter = () => {
  isMouseInside.value = true;
};

// 鼠标离开联想框或推荐框
const onMouseLeave = () => {
  isMouseInside.value = false;
  if (!isFocus.value){
    downListShow.value = false;
    recommendShow.value = false;
  }
};
const historyKW = ref([])
onMounted(() => {
  // 组件挂载后的操作
  // 从 localStorage 获取并解析数组
  const storedHistory = JSON.parse(localStorage.getItem("historyKW")) || [];
  // 如果存在有效数据，则推入到 historyKW.value 中
  historyKW.value.push(...storedHistory);
});

const inputRef = ref(null); // 引用输入框
const clearInput = () => {
  kw.value = ''; // 清空输入框内容kw.value = ''; // 清空输入框内容
  inputRef.value.focus(); // 重新聚焦输入框，保留光标
};


</script>


<template>
  <div>
    <div class="search-container-box">
      <div class="search-box">
        <div class="search-container">
          <input
              ref="inputRef"
              type="text"
              v-model="kw"
              class="search-input"
              placeholder="商品/店铺/分类"
              @keyup.enter="submitKG"
              @focus="onFocus"
              @blur="onBlur"

          />
          <el-button v-if="kw" class="clear-btn" size="large" type="primary"  @click="clearInput"  link :icon="Close"/>
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
        <!-- 联想框 -->
        <div
            class="search-container-down"
            v-show="downListShow"
            @mouseenter="onMouseEnter"
            @mouseleave="onMouseLeave"
        >
          <div style=" width: 100%;">
            <el-scrollbar>
              <p v-for="item in suggestionList" :key="item"  class="scrollbar-demo-item" @click="searchItem(item)">
                <span style="margin-left: 12px">
                    {{ item }}
                  </span>
              </p>
            </el-scrollbar>
          </div>

        </div>
        <!-- 推荐框 -->
        <div
            class="search-recommend-down"
            v-show="recommendShow"
            @mouseenter="onMouseEnter"
            @mouseleave="onMouseLeave"
        >
          <!-- 排行榜-->
          <div style=" width: 100%; border-right: 1px solid #ddd;position: relative">
            <div class="recommend-scrollbar-title">
              实时热搜
            </div>
            <el-scrollbar height="300px" style="margin-top: 20px">
              <p v-for="item in 20" class="scrollbar-demo-item" @click="searchItem(item)">
                  <span style="margin-left: 12px">
                    {{ item }}
                  </span>
              </p>
            </el-scrollbar>
          </div>
          <!-- 搜索历史-->
          <div style=" width: 100%; border-right: 1px solid #ddd;position: relative">
            <div class="recommend-scrollbar-title">
              搜索历史
            </div>

            <el-scrollbar height="300px" style="margin-top: 20px">
              <p v-for="(item, index) in historyKW" :key="index" class="scrollbar-demo-item">
                 <span style="margin-left: 12px">
                    {{ item }}
                  </span>
              </p>
            </el-scrollbar>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped lang="scss">
.recommend-scrollbar-title {
  position: absolute;
  top: 2px;
  left: 26px;
  font-weight: bold;
  line-height: 20px; /* 使文字垂直居中，如果容器高度固定 */
  height: 20px;
  color: var(--el-color-primary);
}

.scrollbar-demo-item {
  display: flex;
  height: 20px;
  margin: 6px 16px;
  align-items: center; /* 垂直居中子元素（文本） */
  border-radius: 2px;
  color: var(--el-color-primary);
}
.scrollbar-demo-item:hover {
  background-color: #ddd; /* 鼠标悬停时的背景色 */
  cursor: pointer; /* 鼠标悬停时显示为指针 */
}

.search-input {
  height: 18px;
  margin-left: 10px;
}

.clear-btn {
  position: absolute;
  left: 520px;
  top: 14px;
  transform: translateY(-50%);
  padding: 0;
}

.search-recommend-down {
  display: flex;
  position: absolute;
  width: 568px;
  height: 326px;
  margin-left: 25px;
  border-radius: 0 0 5px 5px;
  z-index: 100; /* Higher z-index to float above other elements */
  background-color: white;
}

.search-container-down {
  display: flex;
  position: absolute;
  width: 568px;
  margin-left: 25px;
  border-radius: 0 0 5px 5px;
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
  position: relative;
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

