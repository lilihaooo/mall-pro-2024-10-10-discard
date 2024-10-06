<script setup>
import {ref, reactive, onMounted, computed} from 'vue';
import {useRoute} from "vue-router";
import {
  DeleteImage,

  getGoodsInfo,
  GetTagList,
  SetCoverImage,
  UpdateGoodsBaseInfo,
  UpdateGoodsCouponInfo
} from "@/api/product";
import {ElMessage} from "element-plus";
import ImageCompress from "@/utils/image";
import {Plus} from "@element-plus/icons-vue";


const route = useRoute();
const id = route.query.id;
const goodsInfo = reactive({
  goodsTitle: null,
  goodsDes: null,
  images: [],
  coverImageUrl: "",
  shopName: null,
  shopWlF: null,
  shopFWF: null,
  shopSPF: null,
  shopLogo: null,
  tzName: null,
  tzLevel: null,
  brandName: null,
  tags: [],


  cv: null,
  cr: null,


  saleAll: null,
  sale2h: null,
  saleDay: null,

  coupon: null,
  couponValue: null,
  couponEndTime: null,
  couponBeginTime: null,

  postCouponPrice: null,
  price: null


})


const coverImageID = ref(0)

// 0 无消费券; 1 正常; 2 过期; 3 未开始;
const couponStatus = ref(0)



const getTagList = async () => {
  const res = await GetTagList();
  tagData.value = res.data
}
// 基本编辑抽屉
const showBaseUpdateVisible = () => {
  baseUpdateVisible.value = true
}

const showTags = ref([])
const baseUpdateVisible = ref(false)
const baseUpdate = reactive({
  description: "",
  price: 0,
  commission_rate: 0,
  tags: [],
  brand_id: 0,
  goods_id: 0
})


const updateGoodBaseInfo = async (data) => {
  await UpdateGoodsBaseInfo(data);
}

// 提交
let arr = []
const onBaseUpdateSubmit = async () => {
  arr = []
  showTags.value.forEach(tag => {
    arr.push(tag.id)
  });
  baseUpdate.tags = arr
  baseUpdate.goods_id = Number(id)
  baseUpdate.commission_rate = Number(baseUpdate.commission_rate)

  await updateGoodBaseInfo(baseUpdate)
  baseUpdateVisible.value = false

  await getGoodsDetail()
}


const updateGoodCouponInfo = async (data) => {
  await UpdateGoodsCouponInfo(data);
}
// 提交优惠券更改
const onCouponUpdateSubmit = async () => {
  couponInfo.value.goods_id = Number(id)
  couponInfo.value.amount = Number(couponInfo.value.amount)
  couponInfo.value.coupon_total = Number(couponInfo.value.coupon_total)
  await updateGoodCouponInfo(couponInfo.value)
  couponUpdateVisible.value = false
  await getGoodsDetail()
}

const onConponUpdateReset = () => {
  couponInfo.value = originalCouponInfo.value
  console.log("重置")
}

const originalDescription = ref("")
const originalCommissionRate = ref(0)
const originalTags = ref([])
const originalPrice = ref(0)
const originalBrandID = ref(0)

// 重置
const onBaseUpdateReset = () => {
  showTags.value = originalTags.value
  baseUpdate.price = originalPrice.value
  baseUpdate.description = originalDescription.value
  baseUpdate.commission_rate = originalCommissionRate.value
  baseUpdate.brand_id = originalBrandID.value
}
const delTag = (id) => {
  showTags.value = showTags.value.filter(tag => tag.id !== id);
}
const tagData = ref([])


let selectedIds = []
const originalSelectedIds = ref([])

const changeTagPosition = (id) => {
  const index = selectedIds.indexOf(id);
  if (index !== -1) {
    // 如果 selectedIds 中包含传入的 id，则删除该 id
    selectedIds.splice(index, 1);
  } else {
    // 如果 selectedIds 中不包含传入的 id，则将其加入 selectedIds
    selectedIds.push(id);
  }

  initTagList();
};


