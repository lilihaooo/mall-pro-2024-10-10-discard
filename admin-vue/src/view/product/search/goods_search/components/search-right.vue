<script setup>
import {ref, watch} from 'vue';
import {ApiBatchCancelCollect, ApiCollect, ApiMyCollect, ApiMyPromotion, ApiPromotion} from "@/api/product";
import {ElMessage} from "element-plus";

const truncateTitle = (title) => {
  if (title.length > 15) {
    return title.slice(0, 15) + '...';
  }
  return title;
}
const checkCollectIDList = ref([])
const drawerCollect = ref(false);
const openDrawerCollect = () => {
  drawerCollect.value = true
  checkCollectIDList.value = []
}


const drawerPromotion = ref(false);
const openDrawerPromotion = () => {
  drawerPromotion.value = true
}


const drawerClose = () => {
  drawerCollect.value = false;
  drawerPromotion.value = false;
}


const myCollectCount = ref(0)
const myPromotionCount = ref(0)

const handleCollectCurrentChange = (val) => {
  getMyCollect(val)
}


const handlePromotionCurrentChange = (val) => {
  getMyPromotion(val)
}


const myCollect = ref([])

const getMyCollect = async (page) => {
  const res = await ApiMyCollect(page)
  if (res.code === 0) {
    myCollect.value = res.data.list
    myCollectCount.value = res.data.total

  }
}

const myPromotion = ref([])
const getMyPromotion = async (page) => {
  const res = await ApiMyPromotion(page)
  if (res.code === 0) {
    myPromotion.value = res.data.list
    myPromotionCount.value = res.data.total

  }
}

// 批量取消收藏
const batchCancelGoodsCollect = async (ids) => {
  const res = await ApiBatchCancelCollect(ids)
  if (res.code === 0) {
    ElMessage.success('取消成功')
    // 过滤掉 id 在 ids 数组中的元素
    myCollect.value = myCollect.value.filter(item => !ids.includes(item.id));
  }
}

const goodsPromotion = async (goodsID) => {
  const res = await ApiPromotion(goodsID)
  if (res.code === 0) {
    ElMessage.success('推广成功')
  }
}
const handlePromotion = (id) => {
  goodsPromotion(id)
}


watch(drawerCollect, (nValue) => {
  if (nValue) {
    // 获取我的收藏
    getMyCollect(1)
  }
})


watch(drawerPromotion, (nValue) => {
  if (nValue) {
    // 获取我的收藏
    getMyPromotion(1)
  }
})


const formattedCouponEndTime = (endTime) => {
  return endTime.split('T')[0]; // 只取日期部分
}

const timeAgoCollect = (createdAt) => {
  const now = new Date(); // 获取当前时间
  const createdTime = new Date(createdAt); // 将字符串时间转换为 Date 对象
  const diffInMs = now - createdTime; // 时间差，单位为毫秒
  const diffInSeconds = diffInMs / 1000; // 转换为秒
  const diffInMinutes = diffInSeconds / 60;
  const diffInHours = diffInMinutes / 60;
  const diffInDays = diffInHours / 24;
  const diffInMonths = diffInDays / 30; // 粗略地将一个月视为 30 天
  const diffInYears = diffInMonths / 12;

  if (diffInMinutes < 1) {
    return '刚刚收藏';
  } else if (diffInMinutes < 5) {
    return '5分钟前收藏';
  } else if (diffInMinutes < 10) {
    return '10分钟前收藏';
  } else if (diffInMinutes < 30) {
    return '30分钟前收藏';
  } else if (diffInHours < 1) {
    return '1小时前收藏';
  } else if (diffInDays < 1) {
    return `${Math.floor(diffInHours)}小时前收藏`;
  } else if (diffInDays < 7) {
    return `${Math.floor(diffInDays)}天前收藏`;
  } else if (diffInDays < 30) {
    return `${Math.floor(diffInDays / 7)}周前收藏`;
  } else if (diffInMonths < 12) {
    return `${Math.floor(diffInMonths)}个月前收藏`;
  } else {
    return `${Math.floor(diffInYears)}年前收藏`;
  }
}

