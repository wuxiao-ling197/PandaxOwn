<template>
    <div class="app_container">
        <div>
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
                    <!-- 新建按钮 -->
                    <el-button type="primary" @click="handleAdd()"><el-icon style="margin-right: 3px;">
                            <EditPen />
                        </el-icon>添加</el-button>
                    <!-- </template>
                        </el-popover> -->
                    <!-- 表格筛选列 -->
                    <el-popover placement="right-start" title="筛选列" :width="40" trigger="click">
                        <template #reference>
                            <el-button><el-icon>
                                    <Grid />
                                </el-icon></el-button>
                        </template>
                        <!-- <div class="scrollable-checkbox-list">
                                <el-checkbox v-for="col in colData" :key="col.value" v-model="col.istrue"
                                    :label="col.title" style="display: block; margin-bottom: 5px;"></el-checkbox>
                            </div> -->
                    </el-popover>
                </div>
                <el-table :data="state.tableData.data" v-loading="state.loading" :key="reload" border>
                    <el-table-column label="编码" prop="id" type="selection" :selectable="selectable"></el-table-column>
                    <el-table-column label="名称" prop="name"></el-table-column>
                    <el-table-column label="标识符" prop="slug"></el-table-column>
                    <el-table-column label="描述" prop="description"></el-table-column>
                    <el-table-column label="lft" prop="lft"></el-table-column>
                    <el-table-column label="rght" prop="rght"></el-table-column>
                    <el-table-column label="tree_id" prop="tree_id"></el-table-column>
                    <el-table-column label="级别" prop="level"></el-table-column>
                    <el-table-column label="上级" prop="parent_id"></el-table-column>
                    <el-table-column label="自定义配置数据" prop="custom_field_data"></el-table-column>
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
                                        @click="handleAddSite(scope.row)">
                                        <SvgIcon name="elementEdit" />
                                        添加站点
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
                    <el-table-column label="创建时间" prop="created"></el-table-column>
                    <el-table-column label="更新时间" prop="last_updated"></el-table-column>
                </el-table>
            </el-card>
            <!-- 添加站点组实例对话框 -->
            <el-dialog v-model="state.isAddGroup" :title="state.title">
                <!-- :rules="state.formRules" -->
                <el-form :model="state.formData" label-width="auto" style="max-width: 600px">
                    <el-form-item label="名称" required>
                        <el-input v-model="state.formData.name" />
                    </el-form-item>
                    <el-form-item label="标识符" required>
                        <el-input v-model="state.formData.slug" />
                    </el-form-item>
                    <el-form-item label="上级站点组">
                        <!-- <el-select v-model="state.formData.parent_id" placeholder="请选择上级组织">
                        <el-option v-for="item in state.groupOptions" :label="item.name" :value="item.id"  />
                    </el-select> -->
                        <el-tree-select v-model="state.formData.parent_id" :data="state.groupOptions"
                            :props="{ value: 'id', label: 'name', children: 'children' }" node-key="id" check-strictly
                            placeholder="请选择上级站点组" />
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

            <!-- 添加站点对话框 -->
            <el-dialog v-model="state.isAddSite" :title="state.title">
                <template #header>
                    <div class="custom-dialog-title">
                        <span>站点组 </span>
                        <span class="dynamic-name-style">{{ currentRowName }}</span>
                        <span> 添加站点</span>
                    </div>
                </template>
                <el-form :model="state.formData" label-width="auto" style="max-width: 600px">
                    <el-form-item label="名称">
                        <el-select v-model="state.groupId" collapse-tags collapse-tags-tooltip placeholder="选择站点组"
                            node-key="id" disabled>
                            <el-option v-for="item in state.groupOptions" :label="item.name" :value="item.id" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="站点">
                        <el-select v-model="state.siteIds" multiple collapse-tags collapse-tags-tooltip
                            :max-collapse-tags="3" placeholder="选择站点">
                            <el-option v-for="item in state.siteOptions" :label="item.name" :value="item.id" />
                        </el-select>
                    </el-form-item>
                </el-form>
                <template #footer>
                    <span class="dialog-footer">
                        <el-button @click="onCancel">取 消</el-button>
                        <el-button type="primary" @click="onSubmitSite" :loading="state.loading">保 存</el-button>
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


    </div>

