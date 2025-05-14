import request from "@/utils/request";

// 查询机柜列表
export function listRacks(query:any) {
    return request({
        url: '/dicm/racks/list',
        method: 'get',
        params: query
    })
}

// 查询机柜类型列表
export function listRackRole(query:any) {
    return request({
        url: '/dicm/rackrole/list',
        method: 'get',
        params: query
    })
}

// 查询机柜预留列表
export function listRackReserve(query:any) {
    return request({
        url: '/dicm/rackreserve/list',
        method: 'get',
        params: query
    })
}

// 新增机柜实例
export function addRacks(data: any) {
	return request({
		url: '/dicm/racks',
		method: 'post',
		data: data
	})
}

// 修改机柜实例
export function updateRacks(data:any) {
	return request({
		url: '/dicm/racks',
		method: 'put',
		data: data
	})
}

// 查看机柜详情
export function listRackInfo(name:any) {
    return request({
        url: '/dicm/racks/'+name,
        method: 'get',
    })
}