const tagList = ref([])
// 显示标签对话框
const tagDialogVisible = ref(false)
const initTagList = () => {
  // 添加 selected 字段
  const updatedData = computed(() => {
    return tagData.value.map(item => ({
      ...item,
      selected: selectedIds.includes(item.id),
    }));
  });

  // 按 group 分组
  const groupedData = computed(() => {
    return updatedData.value.reduce((acc, item) => {
      if (!acc[item.group]) {
        acc[item.group] = [];
      }
      acc[item.group].push(item);
      return acc;
    }, {});
  });

  tagList.value = groupedData.value
}
const showTag = async () => {
  tagDialogVisible.value = true
  await getTagList()
  //获取数据
  initTagList()
}
const tagCommit = () => {
  const filteredData = ref(tagData.value.filter(item => selectedIds.includes(item.id)));
  showTags.value = filteredData.value;
  tagDialogVisible.value = false
}
const tagReset = () => {
  selectedIds = originalSelectedIds.value
  initTagList()
}
// 优惠券抽屉
const couponUpdateVisible = ref(false)
const originalCouponInfo = ref({})  // 保留一份原始数据
const couponInfo = ref({})
const showCouponUpdateVisible = () => {
  couponUpdateVisible.value = true
  console.log(couponInfo.value)
}


const getGoodsDetail = async () => {
  const res = await getGoodsInfo({id: id});

  let coverImage = res.data.images[0].url
  if (res.data.cover_image_id !== 0) {
    coverImage = res.data.image.url;
    coverImageID.value = res.data.cover_image_id
  }
  goodsInfo.coverImageUrl = coverImage;

  goodsInfo.shopName = res.data.shop.name;
  goodsInfo.shopFWF = res.data.shop.service_score;
  goodsInfo.shopSPF = res.data.shop.product_score;
  goodsInfo.shopWlF = res.data.shop.logistics_score;
  goodsInfo.shopLogo = res.data.shop.logo;
  goodsInfo.tzName = res.data.media.media_mame;
  goodsInfo.tzLevel = res.data.media.user_level;
  goodsInfo.tzLogo = res.data.media.media_head;


  goodsInfo.brandName = res.data.brand.name

  if (res.data.title.length > 50) {
    goodsInfo.goodsTitle = res.data.title.substring(0, 50) + '...';
  } else {
    goodsInfo.goodsTitle = res.data.title;
  }
  goodsInfo.goodsDes = res.data.description

  goodsInfo.coverImages = res.data.images;
  goodsInfo.images = res.data.images
  goodsInfo.tags = res.data.tags;

  const cEndTime = new Date(res.data.coupon.end_time)
  const cBeginTime = new Date(res.data.coupon.begin_time)
  const nowTime = new Date()

  if (res.data.coupon.ID === null) {
    couponStatus.value = 0;
  } else {
    if (cEndTime >= nowTime && cBeginTime <= nowTime) {
      couponStatus.value = 1;
    } else if (cEndTime < nowTime) {
      couponStatus.value = 2;
    } else if (cBeginTime > nowTime) {
      couponStatus.value = 3;
    }
  }

  goodsInfo.coupon = res.data.coupon;
  goodsInfo.postCouponPrice = res.data.post_coupon_price
  goodsInfo.couponEndTime = formatDateTime(res.data.coupon.end_time);
  goodsInfo.couponBeginTime = formatDateTime(res.data.coupon.begin_time);
  goodsInfo.couponValue = res.data.coupon.amount

  goodsInfo.price = res.data.price;
  baseUpdate.price = res.data.price;
  baseUpdate.description = res.data.description;
  baseUpdate.commission_rate = res.data.commission_rate;

  if (res.data.tags.length > 0) {
    for (let i = 0; i < res.data.tags.length; i++) {
      let tag = res.data.tags[i];
      baseUpdate.tags = tag.id
    }
  }


  goodsInfo.cv = res.data.commission_value;
  goodsInfo.cr = res.data.commission_rate;

  goodsInfo.sale2h = res.data.sales_2_hour
  goodsInfo.saleDay = res.data.sales_day
  goodsInfo.saleAll = res.data.sales_all

  showTags.value = res.data.tags;
  originalTags.value = res.data.tags;
  originalDescription.value = res.data.description;
  originalCommissionRate.value = res.data.commission_rate;
  originalPrice.value = res.data.price;

  originalBrandID.value = res.data.brand_id;
  baseUpdate.brand_id = res.data.brand_id;

  originalSelectedIds.value = []
  res.data.tags.forEach(tag => originalSelectedIds.value.push(tag.id));
  res.data.tags.forEach(tag => selectedIds.push(tag.id));

  // 优惠券初始化
  couponInfo.value = res.data.coupon;
  originalCouponInfo.value = res.data.coupon;
}

function formatDateTime(dateString) {
  const date = new Date(dateString);

  // 获取年、月、日、时、分、秒
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始计数，所以需要+1
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');

  // 格式化输出
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}


const computedWidth = computed(() => {
  const imageCount = goodsInfo.images.length;
  return imageCount > 4 ? '296px' : `${imageCount * 70}px`;
});

