<template>
  <div class="app-container">
    <el-card shadow="always">
      <!-- 查询 -->
      <el-form :model="state.queryParams" ref="queryForm" :inline="true" label-width="68px">
        <el-form-item label="组织名称" prop="departmentName">
          <el-input placeholder="请输入组织名称模糊查询" clearable @keyup.enter.native="handleQuery"
            v-model="state.queryParams.departmentName" />
        </el-form-item>
        <el-form-item label="部门代码" prop="departmentId">
          <el-input placeholder="请输入组织名称模糊查询" clearable @keyup.enter.native="handleQuery"
            v-model="state.queryParams.departmentId" />
        </el-form-item>
        <el-form-item label="公司代码" prop="companyId">
          <el-input placeholder="请输入组织名称模糊查询" clearable @keyup.enter.native="handleQuery"
            v-model="state.queryParams.companyId" />
        </el-form-item>
        <el-form-item label="状态" prop="active">
          <el-select v-model="state.queryParams.active" placeholder="状态" clearable style="width: 240px">
            <el-option :key="true" label="正常" :value="true"></el-option>
            <el-option :key="false" label="停用" :value="false"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" plain @click="handleQuery">
            <SvgIcon name="elementSearch" />
            搜索
          </el-button>
          <el-button @click="resetQuery">
            <SvgIcon name="elementRefresh" />
            重置
          </el-button>
        </el-form-item>

      </el-form>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span class="card-header-text">组织列表</span>
          <div>
            <el-button type="primary" plain v-auth="'system:organization:list'" @click="onOpenCompanyModule">
              <el-icon>
                <House />
              </el-icon>
              公司信息
            </el-button>
            <el-button type="primary" plain v-auth="'system:organization:add'" @click="onOpenAddModule">
              <SvgIcon name="elementPlus" />
              新增
            </el-button>
          </div>
        </div>
      </template>
      <!--数据表格-->
      <el-table v-loading="state.loading" :data="state.tableData.data" row-key="id" border default-expand-all
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
        <el-table-column prop="id" label="部门编码" width="100" fixed>
          <!-- <template #default="{row}">
                {{ row.departments?.id || row.departments?.id || row.id }}
              </template> -->
        </el-table-column>
        <el-table-column prop="name" label="组织名称" width="160">
          <template #default="{ row }">
            {{ JSON.parse(row.name).zh_CN || JSON.parse(row.name).en_US }}
          </template>
        </el-table-column>
        <!-- <el-table-column label="部门" prop="departments">
            <el-table-column prop="departments.id" label="部门编码" width="120" >
              <template #default="{row}">
                {{ row.departments?.id || row.departments?.id || row.id }}
              </template>
            </el-table-column>
            <el-table-column prop="departments.manager_id" label="管理员" width="120">
              <template #default="{row}">
                {{ row.manager_id || row.manager_id || '无' }}
              </template>
            </el-table-column>
            <el-table-column prop="departments.name" label="部门名称" >
              <template #default="{row}">
                {{ JSON.parse(row.name).zh_CN || JSON.parse(row.name).en_US }}
              </template>
            </el-table-column>
            <el-table-column prop="departments.active" label="部门状态" width="120" >
              <template #default="scope">
                <el-tag :type="scope.row.active === true ? 'success' : 'danger'" disable-transitions>{{ scope.row.active ?
                  "启用" : "已停用" }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table-column> -->
        <!-- </el-table-column> -->
        <el-table-column prop="manager_id" label="管理员" width="120">
          <template #default="{ row }">
            <el-link v-if="row.manager_id" type="primary" :underline="false" @click="handleClick(row.manager_id)">
              {{ row.manager_id }}
            </el-link>
            <span v-else>-无-</span>
          </template>
        </el-table-column>
        <!-- <el-table-column prop="departments.name" label="部门名称" >
              <template #default="{row}">
                {{ JSON.parse(row.name).zh_CN || JSON.parse(row.name).en_US }}
              </template>
            </el-table-column> -->
        <!-- <el-table-column prop="email" label="组织邮箱" width="200"></el-table-column>
        <el-table-column prop="phone" label="联系方式" width="200"></el-table-column> -->
        <el-table-column prop="parent_id" label="上级部门" width="120" sortable>
          <template #defult="scope">
            {{ scope.row.parent_id || '无' }}
          </template>
        </el-table-column>
        <el-table-column prop="master_department_id" label="顶级部门" width="120" sortable>
          <template #defult="scope">
            {{ scope.row.master_department_id || '无' }}
          </template>
        </el-table-column>
        <el-table-column prop="active" label="状态" width="60">
          <template #default="scope">
            <el-tag :type="scope.row.active === true ? 'success' : 'danger'" disable-transitions>{{ scope.row.active ?
              "启用" : "已停用" }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="虚拟" align="center" prop="departments.manager_id" width="100">
        </el-table-column>
        <el-table-column label="员工" align="center" prop="employees" width="120">
          <template #default="{ row }">
            <el-button size="small" type="info" @click="handleClick(row.employees.department_id)">
              查看员工 {{ Array.isArray(row.employees) ? row.employees.length : 0  }}
            </el-button>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
          <template #default="scope">
            <el-popover placement="left">
              <template #reference>
                <el-button type="primary" circle>
                  <SvgIcon name="elementStar" />
                </el-button>
              </template>
              <div>
                <el-button text type="primary" v-auth="'system:organization:edit'" @click="onOpenEditModule(scope.row)">
                  <SvgIcon name="elementEdit" />
                  修改
                </el-button>
              </div>
              <div>
                <el-button text type="primary" v-auth="'system:organization:add'" @click="onOpenAddModule(scope.row)">
                  <SvgIcon name="elementPlus" />
                  新增
                </el-button>
              </div>
              <div>
                <el-button v-if="scope.row.parentId != 0" text type="primary" v-auth="'system:organization:delete'"
                  @click="onTabelRowDel(scope.row)">
                  <SvgIcon name="elementDelete" />
                  删除
                </el-button>
              </div>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <!-- 添加或修改组织对话框 -->
    <EditModule ref="editModuleRef" :title="state.title" />
    <CompanyInfo ref="companyFormRef" :title="state.title" />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, getCurrentInstance, onUnmounted, } from "vue";
import { ElMessageBox, ElMessage } from "element-plus";
import { listOrganization, delOrganization, departmentTree } from "@/api/system/organization";
import EditModule from "./component/editModule.vue";
import CompanyInfo from './component/companyInfo.vue';
import { parseTime } from "@/utils";
import router from "@/router";

const { proxy } = getCurrentInstance() as any;
const editModuleRef = ref();
const companyFormRef = ref();
const state = reactive({
  // 遮罩层
  loading: true,
  // 弹出层标题
  title: "",
  // 组织表格树数据
  tableData: {
    data: [],
    total: 0,
  },
  child: undefined,
  // 状态数据字典
  statusOptions: [],
  // 员工数
  empTotal:undefined,
  // 查询参数
  queryParams: {
    pageNum: undefined,
    pageSize: undefined,
    departmentName: undefined,
    active: undefined,
    departmentId: undefined,
    companyId: undefined
  },
});

/** 查询组织列表 */
const handleQuery = () => {
  state.loading = true;
  listOrganization(state.queryParams).then((response: any) => {
    if (response.code != 200) {
      state.loading = false;
    }
    state.tableData.data = response.data.data;
    state.tableData.total = response.data.total;
    console.log("组织table：", state.tableData.data);
    state.loading = false;
  });
};

/** 重置按钮操作 */
const resetQuery = () => {
  state.queryParams.departmentName = undefined;
  state.queryParams.active = undefined;
  handleQuery();
};

// 单元格跳转
const handleClick = (id: any) => {
  console.log('组织点击跳转路由携带参数：', id);
  // /system/user/list
  // router.push({ path: '/system/user/list', query: { id: id } })
  router.push({name:'/system/user', query: { id: id }}) //http://192.168.0.5:7789/#/system/organization?department_id=2
}

const handleTable = (data: any) => {
  if (data.children.Array.length != 0) {
    state.child = data.children
  }
  if (data.departments.Array.length != 0) {
    state.child = data.departments
  }
}
//打开公司信息数据框
const onOpenCompanyModule = () => {
  state.title = "公司信息";
  companyFormRef.value.openDialog({});
};

// 打开新增组织弹窗
const onOpenAddModule = (row: any) => {
  let parentId = row.organizationId;
  row = {};
  row.parentId = parentId;
  state.title = "添加组织";
  editModuleRef.value.openDialog(row);
};
// 打开编辑组织弹窗
const onOpenEditModule = (row: object) => {
  state.title = "修改组织";
  editModuleRef.value.openDialog(row);
};
/** 删除按钮操作 */
const onTabelRowDel = (row: any) => {
  ElMessageBox({
    message: '是否确认删除名称为"' + row.organizationName + '"的数据项?',
    title: "警告",
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(function () {
    return delOrganization(row.organizationId).then((res: any) => {
      if (res.code === 200) {
        handleQuery();
        ElMessage.success("删除成功");
      } else {
        ElMessage.error("删除失败");
      }
    });
  });
};
// 页面加载时
onMounted(() => {
  // 查询组织信息
  handleQuery();
  // handleTable(state.tableData);
  // 查询组织状态数据字典
  proxy.getDicts("sys_normal_disable").then((response: any) => {
    state.statusOptions = response.data;
  });
  proxy.mittBus.on("onEditOrganizationModule", (res: any) => {
    handleQuery();
  });
});
// 页面卸载时
onUnmounted(() => {
  proxy.mittBus.off("onEditOrganizationModule");
});
</script>
