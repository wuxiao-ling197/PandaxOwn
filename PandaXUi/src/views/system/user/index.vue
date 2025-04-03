<template>
  <div class="system-user-container app-container">
    <el-row :gutter="20">
      <!--组织树——组织数据-->
      <el-col :span="5" :xs="24">
        <el-card shadow="always">
          <!-- 顶部搜索 -->
          <div class="head-container">
            <el-input v-model="state.departmentName" placeholder="请输入部门名称" clearable prefix-icon="el-icon-search"
              style="margin-bottom: 20px" />
          </div>
          <div class="head-container">
            <span style="color:#7a7a7f;font-weight: 800;font-stretch: condensed;margin-bottom: 15px;">公司</span>
            <!-- 渲染公司结构 -->
            <el-tree :data="state.companyOptions" 
            :props="state.companyProps" 
            node-key="id"
              :expand-on-click-node="false" 
              :filter-node-method="filterNode" 
              ref="tree" 
              default-expand-all
              :highlight-current="true"
              @node-click="handleCompanyClick" class="main-tree">
              <!-- :default-expanded-keys="state.defaultExpandedKeys"
              :default-checked-keys="state.defaultCheckedKeys" -->
              <!-- <template #default="{ node, data }">
                <div class="custom-tree-node">
                  公司/部门图标
                  <i :class="data.id ? 'el-icon-s-home' : 'el-icon-office-building'"></i>

                  节点名称
                  <span class="node-name">{{ getNodeName(data) }}</span>

                  节点类型标签
                  <span class="node-type">
                    {{ data.id ? '部门' : '公司' }}
                  </span>

                  如果有部门数据且当前是公司节点，渲染部门子节点
                  <div v-if="!data.id && data.Departments && data.Departments.length > 0" class="dept-container">
                    <el-tree :data="data.Departments" :props="state.deptProps" node-key="id" :show-checkbox="false"
                      default-expand-all class="dept-subtree">
                      <template #default="{ node: deptNode, data: deptData }">
                        <div class="custom-dept-node">
                          <i class="el-icon-s-home"></i>
                          <span class="node-name">{{ getDeptName(deptData) }}</span>
                          <span class="node-type">部门</span>
                        </div>
                      </template>
                    </el-tree>
                  </div>
                </div>
              </template> -->
            </el-tree>
          </div>
        </el-card>
        <el-card shadow="always">
          <span style="color:#7a7a7f;font-weight: 800;font-stretch: condensed;margin-bottom: 15px;">部门</span>
          <div class="head-container">
            <!-- 渲染部门 :empty-text="departmentText" -->
            <el-tree :data="state.departmentOptions" :props="state.departmentProps" node-key="id"
              :expand-on-click-node="false" :filter-node-method="filterNode" ref="tree" 
              default-expand-all class="main-tree">
              <!-- <template #default="{ node, data }">
                <div class="custom-tree-node">
                  公司/部门图标
                  <i :class="data.id ? 'el-icon-s-home' : 'el-icon-office-building'"></i>

                  节点名称
                  <span class="node-name">{{ getNodeName(data) }}</span>

                  节点类型标签
                  <span class="node-type">
                    {{ data.id ? '部门' : '公司' }}
                  </span>

                  如果有部门数据且当前是公司节点，渲染部门子节点
                  <div v-if="!data.id && data.Departments && data.Departments.length > 0" class="dept-container">
                    <el-tree :data="data.Departments" :props="state.deptProps" node-key="id" :show-checkbox="false"
                      default-expand-all class="dept-subtree">
                      <template #default="{ node: deptNode, data: deptData }">
                        <div class="custom-dept-node">
                          <i class="el-icon-s-home"></i>
                          <span class="node-name">{{ getDeptName(deptData) }}</span>
                          <span class="node-type">部门</span>
                        </div>
                      </template>
                    </el-tree>
                  </div>
                </div>
              </template> -->
            </el-tree>
          </div>
        </el-card>
      </el-col>

      <el-col :span="19" :xs="24">
        <el-card shadow="always">
          <!-- 搜索框-->
          <el-form :model="state.queryParams" ref="queryForm" :inline="true" label-width="78px">
            <el-form-item label="用户名称" prop="login">
              <el-input placeholder="用户名称模糊查询" clearable @keyup.enter="handleQuery" style="width: 240px"
                v-model="state.queryParams.login" />
            </el-form-item>
            <el-form-item label="员工名称" prop="name">
              <el-input placeholder="员工名称模糊查询" clearable @keyup.enter="handleQuery" style="width: 240px"
                v-model="state.queryParams.name" />
            </el-form-item>
            <el-form-item label="手机号码" prop="work_phone">
              <el-input v-model="state.queryParams.work_phone" placeholder="请输入手机号码" clearable style="width: 240px"
                @keyup.enter="handleQuery" />
            </el-form-item>
            <el-form-item label="所属部门" prop="department_id">
              <el-input v-model="state.queryParams.department_id" placeholder="请选择所属部门" clearable style="width: 240px"
                @keyup.enter="handleQuery" />
            </el-form-item>
            <el-form-item label="状态" prop="active">
              <el-select v-model="state.queryParams.active" placeholder="用户状态" clearable style="width: 240px">
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
        <!-- 操作按钮 -->
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span class="card-header-text">用户列表</span>
              <div>
                <el-button type="primary" plain @click="handleAdd">
                  <SvgIcon name="elementPlus" />
                  新增
                </el-button>
                <el-button type="danger" plain :disabled="state.multiple" @click="handleDelete">
                  <SvgIcon name="elementDelete" />
                  删除
                </el-button>
                <el-button type="warning" plain @click="handleExport">
                  <SvgIcon name="elementDownload" />
                  导出
                </el-button>
              </div>
            </div>
          </template>
          <!-- 表格数据 -->
          <el-table v-loading="state.loading" :data="state.tableData.data" stripe
            @selection-change="handleSelectionChange" style="width: 100%">
            <!-- 多选按钮 -->
            <el-table-column type="selection" width="45" align="center" />
            <el-table-column label="用户编号" align="center" key="id" prop="id" />
            <el-table-column label="用户名" prop="login" show-overflow-tooltip></el-table-column>
            <el-table-column label="绑定员工" prop="employee.name" show-overflow-tooltip></el-table-column>
            <el-table-column label="部门" prop="employee.department_id" show-overflow-tooltip>
            </el-table-column>
            <el-table-column label="员工类型" key="employee.employee_type" prop="employee.employee_type"
              show-overflow-tooltip>
            </el-table-column>
            <el-table-column label="头像" show-overflow-tooltip>
              <template #default="scope">
                <el-image class="system-user-photo" :src="scope.row.avatar
                  ? scope.row.avatar
                  : letterAvatar(scope.row.username)
                  "></el-image>
              </template>
            </el-table-column>
            <el-table-column label="用户性别" align="center" prop="employee.gender">
              <template #default="scope">
                <span>{{ scope.row.employee?.gender === "male" ? "男性" : scope.row.employee?.gender === "female" ? "女性" : "未知" }}</span>
              </template></el-table-column>
            <el-table-column label="婚姻状态" align="center" prop="employee.marital" />
            <el-table-column label="工作岗位" align="center" prop="employee.job_title"></el-table-column>
            <el-table-column prop="employee.work_phone" label="手机" show-overflow-tooltip></el-table-column>
            <el-table-column prop="employee.work_email" label="邮箱" show-overflow-tooltip></el-table-column>
            <el-table-column label="状态" align="center" key="active">
              <template #default="scope">
                <span>{{scope.row.active === "false" ? "停用" : "正常"}}</span>
              </template>
            </el-table-column>
            <el-table-column prop="create_date" label="创建时间" show-overflow-tooltip>
              <template #default="scope">
                {{ dayjs(scope.row.create_date).format('YYYY-MM-DD HH:mm:ss') }}
              </template>
            </el-table-column>
            <el-table-column prop="path" align="center" label="操作">
              <template #default="scope">
                <el-popover placement="left">
                  <template #reference>
                    <el-button type="primary" circle>
                      <SvgIcon name="elementStar" />
                    </el-button>
                  </template>
                  <div>
                    <!-- v-auth="'system:user:edit'"  -->
                    <el-button text type="primary" @click="handleUpdate(scope.row)">
                      <SvgIcon name="elementEdit" />
                      修改
                    </el-button>
                  </div>
                  <div>
                    <el-button text type="primary" icon="cpu" @click="handleResetPwd(scope.row)">重置密码</el-button>
                  </div>
                  <div>
                    <!-- v-auth="'system:user:edit'" -->
                    <el-button text type="primary" icon="Key" @click="handleEnableTotp(scope.row)">
                      启用TOTP
                    </el-button>
                  </div>
                  <div>
                    <el-button text type="primary" v-auth="'system:user:delete'" @click="handleDelete(scope.row)">
                      <SvgIcon name="elementDelete" />
                      删除
                    </el-button>
                  </div>
                </el-popover>
              </template>
            </el-table-column>
          </el-table>
          <div v-show="state.tableData.total > 0">
            <el-divider></el-divider>
            <el-pagination background :total="state.tableData.total" :page-sizes="[10, 20, 30, 50, 100]"
              :current-page="state.queryParams.pageNum" :page-size="state.queryParams.pageSize"
              layout="total, sizes, prev, pager, next, jumper" @size-change="onHandleSizeChange"
              @current-change="onHandleCurrentChange" />
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog v-model="state.isTOTP" title="请确认密码" center>
      <el-form-item>
        <el-input placeholder="请输入用户密码" clearable v-model="state.isPassword"></el-input>
      </el-form-item>
      <el-button type="primary" plain @click="handleOK">确认</el-button>
    </el-dialog>
    <!-- 添加或修改参数配置对话框 -->
    <EditModule ref="userFormRef" :title="state.title" />

  </div>
