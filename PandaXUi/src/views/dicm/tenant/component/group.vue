<template>
    <div class="app-container">
        <el-card>
            <!--  表格上方搜索按钮 -->
            <div class="row" style="display: flex; justify-content: flex-end; margin-bottom: 3px;">
                <el-input class="search_input" v-model="search" placeholder="请输入搜索数据" clearable>
                    <template #prepend>
                        <el-select v-model="colData" placeholder="搜  索" style="width: 80px">
                            <el-option v-for="item in colData" :key="item.title" :label="item.title"
                                :value="item.title" />
                        </el-select>
                    </template>
                    <template #append>
                        <el-button><el-icon>
                                <Search />
                            </el-icon></el-button>
                    </template></el-input>
                <!-- 新建按钮 -->
                <el-button type="primary" v-auth="'system:organization:edit'" @click="handleAdd()"
                    style="margin-right: 5px;">
                    <SvgIcon name="elementEdit" style="margin-right: 3px;" />
                    添加
                </el-button>

            </div>
            <el-table v-loading="state.loading" :data="state.tableData.data" row-key="id" border default-expand-all>
                <el-table-column prop="id" label="编码" width="40" type="selection" :selectable="selectable"/>
                <el-table-column prop="name" label="名称" width="100" align="center" />
                <el-table-column prop="slug" label="短标识符" width="100" align="center" />
                <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" align="center" />
                <el-table-column prop="color" label="颜色标签" width="100" align="center" />
                <el-table-column prop="parent_id" label="上级租户组" width="100" align="center">
                    <template #default="{ row }">
                        <el-link type="primary" :underline="false" style="font-weight: 600;"
                            @click="hanlder2Parent(row.parent_id)">
                            {{ row.parent_id == 0 ? "-----" : row.parent_id }}
                        </el-link>
                    </template>
                </el-table-column>
                <el-table-column prop="description" label="描述" width="150" align="center" />
                <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                    <template #default="scope">
                        <el-popover placement="left">
                            <template #reference>
                                <el-button type="primary" circle>
                                    <SvgIcon name="elementStar" />
                                </el-button>
                            </template>
                            <div>
                                <el-button text type="primary" v-auth="'system:organization:edit'"
                                    @click="handleAddTenant(scope.row)">
                                    <SvgIcon name="elementEdit" />
                                    添加租户
                                </el-button>
                            </div>
                            <div>
                                <el-button text type="primary" v-auth="'system:organization:add'"
                                    @click="handleUpdate(scope.row)">
                                    <SvgIcon name="elementPlus" />
                                    编辑
                                </el-button>
                            </div>
                            <div>
                                <el-button v-if="scope.row.parentId != 0" text type="primary"
                                    @click="handleDelete(scope.row.id)">
                                    <SvgIcon name="elementDelete" />
                                    删除
                                </el-button>
                            </div>
                        </el-popover>
                    </template>
                </el-table-column>
                <el-table-column prop="created" label="添加" width="180" align="center">
                    {{ dayjs(state.tableData.data.created).format('YYYY-MM-DD HH:mm:ss')
                    }}
                </el-table-column>
                <el-table-column prop="last_updated" label="更新" width="180" align="center">
                    {{ dayjs(state.tableData.data.last_updated).format('YYYY-MM-DD HH:mm:ss')
                    }}
                </el-table-column>
                <el-table-column prop="deleted" label="归档" width="180" align="center">
                    {{ state.tableData.data.deleted === "0001-01-01T00:00:00Z" ? "使用中" :
                        dayjs(state.tableData.data.deleted).format('YYYY-MM-DD HH:mm:ss')
                    }}
                </el-table-column>

            </el-table>
        </el-card>
        <!-- 添加租户组实例对话框 -->
        <el-dialog v-model="state.isAddGroup" :title="state.title">
            <!-- :rules="state.formRules" -->
            <el-form :model="state.formData" label-width="auto" style="max-width: 600px">
                <el-form-item label="名称" required>
                    <el-input v-model="state.formData.name" />
                </el-form-item>
                <el-form-item label="标识符" required>
                    <el-input v-model="state.formData.slug" />
                </el-form-item>
                <el-form-item label="上级租户组">
                    <!-- <el-select v-model="state.formData.parent_id" placeholder="请选择上级组织">
                        <el-option v-for="item in state.groupOptions" :label="item.name" :value="item.id"  />
                    </el-select> -->
                    <el-tree-select v-model="state.formData.parent_id" :data="state.groupOptions"
                        :props="{ value: 'id', label: 'name', children: 'children' }" node-key="id" check-strictly
                        placeholder="请选择上级租户组" />
                </el-form-item>
                <el-form-item label="lft" required>
                    <el-input v-model="state.formData.lft" required />
                </el-form-item>
                <el-form-item label="rght" required>
                    <el-input v-model="state.formData.rght" required />
                </el-form-item>
                <el-form-item label="tree_id" required>
                    <el-input v-model="state.formData.tree_id" required />
                </el-form-item>
                <el-form-item label="级别" required>
                    <el-input v-model="state.formData.level" required />
                </el-form-item>
                <el-form-item label="自定义配置数据" required>
                    <el-input v-model="state.formData.custom_field_data" type="textarea" :rows="5" placeholder='例如：
{
    "ids":[5,6],
    "group_id":8
}' required />
                </el-form-item>
                <el-form-item label="描述" required>
                    <el-input v-model="state.formData.description" type="textarea" required />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="onCancel">取 消</el-button>
                    <el-button type="primary" @click="onSubmitGroup" :loading="state.loading">保 存</el-button>
                </span>
            </template>
        </el-dialog>

        <!-- 添加租户对话框 -->
        <el-dialog v-model="state.isAddTenant" :title="state.title">
            <template #header>
                <div class="custom-dialog-title">
                    <span>租户组 </span>
                    <span class="dynamic-name-style">{{ currentRowName }}</span>
                    <span> 添加租户</span>
                </div>
            </template>
            <el-form :model="state.formData" label-width="auto" style="max-width: 600px">
                <el-form-item label="名称">
                    <el-select v-model="state.groupId" collapse-tags collapse-tags-tooltip placeholder="选择租户组"
                        node-key="id" disabled>
                        <el-option v-for="item in state.groupOptions" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="租户">
                    <el-select v-model="state.tenantIds" multiple collapse-tags collapse-tags-tooltip
                        :max-collapse-tags="3" placeholder="选择租户">
                        <el-option v-for="item in state.tenantOptions" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="onCancel">取 消</el-button>
                    <el-button type="primary" @click="onSubmitTenant" :loading="state.loading">保 存</el-button>
                </span>
            </template>
        </el-dialog>
        <!-- 分页 -->
        <div v-show="state.tableData.total > 0">
            <el-divider></el-divider>
            <el-pagination background :total="state.tableData.total" :page-sizes="[10, 20, 30, 50, 100]"
                :current-page="state.queryParams.pageNum" :page-size="state.queryParams.pageSize"
                layout="total, sizes, prev, pager, next, jumper" @size-change="onHandleSizeChange"
                @current-change="onHandleCurrentChange" />
        </div>
    </div>