// 图片上传相关
const path = ref(import.meta.env.VITE_BASE_API)
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
}
const uploadSuccess = (res) => {
  console.log("上传成功")
  getGoodsDetail()
}
const beforeImageUpload = (file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    return false
  }
  const isRightSize = file.size / 1024 < 512   // 图片高于512K 会被压缩
  if (!isRightSize) {
    // 压缩
    const compress = new ImageCompress(file, 2048, 1920)
    return compress.compress()
  }
  return isRightSize
}

// 图片点击点击事件
const setCoverImage = async (goodsID, imageID) => {
  await SetCoverImage(goodsID, imageID);
  await getGoodsDetail()
}
const deleteImage = async (imageID) => {
  await DeleteImage({id: imageID});
  await getGoodsDetail()
}


onMounted(() => {
  getGoodsDetail()
})
</script>


<template>
  <!--基础抽屉-->
  <el-drawer
      v-model="baseUpdateVisible"
      title="基本信息修改"
      :show-close="false"
  >

    <el-form :inline="false" :model="baseUpdate" label-width="68px">
      <el-form-item label="价格">
        <el-input type="number" v-model="baseUpdate.price"/>
      </el-form-item>

      <el-form-item label="佣金率%">
        <el-input type="number" v-model="baseUpdate.commission_rate"/>
      </el-form-item>

      <el-form-item label="标签">
        <div style="display: flex; align-items: center;">
          <!-- 标签列表 -->
          <div style="display: flex; flex-wrap: wrap; gap: 2px;">
            <el-tag v-for="item in showTags"
                    :key="item.id"
                    type="danger"
                    closable
                    @close="delTag(item.id)"
            >{{ item.title }}
            </el-tag>
            <!-- 按钮 -->
            <el-button
                size="small" type="primary"
                @click="showTag"
                icon="plus"
            >
            </el-button>
          </div>
        </div>
      </el-form-item>


      <el-form-item label="描述">
        <el-input type="textarea"
                  :autosize="{ minRows: 2, maxRows: 6 }"
                  v-model="baseUpdate.description"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onBaseUpdateSubmit">保存</el-button>
        <el-button @click="onBaseUpdateReset">重置</el-button>
      </el-form-item>


    </el-form>
  </el-drawer>
  <!--优惠券抽屉-->
  <el-drawer
      v-model="couponUpdateVisible"
      title="优惠券信息修改"
      :show-close="false"
  >
    <el-form :inline="false" :model="couponInfo" label-width="74px">
      <el-form-item label="金额">
        <el-input type="number" v-model="couponInfo.amount"/>
      </el-form-item>
      <el-form-item label="券总数">
        <el-input type="number" v-model="couponInfo.coupon_total"/>
      </el-form-item>
      <el-form-item label="启用金额">
        <el-input type="number" v-model="couponInfo.min_amount"/>
      </el-form-item>
      <el-form-item label="开始时间">
        <el-date-picker
            v-model="couponInfo.begin_time"
            type="date"
            placeholder="Pick a day"
        />
      </el-form-item>
      <el-form-item label="结束时间">
        <el-date-picker
            v-model="couponInfo.end_time"
            type="date"
            placeholder="Pick a day"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onCouponUpdateSubmit">保存</el-button>
        <el-button @click="onConponUpdateReset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-drawer>


  <!--标签修改对话框-->
  <el-dialog
      v-model="tagDialogVisible"
      title="添加标签"
      width="80%"
  >
    <div v-for="(tags, groupName) in tagList" :key="groupName">
      <div class="tag-group-name">{{ groupName }}</div>
      <div>
        <el-button @click="changeTagPosition(tag.id)" plain :type="tag.selected ? 'primary' : 'info'"
                   style="margin-right: 4px"
                   size="small" v-for="tag in tags"
        >{{ tag.title }}
        </el-button>
      </div>
    </div>


    <template #footer>
      <span class="dialog-footer">
        <el-button @click="tagReset">重置</el-button>
        <el-button type="primary" @click="tagCommit">确认</el-button>
      </span>
    </template>
  </el-dialog>

  <div class="fixed-layout">
    <el-container>
      <el-main style="padding-top: 4px">
        <div class="content">
          <el-row>
            <el-col :span="24">
              <el-card body-style="padding:10px" style="margin-bottom: 2px" class="box-card">
                <el-button @click="showBaseUpdateVisible">
                  基础
                </el-button>

                <el-button @click="showCouponUpdateVisible">
                  优惠券
                </el-button>


              </el-card>

            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="8">
              <el-card class="box-card" :body-style="{ padding: '5px' }">
                <div class="fixed-image-container">
                  <img class="fixed-image" :src="goodsInfo.coverImageUrl" alt="Product Image"/>
                </div>
                <div style="margin-top: 10px; display: flex; align-items: center;">
                  <el-scrollbar :style="{ width: computedWidth, overflowX: 'auto' }">
                    <div style="height: 74px; display: flex; min-width: max-content;">
                      <el-dropdown v-for="item in goodsInfo.images" :key="item.id">
                        <div class="fixed-image-container_sub">
                          <img
                              :style="item.id === coverImageID ? 'border: 2px var(--el-color-primary) solid; box-sizing: border-box;' : ''"
                              class="fixed-image-sub"
                              :src="item.url"
                          />
                        </div>
                        <template #dropdown>
                          <el-dropdown-menu>
                            <el-dropdown-item icon="edit" @click="setCoverImage(id, item.id)">封面</el-dropdown-item>
                            <el-dropdown-item icon="delete" @click="deleteImage(item.id, item.url)">删除</el-dropdown-item>
                          </el-dropdown-menu>
                        </template>
                      </el-dropdown>
                    </div>
                  </el-scrollbar>

                  <div style="height: 74px; margin-left: 8px">
                    <el-upload
                        :show-file-list="false"
                        :action="`${path}/goods/image/upload?id=${id}`"
                        :on-error="uploadError"
                        :on-success="uploadSuccess"
                        :before-upload="beforeImageUpload"
                        :multiple="false"
                    >
                      <el-icon style="width: 62px; height: 62px;background-color: #eee; border-radius: 6px;">
                        <Plus/>
                      </el-icon>
                    </el-upload>
                  </div>
                </div>
              </el-card>


              <el-card class="box-card" :body-style="{ padding: '15px' }">
                <div>
                  <div>店铺信息</div>
                  <div class="shop-info">
                    <img :src="goodsInfo.shopLogo"
                         style="display: inline; width: 18px; height: 18px; margin-right: 10px"
                    />
                    <span style="font-size: 16px; width: 116px">{{ goodsInfo.shopName }}</span>
                    <div class="shop-info-score">
                      <span style="margin-left: 10px">商品:{{ goodsInfo.shopSPF }}</span>
                      <span style="margin-left: 10px">服务:{{ goodsInfo.shopFWF }}</span>
                      <span style="margin-left: 10px">物流:{{ goodsInfo.shopWlF }}</span>
                    </div>
                  </div>
                </div>
                <div style="margin-top: 25px">
                  <div>团长信息</div>
                  <div class="tz-info">
                    <img :src="goodsInfo.tzLogo"
                         style="display: inline; width: 18px; height: 18px; margin-right: 10px"
                    />
                    <span style="font-size: 16px; width: 116px">{{ goodsInfo.tzName }}</span>
                    <div class="tz-info-level">
                      <span style="margin-left: 10px">Lv.{{ goodsInfo.tzLevel }}</span>
                    </div>
                  </div>
                </div>
              </el-card>
            </el-col>
            <el-col :span="16">
              <el-card class="box-card" style="width: 100%; height: 644px">
                <div>
                  <div class="product-details-bt">
                    <el-button
                        color="#FF6545"
                        size="small"
                        class="goods-brand"
                        v-if="goodsInfo.brandName"
                    >{{ goodsInfo.brandName }}
                    </el-button>
                    <el-button link class="goods-title">{{ goodsInfo.goodsTitle }}</el-button>
                  </div>
                  <div class="product-details-des">
                    <span>
                      {{ goodsInfo.goodsDes }}
                    </span>
                  </div>
                  <div class="product-details-tag">
                    <el-tag v-for="item in goodsInfo.tags" :key="item.id" size="small" type="danger"
                            style="margin-top: 3px; margin-left: 0"
                    >{{ item.title }}
                    </el-tag>
                  </div>
                  <div class="goods-pcc">
                    <table>
                      <tr class="goods-pcc-tr">
                        <td class="goods-pcc-td-1">原价:</td>
                        <td class="goods-pcc-td-2"
                            :style="{ textDecoration: couponStatus === 1 ? 'line-through' : 'none' }"
                        >￥{{ goodsInfo.price }}
                        </td>
                      </tr>
                      <tr class="goods-pcc-tr" v-if="goodsInfo.coupon != null">
                        <td class="goods-pcc-td-1"
                            :style="{ textDecoration: couponStatus !== 1 ? 'line-through' : 'none' }"
                        >券后价:
                        </td>
                        <td class="goods-pcc-td-2">￥{{ goodsInfo.postCouponPrice }}</td>
                      </tr>
                      <tr class="goods-pcc-tr">
                        <td class="goods-pcc-td-1">优惠券:</td>
                        <td class="goods-pcc-td-2">
                          <span v-if="couponStatus !==0"
                                :style="{ textDecoration: couponStatus !== 1 ? 'line-through' : 'none' }"
                          >￥ {{ goodsInfo.couponValue }}</span>

                          <span v-else>none</span>

                        </td>
                        <td class="goods-pcc-td-3">
                          {{ goodsInfo.couponBeginTime }} ~ {{ goodsInfo.couponEndTime }}
                          <span v-if="couponStatus === 2">(过期)</span>
                          <span v-if="couponStatus === 3">(未开始)</span>
                        </td>
                      </tr>

                      <tr class="goods-pcc-tr">
                        <td class="goods-pcc-td-1">佣金:</td>
                        <td class="goods-pcc-td-2">{{ goodsInfo.cr }}%</td>
                        <td class="goods-pcc-td-3-cv">约 {{ goodsInfo.cv }} 元</td>
                      </tr>
                    </table>
                  </div>
                  <div class="goods-sale">
                    <table>
                      <tr class="goods-sale-tr">
                        <td class="goods-sale-td-1">
                          <span style="margin-right: 10px;">2小时销量:</span>
                          <span class="goods-sale-span">{{ goodsInfo.sale2h }}</span>
                        </td>

                        <td class="goods-sale-td-1">
                          <span style="margin-right: 10px;">今日销量:</span>
                          <span class="goods-sale-span">{{ goodsInfo.saleDay }}</span>
                        </td>

                        <td class="goods-sale-td-1">
                          <span style="margin-right: 10px;">销量:</span>
                          <span class="goods-sale-span">{{ goodsInfo.saleAll }}</span>
                        </td>
                      </tr>

                    </table>

                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </div>
      </el-main>
    </el-container>
  </div>
