<template>
  <div>
    <!--    搜索框-->
    <el-form ref="elSearchFormRef" :inline="true" @keyup.enter="submitKG" style="white-space: nowrap;">
      <el-form-item>
        <el-input v-model.number="goodsID" placeholder="商品ID"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="search" @click="submitKG">ID查询</el-button>
      </el-form-item>
    </el-form>
    <!--    分类/筛选/排序-->
    <table style="white-space: nowrap;width: 90%">
      <!--      分类-->
      <tr style="height: 35px; border-top: 1px solid">
        <td>&nbsp&nbsp分类:</td>
        <td>
          <el-popover
              v-for="(category, i) in categoryData" :key="i"
              placement="bottom-start"
              width="50%"
              popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
              :show-after="500"
          >
            <template #reference>
              <el-button
                  link
                  style="margin-left: 20px"
                  :style="{ color: isCategorySelected(category.id) ? 'var(--el-color-primary)' : '' }"
                  @click="setCategoryId(category.id)"
              >
                {{ category.name }}
              </el-button>
            </template>
            <template #default>
              <div v-for="(subCategory, subIndex) in category.children" :key="subIndex" style="margin-bottom: 15px">
                <el-button
                    link
                    @click="setCategoryId(subCategory.id)"
                    :style="{ color: isCategorySelected(subCategory.id) ? 'var(--el-color-primary)' : '' }"
                >
                  <b>{{ subCategory.name }}</b>
                </el-button>
                <div style="width: 90%">
            <span v-for="(subSubCategory, subSubIndex) in subCategory.children" :key="subSubIndex">
              <el-button
                  link
                  @click="setCategoryId(subSubCategory.id)"
                  size="small"
                  style="margin-left: 2px"
                  :style="{ color: isCategorySelected(subSubCategory.id) ? 'var(--el-color-primary)' : '' }"
              >
                {{ subSubCategory.name }}&nbsp;
              </el-button>
            </span>
                </div>
              </div>
            </template>
          </el-popover>
        </td>
      </tr>
      <!--      筛选-->
      <tr style="height: 35px;border-top: 1px solid">
        <td>&nbsp&nbsp筛选:</td>
        <td>
            <div style="height:35px; display: flex; align-items: center; color: #606266; font-size: 14px">
            <span style="margin-left: 22px">
            券后价
            <el-input size="small" v-model="filt.qhj_f" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
            ~
            <el-input size="small" v-model="filt.qhj_t" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
            </span>
              <span style="margin-left: 10px">
           佣金比例
            <el-input placeholder="%" size="small" v-model="filt.cr_f" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
            ~
            <el-input placeholder="%" size="small" v-model="filt.cr_t" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          </span>
              <span style="margin-left: 10px">
          佣金  ≥
          <el-input size="small" v-model="filt.cv" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          </span>
              <span style="margin-left: 10px">
          销量  ≥
          <el-input size="small" v-model="filt.sale_num" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          </span>
              <span style="margin-left: 10px">
          券面额
          <el-input size="small" v-model="filt.coupon_v_f" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          ~
          <el-input size="small" v-model="filt.coupon_v_t" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          </span>

              <span style="margin-left: 10px">
          券到期时间  ≥
          <el-input size="small" v-model="filt.qgt" type="number" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          </span>

              <span style="margin-left: 10px">
            体验分  ≥
          <el-input size="small" v-model="filt.experience_score" :class="config.side_mode ==='normal'? 'input_tight':'input_loose'"/>
          </span>

              <span style="margin-left: 10px">
            <el-button type="primary" size="small" @click="onFilt">筛选</el-button>
          </span>
            </div>



        </td>
      </tr>

      <tr style="height: 35px; border-top: 1px solid;">
        <td>&nbsp;&nbsp;排序:</td>
        <td>
          <el-button
              link
              v-for="(sortMethod, index) in sortMethods"
              :key="index"
              @click="setActiveSort(sortMethod)"
              :style="{ marginLeft: '20px', color: activeSort === sortMethod.method ? 'var(--el-color-primary)' : null }"
          >
            {{ sortMethod.label }}
          </el-button>

          <el-button
              link
              @click="setActiveSort2(  {method: 'sortQHJ'})"
          >
            <span :style="{ marginLeft: '20px', color: activeSort === 'sortQHJ' ? 'var(--el-color-primary)' : null }">券后价</span>
            <div class="icon-container">
              <el-icon :style="{ color: sortType === 'asc' ? 'var(--el-color-primary)' : null }">
                <CaretTop/>
              </el-icon>
              <el-icon :style="{ color: sortType === 'desc' ? 'var(--el-color-primary)' : null }">
                <CaretBottom/>
              </el-icon>
            </div>
          </el-button>


        </td>
      </tr>
    </table>
    <!--    多选框-->
    <table style="white-space: nowrap;">
      <tr style="height: 35px; border-top: 1px solid;">
        <td>
          <div style="display: flex; align-items: center; width: 100%">
            <el-checkbox :true-value="1"  v-model="condition.is_choice"   label="精选推荐" size="large"/>
            <el-checkbox :true-value="1"  v-model="condition.jrsx"   label="今日上新" size="large"/>
            <el-checkbox :true-value="0.01" v-model="condition.couv_f"  label="只看有券" size="large"/>
            <el-checkbox :true-value="1" v-model="condition.is_brand" label="品牌库" size="large"/>
            <el-checkbox :true-value="1" v-model="condition.absolute_low_price" label="低价好卖" size="large"/>
            <el-checkbox :true-value="1" v-model="condition.buyer_reviews" label="买家好评" size="large"/>



            <!-- 数据来源-->
            <el-popover
                placement="bottom"
                :width="140"
                trigger="hover"
                :show-after="500"
            >
              <template #reference>
                <span class="popover-reference" :style="{color: dataFromShow ? 'var(--el-color-primary)' : null }"
                > {{ dataFromName }}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-radio-group v-model="condition.data_from" class="ml-4">
                <el-radio :value="1">爬取</el-radio>
                <el-radio :value="2">模拟</el-radio>
                <el-radio :value="3">输入</el-radio>
              </el-radio-group>
              <el-button icon="Delete"  link style="margin-left: 20px" type="primary" size="small" @click="condition.data_from=null">
                清除
              </el-button>
            </el-popover>

            <!-- 平台-->
            <el-popover
                placement="bottom"
                :width="140"
                trigger="hover"
                :show-after="500"
            >
              <template #reference>
                <span class="popover-reference" :style="{color: platformShow ? 'var(--el-color-primary)' : null }"
                > {{ platformName }}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-radio-group v-model="condition.platform" class="ml-4">
                <el-radio :value="1">抖音</el-radio>
                <el-radio :value="2">快手</el-radio>
                <el-radio :value="3">微信</el-radio>
              </el-radio-group>
              <el-button icon="Delete"  link style="margin-left: 20px" type="primary" size="small" @click="condition.platform=null">
                清除
              </el-button>
            </el-popover>
            <!-- 推广素材-->
            <el-popover
                placement="bottom"
                :width="140"
                trigger="hover"
                :show-after="500"
            >
              <template #reference>
                <span class="popover-reference" :style="{color: pushSucaiShow ? 'var(--el-color-primary)' : null }"
                > {{ pushSucaiName }}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-checkbox-group v-model="pushSucai">
                <el-checkbox :value="1" size="large">精推素材</el-checkbox>
                <el-checkbox :value="2" size="large">朋友圈素材</el-checkbox>
                <el-checkbox :value="3" size="large">短视频</el-checkbox>
              </el-checkbox-group>


            </el-popover>
            <!-- 服务保障-->
            <el-popover
                placement="bottom"
                :width="160"
                trigger="hover"
                :show-after="500"
                persistent
            >
              <template #reference>
                <span class="popover-reference"
                      :style="{color: serviceGuaranteeShow ? 'var(--el-color-primary)' : null }"
                > {{ serviceGuaranteeName }}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-checkbox :true-value="1" v-model="condition.yfx" label="运费险" size="large"/>
              <el-checkbox :true-value="1" v-model="condition.hbmx" label="花呗免息" size="large"/>
              <el-checkbox :true-value="1" v-model="condition.psbz" label="破损保障" size="large"/>

              <!-- 偏远地区包邮-->
              <el-popover
                  placement="bottom"
                  :width="60"
                  trigger="hover"
                  :show-after="500"
              >
                <template #reference>
                  <span class="popover-reference" :style="{color: pydqbyShow ? 'var(--el-color-primary)' : null }"> 偏远地区包邮&nbsp;<el-icon><ArrowDown/></el-icon></span>
                </template>
                <el-checkbox-group v-model="area">
                  <el-checkbox :value="1" size="large">西藏</el-checkbox>
                  <el-checkbox :value="2" size="large">新疆</el-checkbox>
                  <el-checkbox :value="3" size="large">内蒙</el-checkbox>
                  <el-checkbox :value="4" size="large">甘肃</el-checkbox>
                  <el-checkbox :value="5" size="large">宁夏</el-checkbox>
                  <el-checkbox :value="6" size="large">海南</el-checkbox>
                  <el-checkbox :value="7" size="large">青海</el-checkbox>
                </el-checkbox-group>
                <el-button size="small" type="primary" @click="areaCommit">确定</el-button>
              </el-popover>
            </el-popover>
            <!-- 店铺类型-->
            <el-popover
                placement="bottom"
                :width="140"
                trigger="hover"
                :show-after="500"
            >
              <template #reference>
                <span class="popover-reference" :style="{color: shopShow ? 'var(--el-color-primary)' : null }"> {{shopName}}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-radio-group v-model="condition.shop" class="ml-4">
                <el-radio :value="1">官方旗舰店</el-radio>
                <el-radio :value="2">阿里健康大药房</el-radio>
                <el-radio :value="3">天猫国际</el-radio>
                <el-radio :value="4">进口超市</el-radio>
                <el-radio :value="5">淘宝冠店</el-radio>
                <el-radio :value="6">淘宝企业店铺</el-radio>
              </el-radio-group>
              <el-button  icon="Delete"  link style="margin-left: 20px" type="primary" size="small" @click="condition.shop=null">清除
              </el-button>
            </el-popover>

            <!-- 商品活动-->
            <el-popover
                placement="bottom"
                :width="140"
                trigger="hover"
                :show-after="500"
            >
              <template #reference>
                <span class="popover-reference" :style="{color: activeShow ? 'var(--el-color-primary)' : null }"
                > {{ activeName }}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-checkbox-group v-model="activeGroup">
                <el-checkbox :value="1" size="large">定向计划</el-checkbox>
                <el-checkbox :value="2" size="large">聚划算</el-checkbox>
                <el-checkbox :value="3" size="large">拍多件</el-checkbox>
                <el-checkbox :value="4" size="large">前N件</el-checkbox>
                <el-checkbox :value="5" size="large">额外满减</el-checkbox>
                <el-checkbox :value="6" size="large">买赠活动</el-checkbox>
                <el-checkbox :value="7" size="large">首单礼金</el-checkbox>
              </el-checkbox-group>
            </el-popover>

            <!-- 商品表现-->
            <el-popover
                placement="bottom"
                :width="140"
                trigger="hover"
                :show-after="500"
            >
              <template #reference>
                <span class="popover-reference" :style="{color: biaoxianShow ? 'var(--el-color-primary)' : null }"> {{biaoxianName}}&nbsp;<el-icon><ArrowDown/></el-icon></span>
              </template>
              <el-radio-group v-model="condition.bx_bdan" class="ml-4">
                <el-radio :value="1">实时在榜</el-radio>
                <el-radio :value="2">30天上榜</el-radio>
              </el-radio-group>
              <el-button icon="Delete"  link style="margin-left: 20px" type="primary" size="small" @click="condition.bx_bdan=null">清除
              </el-button>

              <el-radio-group v-model="condition.bx_type" class="ml-4">
                <el-radio :value="1">30天新品</el-radio>
                <el-radio :value="2">30天低价</el-radio>
                <el-radio :value="3">30天高佣金</el-radio>
              </el-radio-group>
              <el-button icon="Delete"  link style="margin-left: 20px" type="primary" size="small" @click="condition.bx_type=null">清除
              </el-button>
            </el-popover>


          </div>
        </td>
      </tr>
    </table>
  </div>

