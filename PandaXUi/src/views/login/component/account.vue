<template>
  <el-form
    ref="loginFormRef"
    size="large"
    :model="state.loginForm"
    :rules="state.rules"
    class="login-content-form"
  >
    <el-form-item class="login-animation-one">
      <el-input
        type="text"
        :placeholder="$t('message.account.accountPlaceholder1')"
        v-model="state.loginForm.username"
        clearable
        autocomplete="off"
        @change="handleUsernameChange"
      >
        <template #prefix>
          <el-icon class="el-input__icon"><elementUser /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation-two">
      <el-input
        :type="state.isShowPassword ? 'text' : 'password'"
        :placeholder="$t('message.account.accountPlaceholder2')"
        v-model="state.loginForm.password"
        autocomplete="off"
        @change="handleUsernameChange"
      >
        <template #prefix>
          <el-icon class="el-input__icon"><elementUnlock /></el-icon>
        </template>
        <template #suffix>
          <i
            class="iconfont el-input__icon login-content-password"
            :class="state.isShowPassword ? 'icon-yincangmima' : 'icon-xianshimima'"
            @click="state.isShowPassword = !state.isShowPassword"
          >
          </i>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation-three">
      <!-- <el-row :gutter="15"> -->
        <!-- <el-col :span="16"> -->
          <!--输入框原参数 maxlength="6" -->
          <el-input
            type="text"
            :placeholder="$t('message.account.accountPlaceholder3')"
            v-model="state.loginForm.passcode"
            clearable
            autocomplete="off"
          >
            <template #prefix>
              <el-icon class="el-input__icon"><elementPosition /></el-icon>
            </template>
          </el-input>
          <!-- add 重置密码 以及重置totp按钮 -->
          <el-link type="info">
            <el-icon class="custom_button"><CirclePlusFilled /></el-icon>注册新用户</el-link>
          <el-link type="info">
            <el-icon class="custom_button" @click="handleResetPwd"><Edit /></el-icon>重置密码</el-link>
          <el-link class="two" type="info"  @click="handleResetTotp">
            <el-icon class="custom_button"><Connection /></el-icon>双重验证</el-link>
        <!-- </el-col> -->
        <!-- <el-col :span="8">
          <div class="login-content-code">
            <img
              class="login-content-code-img"
              @click="getCaptcha"
              width="130px"
              height="38px"
              :src="state.captchaImage"
              style="cursor: pointer"
            />
          </div>
        </el-col> -->
      <!-- </el-row> -->
    </el-form-item>
    <el-form-item class="login-animation-four">
      <el-button
        type="primary"
        class="login-content-submit"
        round
        @click="login"
        :loading="state.loading.signIn"
      >
        <span>{{ $t("message.account.accountBtnText") }}</span>
      </el-button>
    </el-form-item>
  </el-form>

  <!-- todo 修改为totp验证码显示 -->
  <el-dialog v-model="state.isFirstValid" title="请扫描验证码" width="400px" center>
    <img
      class="login-content-code-img"
      :src="state.totpImage"
      center
      style="margin-left: 30px;"
    />
  </el-dialog>
</template>

<script setup lang="ts">
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { onMounted, ref, reactive, computed, getCurrentInstance } from "vue";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { useI18n } from "vue-i18n";
import { initBackEndControlRoutes } from "@/router/index";
import { Session } from "@/utils/storage";
import Cookies from 'js-cookie';
import { captcha, getTotp, signIn, totpEnableone, totpReset, valideTotp } from "@/api/login/index";
import { formatAxis } from "@/utils/formatTime";
import rotate from '@/assets/rotate.png'
import { enableTotp, resetUserPwd } from '@/api/system/user';
import { useUserInfosState } from '@/stores/userInfos';

// 旋转图片滑块组件
// import DragVerifyImgRotate from "@/components/dragVerify/dragVerifyImgRotate.vue";

const { t } = useI18n();
const { proxy } = getCurrentInstance() as any;
const loginFormRef: any = ref(null);
const dragRef: any = ref(null);

const route = useRoute();
const router = useRouter();

const userInfos = useUserInfosState();
const state = reactive({
  dialogVerifyVisible: false,
  imgThree: rotate,
  isFirstValid: false,
  totpImage: undefined,
  loginForm: {
    username: "admin",
    password: "admin",
    passcode: "",
    // codeId: "",
  },
  rules: {
    username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
    password: [{ required: true, message: "请输入密码", trigger: "blur" }],
    passcode: [{ required: true, message: "请输入认证码", trigger: "blur" }],
  },
  isShowPassword: false,
  loading: {
    signIn: false,
  },
});
// 页面刷新加载函数
onMounted(() => {
  // getCaptcha();
  // totpEnable();
});

