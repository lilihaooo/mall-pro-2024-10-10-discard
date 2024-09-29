<script setup>
import router from "@/router";
import {ElMessage} from 'element-plus'
import {ApiCancelCollect, ApiCollect, ApiPromotion} from "@/api/product";
import {useUserStore} from "@/pinia";
import {ref, computed} from "vue"


const userStore = useUserStore()

const userID = userStore.userInfo.ID



const props = defineProps({
  data: {
    type: Array,
    required: true
  }
});
const items = computed(() => props.data);


const goodsDetail = (id, isExpire) => {
  if (!isExpire) {
    router.push({
      name: 'goodsDetail',
      query: {id: id},
    })
  } else {
    ElMessage({
      message: `商品为：${id}，请修改商品有效期！`,
      type: 'warning',
    })
  }
}

// 收藏
const goodsCollect = async (goodsID) => {
  const res = await ApiCollect(userID, goodsID)
  if (res.code === 0) {
    ElMessage.success('收藏成功')
    const currentItem = items.value.find(item => item.id === goodsID);
    currentItem.is_collect = true;
  }
}

// 取消收藏
const cancelGoodsCollect = async (goodsID) => {
  const res = await ApiCancelCollect(userID, goodsID)
  if (res.code === 0) {
    ElMessage.success('取消成功')
    const currentItem = items.value.find(item => item.id === goodsID);
    currentItem.is_collect = false;
  }
}


const handleCollect = (id, isCollect) => {
    if (isCollect) {
      // 取消收藏
      cancelGoodsCollect(id)
    }else{
      // 收藏
      goodsCollect(id)
    }
}

const goodsPromotion = async (goodsID) => {
  const res = await ApiPromotion( goodsID)
  if (res.code === 0) {
    ElMessage.success('推广成功')
  }
}
const handlePromotion = (id) => {
  goodsPromotion(id)
}

</script>

<template>
  <div class="outer-container">
    <el-scrollbar>
      <div class="list-container">
        <div v-for="item in items" :key="item.id" class="module"
             @click="goodsDetail(item.id, item.is_expire)"
        >
          <div class="module_up">
            <div class="image_container">
              <el-image :src="item.cover_image" class="image"/>
              <!-- 条件渲染，显示过期标志图片 -->
              <img
                  v-if="item.is_expire"
                  src="@/assets/shixiao.png"
                  class="expire-image"
                  alt="Expired"
              />
              <div class="image-bottom-brand" v-if="item.brand_name !== ''">
                {{ item.brand_name }}
              </div>
              <div class="image-bottom">
                <div
                    class="bottom-item"
                    @click.stop="handleCollect(item.id, item.is_collect)"
                >
                  {{ item.is_collect ? '取消收藏' : '收藏' }}
                </div>
                <div
                    class="bottom-item"
                    @click.stop="handlePromotion(item.id)"
                >
                  推广
                </div>
              </div>


            </div>
            <div class="goods-base-info">
              <div class="goods-title">{{ item.title }}</div>
              <div class="goods-two">
                <div class="two-item">
                  <div v-if="item.coupon_id !== 0">
                    <div class="info-bt">券后价</div>
                    <div class="qhj">¥{{ item.post_coupon_price }}</div>
                  </div>
                  <div v-else>
                    <div class="info-bt">原价</div>
                    <div class="qhj">¥{{ item.price }}</div>
                  </div>

                </div>
                <div class="two-item">
                  <div>
                    <div class="info-bt">佣金</div>
                    <div style="margin-top: 6px">
                      <div>
                        <span class="cr">%{{ item.commission_rate }}</span>
                        <span class="cv">{{ item.commission_value }}元</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="goods-three">
                <div class="three-item-tr">
                  <div class="three-co">
                    <span class="info-bt">券</span>
                    <span class="three-v">¥{{ item.coupon_amount }}</span>
                  </div>
                  <div class="three-sales">
                    <span class="info-bt">销量</span>
                    <span class="three-v">{{ item.sales_all }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="module_down">
            <div class="more-sales">
              <span class="info-bt-2">实时销量</span>
              <span class="three-v">24+</span>
              <span class="info-bt-2" style="margin-left: 28px">全天销量</span>
              <span class="three-v">445+</span>
            </div>
            <div class="more-tag">
              <div class="more-tag-item">
                <div style="color: #999; margin-right: 6px">
                  <el-icon>
                    <Ticket/>
                  </el-icon>
                </div>
                <div style="color: #FF6545;font-size: 13px"> {{ item.coupon_cover }}</div>
                <div style="margin: 0 2px; color: #999"> /</div>
                <div style="color: #999; font-size: 12px"> {{ item.coupon_total }}</div>
              </div>

              <div class="more-tag-item" style="margin-left: 24px">
                <div style="color: #999; font-size: 12px; margin-right: 12px">开始时间</div>
                <div style="color: #999; font-size: 12px"> {{ item.coupon_begin_time }}</div>
              </div>
              <div class="more-tag-item" style="margin-left: 24px">
                <div style="color: #999; font-size: 12px; margin-right: 12px">结束时间</div>
                <div style="color: #999; font-size: 12px"> {{ item.coupon_end_time }}</div>
              </div>
              <div class="more-tag-item">
                <div style="color: #999; margin-right: 6px">
                  <el-icon>
                    <House/>
                  </el-icon>
                </div>
                <div style="color: #999; font-size: 12px"> {{ item.shop_name }}</div>
              </div>

              <div class="more-tag-item">
                <div style="color: #999; margin-right: 6px">
                  <el-icon>
                    <Place/>
                  </el-icon>
                </div>
                <div style="color: #999; font-size: 12px"> {{ item.media_name }}</div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </el-scrollbar>
  </div>
</template>
<style scoped lang="scss">
.outer-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  overflow-x: auto;
  background-color: #c0deef;
  height: 100%;
}