</template>

<script setup>
import {ref, reactive, watch} from "vue";
import {getCategoryTree} from "@/api/product";
import {storeToRefs} from "pinia";
import {useAppStore} from "@/pinia";

const appStore = useAppStore();
const {config} = storeToRefs(appStore);

// 父子通信
const emit = defineEmits(['search']);
const onSearch = () => {
  emit('search', condition);
};


const goodsID = ref("");
const submitKG = () => {
  if (typeof goodsID.value === "number" && goodsID.value > 0) {
    condition.g_id = goodsID.value;
  }
};

// 多选
const area = ref([])  // 偏远地区包邮
const areaCommit = () => {
  if (area.value.length > 0) {
    condition.area = area.value.join(',');
  }else{
    condition.area = null
  }
}
// 数组参数
const pushSucai = ref([])
const activeGroup = ref([])

watch(pushSucai, val => {
  if (val.length > 0) {
    condition.push_sucai = val.join(',');
  }else{
    condition.push_sucai = null
  }
})
watch(activeGroup, val => {
  if (val.length > 0) {
    condition.av_group = val.join(',');
  }else{
    condition.av_group = null
  }
})




// 分类
const categoryData = ref([])
const getCategoryData = async () => {
  const res = await getCategoryTree()
  if (res.code === 0) {
    categoryData.value = res.data
  }
}
getCategoryData()
const selectedCategoryPath = ref([]);
const setCategoryId = (id) => {
  condition.category_id = id;
  selectedCategoryPath.value = getCategoryPath(id);
};
const isCategorySelected = (id) => {
  return selectedCategoryPath.value.includes(id);
};
const getCategoryPath = (id) => {
  const path = [];
  const findCategory = (categories, id) => {
    for (const category of categories) {
      if (category.id === id) {
        path.push(category.id);
        return true;
      }
      if (category.children && findCategory(category.children, id)) {
        path.push(category.id);
        return true;
      }
    }
    return false;
  };
  findCategory(categoryData.value, id);
  return path.reverse();
};

