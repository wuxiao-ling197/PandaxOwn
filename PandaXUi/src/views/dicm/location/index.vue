<template>
    <div class="app-container">
        <div style="margin: 3px;">
            <el-button-group>
                <el-button type="primary" @click="currentView = 'List'" round><el-icon style="margin-right: 3px;">
                        <Menu />
                    </el-icon>查看列表</el-button>
                <el-button type="primary" @click="currentView = 'Reserve'" round><el-icon style="margin-right: 3px;">
                        <DocumentAdd />
                    </el-icon>机柜预留</el-button>
                <el-button type="primary" @click="currentView = 'Role'" round><el-icon style="margin-right: 3px;">
                        <Files />
                    </el-icon>机柜类型</el-button>
                <el-button type="primary" @click="currentView = '3D'" round><el-icon style="margin-right: 3px;">
                        <OfficeBuilding />
                    </el-icon> 立体视图</el-button>
                <!-- <el-button type="primary" @click="currentView = 'Edit'" round><el-icon style="margin-right: 3px;">
                        <EditPen />
                    </el-icon>编辑</el-button> -->
            </el-button-group>
        </div>
        <div style="margin: 3px;">
            <Reserve v-if="currentView === 'Reserve'"></Reserve>
            <Role v-else-if="currentView === 'Role'"></Role>
            <List v-else-if="currentView === '3D'"></List>
            <Info v-else-if="currentView === 'Info'" ref="editFormRef"></Info>
            <!-- ref="state.rowData" -->
            <Editt v-else-if="currentView === 'Editt'"></Editt>
            <div v-else>
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
                        <!-- 新建按钮 @click="currentView = 'Editt'"-->
                        <el-button type="primary" @click="handleAdd()"><el-icon style="margin-right: 3px;">
                                <EditPen />
                            </el-icon>添加实例</el-button>
                        <!-- </template>
                        </el-popover> -->
                        <!-- 表格筛选列 -->
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
                    <el-table v-loading="state.loading" :key="reload" :data="state.tableData.data" row-key="id" border
                        default-expand-all>
                        <el-table-column prop="id" label="编码" width="40" fixed type="selection"  :selectable="selectable"/>
                        <el-table-column prop="name" label="名称" width="100" v-if="colData[1].istrue" fixed>
                            <template #default="{ row }">
                                <el-link :type="nameButton(row.deleted)" :underline="false" @click="handleClick(row.name)"
                                    style="font-weight: 600;">
                                    {{ row.name }}
                                </el-link>
                            </template>
                        </el-table-column>
                        <el-table-column prop="site_id" label="站点" width="100" v-if="colData[2].istrue">
                            <template #default="{ row }">
                                <el-link type="primary" :underline="false" @click="handle2Site(row.site_id)"
                                    style="font-weight: 600;">
                                    {{ row.site_id === null ? '-----' : row.site_id }}</el-link>
                            </template>
                        </el-table-column>
                        <el-table-column prop="slug" label="标识符" width="100" v-if="colData[3].istrue" />
                        <el-table-column prop="status" label="状态" width="100" v-if="colData[4].istrue" >
                            <template #default="scope">
                                <el-tag :type="scope.row.status === 'abandon' ? 'danger' : 'success'"
                                    disable-transitions>{{ scope.row.status }}
                                </el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="tenant_id" label="租户" width="100" v-if="colData[5].istrue">
                            <template #default="{ row }">
                                <el-link type="primary" :underline="false" @click="handle2Tenant(row.tenant_id)"
                                    style="font-weight: 600;">{{ row.tenant_id === null ? '-----' : row.tenant_id
                                    }}</el-link>
                            </template>
                        </el-table-column>
                        <el-table-column prop="tree_id" label="tree_id" width="100" v-if="colData[6].istrue">
                            <!-- <template #header>
                                <span>U高度
                                    <el-tooltip content="机架单元高度（垂直）,单位：U（1U=1.75英寸）,例如：42U表示标准机柜高度为42个机架单元"
                                        placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template> -->
                        </el-table-column>
                        <el-table-column prop="parent_id" label="上级" width="150" v-if="colData[7].istrue">
                            <!-- <template #header>
                                <span>起始U位
                                    <el-tooltip content="标识设备在机柜中安装的起始U位置（从下往上计数）。例如：设备从第5U开始安装，占用u_height=2，则占据5U-6U。"
                                        placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template> -->
                            <template #default="{ row }">
                                <el-link type="primary" :underline="false" @click="handleClick(row.name)"
                                    style="font-weight: 600;">
                                    {{ row.parent_id === null ? '-----' : row.parent_id }}</el-link>
                            </template>
                        </el-table-column>
                        <el-table-column prop="lft" label="lft" width="100" v-if="colData[8].istrue">
                            <!-- <template #header>
                                <span>安装深度
                                    <el-tooltip content="设备在机柜中的安装深度（如前后导轨间距），单位需与外部尺寸单位一致。" placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template> -->
                        </el-table-column>
                        <el-table-column prop="rght" label="rght" width="100" v-if="colData[9].istrue">
                            <!-- <template #header>
                                <span>外部深度
                                    <el-tooltip content="机柜整体的物理宽度（如标准19英寸机柜宽度通常为600mm）。" placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template> -->
                        </el-table-column>
                        <el-table-column prop="level" label="级别" width="100" v-if="colData[10].istrue" />
                        <el-table-column prop="facility" label="设备" width="100" v-if="colData[11].istrue" />
                        <!-- <el-table-column label="站点" align="center" prop="site_id" width="100" v-if="colData[17].istrue">
                            <template #default="{ row }">
                                <el-button size="small" type="info" @click="handleClick(row.site_id)">
                                    站点
                                    {{ Array.isArray(row.employees) ? row.employees.length : 0  }}
                                </el-button>
                            </template>
                        </el-table-column> -->
                        <el-table-column prop="custom_field_data" label="自定义配置数据" width="150"
                            v-if="colData[12].istrue" />
                        <el-table-column prop="created" label="添加时间" width="180" v-if="colData[13].istrue">
                            <template #default="{ row }">
                                {{ dayjs(row.created).format('YYYY-MM-DD HH:mm:ss') }}
                            </template>
                        </el-table-column>
                        <el-table-column prop="last_updated" label="更新时间" width="180" v-if="colData[14].istrue">
                            <template #default="{ row }">{{ dayjs(row.last_updated).format('YYYY-MM-DD HH:mm:ss')
                                }}</template>
                        </el-table-column>
                        <el-table-column prop="deleted" label="归档时间" width="180" v-if="colData[15].istrue">
                            <template #default="{ row }">{{ dayjs(row.deleted).format('YYYY-MM-DD HH:mm:ss')
                                }}</template>
                        </el-table-column>
                        <!-- <el-table-column fixed label="操作" align="center" class-name="small-padding fixed-width">
                            <template #default="scope">--!
                                <el-popover placement="left">
                                    <template #reference>
                                        <el-button text type="primary" @click="handleUpdate(scope.row)" round><el-icon
                                                style="margin-right: 3px;">
                                                <EditPen />
                                            </el-icon>编辑</el-button>
                                    </template>
                                    <div>
                                        <el-button text type="primary" v-auth="'system:organization:edit'"
                                            @click="onOpenEditModule(scope.row)">
                                            <SvgIcon name="elementEdit" />
                                            修改
                                        </el-button>
                                    </div>
                                    <div>
                                        <el-button text type="primary" v-auth="'system:organization:add'"
                                            @click="onOpenAddModule(scope.row)">
                                            <SvgIcon name="elementPlus" />
                                            新增
                                        </el-button>
                                    </div>
                                    <div>
                                        <el-button v-if="scope.row.parentId != 0" text type="primary"
                                            v-auth="'system:organization:delete'" @click="onTabelRowDel(scope.row)">
                                            <SvgIcon name="elementDelete" />
                                            删除
                                        </el-button>
                                    </div>
                                </el-popover>
                            </template>
                        </el-table-column> -->
                    </el-table>
                </el-card>
                <!-- 分页组件 -->
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
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { ElMessage, type ComponentSize, dayjs } from 'element-plus'