.list-container {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 10px;
  width: calc(5 * 266px + 4 * 5px); /* 5 modules width + 4 gaps */
  position: relative; /* To contain the absolute positioned modules */
  padding-bottom: 146px;
}

.module {
  margin-left: 1px;
  margin-right: 1px;
  margin-top: 1px;
  border-radius: 10px 10px 0 0;
  position: relative; /* To position the goods-more-info absolutely */
}

.module_up {
  width: 260px;
  height: 420px;
  border-radius: 10px 10px 0 0;
}

.module:hover {
  box-shadow: 0 -1px 0 0 var(--el-color-primary), /* top */
  -1px 0 0 0 var(--el-color-primary), /* left */
  1px 0 0 0 var(--el-color-primary); /* right */
}

.module_down {
  width: 260px;
  height: 144px;
  background-color: white;
  position: absolute;
  top: 100%; /* Position below the module */
  display: none; /* Hidden by default */
  z-index: 100; /* Higher z-index to float above other elements */
  border-radius: 0 0 10px 10px;
}

.module:hover .module_down {
  display: block;
  box-shadow: 0 1px 0 0 var(--el-color-primary), /* top */
  -1px 0 0 0 var(--el-color-primary), /* left */
  1px 0 0 0 var(--el-color-primary); /* right */
}

.module:hover .goods-base-info {
  border-radius: 0;
}

.goods-base-info {
  width: 100%;
  height: 160px;
  background-color: white;
  border-radius: 0 0 10px 10px;
}

.image_container {
  position: relative; /* 使子元素的绝对定位相对于该容器 */
  width: 100%;
  height: 260px;
}

.expire-image {
  position: absolute; /* 绝对定位 */
  top: 50px;
  left: 50px;
  width: 60%; /* 覆盖整个容器 */
  height: 60%;
  object-fit: cover; /* 保持图片的纵横比 */
  z-index: 10; /* 确保图片在上层显示 */
  pointer-events: none; /* 防止覆盖图片影响底层图片的点击事件 */
}

.image-bottom {
  display: none; /* Hidden by default */
  justify-content: space-between;
  position: absolute; /* 绝对定位 */
  top: 220px;
  width: 100%; /* 覆盖整个容器 */
  height: 40px;
  object-fit: cover; /* 保持图片的纵横比 */
  z-index: 10; /* 确保图片在上层显示 */

  background-color: white; /* 设置父容器背景颜色为白色，确保间隙显示为白色 */
  gap: 1px; /* 设置两个子元素之间的间隙 */
}


.image-bottom-brand {
  display: flex; /* Hidden by default */
  align-items: center; /* 确保每个子 div 的内容也垂直居中 */
  position: absolute; /* 绝对定位 */
  top: 240px;
  right: 0;
  height: 20px;
  pointer-events: none; /* 防止覆盖图片影响底层图片的点击事件 */
  padding: 0 8px;
  color: #fff;
  background: linear-gradient(223deg, #fe8721, #ed70aa);
  border-radius: 6px 0 0 0;
  text-align: right;
}

.module:hover .image-bottom-brand {
  top: 200px;
}


.bottom-item {
  flex: 1; /* 平分父容器 */
  display: flex; /* 使用 flex 布局 */
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  text-align: center; /* 文字居中 */
  background-color: #f70;
  color: #fff;
  height: 100%; /* 确保子元素高度继承父元素 */
}


.module:hover .image-bottom {
  display: flex;
}


.image {
  width: 100%;
  display: block;
  border-radius: 10px 10px 0 0;
  object-fit: cover; /* 保持图片的纵横比 */
}

.goods-title {
  padding: 6px 0 3px 3px;
  color: #777;
  font-size: 14px;
  height: 32px;
}

.goods-two {
  padding: 3px 0 3px 3px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.two-item {
  margin: 25px 0 0 0;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.info-bt {
  color: #999999;
  font-size: 12px;
}

.qhj {
  margin-top: 6px;
  color: #FF6545FF;
  font-size: 18px;
}

.cr {
  color: #0841FBFF;
  font-size: 18px;
}

.cv {
  color: #666;
  font-size: 13px;
  background-color: #ddd;
  margin-left: 2px;
  border-radius: 2px;
  padding: 1px 3px;
}

.goods-three {
  padding: 26px 0 3px 3px;
}

.three-item-tr {
  display: flex;
}

.three-sales {
  margin-left: 26px;
}

.three-v {
  margin-left: 3px;
  font-size: 13px;
  color: #FF6545;
}

.more-sales {
  margin: 12px 0 3px 3px;

}

.info-bt-2 {
  font-size: 13px;
  color: #777;
}

.more-tag {
  margin-top: 22px;
}

.more-tag-item {
  margin: 3px 5px 3px 3px;
  display: flex;
  justify-content: flex-start; /* 子元素左对齐 */
  align-items: center; /* 垂直居中对齐 */
}

.more-tag-item > div {
  display: flex;
  align-items: center; /* 确保每个子 div 的内容也垂直居中 */
}


</style>