</template>

<script lang="ts" setup>
import {
  toRefs,
  reactive,
  onMounted,
  ref,
  watch,
  getCurrentInstance,
  onUnmounted,
  computed,
} from "vue";
import {
  listUser,
  changeUserStatus,
  delUser,
  exportUser,
  resetUserPwd,
  enableTotp,
} from "@/api/system/user";
import { departmentTree, treeselect } from "@/api/system/organization";
import { ElMessageBox, ElMessage, dayjs } from "element-plus";
import { getDicts } from "@/api/system/dict/data";
import MDInput from "@/components/panda/MDInput.vue";
import EditModule from "./component/editModule.vue";
import { letterAvatar } from '@/utils/string';
import { handleFileError } from "@/utils/export";

const { proxy } = getCurrentInstance() as any;
const userFormRef = ref();
const state: any = reactive({
  tableData: {
    data: [],
    total: 0,
  },
  isTOTP: false, //启用totp时需要输入密码验证
  isPassword: undefined,
  value: undefined,
  loading: false,
  // 岗位选项
  postOptions: [],
  // 组织树选项
  // defaultCheckedKeys:[1],
  // defaultExpandedKeys:[1],
  FindList: [], //api返回数据
  companyProps: {
    children: "children",
    label: "name",
  }, //公司数据
  departmentProps: {
    label: (data: any) => getDeptName(data),
    children: 'children' 
  }, // 部门数据
  companyOptions: undefined, //接收公司数据
  departmentOptions: undefined,
  departmentText: '暂无部门数据', //没有部门数据提示
  currentCompany: null, //当前选中公司节点
  // 性别状态字典
  sexOptions: [],
  // 角色选项
  roleOptions: [],
  // 状态数据字典
  // statusOptions: [],
  // 组织名称
  departmentName: undefined,
  // 非单个禁用
  single: true,
  // 非多个禁用
  multiple: true,
  // 选中数组
  ids: [],
  // 弹出层标题
  title: "",
  // 查询参数
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    login: undefined,
    name: undefined,
    work_phone: undefined,
    active: undefined,
    department_id: undefined,
  },
});

