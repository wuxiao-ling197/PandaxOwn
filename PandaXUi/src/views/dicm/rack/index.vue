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
                        <!-- <el-table-column prop="id" label="编码" width="100" fixed v-if="colData[0].istrue"/> -->
                        <el-table-column prop="name" label="名称" width="100" v-if="colData[1].istrue" fixed>
                            <template #default="{ row }">
                                <el-link type="primary" :underline="false" @click="handleClick(row.name)"
                                    style="font-weight: 600;">
                                    {{ row.name }}
                                </el-link>
                            </template>
                        </el-table-column>
                        <el-table-column prop="_name" label="别名" width="100" v-if="colData[2].istrue" />
                        <el-table-column prop="type" label="机架类型" width="100" v-if="colData[3].istrue" />
                        <el-table-column prop="serial" label="序列" width="100" v-if="colData[4].istrue" />
                        <el-table-column prop="width" label="内部宽度" width="100" v-if="colData[5].istrue" />
                        <el-table-column prop="u_height" label="U高度" width="100" v-if="colData[6].istrue">
                            <template #header>
                                <span>U高度
                                    <el-tooltip content="机架单元高度（垂直）,单位：U（1U=1.75英寸）,例如：42U表示标准机柜高度为42个机架单元"
                                        placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="starting_unit" label="起始U位" width="150" v-if="colData[7].istrue">
                            <template #header>
                                <span>起始U位
                                    <el-tooltip content="标识设备在机柜中安装的起始U位置（从下往上计数）。例如：设备从第5U开始安装，占用u_height=2，则占据5U-6U。"
                                        placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="mounting_depth" label="安装深度" width="100" v-if="colData[8].istrue">
                            <template #header>
                                <span>安装深度
                                    <el-tooltip content="设备在机柜中的安装深度（如前后导轨间距），单位需与外部尺寸单位一致。" placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="outer_depth" label="外部深度" width="100" v-if="colData[9].istrue">
                            <template #header>
                                <span>外部深度
                                    <el-tooltip content="机柜整体的物理宽度（如标准19英寸机柜宽度通常为600mm）。" placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="outer_width" label="外部宽度" width="100" v-if="colData[10].istrue" />
                        <el-table-column prop="weight_unit" label="重量单位" width="100" v-if="colData[11].istrue" />
                        <el-table-column prop="outer_unit" label="外部尺寸单位" width="100" v-if="colData[12].istrue" />
                        <el-table-column prop="desc_units" label="描述性单位" width="150" v-if="colData[13].istrue">
                            <template #header>
                                <span>描述性单位
                                    <el-tooltip content="补充说明机柜 / 设备单位信息，如：自定义标签、备注等。" placement="top">
                                        <el-icon>
                                            <QuestionFilled />
                                        </el-icon>
                                    </el-tooltip>
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="_abs_weight" label="实际总重量" width="150" v-if="colData[14].istrue" />
                        <el-table-column prop="_abs_max_weight" label="理论最大承重" width="100"
                            v-if="colData[15].istrue"></el-table-column>
                        <el-table-column prop="status" label="状态" width="100">
                            <!-- <template #default="scope">
            <el-tag :type="scope.row.active === 'online' ? 'success' : 'danger'" disable-transitions>{{ scope.row.active ?
              "在线" : "已停用" }}
            </el-tag>
          </template> -->
                        </el-table-column>
                        <el-table-column label="描述" align="center" prop="description" width="100"
                            v-if="colData[16].istrue" />
                        <el-table-column label="站点" align="center" prop="site_id" width="100" v-if="colData[17].istrue">
                            <template #default="{ row }">
                                <el-button size="small" type="info" @click="handleClick(row.site_id)">
                                    站点
                                    <!-- {{ Array.isArray(row.employees) ? row.employees.length : 0  }} -->
                                </el-button>
                            </template>
                        </el-table-column>
                        <el-table-column prop="facility_id" label="设备" width="100" v-if="colData[18].istrue" />
                        <el-table-column prop="tenant_id" label="租户" width="100" v-if="colData[19].istrue" />
                        <el-table-column prop="weight" label="设备重量" width="100" v-if="colData[20].istrue" />
                        <el-table-column prop="location_id" label="物理位置" width="100" v-if="colData[21].istrue" />
                        <el-table-column prop="max_weight" label="最大承重限制" width="100" v-if="colData[22].istrue" />
                        <el-table-column prop="asset_tag" label="标签" width="100" v-if="colData[23].istrue" />
                        <el-table-column prop="role_id" label="机柜用途" width="100" v-if="colData[24].istrue" />
                        <el-table-column prop="custom_field_data" label="自定义配置数据" width="150"
                            v-if="colData[25].istrue" />
                        <el-table-column prop="comments" label="评价" width="150" v-if="colData[26].istrue" />
                        <el-table-column prop="created" label="添加时间" width="180" v-if="colData[27].istrue" />
                        <el-table-column prop="last_updated" label="更新时间" width="180" v-if="colData[28].istrue" />
                        <!-- <el-table-column fixed label="操作" align="center" class-name="small-padding fixed-width">
                            <template #default="scope">
                                <!-- <el-popover placement="left"> -->
                        <!-- <template #reference> -->
                        <!-- <el-button text type="primary" @click="handleUpdate(scope.row)" round><el-icon
                                                style="margin-right: 3px;">
                                                <EditPen />
                                            </el-icon>编辑</el-button> -->
                        <!-- </template> -->
                        <!-- <div>
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
              </div> -->
                        <!-- </el-popover> -->
                        <!-- </template>
                        </el-table-column> -->
                    </el-table>
                </el-card>
            </div>
        </div>
        <div v-show="state.tableData.total > 0">
            <el-divider></el-divider>
            <el-pagination background :total="state.tableData.total" :page-sizes="[10, 20, 30, 50, 100]"
                :current-page="state.queryParams.pageNum" :page-size="state.queryParams.pageSize"
                layout="total, sizes, prev, pager, next, jumper" @size-change="onHandleSizeChange"
                @current-change="onHandleCurrentChange" />
        </div>
    </div>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import List from './component/list.vue';
