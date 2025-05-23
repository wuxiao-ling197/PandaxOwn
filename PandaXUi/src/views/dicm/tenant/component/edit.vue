<template>
  <div class="system-menu-container">
    <el-dialog v-model="state.isShowDialog" :title="state.title">
      <template #header>
        <div style="font-size: large"
          v-drag="['.system-menu-container .el-dialog', '.system-menu-container .el-dialog__header']">{{ title }}</div>
      </template>
      <!-- <el-card shadow="always" style="height: 100%;"> -->
      <el-form :model="state.data" ref="ruleFormRef" label-width="auto" >
         <el-form-item v-if="state.tID"  label="编码">
          <el-input v-if="state.tID" v-model="state.tID" disabled/>
        </el-form-item>
        <el-form-item label="名称" required>
          <el-input v-model="state.data.name"  placeholder="请输入名称" required/>
        </el-form-item>
        <el-form-item label="租户组" >
          <el-tree-select v-model="state.data.group_id" :data="state.groupOptions"
                        :props="{ value: 'id', label: 'name', children: 'children' }" check-strictly
                        placeholder="请选择租户组" />
        </el-form-item>
        <el-form-item label="标识符" required>
          <el-input v-model="state.data.slug" type="textarea" placeholder="请输入唯一标识符" required/>
        </el-form-item>
        <el-form-item label="描述"  required>
          <el-input v-model="state.data.description"   type="textarea"  required/> </el-form-item>
        <el-form-item  label="自定义配置数据" required>
          <el-input v-model="state.data.custom_field_data"  :rows="5"  
            type="textarea" placeholder='例如：
{
    "ids":[5,6],
    "group_id":8
}' />
        </el-form-item>
        <el-form-item label="评价"  >
          <el-input v-model="state.data.comments" type="textarea"  required/>
        </el-form-item>
      </el-form>
      <!-- </el-card> -->

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
import { reactive, ref, unref, getCurrentInstance, onMounted } from "vue";
import { ElMessage, FormItemProps, FormProps } from "element-plus";
import { addTenant, tenantGroupTree, updateTenant } from "@/api/dicm/tenant";

const props = defineProps({
  title: {
    type: String,
    default: () => "",
  },
})

const { proxy } = getCurrentInstance() as any;
const ruleFormRef = ref<HTMLElement | null>(null);
const state = reactive({
  isAdd: true,
  // 是否显示弹出层
  isShowDialog: false,
  loading: false,
  // 组织树选项
  organizationOptions: [],
  groupOptions: [],
  data: {
    slug: undefined, 
    name: undefined, 
    group_id:undefined, 
    comments: undefined, 
    description: undefined,
    custom_field_data: undefined,
  },
  tID:undefined,
});

onMounted(()=>{
  // 加载租户组层级
    tenantGroupTree().then((res: any) => {
        // console.log("租户层级：", res.data);
        state.groupOptions = res.data
    })
})

// 打开弹窗
const openDialog = (row: any) => {
  // console.log("接收参数=", row.name);
  if (row.name==undefined) {
    state.isAdd=true
    state.tID=undefined
  }else{
    state.tID=row.id
    state.isAdd=false
  }
  state.data = JSON.parse(JSON.stringify(row));
  state.isShowDialog = true;
  state.loading = false;
}
// 关闭弹窗
const closeDialog = (row?: object) => {
  proxy.mittBus.emit("onEditUserModule", row);
  state.tID=undefined;
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  state.tID=undefined;
  state.isShowDialog = false;
};

// 提交按钮
const onSubmit = () => {
  const formWrap = unref(ruleFormRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      state.loading = true;
      // true为新增
      if (state.isAdd) {
        // 添加用户
        addTenant(state.data).then((res: any) => {
          if (res.code == 200) {
            ElMessage.success("新增成功");
            closeDialog(); // 关闭弹窗 sb:pq: 类型json的输入语法无效
          }
          state.loading = false;
        });
        
      } else {
        updateTenant(state.data).then((res: any) => {
          if (res.code == 200) {
            ElMessage.success("修改成功");
            closeDialog(); // 关闭弹窗 cg
          }
          state.tID=undefined
          state.loading = false;
        })
      }
    }
  });
};
defineExpose({
  openDialog,
});
</script>

<style scoped lang="scss">
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