watch(
  () => state.departmentName,
  (newValue) => {
    proxy.$refs.tree.filter(newValue);
  }
);
/** 查询用户列表 cg*/
const getList = async () => {
  state.loading = true;
  listUser(state.queryParams).then(
    (response: any) => {
      if (response.code != 200) {
        state.loading = false;
      }
      console.log('user-list=', response.data);
      state.tableData.data = response.data.data;
      state.tableData.total = response.data.total;
      state.loading = false;
    }
  );
};
// 多选框选中数据
const handleSelectionChange = (selection: any) => {
  console.log("用户列表多选框传入数据=", selection);
  state.ids = selection.map((item: any) => item.userId);
  state.single = selection.length != 1;
  state.multiple = !selection.length;
};
/** 用户搜索按钮操作 cg*/
const handleQuery = async () => {
  // console.log("查询用户列表", state.queryParams.userName);
  state.queryParams.pageNum = 1;
  await getList();
};
/** 重置按钮操作 cg*/
const resetQuery = async () => {
  // 表单初始化，方法：`resetFields()` 无法使用
  state.queryParams.pageNum = 1;
  state.queryParams.pageSize = 10;
  state.queryParams.login = "";
  state.queryParams.work_phone = "";
  state.queryParams.active = "";
  state.queryParams.departmentId = 0;
  handleQuery();
};
/** 新增按钮操作 */
const handleAdd = () => {
  state.title = "添加用戶";
  userFormRef.value.openDialog({});
};
/** 修改按钮操作 */
const handleUpdate = (row: any) => {
  state.title = "修改用户";
  userFormRef.value.openDialog(row);
};