</template>
<script lang="ts" setup>
import { addTenantGroup, deleteTenantGroup, groupAddTenant, listTenant, listTenantGroup, tenantGroupTree, updateTenantGroup } from '@/api/dicm/tenant';
import { onMounted, reactive, ref } from 'vue';
import { dayjs, ElMessage } from 'element-plus'
import { useRouter } from 'vue-router';

const selectable=ref()

// 1. 获取 router 实例
const router = useRouter();
// 表格上方搜索
var search = ref()
const colData = reactive([
    { title: "编码", istrue: true, value: "id" },
    { title: "名称", istrue: true, value: "name" },
    { title: "序列", istrue: true, value: "slug" },
    { title: "颜色", istrue: true, value: "color" },
    { title: "描述", istrue: true, value: "description" },
])

// 自定义 JSON 校验函数
const isValidJson = () => {
    const str = state.formData.custom_field_data
    // 首先，确保输入的是字符串类型
    if (typeof str !== 'string') {
        return false;
    }
    // 尝试解析 JSON 字符串
    try {
        JSON.parse(str);
    } catch (e) {
        // 如果解析过程中抛出错误 (SyntaxError)，则说明不是合法的 JSON
        return false;
    }
    // 如果没有抛出错误，则是合法的 JSON
    return true;
}
const currentRowName = ref()
const state = reactive({
    loading: true,
    // 表格渲染数据
    tableData: {
        data: [],
        total: 0,
    },
    title: "",
    tenantIds: [],
    groupId: "",
    // 表单-添加|编辑 定义为json类型
    formData: {
        name: undefined,
        slug: undefined,
        description: undefined,
        custom_field_data: undefined,
        lft: undefined,
        rght: undefined,
        tree_id: undefined,
        level: undefined,
        parent_id: undefined
    },
    // 表单校验需要完善
    formRules: {
        custom_field_data: [
            { required: true, message: 'JSON 数据不能为空', trigger: 'blur' },
            { validator: isValidJson, trigger: 'blur' } // 在失焦时触发 JSON 格式校验
            // 你也可以设置 trigger: 'change' 来在内容改变时实时校验，但对性能稍有影响
        ],
    },
    // 弹出框
    isAddGroup: false,
    isAddTenant: false,
    // 查询参数
    queryParams: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
        slug: undefined,
        level: undefined,
        parent_id: undefined,
        id: undefined
    },
    //多选数据项
    groupOptions: [
        // {id:2,name:"测试post"},
        // {id:3,name:"测试外键"},
        // {id:4,name:"1"},
    ],
    tenantOptions: [],
})

