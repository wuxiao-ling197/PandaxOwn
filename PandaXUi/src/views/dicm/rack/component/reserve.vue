<template>
  <div class="app-container">
    <div>
      <el-card>
        <!-- 表格顶部搜索筛选按钮 -->
        <div class="row" style="display: flex; justify-content: flex-end; margin-bottom: 3px;">
          <el-input class="search_input" v-model="search" style="width: 30%;margin-right: 10px;" placeholder="请输入搜索数据"
            clearable>
            <template #prepend>
              <el-select v-model="colData" placeholder="搜  索" style="width: 80px">
                <el-option v-for="item in colData" :key="item.title" :label="item.title" :value="item.title" />
              </el-select>
            </template>
            <template #append>
              <!--  搜索按钮 -->
              <el-button><el-icon>
                  <Search />
                </el-icon></el-button>
            </template></el-input>
          <!-- 表格筛选列 -->
          <!-- <el-popover placement="right-start" title="筛选列" :width="40" trigger="click">
                            <template #reference>
                                <el-button><el-icon>
                                        <Grid />
                                    </el-icon></el-button>
                            </template>
                            <div>
                                <el-checkbox v-for="col in colData" :key="col.title" v-model="col.istrue"
                                    :label="col.title" style="display: block; margin-bottom: 5px;"></el-checkbox>
                            </div>
                        </el-popover> -->
        </div>
        <!--数据表格-->
        <el-table v-loading="state.loading" :data="state.tableData.data" row-key="id" border default-expand-all>
          <!-- <el-table-column prop="id" label="编码" width="100" fixed /> -->
          <el-table-column prop="units" label="单位" width="100" />
          <el-table-column prop="rack_id" label="相关机柜" width="100" />
          <el-table-column prop="tenant_id" label="租户" width="100" />
          <el-table-column prop="user_id" label="用户" width="100" />
          <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" />
          <el-table-column prop="description" label="描述" width="150" />
          <el-table-column prop="comments" label="评价" width="150" />
          <el-table-column prop="created" label="添加" width="180" />
          <el-table-column prop="last_updated" label="更新" width="180" />
          <!-- <template #default="scope">
            <el-tag :type="scope.row.active === 'online' ? 'success' : 'danger'" disable-transitions>{{ scope.row.active ?
              "在线" : "已停用" }}
            </el-tag>
          </template> -->
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <!-- <template #default="scope">
            <el-popover placement="left">
              <template #reference>
                <el-button type="primary" circle>
                  <SvgIcon name="elementStar" />
                </el-button>
              </template>
              <div>
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
              </div>
            </el-popover>
          </template> -->
          </el-table-column>
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

</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { listRackReserve } from '@/api/dicm/rack';

var search = ref()
// colData中列出表格中的每一列，默认都展示
const colData = reactive([
  { title: "编码", istrue: true, value: "id" },
  { title: "租户", istrue: true, value: "tenant_id" },
  { title: "相关机柜", istrue: true, value: "rack_id" },
  { title: "单位", istrue: true, value: "units" },
  { title: "自定义配置数据", istrue: true, value: "custom_field_data" },
  { title: "评价", istrue: true, value: "comments" },
  { title: "用户", istrue: true, value: "user_id" },
  { title: "描述", istrue: true, value: "description" },
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

const rackreserveList = () => {
  state.loading = true;
  listRackReserve(state.queryParams).then((response: any) => {
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

// 分页改变 size
const onHandleSizeChange = (val: number) => {
  state.queryParams.pageSize = val;
  rackreserveList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
  state.queryParams.pageNum = val;
  rackreserveList();
};

// 页面加载时
onMounted(() => {
  // 查询机柜列表
  rackreserveList();

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