import Info from './component/info.vue';
import Reserve from './component/reserve.vue';
import Role from './component/role.vue'

let currentView = ref<'List' | 'Info' | 'Edit' | '3D' | 'Reserve' | 'Role'>('List');
import { listRacks } from '@/api/dicm/rack';
import router from '@/router';

const editFormRef = ref()
var reload = ref()
var search = ref()
// colData中列出表格中的每一列，默认都展示
const colData = reactive([
    { title: "编码", istrue: true, value: "id" },
    { title: "名称", istrue: true, value: "name" },
    { title: "别名", istrue: true, value: "_name" },
    { title: "类型", istrue: true, value: "type" },
    { title: "序列", istrue: true, value: "serial" },
    { title: "内部宽度", istrue: true, value: "width" },
    { title: "U高度", istrue: true, value: "u_height" },
    { title: "起始U位", istrue: true, value: "starting_unit" },
    { title: "安装深度", istrue: true, value: "mounting_depth" },
    { title: "外部深度", istrue: true, value: "outer_depth" },
    { title: "外部宽度", istrue: true, value: "outer_width" },
    { title: "重量单位", istrue: true, value: "weight_unit" },
    { title: "外部尺寸单位", istrue: true, value: "outer_unit" },
    { title: "描述性单位", istrue: true, value: "desc_units" },
    { title: "实际总重量", istrue: true, value: "_abs_weight" },
    { title: "理论最大承重", istrue: true, value: "_abs_max_weight" },
    { title: "状态", istrue: true, value: "status" },
    { title: "描述", istrue: true, value: "description" },
    { title: "站点", istrue: true, value: "site_id" },
    { title: "设备", istrue: true, value: "facility_id" },
    { title: "租户", istrue: true, value: "tenant_id" },
    { title: "设备重量", istrue: true, value: "weight" },
    { title: "物理位置", istrue: true, value: "location_id" },
    { title: "最大承重限制", istrue: true, value: "max_weight" },
    { title: "标签", istrue: true, value: "asset_tag" },
    { title: "机柜用途", istrue: true, value: "role_id" },
    { title: "自定义配置数据", istrue: true, value: "custom_field_data" },
    { title: "评价", istrue: true, value: "comments" },
    { title: "添加时间", istrue: false, value: "created" },
    { title: "更新时间", istrue: false, value: "last_updated" },
])
// 多选框的列表，列出表格的每一列
const checkBoxGroup = ref(["编码", "名称", "别名", "序列", "物理地址", "自定义配置数据", "状态", "站点组", "GPS坐标", "物流地址", "租户", "时区", "地区", "创建时间", "更新时间"])
// 当前选中的多选框，代表当前展示的列
const checkedColumns = ref(["编码", "名称", "别名", "序列", "物理地址", "自定义配置数据", "状态", "站点组", "GPS坐标", "物流地址", "租户", "时区", "地区", "创建时间", "更新时间"])
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
});

const rackList = () => {
    state.loading = true;
    listRacks(state.queryParams).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        console.log("机柜列表数据：", response.data);
        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
        // console.log("机柜列表数据：", state.tableData.data);
        state.loading = false;
    })
}

// 视图对象跳转
// const handleClick = (id: any) => {
//     //   console.log('组织点击跳转路由携带参数：', id);
//     // /system/user/list
//     // router.push({ path: '/system/user/list', query: { id: id } })
//     //   router.push({name:'/system/user', query: { id: id }}) //http://192.168.0.5:7789/#/system/organization?department_id=2
// }

// 跳转详情
const handleClick = (name: any) => {
    router.push({ path: `/dicm/racks/${encodeURIComponent(name)}` })
}



// 分页改变 size
const onHandleSizeChange = (val: number) => {
    state.queryParams.pageSize = val;
    rackList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
    state.queryParams.pageNum = val;
    rackList();
};

// 页面加载时
onMounted(() => {
    // 查询机柜列表
    rackList();
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