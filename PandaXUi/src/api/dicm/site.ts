import request from "@/utils/request";

// 查询站点列表
export function listSites(query:any) {
    return request({
        url: '/dicm/sites/list',
        method: 'get',
        params: query
    })
}

// 查看站点详情
export function listSiteInfo(name:any) {
    return request({
        url: '/dicm/sites/'+name,
        method: 'get',
    })
}

// 添加站点
export function addSite(data: any) {
    return request({
        url: '/dicm/sites',
        method: 'post',
        data: data
    })
}

// 修改站点实例
export function updateSite(data:any) {
    return request({
        url: '/dicm/sites',
        method: 'put',
        data: data
    })
}

// 站点归档
export function deleteSite(data:any) {
    return request({
        url: '/dicm/sites/'+data,
        method: 'put',
    })
}

// 站点归属到站点组
export function groupAddSite(data:any) {
	return request({
		url: '/dicm/sites/group/'+data.name,
		method: 'put',
		data: data
	})
}


/** 站点组API */
export function listSiteGroup(query:any) {
    return request({
        url: '/dicm/sites/group/list',
        method: 'get',
        params: query
    })
}

// 获取站点组组组织树
export function siteGroupTree() {
    return request({
        url: '/dicm/sites/groupTree',
        method: 'get',
    })
}

// 添加站点组
export function addSiteGroup(data: any) {
    return request({
        url: '/dicm/sites/group',
        method: 'post',
        data: data
    })
}

// 修改站点组实例
export function updateSiteGroup(data:any) {
    return request({
        url: '/dicm/sites/group',
        method: 'put',
        data: data
    })
}

// 站点组————归档
export function deleteSiteGroup(data:any) {
    return request({
        url: '/dicm/sites/group/'+data,
        method: 'put',
    })
}

/** Location */
// 物理位置分页列表
export function listLocations(query:any) {
    return request({
        url: '/dicm/locations/list',
        method: 'get',
        params: query
    })
}

// 渲染物理位置层级结构
export function locationTree() {
    return request({
        url: '/dicm/locations/tree',
        method: 'get'
    })
}