// zc
const totpEnable = async () =>{
  console.log("表单参数",state.loginForm);
  let res: any = await totpEnableone(state.loginForm);
  console.log("totpEnableone获取totp对象：",res);
  if (res.data!="用户已激活双重验证") {
    state.totpImage = res.data
    state.isFirstValid = true
  }
}
const handleUsernameChange =async () =>{
  totpEnable();
}

const handleResetTotp = async () =>{
  let res:any = await totpReset(state.loginForm);
  if(res.data=="用户双重验证已重置"){
    state.loginForm.username= ""
  state.loginForm.password= ""
  ElMessage.success("已重置双重验证，请再次输入信息并登录！！！");
  }
}

/** 重置密码按钮操作 */
const handleResetPwd = async (value: any) => {

  resetUserPwd(value).then((res: any) => {
    ElMessage.success("重置密码邮件已发送，请注意查收！！！");
  });
};

// 时间获取
const currentTime = computed(() => {
  return formatAxis(new Date());
});
// 校验登录表单并登录 登录按钮
const login = async () => {
  loginFormRef.value.validate((valid: boolean) => {
    if (valid) {
      onSignIn();
    } else {
      return false;
    }
  });
};

const onSignIn = async () => {
  let res: any = await valideTotp(state.loginForm);
  console.log("valideTotp获取验证totp返回结果：",res);
  if (res.data) {
    userInfos.userInfos.isTOTP=true
    state.loading.signIn = true;
  let loginRespon;
  try {
    loginRespon = await signIn(state.loginForm); //调用登录路由
    state.loading.signIn = false
  } catch (e) {
    dragRef.value.reset();
    state.loading.signIn = false;
    state.loginForm.passcode = "";
    state.loading.signIn = false
    return;
  }
  let loginRes = loginRespon.data;
  Session.set("token", loginRes.token);
  Cookies.set('userName', state.loginForm.username);

  // 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
  // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
  await initBackEndControlRoutes();
  // 执行完 initBackEndControlRoutes，再执行 signInSuccess
  signInSuccess();
  }
  
};
// const openVerify = () => {
//   state.dialogVerifyVisible = true;
// };
// const passVerify = () => {
//   login();
// };

// 登录成功后的跳转
const signInSuccess = () => {
  // 初始化登录成功时间问候语
  let currentTimeInfo = currentTime.value;
  // 登录成功，跳到转首页
  // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
  // 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
  if (route.query?.redirect) {
    router.push({
      path: route.query?.redirect,
      query:
        Object.keys(route.query?.params).length > 0
          ? JSON.parse(route.query?.params)
          : "",
    });
  } else {
    router.push("/");
  }
  //登录成功提示
  setTimeout(() => {
    // 关闭 loading
    state.loading.signIn = true;
    const signInText = t("message.signInText");
    ElMessage.success(`${currentTimeInfo}，${signInText}`);
    // 修复防止退出登录再进入界面时，需要刷新样式才生效的问题，初始化布局样式等(登录的时候触发，目前方案)
    proxy.mittBus.emit("onSignInClick");
  }, 300);
};
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
  }
  .login-animation-four {
    animation-delay: 0.4s;
    margin-bottom: 5px;
  }

  .login-content-password {
    display: inline-block;
    width: 25px;
    cursor: pointer;
    &:hover {
      color: #909399;
    }
  }
  .login-content-code {
    display: flex;
    align-items: center;
    justify-content: space-around;
    .login-content-code-img {
      width: 100%;
      height: 40px;
      line-height: 40px;
      background-color: #ffffff;
      border: 1px solid rgb(220, 223, 230);
      color: #333;
      font-size: 16px;
      font-weight: 700;
      letter-spacing: 5px;
      text-indent: 5px;
      text-align: center;
      cursor: pointer;
      transition: all ease 0.2s;
      border-radius: 4px;
      user-select: none;
      &:hover {
        border-color: #c0c4cc;
        transition: all ease 0.2s;
      }
    }
  }
  .login-content-submit {
    width: 100%;
    letter-spacing: 2px;
    font-weight: 300;
    margin-top: 15px;
  }
  .el-link{
    font-size:14px;
    // font-style:italic;
    // color: #49171b;
    margin-left: 30px;
  }
  // .two {
  //     margin-left: 120px;
  //   }
  .custom_button {
    margin-right: 5px;
  }
}
</style>