const timeAgoPromotion = (createdAt) => {
  const now = new Date(); // 获取当前时间
  const createdTime = new Date(createdAt); // 将字符串时间转换为 Date 对象
  const diffInMs = now - createdTime; // 时间差，单位为毫秒
  const diffInSeconds = diffInMs / 1000; // 转换为秒
  const diffInMinutes = diffInSeconds / 60;
  const diffInHours = diffInMinutes / 60;
  const diffInDays = diffInHours / 24;
  const diffInMonths = diffInDays / 30; // 粗略地将一个月视为 30 天
  const diffInYears = diffInMonths / 12;

  if (diffInMinutes < 1) {
    return '刚刚推广';
  } else if (diffInMinutes < 5) {
    return '5分钟前推广';
  } else if (diffInMinutes < 10) {
    return '10分钟前推广';
  } else if (diffInMinutes < 30) {
    return '30分钟前推广';
  } else if (diffInHours < 1) {
    return '1小时前推广';
  } else if (diffInDays < 1) {
    return `${Math.floor(diffInHours)}小时前推广`;
  } else if (diffInDays < 7) {
    return `${Math.floor(diffInDays)}天前推广`;
  } else if (diffInDays < 30) {
    return `${Math.floor(diffInDays / 7)}周前推广`;
  } else if (diffInMonths < 12) {
    return `${Math.floor(diffInMonths)}个月前推广`;
  } else {
    return `${Math.floor(diffInYears)}年前推广`;
  }
}

</script>

