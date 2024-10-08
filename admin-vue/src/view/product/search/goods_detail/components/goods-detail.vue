<script setup>
import {ref,reactive, onMounted} from 'vue';
import {useRoute} from "vue-router";
import {getGoodsInfo} from "@/api/product";


const route = useRoute();
const id = route.query.id;
const goodsInfo = reactive({
  goodsTitle: null,
  goodsDes: null,
  images: [],
  coverImageUrl:"",
  shopName: null,
  shopWlF: null,
  shopFWF: null,
  shopSPF: null,
  shopLogo: null,
  tzName: null,
  tzLevel: null,
  brandName:null,
  tags:[],

  cv: null,
  cr: null,

  saleAll: null,
  sale2h: null,
  saleDay: null,

  coupon: null,
  couponValue: null,
  couponEndTime: null
})

const priceName = ref("原价:")
const priceValue= ref(0)

const getGoodsDetail = async () => {
  const res = await getGoodsInfo({id: id});
  goodsInfo.coverImageUrl = res.data.image.url
  goodsInfo.shopName = res.data.shop.name;
  goodsInfo.shopFWF = res.data.shop.service_score;
  goodsInfo.shopSPF = res.data.shop.product_score;
  goodsInfo.shopWlF = res.data.shop.logistics_score;
  goodsInfo.shopLogo = res.data.shop.logo;

  goodsInfo.tzName = res.data.media.media_name;
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
  goodsInfo.images = res.data.images;
  goodsInfo.tags = res.data.tags;

  if (res.data.coupon.ID !== 0) {
    goodsInfo.coupon = res.data.coupon;
    priceName.value ="券后价:"
    priceValue.value = res.data.post_coupon_price
    goodsInfo.couponEndTime = formatDateTime(res.data.coupon.end_time) ;
    goodsInfo.couponValue = res.data.coupon.amount
  }else {
    priceValue.value = res.data.price;
  }



  goodsInfo.cv = res.data.commission_value;
  goodsInfo.cr= res.data.commission_rate;

  goodsInfo.sale2h = res.data.sales_2_hour
  goodsInfo.saleDay = res.data.sales_day
  goodsInfo.saleAll = res.data.sales_all
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



onMounted(()=>{
  getGoodsDetail()
})
</script>


<template>
  <div class="fixed-layout">
    <el-container>
      <el-main>
        <div class="content">
          <el-row :gutter="10">
            <el-col :span="8">
              <el-card class="box-card" :body-style="{ padding: '5px' }">
                <div class="fixed-image-container">
                  <img class="fixed-image" :src="goodsInfo.coverImageUrl" alt="Product Image" />
                </div>
                <div style="margin-top: 10px">
                  <el-scrollbar>
                    <div style="height: 74px">
                      <div class="fixed-image-container_sub">
                        <img v-for="item in goodsInfo.images" :key="item.id" :src="item.url" class="fixed-image-sub"/>
                      </div>
                    </div>
                  </el-scrollbar>
                </div>
              </el-card>
              <el-card class="box-card" :body-style="{ padding: '15px' }">
                <div>
                  <div>店铺信息</div>
                  <div class="shop-info">
                    <img :src="goodsInfo.shopLogo"
                         style="display: inline; width: 18px; height: 18px; margin-right: 10px"
                    />
                    <span style="font-size: 16px; width: 116px">{{goodsInfo.shopName}}</span>
                    <div class="shop-info-score">
                      <span style="margin-left: 10px">商品:{{ goodsInfo.shopSPF}}</span>
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
                    <span style="font-size: 16px; width: 116px">{{goodsInfo.tzName}}</span>
                    <div class="tz-info-level">
                      <span style="margin-left: 10px">Lv.{{goodsInfo.tzLevel}}</span>
                    </div>
                  </div>
                </div>
              </el-card>
            </el-col>
            <el-col :span="16">
              <el-card class="box-card" style="width: 98%; height: 624px">
                <div>
                  <div class="product-details-bt">
                    <el-button
                        color="#FF6545"
                        size="small"
                        class="goods-brand"
                        v-if='goodsInfo.brandName'
                    >{{goodsInfo.brandName}}</el-button>
                    <el-button link class="goods-title">{{goodsInfo.goodsTitle}}</el-button>
                  </div>
                  <div class="product-details-des">
                    <span>
                      {{goodsInfo.goodsDes}}
                    </span>
                  </div>

                  <div class="product-details-tag">
                    <el-tag v-for="item in goodsInfo.tags" :key="item.id" class="ml-2" size="small" type="danger" style="margin-top: 3px">{{item.title}}</el-tag>
                  </div>


                  <div class="goods-pcc">
                    <table>
                      <tr class="goods-pcc-tr">
                        <td class="goods-pcc-td-1">{{priceName}}</td>
                        <td class="goods-pcc-td-2">￥{{priceValue}}</td>
                        <!--                        <td class="goods-pcc-td-3">333</td>-->
                      </tr>

                      <tr class="goods-pcc-tr">
                        <td class="goods-pcc-td-1">优惠券:</td>
                        <td class="goods-pcc-td-2">
                          <span v-if="goodsInfo.coupon != null">￥ {{goodsInfo.couponValue}}</span>
                          <span v-else>none</span>

                        </td>
                        <td class="goods-pcc-td-3">
                          {{ goodsInfo.couponEndTime }}
                        </td>
                      </tr>

                      <tr class="goods-pcc-tr">
                        <td class="goods-pcc-td-1">佣金:</td>
                        <td class="goods-pcc-td-2">{{goodsInfo.cr}}%</td>
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
                          <span class="goods-sale-span">{{goodsInfo.saleDay}}</span>
                        </td>

                        <td class="goods-sale-td-1">
                          <span style="margin-right: 10px;">销量:</span>
                          <span class="goods-sale-span">{{goodsInfo.saleAll}}</span>
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
  display: flex;
  justify-content: flex-start;
  align-items: center;
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
  object-fit: cover;
  margin-right: 14px;
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
}

.goods-title {
  font-size: 18px;
  margin-left: 16px;
  max-width: 500px;
  font-family: "Microsoft YaHei", "Helvetica Neue", Arial, sans-serif; /* 微软雅黑在某些情况下看起来可能类似于黑体，然后是Helvetica Neue和Arial等无衬线字体 */
  color: #000;
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

</style>
