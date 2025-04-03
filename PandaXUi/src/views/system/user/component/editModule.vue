<template>
  <div class="system-menu-container">
    <el-dialog v-model="state.isShowDialog" width="769px" center>
      <template #header>
        <div style="font-size: large"
          v-drag="['.system-menu-container .el-dialog', '.system-menu-container .el-dialog__header']">{{ title }}</div>
      </template>
      <!-- <el-form :model="state.ruleForm" label-width="100px"> -->
      <el-row :gutter="20">
        <!-- 左侧员工数据  :rules="ruleRules" -->
        <el-col :span="12">
          <el-card shadow="always" style="height: 100%;">
            <h4>员工信息</h4>
            <el-form :model="state.ruleForm.employee" ref="ruleFormRef" label-width="100px">
              <!-- 等待修改 这里应该是选项 -->
              <el-form-item label="员工姓名" prop="state.ruleForm.employee.name">
                <el-input v-model="state.ruleForm.employee.name" placeholder="请输入员工姓名" maxlength="11" required />
              </el-form-item>
              <el-form-item label="相关用户" prop="state.ruleForm.employee.user_id">
                <el-select v-model="state.ruleForm.employee.user_id" placeholder="请选择相关用户" maxlength="11" />
              </el-form-item>
              <!-- 应该有个专属功能来进行部门调动 -->
              <el-form-item label="所属部门" prop="state.ruleForm.employee.department_id">
                <!-- <el-select v-model="state.ruleForm.employee.department_id" placeholder="请选择所属部门" maxlength="11" /> -->
                <el-tree-select
                v-model="state.ruleForm.employee.department_id"
                :data="deptOptions"
                :props="{ value: 'id', label: 'label', children: 'children' }"
                value-key="id"
                placeholder="请选择归属部门"
                check-strictly
              />
              </el-form-item>
              <el-form-item label="工作岗位" prop="state.ruleForm.employee.job_id">
                <el-select v-model="state.ruleForm.employee.job_id" placeholder="请选择工作岗位" maxlength="11" />
              </el-form-item>
              <el-form-item label="员工性别" prop="state.ruleForm.employee.gender">
                <el-select v-model="state.ruleForm.employee.gender" placeholder="请选择员工性别" maxlength="11" />
              </el-form-item>
              <el-form-item label="员工类型" prop="state.ruleForm.employee.employee_type">
                <el-select v-model="state.ruleForm.employee.employee_type" placeholder="请选择员工类型" maxlength="11" />
              </el-form-item>
              <el-form-item label="员工生日" prop="state.ruleForm.employee.birthday">
                <el-date-picker v-model="state.ruleForm.employee.birthday" type="date" placeholder="请选择员工生日"
                  maxlength="11" />
              </el-form-item>
              <el-form-item label="工作电话" prop="stateruleForm.employee.work_phone">
                <el-input v-model="state.ruleForm.employee.work_phone" placeholder="请输入工作电话" maxlength="11" />
              </el-form-item>
              <el-form-item label="工作邮箱" prop="state.ruleForm.employee.work_email">
                <el-input v-model="state.ruleForm.employee.work_email" placeholder="请输入工作邮箱" maxlength="11" required />
              </el-form-item>
              <el-form-item label="已辞职" prop="state.ruleForm.employee.resigned">
                <el-select v-model="state.ruleForm.employee.resigned" placeholder="请输入辞职原因" maxlength="11" />
              </el-form-item>
              <el-form-item label="已解雇" prop="state.ruleForm.employee.fired">
                <el-select v-model="state.ruleForm.employee.fired" placeholder="请输入解雇原因" maxlength="11" />
              </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                  <el-button @click="onCancel">取 消</el-button>
                  <el-button type="primary" @click="employeeSubmit" :loading="state.loading" style="margin-left: 50px;">保 存</el-button>
                </span>
              </template>
          </el-card>
        </el-col>
        <!-- 右侧用户数据 :rules="userRules"-->
        <el-col :span="12">
          <el-card shadow="always" style="height: 100%;">
            <h4>相关用户</h4>
            <el-form :model="state.ruleForm.user" ref="ruleFormRef" label-width="100px">
              <el-form-item label="登录名" prop="state.ruleForm.user.login">
                <el-input v-model="state.ruleForm.user.login" placeholder="请输入登录名" maxlength="11" required />
              </el-form-item>
              <el-form-item label="所属公司" prop="state.ruleForm.user.company_id">
                <!-- <el-input v-model="state.ruleForm.user.company_id" placeholder="请选择所属公司" maxlength="11" /> -->
                <el-tree-select
                v-model="state.ruleForm.user.company_id"
                :data="compOptions"
                :props="{ value: 'id', label: 'label', children: 'children' }"
                value-key="id"
                placeholder="请选择所属公司"
                check-strictly
              />
              </el-form-item>
              <el-form-item label="用户状态" prop="state.ruleForm.user.active">
                <el-input v-model="state.ruleForm.user.active" placeholder="请输入用户状态" maxlength="11" />
              </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                  <el-button @click="onCancel">取 消</el-button>
                  <el-button type="primary" @click="userSubmit" :loading="state.loading" style="margin-left: 50px;">保 存</el-button>
                </span>
              </template>
          </el-card>
        </el-col>
      </el-row>
      <!-- </el-form> -->

      <!-- <el-form
        ref="ruleFormRef"
        :model="state.ruleForm"
        :rules="state.ruleRules"
        label-width="80px"
      >
        <el-row :gutter="35">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" >
            <el-form-item label="用户登录名" prop="login">
              <el-input
                v-model="state.ruleForm.login"
                placeholder="请输入用户登录名"
              />
            </el-form-item>
          </el-col>
          <el-col v-if="state.ruleForm.userId == undefined" :xs="24" :sm="12" :md="12" :lg="12" :xl="12" >
            <el-form-item

                    label="用户名称"
                    prop="username"
            >
              <el-input
                      v-model="state.ruleForm.username"
                      placeholder="请输入用户名称"
              />
            </el-form-item>
          </el-col>
          <el-col v-if="state.ruleForm.userId == undefined" :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <el-form-item
                    label="用户密码"
                    prop="password"
            >
              <el-input
                      v-model="state.ruleForm.password"
                      placeholder="请输入用户密码"
                      type="password"
              />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <el-form-item label="手机号码" prop="phone">
              <el-input
                v-model="state.ruleForm.phone"
                placeholder="请输入手机号码"
                maxlength="11"
              />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <el-form-item label="邮箱" prop="email">
              <el-input
                v-model="state.ruleForm.email"
                placeholder="请输入邮箱"
                maxlength="50"
              />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" >
            <el-form-item label="用户性别" prop="sex">
              <el-select v-model="state.ruleForm.sex" placeholder="请选择">
                <el-option
                  v-for="dict in state.sexOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                ></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <el-form-item label="角色" prop="roleIds">
              <el-select v-model="state.ruleForm.roleIds" multiple placeholder="请选择">
                <el-option
                        v-for="item in state.roleOptions"
                        :key="item.roleId"
                        :label="item.roleName"
                        :value="item.roleId"
                        :disabled="item.status == 1"
                ></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" >
            <el-form-item label="归属组织" prop="organizationId">
              <el-cascader
                  v-model="state.ruleForm.organizationId"
                  :options="state.organizationOptions"
                  :props="{
                  label: 'organizationName',
                  value: 'organizationId',
                  checkStrictly: true,
                  emitPath: false,
                }"
                  class="w100"
                  clearable
                  filterable
                  placeholder="请选择归属组织"
                  :show-all-levels="false"
              ></el-cascader>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <el-form-item label="岗位" prop="postIds">
              <el-select v-model="state.ruleForm.postIds" multiple placeholder="请选择">
                <el-option
                  v-for="item in state.postOptions"
                  :key="item.postId"
                  :label="item.postName"
                  :value="item.postId"
                  :disabled="item.status == 1"
                ></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <el-form-item label="状态">
              <el-radio-group v-model="state.ruleForm.status">
                <el-radio
                        v-for="dict in state.statusOptions"
                        :key="dict.dictValue"
                        :label="dict.dictValue"
                >{{ dict.dictLabel }}
                </el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" >
            <el-form-item label="备注">
              <el-input
                v-model="state.ruleForm.remark"
                type="textarea"
                placeholder="请输入内容"
              ></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form> -->

      <!-- <template #footer>
        <span class="dialog-footer">
          <el-button @click="onCancel">取 消</el-button>
          <el-button type="primary" @click="onSubmit" :loading="state.loading">保 存</el-button>
        </span>
      </template> -->
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, unref, getCurrentInstance } from "vue";
import { treeselect } from "@/api/system/organization";
import { updateUser, addUser, getUser, getUserInit, updateEmployee, deptTreeSelect, compTreeSelect } from "@/api/system/user";
import { ElMessage } from "element-plus";
import { spawn } from "child_process";