</template>


<style scoped>


.fixed-layout {
  width: 1200px; /* 固定总宽度 */
  height: 75vh; /* 固定总高度 */
  overflow: auto; /* 允许滚动 */
  margin: 0 auto; /* 居中显示 */
}

.content {
  display: flex;
  flex-direction: column;

}

.box-card {
  border: 1px solid #EBEEF5;
  border-radius: 4px;
  margin-bottom: 10px;
}

.fixed-image-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 365px; /* 固定宽度 */
  height: 365px; /* 固定高度 */
  margin: 0 auto; /* 居中显示 */
}

.fixed-image-container_sub {
  margin-right: 8px;
  width: 62px; /* 固定宽度 */
  height: 62px; /* 固定高度 */
}

.fixed-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 10px;
}


.fixed-image-sub {
  width: 100%;
  height: 100%;

  margin-right: 80px;
  border-radius: 6px;
}


.shop-info {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin-top: 15px;
}

.shop-info-score {
  margin-left: 40px;
  font-size: 13px;
}

.tz-info {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin-top: 15px;
}

.tz-info-level {
  margin-left: 40px;
  font-size: 13px;
}

.goods-brand {
  font-size: 18px;
  color: #fff;
  max-width: 140px;
  margin-right: 6px;
  padding: 4px;
  margin-left: 0;
}

.goods-title {
  font-size: 18px;
  max-width: 500px;
  font-family: "Microsoft YaHei", "Helvetica Neue", Arial, sans-serif; /* 微软雅黑在某些情况下看起来可能类似于黑体，然后是Helvetica Neue和Arial等无衬线字体 */
  color: #000;
  padding-left: 0;
}


.product-details-bt {
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.product-details-des {
  font-size: 13px;
  margin-top: 10px;
  color: #999999;
}


.product-details-tag {
  margin-top: 16px;
}

.goods-pcc {
  margin-top: 16px;
  color: #999999;
}

.goods-pcc-tr {
  height: 40px;
}

.goods-pcc-td-1 {
  width: 80px;
}

.goods-pcc-td-2 {
  width: 60px;
  font-size: 20px;
  color: #FF6545;
}

.goods-pcc-td-3-cv {
  color: #FF6545;
}

.goods-sale {
  margin-top: 10px;
  color: #999999;
}

.goods-sale-td-1 {
  width: 160px;
}

.goods-sale-span {
  color: #FF6545;
}

.tag-group-name {
  margin-bottom: 6px;
  font-weight: bold;
  margin-top: 6px;
}

</style>
