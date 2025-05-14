<template>
    <div class="app-container">
        <div style="margin: 3px;">
            <el-button-group>
                <el-button type="primary" @click="currentView = 'null'" round><el-icon style="margin-right: 3px;">
                        <Menu />
                    </el-icon>查看列表</el-button>
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
            <Edit v-else-if="currentView === 'Edit'" ref="editFormRef" :title="state.title"></Edit>
            <div v-else>
                <el-card>
                    <!-- 筛选按钮 -->
                    <div class="row" style="display: flex; justify-content: flex-end; margin-bottom: 3px;">
                        <!-- <el-popover placement="left-start" title="搜索" :width="400">
                            <template #reference> -->
                                <el-input v-model="search" style="width: 30%;margin-right: 10px;" placeholder="请输入搜索数据" clearable>
                                    <template #prepend>
                                        <el-select v-model="colData" placeholder="搜  索" style="width: 80px">
                                            <el-option v-for="item in colData" :key="item.title"
                                                :label="item.title" :value="item.title" />
                                        </el-select>
                                    </template>
                                    <template #append>
                                        <el-button @click="SearchTable"><el-icon>
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
                                <el-checkbox v-for="col in colData" :key="col.title" v-model="col.istrue"
                                    :label="col.title" style="display: block; margin-bottom: 5px;"></el-checkbox>
                            </div>
                        </el-popover>
                    </div>
                    <!--数据表格-->
                    <el-table v-loading="state.loading" :key="reload" :data="state.tableData.data" row-key="id" border
                        default-expand-all>
                        <el-table-column v-if="colData[0].istrue" prop="id" label="编码" width="100" fixed
                            key="Math.random()" />
                        <el-table-column v-if="colData[1].istrue" prop="name" label="名称" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[2].istrue" prop="_name" label="别名" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[3].istrue" prop="slug" label="短标识符" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[4].istrue" prop="physical_address" label="物理地址" width="180"
                            key="Math.random()" />
                        <el-table-column v-if="colData[5].istrue" prop="custom_field_data" label="自定义配置数据" width="180"
                            key="Math.random()" />
                        <el-table-column v-if="colData[6].istrue" prop="status" label="状态" width="100"
                            key="Math.random()">
                            <!-- <template #default="scope">
            <el-tag :type="scope.row.active === 'online' ? 'success' : 'danger'" disable-transitions>{{ scope.row.active ?
              "在线" : "已停用" }}
            </el-tag>
          </template> -->
                        </el-table-column>
                        <el-table-column v-if="colData[7].istrue" prop="group_id" label="站点组" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[8].istrue" prop="latitude" label="GPS坐标" width="100"
                            key="Math.random()">
                            <template #default="scope">
                                {{ [Number(scope.row.latitude), Number(scope.row.longitude)] }}
                            </template>
                        </el-table-column>
                        <el-table-column v-if="colData[9].istrue" prop="shipping_address" label="物流地址" width="180"
                            key="Math.random()" />
                        <el-table-column v-if="colData[10].istrue" prop="tenant_id" label="租户" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[11].istrue" prop="time_zone" label="时区" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[12].istrue" prop="region_id" label="地区" width="100"
                            key="Math.random()" />
                        <el-table-column v-if="colData[13].istrue" label="创建时间" align="center" prop="created"
                            key="Math.random()" width="180" />
                        <el-table-column v-if="colData[14].istrue" label="更新时间" align="center" prop="last_updated"
                            key="Math.random()" width="180">
                            <template #default="{ row }">
                                <el-button size="small" type="info" @click="handleClick(row.site_id)">
                                    站点
                                    <!-- {{ Array.isArray(row.employees) ? row.employees.length : 0  }} -->
                                </el-button>
                            </template>
                        </el-table-column>
                        <el-table-column v-if="colData[15].istrue" label="操作" align="center" key="Math.random()"
                            class-name="small-padding fixed-width">
                            <template #default="scope">
            <el-popover placement="left">
              <template #reference>
                <el-button type="primary" circle>
                  <SvgIcon name="elementStar" />
                </el-button>
              </template>
              <div>
                <!-- <el-button text type="primary" v-auth="'system:organization:edit'" @click="onOpenEditModule(scope.row)">
                  <SvgIcon name="elementEdit" />
                  修改
                </el-button> -->
                <el-button type="primary" @click="currentView = 'Edit'" round><el-icon style="margin-right: 3px;">
                        <EditPen />
                    </el-icon>编辑</el-button>
              </div>
              <!-- <div>
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
            </el-popover>
          </template>
                        </el-table-column>
                    </el-table>
                </el-card>
            </div>
        </div>
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
<script setup lang="ts">
import { ref } from 'vue';
import Location from './component/location.vue';
import Region from './component/region.vue';
import Edit from './component/edit.vue';