// 加载租户组列表 cg
const tenantGroupList = () => {
    state.loading = true;
    listTenantGroup(state.queryParams).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        // console.log("list=", response.data.data);
        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
        state.loading = false;
    })
    // 加载租户组层级
    tenantGroupTree().then((res: any) => {
        // console.log("租户层级：", res.data);
        state.groupOptions = res.data
    })
}


/**按钮操作 */
// 跳转父级 cg
const hanlder2Parent = (parent: any) => {
    const payload = { id: parent }
    listTenantGroup(payload).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
        state.loading = false;
    })
}
// 租户组添加租户 cg
const handleAddTenant = (row: any) => {
    state.isAddTenant = true
    currentRowName.value = row.name
    state.groupId = row.id
    listTenant({}).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        state.tenantOptions = response.data.data;
    })

}
// 添加租户组
const handleAdd = () => {
    console.log("添加参数=", state.formData);
    state.isAddGroup = true
    state.title = "添加租户组实例"
}

const handleUpdate = (data: any) => {
    state.isAddGroup = true;
    state.title = "编辑租户组信息"
    state.formData = data
    console.log("编辑参数：", state.formData);

}
/** 删除按钮操作 */
const handleDelete = (id: any) => {
    // console.log([id], typeof id);
    deleteTenantGroup([id]).then((res: any) => {
        if (res.code === 200) {
            ElMessage({
                type: "success",
                message: "归档成功",
            });
        }
    }),
        tenantGroupList();
}
const onSubmitGroup = () => {
    console.log("对话框参数=", state.formData);
    // 假设 formData 是你组件中的状态
    let payload = { ...state.formData }; // 复制一份，避免直接修改状态
    if (payload.parent_id === undefined || payload.parent_id === "") {
        // payload.parent_id = null; // 显式发送 null
        // 或者直接省略该字段 (sql.NullInt64 也能正确处理省略的情况):
        delete payload.parent_id;
        console.log(payload);
    }
    if (state.title === "添加租户组实例") {
        addTenantGroup(payload).then((res: any) => {
            if (res.code === 200) {
                ElMessage({
                    type: "success",
                    message: "新增成功",
                });
            }
            tenantGroupList();

        })
        state.isAddGroup = false
    } else {
        updateTenantGroup(payload).then((res: any) => {
            console.log("修改=",res);
            if (res.code === 200) {
                ElMessage({
                    type: "success",
                    message: "修改成功",
                });
            }
            tenantGroupList();
        })
    }
}
// cg
const onSubmitTenant = () => {
    const data = { group_id: state.groupId, tenant_ids: state.tenantIds }
    groupAddTenant(data).then((res: any) => {
        if (res.code === 200) {
            ElMessage({
                type: "success",
                message: "操作成功",
            });
        }
    })
    state.isAddTenant = false
}
const onCancel = () => {
    if (state.isAddGroup == true) {
        state.isAddGroup = false
    }
    else if (state.isAddTenant == true) {
        state.isAddTenant = false
    }
}


// 分页改变 size
const onHandleSizeChange = (val: number) => {
    state.queryParams.pageSize = val;
    tenantGroupList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
    state.queryParams.pageNum = val;
    tenantGroupList();
};
onMounted(() => {
    tenantGroupList();
})

</script>
<style lang="css" scoped>
.search_input {
    position: absolute;
    width: 60%;
    left: 0;
    margin-left: 45px;
}

.custom-dialog-title {
    font-size: 16px;
    /* Element Plus 默认标题字体大小可能不同，按需调整 */
    font-weight: 500;
    /* Element Plus 默认标题字体粗细可能不同，按需调整 */
    color: #303133;
    /* Element Plus 默认标题颜色 */
    /* 你可以添加其他需要的样式，如行高、边距等 */
}

.dynamic-name-style {
    color: #4382c1;
    /* 蓝色字体 (Element Plus 主题蓝) */
    font-weight: 600;
    /* 下划线 */
    /* 你可以添加其他样式，如字体粗细、内外边距等 */
    margin: 0 4px;
    /* 例如，在动态名称两侧添加一点间距 */
}

.dialog-footer {
    margin-left: 60px;
}
</style>