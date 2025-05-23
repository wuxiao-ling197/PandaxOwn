<template>
    <div class="app-container">
        <div style="margin: 3px;">
            <el-button-group>
                <el-button type="primary" @click="currentView = 'List'" round><el-icon style="margin-right: 3px;">
                        <Menu />
                    </el-icon>查看列表</el-button>
                <!-- <el-button type="primary" @click="currentView = 'Reserve'" round><el-icon style="margin-right: 3px;">
                        <DocumentAdd />
                    </el-icon>机柜预留</el-button>
                <el-button type="primary" @click="currentView = 'Role'" round><el-icon style="margin-right: 3px;">
                        <Files />
                    </el-icon>机柜类型</el-button> -->
                <el-button type="primary" @click="currentView = 'Group'" round><el-icon style="margin-right: 3px;">
                        <OfficeBuilding />
                    </el-icon> 租户组</el-button>
                <!-- <el-button type="primary" @click="currentView = 'Edit'" round><el-icon style="margin-right: 3px;">
                        <EditPen />
                    </el-icon>编辑</el-button> -->
            </el-button-group>
        </div>
        <div style="margin: 3px;">
            <Group v-if="currentView === 'Group'"></Group>
            <!-- <Role v-else-if="currentView === 'Role'"></Role>
            <List v-else-if="currentView === '3D'"></List>
            <Info v-else-if="currentView === 'Info'" ref="editFormRef"></Info> -->
            <div v-else-if="currentView === 'List'">
                <el-card>
                    <!-- 表格顶部搜索筛选按钮 -->
                    <div class="row" style="display: flex; justify-content: flex-end; margin-bottom: 3px;">
                        <!-- <el-popover placement="left-start" :width="600">
                            <template #reference> -->
                        <el-input class="search_input" v-model="search" placeholder="请输入搜索数据" clearable>
                            <template #prepend>
                                <el-select v-model="colData" placeholder="搜  索" style="width: 80px">
                                    <el-option v-for="item in colData" :key="item.title" :label="item.title"
                                        :value="item.title" />
                                </el-select>
                            </template>
                            <template #append>
                                <!--  搜索按钮 -->
                                <el-button><el-icon>
                                        <Search />
                                    </el-icon></el-button>
                            </template></el-input>
                        <!-- </template>
                        </el-popover> -->
                        <!-- 表格筛选列 -->
                        <el-button type="primary" @click="handleAdd()">
                            <SvgIcon name="elementPlus" style="margin-right: 3px;" />添加
                        </el-button>
                        <el-button type="primary" @click="state.isAddGroup = true"><el-icon style="margin-right: 3px;">
                                <Connection />
                            </el-icon>归属租户组</el-button>
                        <el-popover placement="right-start" title="筛选列" :width="40" trigger="click">
                            <template #reference>
                                <el-button><el-icon>
                                        <Grid />
                                    </el-icon></el-button>
                            </template>
                            <div class="scrollable-checkbox-list">
                                <el-checkbox v-for="col in colData" :key="col.value" v-model="col.istrue"
                                    :label="col.title" style="display: block; margin-bottom: 5px;"></el-checkbox>
                            </div>
                        </el-popover>
                    </div>
                    <!--数据表格-->
                    <el-table v-loading="state.loading" :data="state.tableData.data" row-key="id" border
                        default-expand-all>
                        <!-- v-if="colData[0].istrue" -->
                        <el-table-column prop="id" label="编码" width="40" type="selection" :selectable="selectable"/>
                        <el-table-column prop="name" label="名称" width="100" align="center"  fixed>
                        </el-table-column>
                        <!-- 点击租户组实现视图跳转，但是仅仅跳转到列表（未携带参数） -->
                        <el-table-column prop="group_id" label="租户组" align="center" width="100">
                            <!-- {{ state.tableData.data.group_id.Int64 == 0 ? "-----": state.tableData.data.group_id.Int64}} -->
                            <template #default="{ row }">
                                <el-link type="primary" :underline="false" @click="currentView = 'Group'"
                                    style="font-weight: 600;">
                                    {{ row.group_id }}
                                </el-link>
                            </template>
                        </el-table-column>
                        <el-table-column prop="slug" label="slug" width="100" align="center">
                            <!-- <template #default="scope">
            <el-tag :type="scope.row.active === 'online' ? 'success' : 'danger'" disable-transitions>{{ scope.row.active ?
              "在线" : "已停用" }}
            </el-tag>
          </template> -->
                        </el-table-column>
                        <el-table-column label="描述" align="center" prop="description" width="100" />
                        <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" align="center" />
                        <el-table-column prop="comments" label="评价" width="150" align="center" />
                        <el-table-column prop="created" label="添加时间" width="180" align="center">
                            {{ dayjs(state.tableData.data.created).format('YYYY-MM-DD HH:mm:ss')
                            }}
                        </el-table-column>
                        <el-table-column prop="last_updated" label="更新时间" width="180" align="center">
                            {{ dayjs(state.tableData.data.last_updated).format('YYYY-MM-DD HH:mm:ss')
                            }}
                        </el-table-column>
                        <el-table-column prop="deleted" label="归档" width="180" align="center">
                    {{ dayjs(state.tableData.data.deleted).format('YYYY-MM-DD HH:mm:ss')
                    }}
                </el-table-column>
                        <el-table-column fixed label="" align="center" class-name="small-padding fixed-width">
                            <template #default="scope">
                                <el-popover placement="left">
                                    <template #reference>
                                        <!--  @click="handleUpdate(scope.row)" -->
                                        <el-button text type="primary" round><el-icon style="margin-right: 3px;">
                                                <EditPen />
                                            </el-icon>操作</el-button>
                                    </template>
                                    <div>
                                        <el-button text type="primary" v-auth="'system:user:edit'"
                                            @click="handleUpdate(scope.row)">
                                            <SvgIcon name="elementEdit" />
                                            修改
                                        </el-button>
                                    </div>
                                    <!-- <div>
                                        <el-button text type="primary" v-auth="'system:user:add'" @click="handleAdd()">
                                            <SvgIcon name="elementPlus" />
                                            新增
                                        </el-button>
                                    </div> -->
                                    <div>
                                        <el-button v-if="scope.row.parentId != 0" text type="primary" @click="handleDelete(scope.row)">
                                            <SvgIcon name="elementDelete" />
                                            删除
                                        </el-button>
                                    </div>
                                </el-popover>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-card>
                <div v-show="state.tableData.total > 0">
                    <el-divider></el-divider>
                    <el-pagination background :total="state.tableData.total" :page-sizes="[10, 20, 30, 50, 100]"
                        :current-page="state.queryParams.pageNum" :page-size="state.queryParams.pageSize"
                        layout="total, sizes, prev, pager, next, jumper" @size-change="onHandleSizeChange"
                        @current-change="onHandleCurrentChange" />
                </div>
            </div>
        </div>
    </div>
    <!-- 添加或修改岗位对话框 -->
    <EditModule ref="editFormRef" :title="state.title" />
    <!-- 添加租户对话框 cg-->
    <el-dialog v-model="state.isAddGroup" title="租户归属到租户组">
        <el-form :model="state.gdata" label-width="auto" style="max-width: 600px">
            <el-form-item label="名称">
                <!-- <el-select v-model="state.gdata.groupId" collapse-tags collapse-tags-tooltip placeholder="选择租户组"
                        node-key="id">
                        <el-option v-for="item in state.groupOptions" :label="item.name" :value="item.id" />
                    </el-select> -->
                <el-tree-select v-model="state.gdata.group_id" :data="state.groupOptions"
                    :props="{ value: 'id', label: 'name', children: 'children' }" node-key="id" check-strictly
                    placeholder="请选择归属租户组" />
            </el-form-item>
            <el-form-item label="租户">
                <el-select v-model="state.gdata.tenant_ids" multiple collapse-tags collapse-tags-tooltip
                    :max-collapse-tags="5" placeholder="选择租户">
                    <el-option v-for="item in state.tableData.data" :label="item.name" :value="item.id" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="onCancel">取 消</el-button>
                <el-button type="primary" @click="handle2Group()" :loading="state.loading">保 存</el-button>
            </span>
        </template>
    </el-dialog>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { dayjs, ElMessage } from 'element-plus'