// 筛选
const filt = reactive({
  "qhj_f": null, //  券后价
  "qhj_t": null, // 券后价
  "cr_f": null,                // 佣金比例
  "cr_t": null,                // 佣金比例
  "cv": null,                  // 佣金
  "sale_num": null,         // 销量
  "coupon_v_f": null,          // 券面额
  "coupon_v_t": null,          // 券面额
  "qgt": null,    // 券过期天数
  "experience_score": null,    // 体验分
})
const onFilt = () => {
  condition.qhj_f = Number(filt.qhj_f) === 0? null: Number(filt.qhj_f)
  condition.qhj_t = Number(filt.qhj_t)=== 0? null: Number(filt.qhj_t)
  condition.cr_f = Number(filt.cr_f)=== 0? null: Number(filt.cr_f)
  condition.cr_t = Number(filt.cr_t)=== 0? null: Number(filt.cr_t)
  condition.cv = Number(filt.cv)=== 0? null: Number(filt.cv)
  condition.sale_num = Number(filt.sale_num)=== 0? null: Number(filt.sale_num)
  condition.couv_f = Number(filt.coupon_v_f)=== 0? condition.couv_f: Number(filt.coupon_v_f)
  condition.couv_t = Number(filt.coupon_v_t)=== 0? null: Number(filt.coupon_v_t)
  condition.qgt =Number(filt.qgt)=== 0? null: Number(filt.qgt)
  condition.tyf = Number(filt.experience_score)=== 0? null: Number(filt.experience_score)
}


