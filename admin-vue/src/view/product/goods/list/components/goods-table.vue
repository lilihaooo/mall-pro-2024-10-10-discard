<script setup>
import router from "@/router";
const props = defineProps({
  data: {
    type: Array,
    required: true
  }
});


const goodsManage =(id)=>{
  router.push({
    name:'goodsManage',
    query:{id:id},
  })
}


</script>

<template>
<div>
  <el-table
      ref="multipleTable"
      style="width: 100%"
      tooltip-effect="dark"
      :data="data"
      row-key="ID"
      max-height="500px"
  >
    <el-table-column type="selection" width="55"/>
    <el-table-column align="center" label="商品ID" width="80" prop="id"/>
    <el-table-column align="center" label="商品名称" width="260" prop="title"/>
    <el-table-column align="center" label="分类" min-width="210" prop="category_name"/>
    <el-table-column align="center" label="图片" width="240" >
      <template #default="scope">
        <el-image
            style="width: 120px;"
            :src="scope.row.cover_image"
        />
      </template>
    </el-table-column>
    <el-table-column align="center" label="店铺名" width="140" prop="shop_name"/>
    <el-table-column align="center" label="价格" width="120" prop="price"/>

    <el-table-column align="center" label="券后价"  min-width="120" fixed="right" >
      <template #default="scope">
        <span v-if="scope.row.coupon_id !== 0">{{scope.row.post_coupon_price}}</span>
        <span v-else>none</span>
      </template>
    </el-table-column>

    <el-table-column align="center" label="状态" width="120" prop="status"/>
    <el-table-column align="center" label="佣金比例" prop="commission_rate" width="120"/>
    <el-table-column align="center" label="佣金" width="120" prop="commission_value"/>
    <el-table-column align="center" label="近2小时销量" min-width="120" prop="sales_2_hour"/>
    <el-table-column align="center" label="今日销量" min-width="120" prop="sales_day"/>
    <el-table-column align="center" label="月销量" min-width="120" prop="sales_month"/>
    <el-table-column align="center" label="全部销量" min-width="120" prop="sales_all"/>
    <el-table-column align="center" label="平台" min-width="120" prop="platform_name"/>
    <el-table-column align="center" label="库存" min-width="120" prop="inventory"/>
    <el-table-column align="center" label="体验分" min-width="120" prop="experience_score"/>

    <el-table-column align="center" label="媒体"  min-width="120"  >
      <template #default="scope">
        <span v-if="scope.row.media_name !== ''">{{scope.row.media_name}}</span>
        <span v-else>none</span>
      </template>
    </el-table-column>

    <el-table-column align="center" label="优惠券ID"  min-width="120"  >
      <template #default="scope">
        <span v-if="scope.row.coupon_id !== 0">{{scope.row.coupon_id}}</span>
        <span v-else>none</span>
      </template>
    </el-table-column>

    <el-table-column align="center" label="券面额" min-width="120" prop="coupon_amount" fixed="right" />
    <el-table-column align="center" label="到期时间" min-width="220" prop="coupon_end_time"/>
    <el-table-column align="center" label="领券量" min-width="120" prop="coupon_cover"/>
    <el-table-column align="center" label="券总量" min-width="120" prop="coupon_total"/>
    <el-table-column align="center" label="推广热度" min-width="120" prop="push_hot"/>

    <el-table-column align="center" label="操作" fixed="right" min-width="180">
      <template #default="scope">
        <el-button type="primary" link @click="goodsManage(scope.row.id)">
          <el-icon><More /></el-icon>
          管理
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</div>

</template>
<style scoped lang="scss">

</style>