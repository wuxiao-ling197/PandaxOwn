import request from '@/utils/request';


/**
 * 获取验证码
 * @param params 要传的参数值
 * @returns 返回接口数据
 */
export function captcha() {
	return request({
		url: '/system/user/getCaptcha',
		method: 'get',
	});
}

export function getTotp() {
	return request({
		url: '/system/user/getTotp',
		method: 'get',
	})
}

// 验证totp
export function valideTotp(params: object) {
	return request({
		url: '/system/user/valideTotp',
		method: 'post',
		data: params
	});
}

// 激活totp
export function totpEnableone(params: object) {
	return request({
		url: '/system/user/enableTotp',
		method: 'post',
		data: params
	});
}

// 激活totp
export function totpReset(params: object) {
	return request({
		url: '/system/user/resetTotp',
		method: 'post',
		data: params
	});
}

// export function getLoginUser(params: object) {
// 	return request({
// 		url: '/system/user/getLoginUser',
// 		method: 'post',
// 		data: params,
// 	})
// }

/**
 * 用户登录路由
 * @param params 要传的参数值
 * @returns 返回接口数据
 */
export function signIn(params: object) {
	return request({
		url: '/system/user/login',
		method: 'post',
		data: params,
	});
}

/**
 * 用户退出登录
 * @param params 要传的参数值
 * @returns 返回接口数据
 */
export function signOut(params: object) {
	return request({
		url: '/system/user/logout',
		method: 'post',
		data: params,
	});
}