import Group from './component/group.vue';

let currentView = ref<'List' | 'Group'>('List');
import { addTenant2Group, deleteTenant, listTenant, tenantGroupTree } from '@/api/dicm/tenant';
import EditModule from './component/edit.vue'

const editFormRef = ref()
var search = ref()
// colData中列出表格中的每一列，默认都展示
const colData = reactive([
    { title: "编码", istrue: true, value: "id" },
    { title: "名称", istrue: true, value: "name" },
    { title: "slug", istrue: true, value: "slug" },
    { title: "描述", istrue: true, value: "description" },
    { title: "自定义配置数据", istrue: true, value: "custom_field_data" },
    { title: "评价", istrue: true, value: "comments" },
    { title: "添加时间", istrue: false, value: "created" },
    { title: "更新时间", istrue: false, value: "last_updated" },
])
// 表格表头多选框
const selectable= ref()
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
    // 租户组组织结构
    groupOptions: [],
    // 查询参数 绑定到表格上方搜索框 todo
    queryParams: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
        slug: undefined,
        group_id: undefined,
        id: undefined
    },
    // 租户归属租户组
    gdata: {
        tenant_ids: [], //目标租户
        group_id: "", //目标组
    },
    isAddGroup: false, //对话框显示控制
});

// 加载租户列表 cg
const tenantList = () => {
    state.loading = true;
    listTenant(state.queryParams).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        // console.log(state.tableData.data);
        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
        state.loading = false;
    })
    // 加载租户组层级 如果需要在侧边添加组织结构
    tenantGroupTree().then((res: any) => {
        state.groupOptions = res.data
    })
}