/** 查询组织下拉树结构 */
const getTreeselect = async () => {
  treeselect().then((response) => {
    // 公司数据
    state.companyOptions = response.data;
    departmentTree().then((response) => {
      console.log('部门组织树结构=', response.data);
    state.departmentOptions = response.data;
  })
  });
};
// 组织节点单击事件
const handleCompanyClick = (data: any) => {
  state.queryParams.departmentId = data.departmentId;
  state.currentCompany = data
  // 部门数据
  // console.log('当前点击公司节点的部门数据为=', state.currentCompany);
    if (state.currentCompany || !state.currentCompany.departments) {
      state.departmentProps=[]
    }
    const dept = state.currentCompany.departments
    state.departmentOptions = dept
    // console.log('部门树结构=', state.departmentOptions);
    getList();
    state.queryParams.departmentId = 0
};
// 处理部门的多语言名称
const getDeptName = (dept: any) => {
  // console.log('解析部门多名称=', dept);
  try {
    const nameObj = JSON.parse(dept.name)
    return nameObj.zh_CN || nameObj.en_US || dept.name
  } catch {
    return dept.name
  }
}
// 组织树筛选节点
const filterNode = (value: string, data: any) => {
  console.log('filterNode数据传入=',value,data);
  if (!value) return true;
  return data.name.includes(value);
};

// update 需要修改 绑定到按钮 用户状态修改
const handleStatusChange = (row: any) => {
  let text = row.active === true ? "启用" : "停用";
  ElMessageBox({
    title: "警告",
    message: '确认要"' + text + '""' + row.username + '"用户吗?',
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    beforeClose: (action: string, instance: any, done: any) => {
      if (action === "confirm") {
        return changeUserStatus(row.id, row.active).then(() => {
          ElMessage.success(text + "成功");
          done();
        });
      } else {
        done();
      }
    },
  }).catch(() => {
    row.active = row.active === true ? "1" : "0";
  });
};
/** 删除按钮操作 */
const handleDelete = (row: any) => {
  const userIds = row.id || state.ids;
  ElMessageBox({
    message: '是否确认删除用户编号为"' + userIds + '"的数据项?',
    title: "警告",
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
  }).then(function () {
    return delUser(userIds).then((res: any) => {
      if (res.code === 200) {
        getList();
        ElMessage.success("删除成功");
      } else {
        ElMessage.error("删除失败");
      }
    });
  });
};

