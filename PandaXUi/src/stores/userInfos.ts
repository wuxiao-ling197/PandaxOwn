import { defineStore } from 'pinia';
import { UserInfosState } from 'storeInterface';
import { Session } from '@/utils/storage';
import Cookies from 'js-cookie';
import {letterAvatar} from "@/utils/string";
import { authUser } from "@/api/system/user";

// 保存当前登录用户信息
export const useUserInfosState = defineStore('userInfos', {
	state: (): UserInfosState => ({
		userInfos: {
			isTotp: false,
			username: "",
			login:"",
			employeeType:"",
			photo: '',
			time: 0,
			authBtnList: [],

			userId: 0,
			roleId: 0,
			organizationId: 0,
			companyId:0,
			postId: 0,
			jobTitle:"",
			lastCheckIn:"",
			workEmail:"",
			workPhone:"",
			signature:"",
			// jobTitle:"",

			lastLoginTime: 0,
			lastLoginIp: "127.0.0.1",
		},
	}),
	actions: {
		// 设置用户信息
		getUserInfos( data: object) {
			this.userInfos = data;
		},
		// 设置用户信息
		async setUserInfos() {
			const userName = Cookies.get('userName');			

			let response = await authUser({"username": userName})
			
			let loginRes = response.data
			Session.set("menus", loginRes.menus);
			let perms = loginRes.permissions;
			perms.push("base");
			// 用户信息模拟数据
			const userInfos = {
				username: loginRes.user.employee.name,
				login: loginRes.login,
				userId: loginRes.user.id,
				companyId:loginRes.user.company_id,
				roleId: loginRes.user.roleId,
				organizationId: loginRes.user.employee.department_id,
				jobTitle: loginRes.user.employee.job_title,
				lastCheckIn: loginRes.user.employee.last_check_in, //最新签到时间
				workEmail: loginRes.user.employee.work_email,
				workPhone: loginRes.user.employee.work_phone,
				employeeType: loginRes.user.employee.employee_type,
				signature: loginRes.user.signature,
				postId: loginRes.user.postId,
				// 头像
				photo: loginRes.user.avatar || letterAvatar(loginRes.user.username),
				time: new Date().getTime(),
				lastLoginTime: new Date().getTime(),
				lastLoginIp: "127.0.0.1",
				//authPageList: perms,
				authBtnList: perms,
			};
			// 存储用户信息到浏览器缓存
			Session.set('userInfo', userInfos);


			if (Session.get('userInfo')) {
				this.getUserInfos(Session.get('userInfo'));
			} else {
				this.getUserInfos(userInfos);
			}
		},
		async setUserInfo() {
			this.getUserInfos(Session.get('userInfo'));
		}

	},
});
