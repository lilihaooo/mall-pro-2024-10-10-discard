import service from '@/utils/request'


export const getGoodsListV2 = (params) => {
    return service({
        url: '/goods/v2/search/list',
        method: 'get',
        params
    })
}


export const GetSuggestionList = (params) => {
    return service({
        url: '/search/suggestion/list',
        method: 'get',
        params
    })
}

export const GetHotSearchList = () => {
    return service({
        url: '/search/hotSearch/list',
        method: 'get',
    })
}

// 商品收藏
export const ApiCollect =  (userID, goodsID) => {
    return service({
        url: '/goods/collect',
        method: 'post',
        data: {
            user_id: userID,
            goods_id: goodsID
        }
    })
}

// 商品取消收藏
export const ApiCancelCollect =  (userID, goodsID) => {
    return service({
        url: '/goods/cancelCollect',
        method: 'post',
        data: {
            user_id: userID,
            goods_id: goodsID
        }
    })
}
export const ApiBatchCancelCollect =  (ids) => {
    return service({
        url: '/goods/batchCancelCollect',
        method: 'post',
        data: {
            ids: ids,
        }
    })
}



export const ApiMyCollect =  (page) => {
    return service({
        url: '/goods/myCollect',
        method: 'get',
        data: {
            page: page,
        }
    })
}

// 商品推广
export const ApiPromotion = (userID, goodsID) => {
    return service({
        url: '/goods/promotion',
        method: 'post',
        data: {
            user_id: userID,
            goods_id: goodsID
        }
    })
}


export const getGoodsListV1 = (params) => {
    return service({
        url: '/goods/search/list',
        method: 'get',
        params
    })
}


export const getGoodsInfo = (params) => {
    return service({
        url: '/goods/info',
        method: 'get',
        params
    })
}
export const SetCoverImage = (goodsID, imageID) => {
    return service({
        url: '/goods/image/setCover',
        method: 'post',
        data: {
            image_id: imageID,
            goods_id: goodsID
        }
    })
}


export const DeleteImage = (params) => {
    return service({
        url: '/goods/image/delete',
        method: 'get',
        params: params
    })
}


export const getCategoryTree = () => {
    return service({
        url: '/category',
        method: 'get',
    })
}

export const getGrabRecord = (params) => {
    return service({
        url: '/grab/record',
        method: 'get',
        params
    })
}


export const getGrabLog = (params) => {
    return service({
        url: '/grab/log',
        method: 'get',
        params
    })
}


export const getShopList = (params) => {
    return service({
        url: '/shop/list',
        method: 'get',
        params
    })
}


export const GetTagList = () => {
    return service({
        url: '/tag/list',
        method: 'get',
    })
}

export const GetBrandList = () => {
    return service({
        url: '/brand/list',
        method: 'get',
    })
}


export const UpdateGoodsBaseInfo = (data) => {
    return service({
        url: '/goods/updateBaseInfo',
        method: 'post',
        data: data
    })
}

export const UpdateGoodsCouponInfo = (data) => {
    return service({
        url: '/goods/updateCouponInfo',
        method: 'post',
        data: data
    })
}