<template>
  <div @click="drawerClose">
    <div style="display: flex; justify-content: center;">
      <div style="margin-top: 10px; width: 70%;">
        <div class="left-bottom" @click.stop="openDrawerCollect"
             :style="{ backgroundColor: drawerCollect ? 'var(--el-color-primary)' : '', color: drawerCollect ? 'white' : ''  }"
        >
          <div>
            <el-icon>
              <Goods/>
            </el-icon>
          </div>
          我的收藏
        </div>

        <div class="left-bottom"
             @click.stop="openDrawerPromotion"
             :style="{ backgroundColor: drawerPromotion ? 'var(--el-color-primary)' : '', color: drawerPromotion ? 'white' : ''  }"
        >
          <div>
            <el-icon>
              <Money/>
            </el-icon>
          </div>
          我的推广
        </div>
      </div>
    </div>

    <!--    收藏抽屉-->
    <el-drawer v-model="drawerCollect"

               style="margin-right: 59px;
               margin-top: 63px;
               border-radius: 14px 0 0 14px;
               height: 92.8vb;
               background-color:#eaf4d5;
               position: absolute;
               right: 0%;"
               size="410"
               modal-class="modal-collect"
               :modal="false"
               :show-close="false"
    >
      <div style="margin-right: 20px;">
        <el-checkbox-group v-model="checkCollectIDList">
          <div v-for="one in myCollect" class="one">
            <div class="multiple-choice-container">
              <el-checkbox :value="one.id"/>
            </div>

            <div class="collect-div">
              <div class="collect-div-up">
                <div class="collect-div-up-one">
                  {{ timeAgoCollect(one.created_at) }}
                </div>

                <div class="collect-div-up-two">
                  <div class="collect-div-up-two-left">
                    <el-image class="collect-image"
                              :src="one.image.url"
                    />
                  </div>
                  <div class="collect-div-up-two-right">
                    <div style="  width: 100%; margin-left: 4px">
                      <div class="goods-info-row" style="font-size: 14px; color: #222">
                        {{ truncateTitle(one.title) }}
                      </div>
                      <div class="goods-info-row">
                        <div class="goods-info-row-one-1">
                          <span v-if="one.coupon_value !== 0">券后 </span>
                          <span v-else>原价 </span>
                          <span style="color: #ff4443;font-size: 13px">¥{{ one.post_coupon_price }}</span>
                        </div>
                        <div class="goods-info-row-one-1">
                          佣金 <span style="color: #1e6fff; font-size: 13px">%{{ one.commission_rate }}</span>
                          (约¥{{ one.commission_value }})
                        </div>
                      </div>
                      <div class="goods-info-row" v-if="one.coupon_value !== 0">
                        <div class="goods-info-row-one-1">
                          券 <span style="color: #ff4443;font-size: 13px">¥{{ one.coupon_value }}</span>
                        </div>
                        <div class="goods-info-row-one-1">
                          有效期至 <span style="color: #999; font-size: 13px">
                      {{ formattedCouponEndTime(one.coupon_end_time) }}
                        </span>
                        </div>
                      </div>
                      <div class="goods-info-row" v-else="one.coupon_value !== 0">

                      </div>
                      <div class="goods-info-row">
                        <div class="goods-info-row-one-1">
                          销量 <span style="color: #444;font-size: 12px">1000+</span>
                        </div>

                      </div>
                    </div>


                  </div>
                </div>


              </div>
              <div class="collect-div-down">
                <div style="width: 100%; text-align: center; color: #1e6fff" @click="batchCancelGoodsCollect([one.id])">
                  取消收藏
                </div>
                <div
                    @click="handlePromotion(one.id)"
                    style="width: 100%; text-align: center; color: #1e6fff"
                >
                  推广
                </div>
              </div>
            </div>

          </div>
        </el-checkbox-group>
      </div>


      <template #header>
        <div style="flex: auto; margin-left: 24px; padding-top: 0px;display: flex; justify-content: center;">
          <div style="width: 100%; display: flex;align-items: center;">我的收藏</div>
          <div style="width: 100%; display: flex; justify-content: right; padding-right: 30px; color: red">
            <el-button
                link
                type="danger"
                :disabled="checkCollectIDList.length === 0"
                @click="batchCancelGoodsCollect(checkCollectIDList)"
            >
              删除
            </el-button>
          </div>

        </div>
      </template>

      <template #footer>
        <div style="flex: auto; margin-left: 24px">
          <el-pagination
              :pager-count=5
              :page-sizes="[10]"
              layout="prev, pager, next"
              :total=myCollectCount
              @current-change="handleCollectCurrentChange"
          />
        </div>
      </template>
    </el-drawer>


    <!--    推广抽屉-->
    <el-drawer v-model="drawerPromotion"
               style="margin-right: 59px;
               margin-top: 63px;
               border-radius: 14px 0 0 14px;
               height: 92.8vb;
               background-color:#eaf4d5;
               position: absolute;
               right: 0;"
               size="410"
               modal-class="modal-collect"
               :modal="false"
               :show-close="false"
    >

      <div v-for="one in myPromotion" class="one">
        <div class="promotion-div">
          <div class="collect-div-up">
            <div class="collect-div-up-one">
              {{ timeAgoPromotion(one.created_at) }}
            </div>

            <div class="collect-div-up-two">
              <div class="collect-div-up-two-left">
                <el-image class="collect-image"
                          :src="one.image_url"
                />
              </div>
              <div class="collect-div-up-two-right">
                <div style="  width: 100%; margin-left: 4px">
                  <div class="goods-info-row" style="font-size: 14px; color: #222">
                    {{ truncateTitle(one.title) }}
                  </div>
                  <div class="goods-info-row">
                    <div class="goods-info-row-one-1">
                      <span v-if="one.coupon_value !== 0">券后 </span>
                      <span v-else>原价 </span>
                      <span style="color: #ff4443;font-size: 13px">¥{{ one.post_coupon_price }}</span>
                    </div>
                    <div class="goods-info-row-one-1">
                      佣金 <span style="color: #1e6fff; font-size: 13px">%{{ one.commission_rate }}</span>
                      (约¥{{ one.commission_value }})
                    </div>
                  </div>
                  <div class="goods-info-row" v-if="one.coupon_value !== 0">
                    <div class="goods-info-row-one-1">
                      券 <span style="color: #ff4443;font-size: 13px">¥{{ one.coupon_value }}</span>
                    </div>
                    <div class="goods-info-row-one-1">
                      有效期至 <span style="color: #999; font-size: 13px">
                  {{ formattedCouponEndTime(one.coupon_end_time) }}
                    </span>
                    </div>
                  </div>
                  <div class="goods-info-row" v-else="one.coupon_value !== 0">

                  </div>
                  <div class="goods-info-row">
                    <div class="goods-info-row-one-1">
                      销量 <span style="color: #444;font-size: 12px">{{one.seals_all}}</span>
                    </div>
                    <div class="goods-info-row-one-1">
                      推广次数 <span style="color: #444;font-size: 12px">{{one.promotion_num}}</span>
                    </div>
                  </div>
                </div>


              </div>
            </div>


          </div>
          <div class="promotion-div-down">
            <div
                @click="handlePromotion(one.id)"
                style="width: 100%; text-align: center; color: #1e6fff"
            >
              推广
            </div>
          </div>
        </div>

      </div>


      <template #header>
        <div style="flex: auto; margin-left: 24px; padding-top: 0px;display: flex; justify-content: center;">
          <div style="width: 100%; display: flex;align-items: center;">我的推广</div>
        </div>
      </template>

      <template #footer>
        <div style="flex: auto; margin-left: 24px">
          <el-pagination
              :pager-count=5
              :page-sizes="[10]"
              layout="prev, pager, next"
              :total=myPromotionCount
              @current-change="handlePromotionCurrentChange"
          />
        </div>
      </template>
    </el-drawer>
  </div>