/** 重置密码按钮操作 */
const handleResetPwd = async (value: any) => {
  resetUserPwd(value).then((res: any) => {
    ElMessage.success("重置密码邮件已发送，请注意查收！！！");
  });
};

/** 启用TOTP按钮操作 */
const handleOK = async () => {
  if (!state.isPassword) {
    ElMessage.error("密码不能为空！");
    return;
  }
  const data = {
    value: state.value, // 需要启用TOTP的用户信息
    password: state.isPassword, // 输入的密码
  };
  console.log("启用TOTP传参=", data);
  try {
    await enableTotp(data);  // 传递用户数据和密码到后端
    ElMessage.success("用户【" + state.value.login + "】双重验证已激活！！！");
    state.isTOTP = false; // 关闭对话框
  } catch (error) {
    ElMessage.error("启用 TOTP 失败！");
  }
}
const handleEnableTotp = async (value: any) => {
  state.isTOTP = true
  state.value = value
};

// 分页改变 size
const onHandleSizeChange = (val: number) => {
  state.queryParams.pageSize = val;
  handleQuery();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
  state.queryParams.pageNum = val;
  handleQuery();
};

/** 导出按钮操作 */
const handleExport = () => {
  const queryParams = state.queryParams;
  let data: any = new Date().getTime() / 1000
  let time = parseInt(data) + '';
  queryParams.filename = "用户表_" + time + ".xlsx"
  ElMessageBox({
    message: "是否确认导出所有用户数据项?",
    title: "警告",
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(function () {
      return exportUser(queryParams);
    })
    .then((response: any) => {
      handleFileError(response, queryParams.filename)
    });
};
// 字典状态字典翻译
// const sexFormat = (row: any, column: any) => {
//   return proxy.selectDictLabel(state.sexOptions, row.gender);
// };
// 页面加载时
onMounted(() => {
  getList();
  getTreeselect();
  // 查询显示状态数据字典
  // getDicts("sys_normal_disable").then((response) => {
  //   state.statusOptions = response.data;
  // });
  // 查询显示性別数据字典
  // proxy.getDicts("sys_user_sex").then((response: any) => {
  //   state.sexOptions = response.data;
  // });

  proxy.mittBus.on("onEditUserModule", (res: any) => {
    handleQuery();
  });
});

// 页面卸载时
onUnmounted(() => {
  proxy.mittBus.off("onEditUserModule");
});
</script>

<style scoped lang="scss">
.system-user-container {
  .system-user-search {
    text-align: left;

    .system-user-search-btn {
      text-align: center;
      margin-left: 70px;
    }
  }

  .system-user-photo {
    width: 40px;
    height: 40px;
    border-radius: 100%;
  }
}

.custom-tree-node {
  width: 100%;
  padding: 5px 0;
}

.head-container {
  width: 100%;
  // padding: 20px;
  background: #ffffff;
}

.main-tree {
  width: 100%;
}

.tree-node-content {
  width: 100%;
  padding: 5px 0;
}

.company-node {
  display: flex;
  align-items: center;
  padding: 8px 0;
}

.department-subtree {
  margin-left: 30px;
  border-left: 1px solid #ebeef5;
  padding-left: 15px;
}

.nested-tree {
  background: transparent;
}

.nested-tree>>>.el-tree-node__content {
  height: 32px;
}

.dept-node {
  display: flex;
  align-items: center;
  padding: 5px 0;
}

.node-name {
  margin: 0 8px;
}

.node-type {
  color: #909399;
  font-size: 12px;
  margin-left: 8px;
}

.el-icon-office-building {
  color: #409EFF;
  margin-right: 6px;
}

.el-icon-s-home {
  color: #67C23A;
  margin-right: 6px;
}

/* 移除嵌套树的默认缩进 */
.nested-tree>>>.el-tree-node__indent {
  width: 0 !important;
}
</style>
