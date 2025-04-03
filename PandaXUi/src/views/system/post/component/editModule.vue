<template>
  <div class="system-menu-container">
    <el-dialog v-model="state.isShowDialog" width="769px" center>
      <template #header>
        <div style="font-size: large" v-drag="['.system-menu-container .el-dialog', '.system-menu-container .el-dialog__header']">{{title}}</div>
      </template>
      <el-form
        :model="state.ruleForm"
        :rules="state.ruleRules"
        ref="ruleFormRef"
        label-width="80px"
      >
        <el-row :gutter="35">
          <el-col :span="24" >
            <el-form-item label="岗位名称" prop="name">
              <el-input
                v-model="state.ruleForm.name"
                placeholder="请输入岗位名称"
              />
            </el-form-item>
          </el-col>
          <el-col :span="24" >
            <el-form-item label="招聘类型" prop="contract_type_id">
              <el-tree-select
                v-model="state.ruleForm.contract_type_id"
                placeholder="请选择招聘类型"
              />
            </el-form-item>
          </el-col>
          <el-col :span="24" >
            <el-form-item label="招聘部门" prop="department_id">
              <el-tree-select
                v-model="state.ruleForm.department_id"
                placeholder="请选择部门"
              />
            </el-form-item>
          </el-col>
          <el-col :span="24" >
            <el-form-item label="岗位顺序" prop="sequence">
              <el-input-number
                v-model="state.ruleForm.sequence"
                controls-position="right"
                :min="0"
              />
            </el-form-item>
          </el-col>
          <el-col :span="24" >
            <el-form-item label="发布状态" prop="is_published">
              <el-radio-group v-model="state.ruleForm.is_published">
                <el-radio value="true" size="large">发布</el-radio>
                <el-radio value="false" size="large">稍后</el-radio> 
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="24" >
            <el-form-item label="招聘人数" prop="no_of_recruitment">
              <el-input-number
                v-model="state.ruleForm.no_of_recruitment"
                controls-position="right"
                :min="0"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="onCancel">取 消</el-button>
          <el-button type="primary" @click="onSubmit" :loading="state.loading">保 存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, unref, getCurrentInstance } from "vue";
import { updatePost, addPost } from "@/api/system/post";
import { ElMessage } from "element-plus";

const props = defineProps({
  title: {
    type: String,
    default: () => "",
  },
})

const { proxy } = getCurrentInstance() as any;
const ruleFormRef = ref<HTMLElement | null>(null);
const state = reactive({
  // 是否显示弹出层
  isShowDialog: false,
  loading: false,
  // 岗位对象
  ruleForm: {
    // postId: 0, // 岗位ID
    name:"", // 岗位名称
    department_id: 0,// 岗位编码
    sequence: 0, // 岗位排序
    no_of_recruitment: 0, //计划招聘人数
    contract_type_id:0,//招聘类型
    is_published: false, // 是否发布
  },
  // 岗位状态数据字典
  statusOptions: [],
  // 岗位树选项
  organizationOptions: [],
  // 表单校验
  ruleRules: {
    name: [{ required: true, message: "岗位名称不能为空", trigger: "blur" }],
    no_of_recruitment: [{ required: true, message: "招聘人数不能为空", trigger: "blur" }],
    // contract_type_id: [ { required: true, message: "招聘类型不能为空", trigger: "blur" }]
  },
});
// 打开弹窗
const openDialog = (row: any) => {
  state.ruleForm = JSON.parse(JSON.stringify(row));

  state.isShowDialog = true;
  state.loading = false;
  // 查询岗位状态数据字典
  proxy.getDicts("sys_normal_disable").then((response: any) => {
    state.statusOptions = response.data;
  });
};

// 关闭弹窗
const closeDialog = (row?: object) => {
  proxy.mittBus.emit("onEditPostModule", row);
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  state.isShowDialog = false;
};

// 保存
const onSubmit = () => {
  const formWrap = unref(ruleFormRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      state.loading = true;
      state.ruleForm.contract_type_id=4
      if (state.ruleForm.contract_type_id != undefined && state.ruleForm.contract_type_id != 0) {
        updatePost(state.ruleForm).then((res:any) => {
          if (res.code == 200) {
            ElMessage.success("修改成功");
            closeDialog(state.ruleForm); // 关闭弹窗
          }
          state.loading = false;
        });
      } else {
        addPost(state.ruleForm).then((res:any) => {
          if (res.code == 200) {
            ElMessage.success("新增成功");
            closeDialog(state.ruleForm); // 关闭弹窗
          }
          state.loading = false;
        });
      }
    }
  });
};

defineExpose({
  openDialog,
});
</script>
