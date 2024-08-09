import service from '@/utils/request'



export const getGoodsListV2 = (params) => {
    return service({
        url: '/goods/v2/search/list',
        method: 'get',
        params
    })
}


export const getSuggestionList = (params) => {
    return service({
        url: '/goods/suggestion/list',
        method: 'get',
        params
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