</template>

<style scoped lang="scss">
//  有一些bug, 也不影响: 推广和收藏共用了一些 收藏的css!!!
.one {
  display: flex;
  justify-content: center;
}

.collect-image {
  border-radius: 5px;
  width: 80px;
  height: 80px
}

.multiple-choice-container {
  width: 24px;
  height: 110px;
  margin: 8px 0;
  display: flex;
  align-items: center;
}


.left-bottom {
  padding: 20px 0 20px 0;
  color: #666;
  text-align: center;
}

.left-bottom:hover {
  background-color: var(--el-color-primary);
  color: white;
}

.modal-collect {
  background-color: rgba(0, 0, 0, 1) !important; /* 设置遮罩透明度为50% */
}

.collect-div {
  border: 1px solid #ddd;
  margin: 6px 0;
  height: 110px;
  width: 315px;
  border-radius: 8px;
  background-color: white;
  position: relative;
  padding-left: 10px;

}


.promotion-div {
  border: 1px solid #ddd;
  margin: 6px 0;
  height: 110px;
  width: 315px;
  border-radius: 8px;
  background-color: white;
  position: relative;
  padding-left: 10px;
}

.promotion-div:hover {
  border-radius: 8px 8px 0 0;
  border: 1px solid #ddd;
  border-bottom: none; /* 去掉下边框 */
  height: 111px;
}

.promotion-div:hover .promotion-div-down {
  display: flex;
  top: 110px;
  left: -1px;
  justify-content: center;
  align-items: center;
  width: 325px;
  position: absolute;
  z-index: 10;
  border: 1px solid #ddd;
  border-radius: 0 0 8px 8px;
  border-top: none; /* 去掉下边框 */
  border-top: 1px dashed #ddd;
}

.promotion-div-down {
  display: none;
  height: 28px;
  background-color: white;
}

.collect-div-up {
  height: 140px;
}



.collect-div-down {
  display: none;
  height: 28px;
  background-color: white;
}

.collect-div:hover {
  border-radius: 8px 8px 0 0;
  border: 1px solid #ddd;
  border-bottom: none; /* 去掉下边框 */
  height: 111px;
}


.collect-div:hover .collect-div-down {
  display: flex;
  top: 110px;
  left: -1px;
  justify-content: center;
  align-items: center;
  width: 325px;
  position: absolute;
  z-index: 10;
  border: 1px solid #ddd;
  border-radius: 0 0 8px 8px;
  border-top: none; /* 去掉下边框 */
  border-top: 1px dashed #ddd;
}


.collect-div-up-two {
  display: flex;
  justify-content: center;
}

.collect-div-up-one {
  height: 24px;
  display: flex;
  align-items: center;
  border-radius: 8px 8px 0 0;
  font-size: 11px;
  color: #999;
}

.collect-div-up-two-right {
  height: 80px;
  width: 100%;
  display: flex;
  align-items: center;


}

.goods-info-row {
  height: 20px;
  display: flex;
  align-items: center;
  font-size: 12px;
  color: #999;
}


.goods-info-row-one-1 {
  font-size: 11px;
  color: #999;
  margin-right: 20px;

}

.drawer-down {
  background-color: blue;
  width: 100%;
}


</style>