// 视图对象跳转
// const handleClick = (id: any) => {
//     //   console.log('组织点击跳转路由携带参数：', id);
//     // /system/user/list
//     // router.push({ path: '/system/user/list', query: { id: id } })
//     //   router.push({name:'/system/user', query: { id: id }}) //http://192.168.0.5:7789/#/system/organization?department_id=2
// }

/** 归属到租户组按钮 */
const handle2Group = () => {
    addTenant2Group(state.gdata).then((res: any) => {
        console.log(state.gdata, res);
        if (res.code === 200) {
            ElMessage({
                type: "success",
                message: "操作成功",
            });
        }
    })
    state.isAddGroup = false
    tenantList();
}

/** 新增按钮操作 */
const handleAdd = () => {
    state.title = "添加租户";
    editFormRef.value.openDialog({});
};
/** 修改按钮操作 */
const handleUpdate = (row: any) => {
    console.log("编辑租户==", row);
    state.title = "修改租户";
    editFormRef.value.openDialog(row);
};
/** 删除按钮操作 deleteTenantGroup*/
const handleDelete = (row:any)=>{
    console.log("删除：",row);
    if(row.deleted !== "" || row.deleted !== undefined || row.deleted !== "0001-01-01T00:00:00Z"){
        ElMessage.warning("【" + row.name + "】已归档, 请勿重复操作");
    }else{
        deleteTenant([row.id]).then((res:any)=>{
        if (res.code === 200) {
            ElMessage({
                type: "success",
                message: "操作成功",
            });
        }
    })
    }
    tenantList();
}
/** 取消按钮 */
const onCancel = () => {
    if (state.isAddGroup == true) {
        state.isAddGroup = false
    }
    // else if (state.isAddTenant == true) {
    //     state.isAddTenant = false
    // }
}

// 分页改变 size
const onHandleSizeChange = (val: number) => {
    state.queryParams.pageSize = val;
    tenantList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
    state.queryParams.pageNum = val;
    tenantList();
};

// 页面加载时
onMounted(() => {
    // 查询列表
    tenantList();
})

</script>
<style lang="css" scoped>
.scrollable-checkbox-list {
    max-height: 300px;
    /* 控制最大高度 */
    width: 400px;
    overflow-y: auto;
    /* 垂直滚动 */
    padding-right: 8px;
    /* 避免滚动条遮挡内容 */
}

.search_input {
    position: absolute;
    width: 60%;
    /* 输入框绝对定位 */
    /* top: -50px;         调整到表格上方 */
    left: 0;
    /* 对齐表格左侧 */
    margin-left: 45px;
}
</style>