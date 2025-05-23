<template>
    <div class="app-container">
        <div style="margin: 3px;" v-if="currentView !== 'Edit'">
            <el-button-group>
                <el-button type="primary" @click="currentView = null" round><el-icon style="margin-right: 3px;">
                        <Menu />
                    </el-icon>查看列表</el-button>
                <el-button type="primary" @click="currentView = 'Group'" round><el-icon style="margin-right: 3px;">
                        <Iphone />
                    </el-icon>
                    站点组
                </el-button>
                <el-button type="primary" @click="currentView = 'Location'" round><el-icon style="margin-right: 3px;">
                        <LocationInformation />
                    </el-icon>物理位置</el-button>
                <el-button type="primary" @click="currentView = 'Region'" round><el-icon style="margin-right: 3px;">
                        <OfficeBuilding />
                    </el-icon> 地区</el-button>

            </el-button-group>
        </div>
        <div style="margin: 10px;">
            <Region v-if="currentView === 'Region'"></Region>
            <Location v-else-if="currentView === 'Location'"></Location>
            <Edit v-else-if="currentView === 'Edit'" :item="null" :title="state.title"></Edit>
            <Group v-else-if="currentView === 'Group'"></Group>
            <div v-else>
                <el-card>
                    <!-- 筛选按钮 -->
                    <div class="row" style="display: flex; justify-content: flex-end; margin-bottom: 3px;">
                        <!-- <el-popover placement="left-start" title="搜索" :width="400">
                            <template #reference> -->
                        <el-input class="search_input" v-model="searchContent" style="width: 40%;margin-right: 10px;"
                            placeholder="请输入搜索数据" clearable>
                            <template #prepend>
                                <el-select v-model="searchFiled" placeholder="搜  索" style="width: 100px">
                                    <el-option v-for="item in colData" :v-if="item.isSearch" :key="item.title" :label="item.title"
                                        :value="item.value" />
                                </el-select>
                            </template>
                            <template #append>
                                <el-button @click="SearchTable()"><el-icon>
                                        <Search />
                                    </el-icon></el-button>
                            </template></el-input>
                        <!-- </template>
                        </el-popover> -->
                        <!-- 新建按钮 -->
                        <el-button type="primary" @click="handleAdd()"><el-icon style="margin-right: 3px;">
                                <EditPen />
                            </el-icon>添加实例</el-button>
                        <!-- 表格筛选列 -->
                        <el-popover placement="right-start" title="筛选列" :width="40" trigger="click">
                            <template #reference>
                                <el-button><el-icon>
                                        <Grid />
                                    </el-icon></el-button>
                            </template>
                            <div class="scrollable-checkbox-list">
                                <el-checkbox v-for="col in colData" :key="col.title" v-model="col.istrue"
                                    :label="col.title" style="display: block; margin-bottom: 5px;"></el-checkbox>
                            </div>
                        </el-popover>
                    </div>
                    <!--数据表格-->
                    <el-table v-loading="state.loading" :key="reload" :data="state.tableData.data" row-key="id" border
                        default-expand-all>
                        <el-table-column v-if="colData[0].istrue" prop="id" label="编码" width="40" type="selection"
                            :selectable="selectable" fixed />
                        <el-table-column v-if="colData[1].istrue" prop="name" label="名称" width="100">
                            <template #default="{ row }">
                                <el-link :type="nameButton(row.deleted)" :underline="false"
                                    @click="handle2Info(row.name)" style="font-weight: 600;">
                                    {{ row.name }}
                                </el-link>
                            </template>
                        </el-table-column>
                        <el-table-column v-if="colData[2].istrue" prop="_name" label="别名" width="100" />
                        <el-table-column v-if="colData[3].istrue" prop="slug" label="短标识符" width="100" />
                        <el-table-column v-if="colData[4].istrue" prop="physical_address" label="物理地址" width="180" />
                        < v-if="colData[5].istrue" prop="custom_field_data" label="自定义配置数据" width="180" />
                        <el-table-column v-if="colData[6].istrue" prop="status" label="状态" width="100">
                            <template #default="scope">
                                <el-tag :type="scope.row.status === 'abandon' ? 'danger' : 'success'"
                                    disable-transitions>{{ scope.row.status }}
                                </el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column v-if="colData[7].istrue" prop="group_id" label="站点组" width="100">
                            <template #default="{ row }">
                                <el-link type="primary" :underline="false" @click="handleClick(row.group_id)"
                                    style="font-weight: 600;">
                                    {{ row.group_id }}
                                </el-link>
                            </template>
                        </el-table-column>
                        <el-table-column v-if="colData[8].istrue" prop="latitude" label="GPS坐标" width="100">
                            <template #default="scope">
                                {{ [Number(scope.row.latitude), Number(scope.row.longitude)] }}
                            </template>
                        </el-table-column>
                        <el-table-column v-if="colData[9].istrue" prop="shipping_address" label="物流地址" width="180" />
                        <el-table-column v-if="colData[10].istrue" prop="tenant_id" label="租户" width="100" />
                        <el-table-column v-if="colData[11].istrue" prop="time_zone" label="时区" width="100" />
                        <el-table-column v-if="colData[12].istrue" prop="region_id" label="地区" width="100" />
                        <el-table-column v-if="colData[13].istrue" label="创建时间" align="center" prop="created"
                            width="180" />
                        <el-table-column v-if="colData[14].istrue" label="更新时间" align="center" prop="last_updated"
                            width="180">
                        </el-table-column>
                        <el-table-column v-if="colData[16].istrue" label="已归档" align="center" prop="deleted"
                            width="180">
                        </el-table-column>
                        <el-table-column v-if="colData[15].istrue" label="操作" align="center"
                            class-name="small-padding fixed-width">
                            <template #default="scope">
                                <el-popover placement="left">
                                    <template #reference>
                                        <el-button type="primary" circle>
                                            <SvgIcon name="elementStar" />
                                        </el-button>
                                    </template>
                                    <!-- <div>
                                        <el-button text type="primary" v-auth="'system:organization:edit'" @click="onOpenEditModule(scope.row)">
                  <SvgIcon name="elementEdit" />
                  修改
                </el-button>
                                        <el-button type="primary" @click="currentView = 'Edit'" round><el-icon
                                                style="margin-right: 3px;">
                                                <EditPen />
                                            </el-icon>编辑</el-button>
                                    </div>
                                    <div>
                <el-button text type="primary" v-auth="'system:organization:add'" @click="onOpenAddModule(scope.row)">
                  <SvgIcon name="elementPlus" />
                  新增
                </el-button>
              </div> -->
                                    <div>
                                        <el-button v-if="scope.row.parentId != 0" text type="primary"
                                            v-auth="'system:organization:delete'" @click="handelDelete(scope.row)">
                                            <SvgIcon name="elementDelete" />
                                            删除
                                        </el-button>
                                    </div>
                                </el-popover>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-card>
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
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import Location from './component/location.vue';
import Region from './component/region.vue';
import Edit from './component/edit.vue';
import Group from './component/group.vue'
import router from '@/router';

