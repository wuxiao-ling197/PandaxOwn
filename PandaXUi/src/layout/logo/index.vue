<template>
	<div class="layout-logo" v-if="setShowLogo" @click="onThemeConfigChange">
		<img :src="logo" class="layout-logo-medium-img" />
		<span >{{ getThemeConfig.globalViceTitle }}</span>
	</div>
	<div class="layout-logo-size" v-else @click="onThemeConfigChange">
		<img :src="logo" class="layout-logo-size-img" />
	</div>
</template>

<script lang="ts">
import { computed, getCurrentInstance } from 'vue';
import {useThemeConfigStateStore} from '@/stores/themeConfig'
import logo from '@/assets/iot.png'
export default {
	name: 'layoutLogo',
	setup() {
		const { proxy } = getCurrentInstance() as any;
        const theme = useThemeConfigStateStore();
		// 获取布局配置信息
		const getThemeConfig = computed(() => {
			return theme.themeConfig;
		});
		// 设置 logo 的显示。classic 经典布局默认显示 logo
		const setShowLogo = computed(() => {
			let { isCollapse, layout } = theme.themeConfig;
			return !isCollapse || layout === 'classic' || document.body.clientWidth < 1000;
		});
		// logo 点击实现菜单展开/收起
		const onThemeConfigChange = () => {
			if (theme.themeConfig.layout === 'transverse') return false;
			proxy.mittBus.emit('onMenuClick');
			theme.themeConfig.isCollapse = !theme.themeConfig.isCollapse;
		};
		return {
			logo,
			setShowLogo,
			getThemeConfig,
			onThemeConfigChange,
		};
	},
};
</script>

<style scoped lang="scss">
.layout-logo {
	width: 220px;
	height: 50px;
	display: flex;
	align-items: center;
	justify-content: center;
	box-shadow: rgb(0 21 41 / 2%) 0px 1px 4px;
	color: var(--color-primary);
	font-size: 20px;
	cursor: pointer;
	animation: logoAnimation 0.3s ease-in-out;
	&:hover {
		span {
			color: var(--color-primary-light-2);
		}
	}
	&-medium-img {
		width: 25px;
		margin-right: 15px;
	}
}
.layout-logo-size {
	width: 100%;
	height: 50px;
	display: flex;
	cursor: pointer;
	animation: logoAnimation 0.3s ease-in-out;
	&-img {
		width: 25px;
		margin: auto;
	}
	&:hover {
		img {
			animation: logoAnimation 0.3s ease-in-out;
		}
	}
}
</style>
