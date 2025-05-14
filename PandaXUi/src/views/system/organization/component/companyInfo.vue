<template>
    <div class="system-menu-container">
        <!-- fullscreen -->
        <el-dialog v-model="state.isShowDialog" center top="40vh" width="90%" draggable>
            <template #header>
                <div style="font-size: large"
                    v-drag="['.system-menu-container .el-dialog', '.system-menu-container .el-dialog__header']">
                    {{ title }}</div>
            </template>
            <!-- 查询 -->
            <el-form :model="state.queryParams" ref="queryForm" :inline="true" label-width="120px">
                <el-form-item label="编码" prop="id" v-model="state.queryParams.id" width="200">
                    <el-input placeholder="请输入组织编码" clearable @keyup.enter.native="listCompanyInfo"
                        v-model="state.queryParams.id" />
                </el-form-item>
                <el-form-item label="组织名称" prop="name" width="200">
                    <el-input placeholder="请输入组织名称" clearable @keyup.enter.native="listCompanyInfo"
                        v-model="state.queryParams.name" />
                </el-form-item>
                <el-form-item label="组织邮箱" prop="email">
                    <el-input placeholder="请输入组织邮箱" clearable @keyup.enter.native="listCompanyInfo"
                        v-model="state.queryParams.email" />
                </el-form-item>
                <el-form-item label="联系方式" prop="phone">
                    <el-input placeholder="请输入组织联系方式" clearable @keyup.enter.native="listCompanyInfo"
                        v-model="state.queryParams.phone" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" plain @click="listCompanyInfo">
                        <SvgIcon name="elementSearch" />
                        搜索
                    </el-button>
                    <el-button @click="resetQuery">
                        <SvgIcon name="elementRefresh" />
                        重置
                    </el-button>
                </el-form-item>

            </el-form>
            <!-- 表格数据 :data="state.tableData"-->
            <el-table v-loading="state.loading" :data="filterTableData" row-key="id" border default-expand-all
                :tree-props="{ children: 'children', hasChildren: 'hasChildren' }" stripe>
                <el-table-column prop="id" label="编码" width="60" fixed />
                <el-table-column prop="name" label="组织名称" width="200"></el-table-column>
                <!-- <el-table-column prop="manager_id" label="负责人" width="120" /> -->
                <el-table-column prop="email" label="组织邮箱" width="180"></el-table-column>
                <el-table-column prop="phone" label="联系方式" width="150"></el-table-column>
                <el-table-column prop="parent_id" label="上级组织" width="120" sortable>
                    <template #defult="scope">
                        {{ scope.row.parent_id || '无' }}
                    </template>
                </el-table-column>
                <el-table-column prop="sequence" label="排序" width="60" sortable>
                    <template #defult="scope">
                        {{ scope.row.sequence || '' }}
                    </template>
                </el-table-column>
                <el-table-column prop="create_date" label="创建时间" width="180">
                    <template #default="scope">
                        {{ dayjs(scope.row.create_date).format('YYYY-MM-DD HH:mm:ss') }}
                    </template>
                </el-table-column>
                <el-table-column prop="active" label="状态" width="60">
                    <template #default="scope">
                        <el-tag :type="scope.row.active === true ? 'success' : 'danger'" disable-transitions>{{
                            scope.row.active ? "启用" : "已停用" }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="货币单位" align="center" prop="currency_id" width="120" />
                <el-table-column align="right" fixed width="200">
                    <template #header>
                        <el-input v-model="search" size="small" placeholder="输入名称" />
                    </template>
                    <template #default="scope">
                        <el-button size="small" type="primary" @click="handelEmp(scope.row.id)">
                            查看员工{{ state.emptotal }}
                        </el-button>
                        <el-button size="small" type="primary" @click="handleDep(scope.row.id)">
                            查看部门{{ state.deptotal }}
                        </el-button>
                    </template>
                 </el-table-column>
            </el-table>

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
import { reactive, onMounted, getCurrentInstance, computed, ref } from "vue";
import { listCompany } from "@/api/system/organization";
import { dayjs } from "element-plus";
import router from "@/router";


const props = defineProps({
    title: {
        type: String,
        default: () => "",
    },
})
const { proxy } = getCurrentInstance() as any;
const state = reactive({
    // 是否显示弹出层
    isShowDialog: false,
    // 遮罩层
    loading: true,
    // 弹出层标题
    title: "",
    // 组织表格树数据
    tableData: [] as any,
    child: undefined,
    // 状态数据字典
    statusOptions: [],
    // 查询参数
    queryParams: {
        id: undefined,
        name: undefined,
        phone: undefined,
        email: undefined
    },
    emptotal: undefined,
    deptotal: undefined
});
const search = ref('')
const filterTableData = computed(() =>
    state.tableData.filter(
        (data: any) =>
            !search.value ||
            data.name.toLowerCase().includes(search.value.toLowerCase())
    )
)

const listCompanyInfo = () => {
    state.loading = true;
    listCompany(state.queryParams).then((response: any) => {
        console.log("公司信息列表：", response.data);
        state.tableData = response.data;
        state.loading = false;
    });
}

const handelEmp=(row: any)=>{
    console.log('公司信息跳转携带参数handelEmp：', row);
    router.push({path:"/system/user", query: {company_id:row}})
}

const handleDep=(row: any)=>{
    console.log('公司信息跳转携带参数handleDep：', row);
    router.push({path:"/system/organization", query: {company_id:row}})
}

/** 重置按钮操作 */
const resetQuery = () => {
    state.queryParams.name = undefined;
    state.queryParams.id = undefined;
    state.queryParams.email = undefined;
    state.queryParams.phone = undefined;
    listCompanyInfo();
};

onMounted(() => {
    listCompanyInfo();
})

const openDialog = () => {
    state.isShowDialog = true
}

// 暴露方法给父组件
defineExpose({
    openDialog
})
</script>