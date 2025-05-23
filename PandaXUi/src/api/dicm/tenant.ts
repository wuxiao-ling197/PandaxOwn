import request from "@/utils/request";

// 查询租户列表
export function listTenant(query:any) {
    return request({
        url: '/tenant/list',
        method: 'get',
        params: query
    })
}

// 查看租户详情
export function listTenantInfo(name:any) {
    return request({
        url: '/tenant/'+name,
        method: 'get',
    })
}

// 新增租户实例
export function addTenant(data: any) {
	return request({
		url: '/tenant',
		method: 'post',
		data: data
	})
}

// 修改租户实例
export function updateTenant(data:any) {
	return request({
		url: '/tenant',
		method: 'put',
		data: data
	})
}

// 租户————归档
export function deleteTenant(data:any) {
	return request({
		url: '/tenant/'+data,
		method: 'put',
		// data: data
	})
}

// 租户————租户归属到租户组
export function addTenant2Group(data:any) {
	return request({
		url: '/tenant/add2group',
		method: 'put',
		data: data
	})
}

// 租户组————租户组添加租户
export function groupAddTenant(data:any) {
	return request({
		url: '/tenant/group/'+data.name,
		method: 'put',
		data: data
	})
}

// 查询租户组列表
export function listTenantGroup(query:any) {
    return request({
        url: '/tenant/group/list',
        method: 'get',
        params: query
    })
}

// 查看租户组详情
export function listTenantGroupInfo(name:any) {
    return request({
        url: '/tenant/group/'+name,
        method: 'get',
    })
}

// 新增租户组实例
export function addTenantGroup(data: any) {
	return request({
		url: '/tenant/group',
		method: 'post',
		data: data
	})
}

// 修改租户组实例
export function updateTenantGroup(data:any) {
	return request({
		url: '/tenant/group',
		method: 'put',
		data: data
	})
}

// 租户组————归档
export function deleteTenantGroup(data:any) {
	return request({
		url: '/tenant/group/'+data,
		method: 'put',
	})
}

// 获取租户组组织树
export function tenantGroupTree() {
    return request({
        url: '/tenant/groupTree',
        method: 'get',
    })
}