let currentView = ref<'Location' | 'Edit' | 'Region' | 'Group' | null>(null);

import { onMounted, reactive } from 'vue';
import { deleteSite, listSites } from '@/api/dicm/site';
import { ElMessage } from 'element-plus';
const selectable = ref()

const editForm = ref()
var reload = ref()
var searchFiled = ref() //顶端搜索字段
var searchContent = ref() //顶端搜索内容
// 筛选列、搜索列
const colData = reactive([
    { title: "编码", istrue: true, value: "id",isSearch: true },
    { title: "名称", istrue: true, value: "name",isSearch: true },
    { title: "别名", istrue: true, value: "_name",isSearch: true },
    { title: "标识符", istrue: true, value: "slug",isSearch: true },
    { title: "物理地址", istrue: true, value: "physical_address",isSearch: true },
    { title: "自定义配置数据", istrue: true, value: "custom_field_data",isSearch: true },
    { title: "状态", istrue: true, value: "status",isSearch: true },
    { title: "站点组", istrue: true, value: "group_id",isSearch: true },
    { title: "GPS坐标", istrue: true, value: "latitude",isSearch: false },
    { title: "物流地址", istrue: true, value: "shipping_address",isSearch: true },
    { title: "租户", istrue: true, value: "tenant_id",isSearch: true },
    { title: "时区", istrue: true, value: "time_zone",isSearch: true },
    { title: "地区", istrue: true, value: "region_id",isSearch: true },
    { title: "创建时间", istrue: true, value: "created",isSearch: false },
    { title: "更新时间", istrue: false, value: "last_updated",isSearch: false },
    { title: "操作", istrue: true, value: "",isSearch: false },
    { title: "归档时间", istrue: false, value: "deleted",isSearch: false },
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
});

const rackList = () => {
    state.loading = true;
    listSites(state.queryParams).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        // console.log("站点列表：", response.data);

        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
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
//<'primary' | 'success' |'info' |'warning'| 'danger'>('primary')
const nameButton = (data: any) => {
    return isValidTime(data) ? 'danger' : 'success'
}

// 跳转到详情页
const handle2Info = (name: any) => {
    router.push({ path: `/dicm/sites/${encodeURIComponent(name)}` })
}
// 跳转到站点组
const handleClick = (id: any) => {
    currentView.value = 'Group';
}
// 表格上方单元格搜索
const SearchTable = () => {
    // console.log("搜索内容：", searchFiled.value, searchContent.value);
    const data = {
        [searchFiled.value]: searchContent.value.trim()
    };
    console.log(data);

    listSites(data).then((res: any) => {
        if (res.code != 200) {
            state.loading = false;
        }
        state.tableData.data = res.data.data;
        state.tableData.total = res.data.total;
        state.loading = false;
    })

}

/** 新增按钮操作 */
const handleAdd = () => {
    state.title = "添加";
    currentView.value = 'Edit';
};
/** 修改按钮操作 */
const handleUpdate = (row: any) => {
    state.title = "修改";
    currentView.value = 'Edit';
};
/** 删除or归档按钮操作 */
const handelDelete = (row: any) => {
    console.log("删除：", row);
    if (row.deleted !== "" || row.deleted !== undefined || row.deleted !== "0001-01-01T00:00:00Z") {
        ElMessage.warning("【" + row.name + "】已归档, 请勿重复操作");
    } else {
        deleteSite(row.id).then((res: any) => {
            if (res.code == 200) {
                ElMessage.success("站点【" + row.name + "】已归档");
            }
        })
    }
    rackList();
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
.search_input {
    position: absolute;
    width: 70%;
    /* 输入框绝对定位 */
    /* top: -50px;         调整到表格上方 */
    left: 0;
    /* 对齐表格左侧 */
    margin-left: 45px;
}

.scrollable-checkbox-list {
    max-height: 300px;
    /* 控制最大高度 */
    width: 400px;
    overflow-y: auto;
    /* 垂直滚动 */
    padding-right: 8px;
    /* 避免滚动条遮挡内容 */
}
</style>