// import List from './component/list.vue';
// import Info from './component/info.vue';
// import Reserve from './component/reserve.vue';
// import Role from './component/role.vue';
// import Editt from './component/edit.vue';
const selectable=ref()

let currentView = ref<'List' | 'Info' | '3D' | 'Reserve' | 'Role' | 'Editt'>('List');
import router from '@/router';
import { listLocations } from '@/api/dicm/site';

const editFormRef = ref() //跳转
var reload = ref()
var search = ref()
// colData中列出表格中的每一列，默认都展示
const colData = reactive([
    { title: "编码", istrue: true, value: "id" },
    { title: "名称", istrue: true, value: "name" },
    { title: "站点", istrue: true, value: "site_id" },
    { title: "标识符", istrue: true, value: "slug" },
    { title: "状态", istrue: true, value: "status" },
    { title: "租户", istrue: true, value: "tenant_id" },
    { title: "tree_id", istrue: true, value: "tree_id" },
    { title: "上级", istrue: true, value: "parent_id" },
    { title: "lft", istrue: true, value: "lft" },
    { title: "rght", istrue: true, value: "rght" },
    { title: "级别", istrue: true, value: "level" },
    { title: "设备", istrue: true, value: "facility" },
    { title: "自定义配置数据", istrue: true, value: "custom_field_data" },
    { title: "创建时间", istrue: false, value: "created" },
    { title: "更新时间", istrue: false, value: "last_updated" },
    { title: "归档时间", istrue: false, value: "deleted" },
    // { title: "状态", istrue: true, value: "status" },
    // { title: "描述", istrue: true, value: "description" },
    // { title: "站点", istrue: true, value: "site_id" },
    // { title: "设备", istrue: true, value: "facility_id" },
    // { title: "租户", istrue: true, value: "tenant_id" },
    // { title: "设备重量", istrue: true, value: "weight" },
    // { title: "物理位置", istrue: true, value: "location_id" },
    // { title: "最大承重限制", istrue: true, value: "max_weight" },
    // { title: "标签", istrue: true, value: "asset_tag" },
    // { title: "机柜用途", istrue: true, value: "role_id" },
    // { title: "自定义配置数据", istrue: true, value: "custom_field_data" },
    // { title: "评价", istrue: true, value: "comments" },
    // { title: "添加时间", istrue: false, value: "created" },
    // { title: "更新时间", istrue: false, value: "last_updated" },
])
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
    empTotal: undefined,
    // 查询参数
    queryParams: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
        status: undefined,
        siteId: undefined,
        id: undefined
    },
    // 携带参数跳转到编辑
    rowData: {},
});