const props = defineProps({
  title: {
    type: String,
    default: () => "",
  },
})

const { proxy } = getCurrentInstance() as any;
const ruleFormRef = ref<HTMLElement | null>(null);
const deptOptions = ref(undefined);
const compOptions = ref(undefined);
const state = reactive({
  // 是否显示弹出层
  isShowDialog: false,
  loading: false,
  // 默认密码
  // 性别状态字典
  sexOptions: [],
  // 角色选项
  roleOptions: [],
  // 状态数据字典
  statusOptions: [],
  // 组织树选项
  organizationOptions: [],
  // 岗位选项
  postOptions: [],
  ruleForm: {
    user: {
      id: "", //用户编号
      login: "", // 用戶登录名
      active: "", //用户状态
      company_id: "", // 所属公司
    },
    employee: {
      id: "", // 员工编号
      name: "", // 员工名称
      department_id: "", // 部门ID
      user_id: "", // 相关用户
      job_id: "", // 岗位ID
      employee_type: "",// 员工类型
      work_phone: "", // 手机号
      work_email: "", // 邮箱
      birthday: "", //用户状态
      // avatar: "", // 用户头像
      gender: "", // 性别
      marital: "", // 婚姻状态
      fired: "", // 解聘
      resigned: "", // 辞职
      // postIds: [], //odoo中好像只能绑定一个工作岗位
    },
    // roleIds: [],
  },

  // 显示状态数据字典
  isHideOptions: [],
  // 菜单类型数据字典
  menuTypeOptions: [],
  // 数字是否数据字典
  yesOrNoOptions: [],
  // 菜单树选项
  menuOptions: [],
  // 表单校验
  userRules: {
    login: [
      { required: true, message: "登录名不能为空", trigger: "blur" },
    ],
  },
  ruleRules: {
    name: [
      { required: true, message: "员工名称不能为空", trigger: "blur" },
    ],
    // login: [
    //   { required: true, message: "登录名不能为空", trigger: "blur" },
    // ],
    // department_id: [
    //   { required: true, message: "所属部门不能为空", trigger: "blur" },
    // ],
    // roleIds: [
    //   { required: true, message: "所属角色不能为空", trigger: "blur" },
    // ],
    work_email: [
      {
        type: "email",
        message: "'请输入正确的邮箱地址",
        trigger: ["blur", "change"],
      },
    ],
    work_phone: [
      {
        pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
        message: "请输入正确的手机号码",
        trigger: "blur",
      },
    ],
  },
});
// 打开弹窗
const openDialog = (row: any) => {
  if (row && row.userId && row.userId != undefined && row.userId != 0) {
    getUser(row.userId).then((response: any) => {
      state.ruleForm.employee = response.data.data.employee;
      state.ruleForm.user = response.data.data.user;
      state.postOptions = response.data.posts;
      state.roleOptions = response.data.roles;
      //state.organizationOptions = response.data.organizations;
      // state.ruleForm.postIds = response.data.postIds.split(",").map((item: string)=>{
      //   return Number(item)
      // });
      // state.ruleForm.roleIds = response.data.roleIds.split(",").map((item: string)=>{
      //   return Number(item)
      // });
      // state.ruleForm.password = ""
    });
  } else {
    getUserInit().then(response => {
      state.postOptions = response.data.posts
      state.roleOptions = response.data.roles
    })
    state.ruleForm.employee = JSON.parse(JSON.stringify(row));
  }
  getTreeselect();
  getCompTree();
  getDeptTree();
  state.isShowDialog = true;
  state.loading = false;
  // 查询显示性別数据字典
  proxy.getDicts("sys_user_sex").then((response: any) => {
    state.sexOptions = response.data;
  });
  // 查询显示狀態数据字典
  proxy.getDicts("sys_normal_disable").then((response: any) => {
    state.statusOptions = response.data;
  });
};
// 关闭弹窗
const closeDialog = (row?: object) => {
  proxy.mittBus.emit("onEditUserModule", row);
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  state.isShowDialog = false;
};
/** 查询组织下拉树结构 */
const getTreeselect = async () => {
  treeselect().then((response) => {
    state.organizationOptions = response.data;
  });
};
/** 查询部门下拉树结构 */
const getDeptTree=() =>{
  deptTreeSelect().then((response) => {
    deptOptions.value = response.data;
  });
}
/** 查询公司下拉树结构 */
const getCompTree=() =>{
  compTreeSelect().then((response) => {
    compOptions.value = response.data;
  });
}