</template>
<script setup lang="ts">
import { addSiteGroup, deleteSiteGroup, groupAddSite, listSiteGroup, listSites, siteGroupTree, updateSiteGroup } from '@/api/dicm/site';
import { ElMessage } from 'element-plus';
import { onMounted, reactive, ref } from 'vue';

const search = ref()
const selectable=ref()
const currentRowName = ref()
var reload = ref()
const colData = reactive([
    { title: "编码", istrue: true, value: "id" },
    { title: "名称", istrue: true, value: "name" },
    { title: "序列", istrue: true, value: "slug" },
    { title: "上级组", istrue: true, value: "parent_id" },
    { title: "级别", istrue: true, value: "level" },
    { title: "描述", istrue: true, value: "description" },
])
const state = reactive({
    loading: true,
    tableData: {
        data: [],
        total: 0,
    },
    groupOptions: [],
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
    siteOptions:[],//站点 todo 筛选出没有归属到站点组的
    siteIds:[],
    groupId: "",
    queryParams: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
        status: undefined,
        parentId: undefined,
        id: undefined
    },
    isAddGroup: false,
    isAddSite:false,
    title: "",
})

const siteGroupList = () => {
    listSiteGroup(state.queryParams).then((res: any) => {
        state.tableData.data = res.data.data;
        state.tableData.total = res.data.total;
        state.loading = false;
    })
    siteGroupTree().then((res: any) => {
        // console.log("树结构：", res.data);
        state.groupOptions=res.data
    })
}

/** 添加站点组 */ 
const handleAdd = () => {
    state.isAddGroup = true;
    state.title = "添加站点组实例";
}
// 提交添加站点组对话框内容
const onSubmitGroup = () => {
    console.log("添加参数=", state.formData);
    // 假设 formData 是你组件中的状态
    let payload = { ...state.formData }; // 复制一份，避免直接修改状态
    if (payload.parent_id === undefined || payload.parent_id === "") {
        // payload.parent_id = null; // 显式发送 null
        // 或者直接省略该字段 (sql.NullInt64 也能正确处理省略的情况):
        delete payload.parent_id;
        console.log(payload);
    }
    if (state.title === "添加站点组实例") {
        addSiteGroup(payload).then((res: any) => {
            if (res.code === 200) {
                ElMessage({
                    type: "success",
                    message: "新增成功",
                });
            }
        })
    } else {
        updateSiteGroup(payload).then((res: any) => {
            console.log("修改=",res);
            if (res.code === 200) {
                ElMessage({
                    type: "success",
                    message: "修改成功",
                });
            }
        })
    }
    state.isAddGroup = false;
    siteGroupList();
}
/** 站点归属到站点组 */
const handleAddSite=(row:any)=>{
    state.isAddSite = true
    currentRowName.value = row.name
    state.groupId = row.id
    listSites({}).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        state.siteOptions = response.data.data;
    })
}

/** 编辑按钮操作 */
const handleUpdate = (data: any) => {
    state.isAddGroup = true;
    state.title = "编辑站点组信息";
    state.formData = data;

}
/** 删除按钮操作 */
const handleDelete = (id:any)=>{
    console.log([id], typeof id);
    deleteSiteGroup([id]).then((res:any)=>{
        if (res.code === 200) {
            ElMessage({
                type: "success",
                message: "归档成功",
            });
        }
    }),
    siteGroupList();
}
// 提交站点组添加站点对话框
const onSubmitSite = () => {
    const data = { group_id: state.groupId, site_ids: state.siteIds }
    groupAddSite(data).then((res: any) => {
        console.log("groupAddSite=",res);
        
        if (res.code === 200) {
            ElMessage({
                type: "success",
                message: "操作成功",
            });
        }
    })
    state.isAddSite = false
}

/** 对话框取消 */
const onCancel = () => {
    if (state.isAddGroup == true) {
        state.isAddGroup = false
    }
    else if (state.isAddSite == true) {
        state.isAddSite = false
    }
}

// 分页改变 size
const onHandleSizeChange = (val: number) => {
    state.queryParams.pageSize = val;
    siteGroupList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
    state.queryParams.pageNum = val;
    siteGroupList();
};

onMounted(() => {
    siteGroupList();
})
</script>
<style lang="css" scoped>
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