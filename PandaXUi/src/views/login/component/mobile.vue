<template>
	<el-form size="large" class="login-content-form">
		<el-form-item class="login-animation-one">
			<el-input type="text" :placeholder="$t('message.mobile.placeholder1')" v-model="state.ruleForm.username" clearable autocomplete="off">
				<template #prefix>
					<i class="iconfont icon-dianhua el-input__icon"></i>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation-two">
			<el-row :gutter="15">
				<el-col :span="16">
					<el-input type="text" maxlength="4" :placeholder="$t('message.mobile.placeholder2')" v-model="state.ruleForm.code" clearable autocomplete="off">
						<template #prefix>
							<el-icon class="el-input__icon"><elementPosition /></el-icon>
						</template>
					</el-input>
				</el-col>
				<el-col :span="8">
					<el-button class="login-content-code">{{ $t('message.mobile.codeText') }}</el-button>
				</el-col>
			</el-row>
		</el-form-item>
		<el-form-item class="login-animation-three">
			<el-button type="primary" class="login-content-submit" round>
				<span>{{ $t('message.mobile.btnText') }}</span>
			</el-button>
		</el-form-item>
	</el-form>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
const state = reactive({
	ruleForm: {
		username: '',
		code: '',
	},
});
</script>

<style scoped lang="scss">
.login-content-form {
	margin-top: 20px;

	.login-animation-one,
	.login-animation-two,
	.login-animation-three,
	.login-animation-four {
		opacity: 0;
		animation-name: error-num;
		animation-duration: 0.5s;
		animation-fill-mode: forwards;
	}
	.login-animation-one {
		animation-delay: 0.1s;
	}
	.login-animation-two {
		animation-delay: 0.2s;
	}
	.login-animation-three {
		animation-delay: 0.3s;
		margin-bottom: 5px;
	}
	.login-animation-four {
		animation-delay: 0.4s;
	}
	.login-content-code {
		width: 100%;
		padding: 0;
	}
	.login-content-submit {
		width: 100%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}
}
</style>