// 排序
const activeSort = ref('');
const sortMethods = [
  {method: 'sortDaySales', label: '今日热销'},
  {method: 'sort2HourSales', label: '近2小时销量'},
  {method: 'sortMonth', label: '月销量'},
  {method: 'sortCR', label: '佣金比例'},
  {method: 'sortNew', label: '最新上架'},
  {method: 'sortHot', label: '推广热度'},
  {method: 'sortCouponNum', label: '领券量'}
];
const setActiveSort = (sortMethod) => {
  activeSort.value = sortMethod.method;
  methods[sortMethod.method]();
};
const methods = {
  sortDaySales: () => {
    if ( condition.order_by !== "sales_day"){
      condition.order_by = "sales_day"
    }

  },
  sort2HourSales: () => {
    if (condition.order_by !== "sales_2_hour"){
      condition.order_by = "sales_2_hour"
    }
  },
  sortMonth: () => {
    if ( condition.order_by !== "sales_month"){
      condition.order_by = "sales_month"
    }
  },
  sortCR: () => {
    if (  condition.order_by !== "cr"){
      condition.order_by = "cr"
    }
  },
  sortNew: () => {
    if ( condition.order_by !== "new"){
      condition.order_by = "new"
    }

  },
  sortHot: () => {
   if ( condition.order_by !== "hot"){
     condition.order_by = "hot"
   }

  },
  sortCouponNum: () => {
    if (    condition.order_by !== "qnum"){
      condition.order_by = "qnum"
    }
  }
};

