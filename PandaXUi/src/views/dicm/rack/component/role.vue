<template>
  <div class="app-container">
    <div>
      <el-card>
        <!-- 表格顶部搜索筛选按钮 -->
        <div class="row" style="display: flex; justify-content: flex-end; margin-bottom: 3px;">
          <el-input class="search_input" v-model="searchContent" placeholder="请输入搜索数据" clearable>
            <template #prepend>
              <el-select v-model="searchFiled" placeholder="搜  索" style="width: 80px">
                <el-option v-for="item in colData" :v-if="item.isSearch" :key="item.title" :label="item.title"
                  :value="item.title" />
              </el-select>
            </template>
            <template #append>
              <!--  搜索按钮 -->
              <el-button @click="SearchTable()"><el-icon>
                  <Search />
                </el-icon></el-button>
            </template></el-input>
          <!-- 表格筛选列 -->
          <el-popover placement="right-start" title="筛选列" :width="40" trigger="click">
            <template #reference>
              <el-button><el-icon>
                  <Grid />
                </el-icon></el-button>
            </template>
            <div>
              <el-checkbox v-for="col in colData" :key="col.title" v-model="col.istrue" :label="col.title"
                style="display: block; margin-bottom: 5px;"></el-checkbox>
            </div>
          </el-popover>
        </div>
        <!--数据表格-->
        <el-table v-loading="state.loading" :data="state.tableData.data" row-key="id" border default-expand-all>
          <el-table-column prop="id" label="编码" width="40" fixed type="selection" :selectable="selectable" />
          <el-table-column prop="name" label="名称" width="100" v-if="colData[1].istrue" />
          <el-table-column prop="slug" label="短标识符" width="100" v-if="colData[2].istrue" />
          <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" v-if="colData[3].istrue" />
          <el-table-column prop="color" label="颜色标签" width="100" v-if="colData[4].istrue" />
          <el-table-column prop="description" label="描述" width="150" v-if="colData[5].istrue" />
          <el-table-column prop="created" label="添加" width="180" v-if="colData[6].istrue" />
          <el-table-column prop="last_updated" label="更新" width="180" v-if="colData[7].istrue" />
          <el-table-column prop="deleted" label="归档" width="180" v-if="colData[8].istrue" />
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
import { listRackRole } from '@/api/dicm/rack';

const selectable = ref()//表格多选
var searchFiled = ref() //顶端搜索字段
var searchContent = ref() //顶端搜索内容
// colData中列出表格中的每一列，默认都展示
const colData = reactive([
  { title: "编码", istrue: true, value: "id", isSearch: true },
  { title: "名称", istrue: true, value: "name", isSearch: true },
  { title: "短标识符", istrue: true, value: "slug", isSearch: true },
  { title: "自定义配置数据", istrue: true, value: "custom_field_data", isSearch: true },
  { title: "颜色", istrue: true, value: "color", isSearch: true },
  { title: "描述", istrue: true, value: "description", isSearch: true },
  { title: "添加时间", istrue: false, value: "created", isSearch: false },
  { title: "更新时间", istrue: false, value: "last_updated", isSearch: false },
  { title: "归档时间", istrue: false, value: "deleted", isSearch: false },
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

const rackroleList = () => {
  state.loading = true;
  listRackRole(state.queryParams).then((response: any) => {
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

// 表格上方单元格搜索
const SearchTable = () => {
  // console.log("搜索内容：", searchFiled.value, searchContent.value);
  const data = {
    [searchFiled.value]: searchContent.value.trim()
  };
  listRackRole(data).then((res: any) => {
    if (res.code != 200) {
      state.loading = false;
    }
    state.tableData.data = res.data.data;
    state.tableData.total = res.data.total;
    state.loading = false;
  })
}
// 分页改变 size
const onHandleSizeChange = (val: number) => {
  state.queryParams.pageSize = val;
  rackroleList();
};
// 分页改变 page
const onHandleCurrentChange = (val: number) => {
  state.queryParams.pageNum = val;
  rackroleList();
};

// 页面加载时
onMounted(() => {
  // 查询机柜列表
  rackroleList();

})
</script>
<style lang="css" scoped>
.search_input {
  position: absolute;
  width: 60%;
  left: 0;
  margin-left: 45px;
}
</style>