/** 提交按钮 */
const employeeSubmit = () => {
  // 更新员工
  updateEmployee(state.ruleForm.employee).then((res: any) => {
            if (res.code == 200) {
              ElMessage.success("修改成功");
              // closeDialog(); // 关闭弹窗
            }
            // state.loading = false;
          })

}

const userSubmit = () => {
  // 更新用户
  updateUser(state.ruleForm.user).then((res: any) => {
          if (res.code == 200) {
            ElMessage.success("修改成功");
            // closeDialog(); // 关闭弹窗
          }
          // state.loading = false;
        })
}
const onSubmit = () => {
  const formWrap = unref(ruleFormRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      // state.ruleForm.employee.job_id = state.ruleForm.employee.postIds[0]
      // state.ruleForm.roleId = state.ruleForm.roleIds[0]
      // state.ruleForm.postIds = state.ruleForm.postIds.join(',')
      // state.ruleForm.roleIds = state.ruleForm.roleIds.join(',')
      state.loading = true;
      if (state.ruleForm.employee.user_id != undefined) {
        // 更新用户
        updateUser(state.ruleForm.user).then((res: any) => {
          if (res.code == 200) {
            ElMessage.success("修改成功");
            // closeDialog(); // 关闭弹窗
          }
          // state.loading = false;
        }),
          // 更新员工
          updateEmployee(state.ruleForm.employee).then((res: any) => {
            if (res.code == 200) {
              ElMessage.success("修改成功");
              closeDialog(); // 关闭弹窗
            }
            state.loading = false;
          }
          )
      } else {
        addUser(state.ruleForm.employee).then((res: any) => {
          if (res.code == 200) {
            ElMessage.success("新增成功");
            closeDialog(); // 关闭弹窗
          }
          state.loading = false;
        });
      }
    }
  });
};

// 头像上传
const handleAvatarSuccess = (file: any) => {
  //   state.imageUrl = URL.createObjectURL(file.raw);
};
// 头像上传前校验
const beforeAvatarUpload = (file: any) => {
  const isJPG = file.type === "image/jpeg";
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isJPG) {
    ElMessage.error("上传头像图片只能是 JPG 格式!");
  }
  if (!isLt2M) {
    ElMessage.error("上传头像图片大小不能超过 2MB!");
  }
  return isJPG && isLt2M;
};

defineExpose({
  openDialog,
});

</script>
<style scoped lang="scss">
.updateUser {
  height: 100%;
  overflow: auto;
  padding-bottom: 53px;
  width: 600px;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 108px;
  height: 108px;
  margin: 8px;
  line-height: 108px;
  border-radius: 4px;
  text-align: center;
  background-color: #fafafa;
  border: 1px dashed #d9d9d9;
}

.avatar {
  width: 108px;
  height: 108px;
  margin: 8px;
  border-radius: 4px;
  display: block;
}
.dialog-footer {
  margin-left: 60px;
}
</style>