// ------排序2 -- asc, desc
const sortType = ref('');
const setActiveSort2 = (sortMethod) => {
  activeSort.value = sortMethod.method;
  if (sortType.value === 'desc') {
    sortType.value = 'asc'
    condition.order_by = 'qhj'
    condition.order_type = 'ASC'
  } else {
    sortType.value = 'desc'
    condition.order_by = 'qhj'
    condition.order_type = 'DESC'
  }
};





// 全部条件
const condition = reactive({
  "g_id": null,            // id
  "category_id": null,         // 分类id
  // 筛选
  "qhj_f": null, //  券后价
  "qhj_t": null, // 券后价
  "cr_f": null,                // 佣金比例
  "cr_t": null,                // 佣金比例
  "cv": null,                  // 佣金
  "sale_num": null,         // 销量
  "couv_f": null,          // 券面额
  "couv_t": null,          // 券面额
  "qgt": null,    // 券过期天数
  "tyf": null,    // 体验分
  // 排序
  "order_by": null,
  "order_type": null,
  // 多选框
  "is_choice": null,   // 精选推荐
  "jrsx":null,  // 今日上新
  "data_from": null, // 数据来源
  "absolute_low_price": null,  // 低价好卖
  "buyer_reviews": null,  // 买家好评
  "is_brand": null,  // 商品库
  "platform": null,    // 平台------单选
  "yfx": null, // 运费险
  "hbmx": null,  // 花呗免息
  "psbz": null,  // 花呗免息
  "area": null, // 地区
  "shop": null, // 店铺---单选
  "av_group": [], // 商品活动
  "bx_bdan": null, // 表现--榜单 30天在榜/实时在榜
  "bx_type": null,     // 表现--key  30天新品/30天高佣金/30天低价
  "push_sucai": null, // 推荐素材
  // 分页
  "page": null,
  "pageSize": null,
})


const dataFromName = ref("数据来源")
const dataFromShow = ref(false)

const platformName = ref("平台")
const platformShow = ref(false)
const pushSucaiName = ref("推广素材")
const pushSucaiShow = ref(false)
const serviceGuaranteeName = ref("服务保障")
const serviceGuaranteeShow = ref(false)
const pydqbyShow = ref(false)   // 偏远地区包邮
const shopName = ref("店铺类型")
const shopShow = ref(false)
const activeName = ref("商品活动")
const activeShow = ref(false)
const biaoxianName = ref("商品表现")
const biaoxianShow = ref(false)

