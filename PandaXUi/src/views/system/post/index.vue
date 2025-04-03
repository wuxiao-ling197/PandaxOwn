<template>
  <div class="app-container">
    <el-card shadow="always">
      <!-- 查询 -->
      <el-form
          :model="state.queryParams"
          ref="queryForm"
          :inline="true"
          label-width="68px"
      >
        <el-form-item label="岗位编码" prop="id">
          <el-input
              placeholder="请输入岗位编码模糊查询"
              clearable
              @keyup.enter="handleQuery"
              style="width: 240px"
              v-model="state.queryParams.jobCode"
          />
        </el-form-item>
        <el-form-item label="岗位名称" prop="name">
          <el-input
              placeholder="请输入岗位名称模糊查询"
              clearable
              @keyup.enter="handleQuery"
              style="width: 240px"
              v-model="state.queryParams.jobName"
          />
        </el-form-item>
        <el-form-item label="部门" prop="departmentId">
          <el-input
              v-model="state.queryParams.departmentId"
              placeholder="请输入部门编码模糊查询"
              clearable
              style="width: 240px"
          >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" plain @click="handleQuery">
            <SvgIcon name="elementSearch"/>
            搜索
          </el-button>
          <el-button @click="resetQuery">
            <SvgIcon name="elementRefresh"/>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span class="card-header-text">岗位列表</span>
          <div>
            <el-button
                type="primary"
                plain
                v-auth="'system:post:add'"
                @click="onOpenAddModule"
            >
              <SvgIcon name="elementPlus"/>
              新增
            </el-button>
            <el-button
                type="danger"
                plain
                v-auth="'system:post:delete'"
                :disabled="state.multiple"
                @click="onTabelRowDel"
            >
              <SvgIcon name="elementDelete"/>
              删除
            </el-button>
            <el-button
                type="warning"
                plain
                v-auth="'system:post:export'"
                @click="onTabelRowDel"
            >
              <SvgIcon name="elementDownload"/>
              导出
            </el-button>
          </div>
        </div>
      </template>
      <!--数据表格-->
      <el-table
          v-loading="state.loading"
          :data="state.tableData"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" align="center"/>
        <el-table-column label="岗位编号" align="center" prop="id" sortable />
        <el-table-column label="部门" align="center" prop="department_id" sortable />
        <el-table-column label="岗位名称" align="center" prop="name">
          <template #default="{row}">
            {{ JSON.parse(row.name).zh_CN || JSON.parse(row.name).en_US }}
          </template>
        </el-table-column>
        <el-table-column label="岗位排序" align="center" prop="sequence" sortable/>
        <!-- <el-table-column label="岗位介绍" align="center" prop="job_details"/> -->
        <el-table-column label="期望招聘人数/人" align="center" prop="no_of_recruitment" sortable/>
        <el-table-column label="在岗员工/人" align="center" prop="no_of_employee" sortable>
          <template #default="scope">
            {{ scope.row?.no_of_employee  ?  scope.row.no_of_employee: 0 }}
          </template>
        </el-table-column>
        <el-table-column label="岗位配置/人" align="center" prop="expected_employees" sortable>
          <template #default="scope">
            {{ scope.row?.expected_employees  ?  scope.row.expected_employees: 0 }}
          </template>
        </el-table-column>
        <!-- <el-table-column label="申请投递/人" align="center" prop="expected_employees" sortable>
          <template #default="scope">
            {{ scope.row?.no_of_hired_employee  ?  scope.row.no_of_hired_employee: 0 }}
          </template>
        </el-table-column> -->
        <el-table-column label="已雇佣/人" align="center" prop="no_of_hired_employee" sortable>
          <template #default="scope">
            {{ scope.row?.no_of_hired_employee  ?  scope.row.no_of_hired_employee: 0 }}
          </template>
        </el-table-column>
        <el-table-column label="招聘类型" align="center" prop="contract_type_id" sortable>
          <template #default="scope">
            {{ scope.row?.contract_type_id  ?  scope.row.contract_type_id: 0 }}
          </template>
        </el-table-column>
        <el-table-column
            label="发布"
            align="center"
            prop="is_published"
        >
          <template #default="scope">
            <el-tag
                effect="dark" size="large"
                :type="scope.row.is_published === true ? 'success' : scope.row.is_published === false ? 'danger' : 'warning' " 
                disable-transitions
            >{{ scope.row.is_published === true ? "已发布": "暂未发布" }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
            label="状态"
            align="center"
            prop="active"
        >
          <template #default="scope">
            <el-tag
                effect="dark" size="large"
                :type="scope.row.active === true ? 'success' : scope.row.active === false ? 'danger' : 'warning' " 
                disable-transitions
            >{{ scope.row.active === true ? "在招": "取消" }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="发布人" align="center" prop="user_id" sortable />
        <el-table-column label="发布日期" align="center" prop="create_date">
          <template #default="scope">
               <span>{{ dayjs(scope.row.create_date).format('YYYY-MM-DD HH:mm:ss')}}</span>
            </template>
        </el-table-column>

        <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
        >
          <template #default="scope">
            <el-popover  placement="left">
              <template #reference>
                <el-button type="primary" circle ><SvgIcon name="elementStar"/></el-button>
              </template>
              <div>
                <el-button text type="primary" v-auth="'system:post:edit'" @click="onPublish(scope.row)">
                  <el-icon><Promotion /></el-icon>
                  发布
                </el-button>
              </div>
              <div>
                <el-button text type="primary" v-auth="'system:post:edit'" @click="onOpenEditModule(scope.row)">
                  <SvgIcon name="elementEdit"/>
                  修改
                </el-button>
              </div>
              <div>
                <el-button text type="primary" v-auth="'system:post:delete'" @click="onTabelRowDel(scope.row)">
                  <SvgIcon name="elementDelete"/>
                  删除
                </el-button>
              </div>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页设置-->
      <div v-show="state.total > 0">
        <el-divider></el-divider>
        <el-pagination
            background
            :total="state.total"
            :current-page="state.queryParams.pageNum"
            :page-size="state.queryParams.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <!-- 添加或修改岗位对话框 -->
    <EditModule ref="editModuleRef" :title="state.title"/>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  toRefs,
  reactive,
  onMounted,
  getCurrentInstance,
  onUnmounted,
} from "vue";
import {ElMessageBox, ElMessage} from "element-plus";
import {listPost, delPost, publishPost} from "@/api/system/post";
import EditModule from "./component/editModule.vue";
import dayjs from 'dayjs';
const {proxy} = getCurrentInstance() as any;
const editModuleRef = ref();
const state = reactive({
  // 遮罩层
  loading: true,
  // 选中数组
  ids: [],
  // 非单个禁用
  single: true,
  // 非多个禁用
  multiple: true,
  // 弹出层标题
  title: "",
  // 岗位表格数据
  tableData: [],
  // 总条数
  total: 0,
  // 状态数据字典
  statusOptions: [],
  // 查询参数
  queryParams: {
    // 页码
    pageNum: 1,
    // 每页大小
    pageSize: 10,
    jobCode: undefined,
    jobName: undefined,
    departmentId: undefined,
  },
});

/** 查询岗位列表 */
const handleQuery = () => {
  state.loading = true;
  listPost(state.queryParams).then((response) => {
    console.log("岗位列表：",response.data.data);
    state.tableData = response.data.data;
    state.total = response.data.total;
    state.loading = false;
  });
};
// 名称
const getDeptName = (dept: any) => {
  // console.log('解析部门多名称=', dept);
  try {
    const nameObj = JSON.parse(dept.name)
    return nameObj.zh_CN || nameObj.en_US || dept.name
  } catch {
    return dept.name
  }
}
/** 重置按钮操作 */
const resetQuery = () => {
  state.queryParams.jobName = undefined;
  state.queryParams.jobCode = undefined;
  state.queryParams.departmentId = undefined;
  handleQuery();
};

const handleCurrentChange = (val: number) => {
  state.queryParams.pageNum = val
  handleQuery()
}
const handleSizeChange = (val: number) => {
  state.queryParams.pageSize = val
  handleQuery()
}


// 判断岗位招满还是取消
// const isFinish = ()

// 打开新增岗位弹窗
const onOpenAddModule = () => {
  state.title = "添加岗位";
  editModuleRef.value.openDialog({});
};
// 打开编辑岗位弹窗
const onOpenEditModule = (row: object) => {
  state.title = "修改岗位";
  editModuleRef.value.openDialog(row);
};
/** 发布按钮操作 */
const onPublish = (row: any) => {
  console.log(row.id);
  
  // const postIds = row.id || state.ids;
  ElMessageBox({
    message: '确认发布岗位编号为"' + row.id + '"的招聘信息吗?',
    title: "提示",
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(function () {
    return publishPost(row.id).then((res:any) => {
      console.log("发布=", res);
      
      if (res.code === 200){
        handleQuery();
        ElMessage.success("发布成功");
      }else {
        ElMessage.error("发布失败");
      }
    });
  });
};
/** 删除按钮操作 */
const onTabelRowDel = (row: any) => {
  const postIds = row.postId || state.ids;
  ElMessageBox({
    message: '是否确认删除岗位编号为"' + postIds + '"的数据项?',
    title: "警告",
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(function () {
    return delPost(postIds).then((res:any) => {
      if (res.code === 200){
        handleQuery();
        ElMessage.success("删除成功");
      }else {
        ElMessage.error("删除失败");
      }
    });
  });
};
// 多选框选中数据
const handleSelectionChange = (selection: any) => {
  state.ids = selection.map((item: any) => item.postId);
  state.single = selection.length != 1;
  state.multiple = !selection.length;
};
// 页面加载时
onMounted(() => {
  // 查询岗位信息
  handleQuery();
  // 查询岗位状态数据字典
  proxy.getDicts("sys_normal_disable").then((response: any) => {
    state.statusOptions = response.data;
  });
  proxy.mittBus.on("onEditPostModule", (res: any) => {
    handleQuery();
  });
});
// 页面卸载时
onUnmounted(() => {
  proxy.mittBus.off("onEditPostModule");
});
</script>