const currentView = ref<'Location' | 'Edit' | 'Region' | null>(null);

import { onMounted, reactive } from 'vue';
import { listSites } from '@/api/dicm/site';

const editFormRef = ref()
var reload = ref()
var search = ref()
// colData中列出表格中的每一列，默认都展示
const colData = reactive([
    { title: "编码", istrue: true, value:"id"},
    { title: "名称", istrue: true, value:"name"},
    { title: "别名", istrue: true, value:"_name"},
    { title: "序列", istrue: true, value:"slug"},
    { title: "物理地址", istrue: true, value:"physical_address"},
    { title: "自定义配置数据", istrue: true, value:"custom_field_data"},
    { title: "状态", istrue: true, value:"status"},
    { title: "站点组", istrue: true, value:"group_id"},
    { title: "GPS坐标", istrue: true , value:"latitude"},
    { title: "物流地址", istrue: true , value:"shipping_address"},
    { title: "租户", istrue: true, value:"tenant_id"},
    { title: "时区", istrue: true, value:"time_zone"},
    { title: "地区", istrue: true, value:"region_id"},
    { title: "创建时间", istrue: true, value:"created" },
    { title: "更新时间", istrue: false, value:"last_updated" },
    { title: "操作", istrue: true, value:"" },
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

// 监听checkedColumns的变化，当checkedColumns发生变化时，重新渲染表格
const watchCheckedColumns = () => {
    // 遍历colData，将colData中的istrue属性设置为false
    colData.forEach((item) => {
        item.istrue = false
    })
    // 遍历checkedColumns，将checkedColumns中的值在colData中找到对应的列，将istrue属性设置为true
    checkedColumns.value.forEach((item) => {
        colData.forEach((col) => {
            if (item === col.title) {
                col.istrue = true
            }
        })
    })
    // 重新渲染表格
    reload.value = Math.random()
}

const rackList = () => {
    state.loading = true;
    listSites(state.queryParams).then((response: any) => {
        if (response.code != 200) {
            state.loading = false;
        }
        console.log("站点列表数据：", response.data);
        state.tableData.data = response.data.data;
        state.tableData.total = response.data.total;
        // console.log("机柜列表数据：", state.tableData.data);
        state.loading = false;
    })
}


// 单元格单击
const handleClick = (id: any) => {
    //   console.log('组织点击跳转路由携带参数：', id);
    // /system/user/list
    // router.push({ path: '/system/user/list', query: { id: id } })
    //   router.push({name:'/system/user', query: { id: id }}) //http://192.168.0.5:7789/#/system/organization?department_id=2
}
// 表格上方单元格搜索
const SearchTable = (data: any)=>{
    console.log("单元格搜索", data);
    
}

/** 新增按钮操作 */
const handleAdd = () => {
  state.title = "添加";
  editFormRef.value.openDialog({});
};
/** 修改按钮操作 */
const handleUpdate = (row: any) => {
  state.title = "修改";
  editFormRef.value.openDialog(row);
};

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
  max-height: 300px;  /* 控制最大高度 */
  width: 400px;
  overflow-y: auto;   /* 垂直滚动 */
  padding-right: 8px; /* 避免滚动条遮挡内容 */
}
</style>