// 改变选中状态
watch(condition, (newVal, oldVal) => {
  // 平台相关
  if (newVal.platform === 1) {
    platformShow.value = true
    platformName.value = "抖音"
  }
  if (newVal.platform === 2) {
    platformName.value = "快手"
    platformShow.value = true
  }
  if (newVal.platform === 3) {
    platformName.value = "微信"
    platformShow.value = true
  }
  if (newVal.platform === null) {
    platformShow.value = false
    platformName.value = "平台"
  }

  //数据来源相关
  if (newVal.data_from === 1) {
    dataFromShow.value = true
    dataFromName.value = "爬取"
  }
  if (newVal.data_from === 2) {
    dataFromName.value = "模拟"
    dataFromShow.value = true
  }
  if (newVal.data_from === 3) {
    dataFromName.value = "输入"
    dataFromShow.value = true
  }
  if (newVal.data_from === null) {
    dataFromShow.value = false
    dataFromName.value = "数据来源"
  }

  // 推广素材
  if (pushSucai.value.length > 0) {
    if (pushSucai.value.length === 1) {
      if (pushSucai.value[0] === 1) {
        pushSucaiName.value = "精选素材"
      }
      if (pushSucai.value[0] === 2) {
        pushSucaiName.value = "朋友圈素材"
      }
      if (pushSucai.value[0] === 3) {
        pushSucaiName.value = "短视频"
      }
    } else {
      pushSucaiName.value = "推广素材"
    }
    pushSucaiShow.value = true
  } else {
    pushSucaiShow.value = false
    pushSucaiName.value = "推广素材"
  }

  // 服务保障
  if ((newVal.yfx == null || !newVal.yfx) && (newVal.hbmx == null || !newVal.hbmx) && (newVal.psbz == null || !newVal.psbz) && area.value.length === 0) {
    serviceGuaranteeShow.value = false
  }
  if (newVal.area !== null) {
    pydqbyShow.value = true
  }else{
    pydqbyShow.value = false
  }
  if (newVal.yfx && (newVal.hbmx == null || !newVal.hbmx) && (newVal.psbz == null || !newVal.psbz) && area.value.length === 0) {
    serviceGuaranteeName.value = "运费险"
    serviceGuaranteeShow.value = true
  } else if (newVal.hbmx && (newVal.yfx == null || !newVal.yfx) && (newVal.psbz == null || !newVal.psbz) && area.value.length === 0) {
    serviceGuaranteeName.value = "花呗免息"
    serviceGuaranteeShow.value = true
  } else if (newVal.psbz && (newVal.hbmx == null || !newVal.hbmx) && (newVal.yfx == null || !newVal.yfx) && area.value.length === 0) {
    serviceGuaranteeName.value = "破损保障"
    serviceGuaranteeShow.value = true
  } else if ((newVal.psbz == null || !newVal.psbz) && (newVal.hbmx == null || !newVal.hbmx) && (newVal.yfx == null || !newVal.yfx) && area.value.length > 0) {
    serviceGuaranteeName.value = "偏远地区包邮"
    serviceGuaranteeShow.value = true
  } else {
    serviceGuaranteeName.value = "服务保障"
  }

  // 店铺类型
  if (newVal.shop === 1) {
    shopShow.value = true
    shopName.value = "官方旗舰店"
  }
  if (newVal.shop === 2) {
    shopShow.value = true
    shopName.value = "阿里健康大药房"
  }
  if (newVal.shop === 3) {
    shopShow.value = true
    shopName.value = "天猫国际"
  }
  if (newVal.shop === 4) {
    shopShow.value = true
    shopName.value = "进口超市"
  }
  if (newVal.shop === 5) {
    shopShow.value = true
    shopName.value = "淘宝冠店"
  }
  if (newVal.shop === 6) {
    shopShow.value = true
    shopName.value = "淘宝企业店铺"
  }
  if (newVal.shop === null) {
    shopShow.value = false
    shopName.value = "店铺类型"
  }

  // 商品活动
  if (activeGroup.value.length > 0) {
    if (activeGroup.value.length === 1) {
      if (activeGroup.value[0] === 1) {
        activeName.value = "定向计划"
      }
      if (activeGroup.value[0] === 2) {
        activeName.value = "聚划算"
      }
      if (activeGroup.value[0] === 3) {
        activeName.value = "拍多件"
      }
      if (activeGroup.value[0] === 4) {
        activeName.value = "前N件"
      }
      if (activeGroup.value[0] === 5) {
        activeName.value = "额外满减"
      }
      if (activeGroup.value[0] === 6) {
        activeName.value = "买赠活动"
      }
      if (activeGroup.value[0] === 7) {
        activeName.value = "首单礼金"
      }
    } else {
      activeName.value = "商品活动"
    }
    activeShow.value = true
  } else {
    activeShow.value = false
    activeName.value = "商品活动"
  }


  // 商品表现
  if (newVal.bx_bdan != null && newVal.bx_type == null) {
    biaoxianShow.value=true
    if (newVal.bx_bdan === 1){
      biaoxianName.value="实时榜单"
    }
    if (newVal.bx_bdan === 2){
      biaoxianName.value="30天上榜"
    }
  }
  if (newVal.bx_bdan == null && newVal.bx_type != null) {
    biaoxianShow.value=true
    if (newVal.bx_type === 1){
      biaoxianName.value="30天新品"
    }
    if (newVal.bx_type === 2){
      biaoxianName.value="30天低价"
    }
    if (newVal.bx_type === 3){
      biaoxianName.value="30天高佣金"
    }
  }
  if (newVal.bx_bdan == null && newVal.bx_type == null) {
    biaoxianShow.value=false
    biaoxianName.value="商品表现"
  }
  if (newVal.bx_bdan != null && newVal.bx_type != null) {
    biaoxianShow.value=true
    biaoxianName.value="商品表现"
  }


  // 排序----券后价
  if (newVal.order_by !== "qhj") {
    condition.order_type= null
    sortType.value = ""
  }

})

// condition = reactive({})
watch(condition, (newVal) => {
  console.log(newVal)
  onSearch()
})

</script>

<style>
/* 排序图标的样式 */
.icon-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.icon-container .el-icon {
  margin-bottom: -6px; /* 调整图标之间的距离 */
}

.icon-container .el-icon:last-child {
  margin-bottom: 0; /* 去掉最后一个图标的负间距 */
}

/* 排序图标的样式 end -----------------------  */
.popover-reference {
  margin-left: 15px;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
}

.input_tight{
  width: 60px;
}
.input_loose{
  width: 80px;
}



</style>
