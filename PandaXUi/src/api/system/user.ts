import request from '@/utils/request';
import { pa } from 'element-plus/es/locale';


// 查询用户列表
export function authUser(query: any) {
	return request({
		url: '/system/user/auth',
		method: 'get',
		params: query
	})
}

// 查询用户列表
export function listUser(query: any) {
	return request({
		url: '/system/user/list',
		method: 'get',
		params: query
	})
}

// 用户状态修改
export function changeUserStatus(userId: number, status: string) {
	const data = {
		userId,
		status
	}
	return request({
		url: '/system/user/changeStatus',
		method: 'put',
		data: data
	})
}

// 查询用户详细
export function getUser(userId: number) {
	return request({
		url: '/system/user/getById/' + userId,
		method: 'get'
	})
}

// 添加时查询用户详细
export function getUserInit() {
	return request({
		url: '/system/user/getInit',
		method: 'get'
	})
}

// 添加时查询用户ROLE
export function getRoPo() {
	return request({
		url: '/system/user/getRoPo',
		method: 'get'
	})
}
// 删除用户
export function delUser(userId: number) {
	return request({
		url: '/system/user/' + userId,
		method: 'delete'
	})
}

// 新增用户
export function addUser(data: any) {
	return request({
		url: '/system/user',
		method: 'post',
		data: data
	})
}

// 新增员工
export function addEmployee(data: any) {
	return request({
		url: '/system/user/employee',
		method: 'post',
		data: data
	})
}

// 修改用户
export function updateUser(data:any) {
	return request({
		url: '/system/user',
		method: 'put',
		data: data
	})
}

// 修改员工
export function updateEmployee(data:any) {
	return request({
		url: '/system/user/employee',
		method: 'put',
		data: data
	})
}

// 修改密码
export function updateUserPwd(data:any) {
	return request({
		url: '/system/user/pwd',
		method: 'put',
		data: data
	})
}

// 创建TOTP记录
// export function createTotp(userId: any) {
// 	return request({
// 	  url: '/auth/totp/enable/'+ userId,
// 	  method: 'get'
// 	})
//   }

// 启用用户TOTP cg
export function enableTotp(params: any) {
	const id = params.value.id
	const data = {
		name:params.value.login,
		password:params.password
	}
	console.log("前端api=%T %v",id,data);
	return request({
	  url: '/system/user/totp/enable/'+ id,
	  method: 'get',
	  params: data
	})
  }
  
  // 关闭用户TOTP
  export function closeTotp(userId: any) {
	return request({
	  url: '/system/user/totp/disable/'+ userId,
	  method: 'delete'
	})
  }
  
  // 用户密码重置 cg
  export function resetUserPwd(data: any) {
	return request({
	  url: '/system/user/resetPwd',
	  method: 'put',
	  data: data
	})
  }

  // 查询部门下拉树结构
export function deptTreeSelect() {
	return request({
	  url: '/system/user/deptTree',
	  method: 'get'
	})
  }
  
  // 查询公司下拉树结构
  export function compTreeSelect() {
	return request({
	  url: '/system/user/compTree',
	  method: 'get'
	})
  }


// 导出
export function exportUser(query: any) {
	return request({
		url: '/system/user/export',
		method: 'get',
		params: query,
		responseType: 'blob'
	})
}

// 下载用户导入模板
export function importTemplate() {
	return request({
		url: '/system/user/importTemplate',
		method: 'get'
	})
}

// 用户头像上传
export function uploadAvatar(data:any) {
	return request({
		url: '/system/user/avatar',
		method: 'post',
		data: data
	})
}