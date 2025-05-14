import request from '@/utils/request';

// 查询组织列表
export function listOrganization(query : any) {
	return request({
		url: '/system/organization/list',
		method: 'get',
		params: query
	})
}

//查询公司信息
export function listCompany(query : any) {
	return request({
		url: '/system/organization/company',
		method: 'get',
		params: query
	})
}

// 查询组织详细
export function getOrganization(organizationId: number) {
	return request({
		url: '/system/organization/' + organizationId,
		method: 'get'
	})
}

// 查询组织下拉树结构
export function treeselect() {
	return request({
		url: '/system/organization/organizationTree',
		method: 'get'
	})
}

// 查询部门组织下拉树结构
export function departmentTree() {
	return request({
		url: '/system/organization/departmentTree',
		method: 'get'
	})
}

// 根据角色ID查询组织树结构
export function roleOrganizationTreeselect(roleId: number) {
	return request({
		url: '/system/organization/roleOrganizationTreeSelect/' + roleId,
		method: 'get'
	})
}

export function searchOrganization(organizationId: number) {
	return request({
		url: '/system/organization/'+organizationId,
		method: 'get'
	})
}

// 新增组织
export function addOrganization(data:any) {
	return request({
		url: '/system/organization',
		method: 'post',
		data: data
	})
}

// 修改组织
export function updateOrganization(data:any) {
	return request({
		url: '/system/organization',
		method: 'put',
		data: data
	})
}

// 删除组织
export function delOrganization(organizationId: number) {
	return request({
		url: '/system/organization/' + organizationId,
		method: 'delete'
	})
}