const locationList = () => {
    state.loading = true;
    listLocations(state.queryParams).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        console.log("列表数据：", response.data);
        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
        // console.log("机柜列表数据：", state.tableData.data);
        state.loading = false;
    })
}

const isValidTime = (time: any): boolean => {
    // 判断是否为 null 或 undefined
    if (time == null) return false

    // 转换为Date对象
    let date: Date
    try {
        date = new Date(time)
    } catch (e) {
        return false
    }

    // 检查是否为无效的Date对象
    if (isNaN(date.getTime())) return false

    // 定义各种数据库可能的零值时间格式
    const zeroTimePatterns = [
        '0001-01-01T00:00:00Z',        // UTC零值
        '0001-01-01T00:00:00.000Z',    // UTC零值带毫秒
        '0001-01-01T08:05:43+08:05',   // 带时区的零值（中国时区）
        '0001-01-01T00:00:00+00:00',   // 带时区的零值
        '0001-01-01',                  // 只有日期的零值
    ]

    // 检查是否是任何形式的零值时间
    const isZeroTime = zeroTimePatterns.some(pattern => {
        const zeroDate = new Date(pattern)
        return !isNaN(zeroDate.getTime()) && date.getTime() === zeroDate.getTime()
    })

    return !isZeroTime
}
//=ref<'primary' | 'success' |'info' |'warning'| 'danger'>('primary')
const nameButton = (data: any) => {
    return isValidTime(data) ? 'danger' : 'success'
}

const handle2Site = (id: any) => {
    router.push({ name: '/dicm/sites', query: { id: id } })
}
const handle2Tenant = (id: any) => {
    router.push({ name: '/tenant', query: { id: id } })
}


// 视图对象跳转
const handleClick = (id: any) => {
    //   console.log('组织点击跳转路由携带参数：', id);
    // /system/user/list
    // router.push({ path: '/system/user/list', query: { id: id } })
    //   router.push({name:'/system/user', query: { id: id }}) //http://192.168.0.5:7789/#/system/organization?department_id=2
}

// 点击机柜名称跳转机柜详情
// const handleClick = (name: any) => {
//     router.push({ path: `/dicm/racks/${encodeURIComponent(name)}` })
// }

// 添加数据
const handleAdd = () => {
    state.title="添加站点实例"

}


// 分页改变 size
const onHandleSizeChange = (val: number) => {
    state.queryParams.pageSize = val;
    locationList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
    state.queryParams.pageNum = val;
    locationList();
};

// 页面加载时
onMounted(() => {
    // 查询列表
    locationList();
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