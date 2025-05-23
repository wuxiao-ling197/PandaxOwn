<template>
  <div class="app-container">
    <div v-if="currentView === null">
      <el-card class="view" shadow="always">
        <!-- <template #extra> -->
        <el-button type="info" size="small" @click="handleExit()" style="margin-bottom: 10px;"><el-icon
            style="margin-right: 5px;">
            <ArrowLeftBold />
          </el-icon>
          返回</el-button>
        <!-- </template> -->
        <template #header>
          <div class="card-header" style="font-weight: 600;margin: 3px 5px;text-justify: auto;">
            <span>{{ state.title }}</span>
          </div>
        </template>
        <el-form class="editform" :model="state.data" label-width="auto">
          <h4>基本信息</h4>
          <el-form-item class="item" label="名称" prop="state.data.name" required><el-input v-model="state.data.name"
              placeholder="请输入姓名"></el-input></el-form-item>
          <el-form-item class="item" label="别名" prop="state.data._name"><el-input v-model="state.data._name"
              aria-required="true"></el-input></el-form-item>
          <el-form-item class="item" label="租户" prop="state.data.tenant_id">
            <template #header>
              <span>租户
                <el-tooltip content="站点的管理人员、维护人员。" placement="top">
                  <el-icon>
                    <QuestionFilled />
                  </el-icon>
                </el-tooltip>
              </span>
            </template>
            <el-select v-model="state.data.tenant_id" collapse-tags collapse-tags-tooltip placeholder="选择租户">
              <el-option v-for="item in state.tenantOptions" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item class="item" label="所属站点组" prop="state.data.group_id"><el-select v-model="state.data.group_id"
              collapse-tags collapse-tags-tooltip placeholder="选择站点组">
              <el-option v-for="item in state.groupOptions" :label="item.name" :value="item.id" />
            </el-select></el-form-item>
          <el-form-item class="item" label="时区" prop="state.data.time_zone" required> <el-select
              v-model="state.data.time_zone" collapse-tags collapse-tags-tooltip placeholder="选择时区">
              <el-option v-for="item in state.tzOptions" :label="item.name" :value="item.value" />
            </el-select></el-form-item>
          <el-form-item class="item" label="设备" prop="state.data.facility" required>
            <template #header>
              <span>安装设备
                <el-tooltip content="该机柜中安装的设备。" placement="top">
                  <el-icon>
                    <QuestionFilled />
                  </el-icon>
                </el-tooltip>
              </span>
            </template>
            <!-- <el-select v-model="state.data.facility" multiple collapse-tags
                                                        collapse-tags-tooltip :max-collapse-tags="999"
                                                        placeholder="选择设备">
                                                        <el-option v-for="item in state.facilityOptions"
                                                                :label="item.name" :value="item.id" />
                                                </el-select> -->
            <el-input v-model="state.data.facility" placeholder="设备"></el-input>
          </el-form-item>
          <el-form-item class="item" label="地区" prop="state.data.region_id">
            <!-- <el-input
                                                v-model="state.data.location_id"
                                                aria-required="true"></el-input> -->
            <el-tree-select v-model="state.data.region_id" :data="state.regionOptions"
              :props="{ value: 'id', label: 'name', children: 'children' }" check-strictly placeholder="选择地区" />
          </el-form-item>
          <el-form-item class="item" label="状态" prop="state.data.status" required><el-input v-model="state.data.status"
              aria-required="true"></el-input></el-form-item>
          <el-form-item class="item" label="经度" prop="state.data.latitude" required><el-input
              v-model="state.data.latitude" aria-required="true"></el-input></el-form-item>
          <el-form-item class="item" label="纬度" prop="state.data.longitude" required><el-input
              v-model="state.data.longitude" aria-required="true"></el-input></el-form-item>
          <el-form-item class="item" label="自定义配置数据" required>
            <el-input v-model="state.data.custom_field_data" :rows="5" type="textarea" placeholder='例如：
{
    "ids":[5,6],
    "group_id":8
}' />
          </el-form-item>
          <el-form-item class="item" label="物理地址" prop="state.data.physical_address" required><el-input
              v-model="state.data.physical_address" aria-required="true"></el-input></el-form-item>
          <el-form-item class="item" label="物流地址" prop="state.data.shipping_address" required><el-input
              v-model="state.data.shipping_address" aria-required="true"></el-input></el-form-item>
        </el-form>
        <el-form class="editform" :model="state.data" label-width="auto">
          <h4>其他</h4>
          <el-form-item class="item" label="评价" required>
            <el-input v-model="state.data.comments" :rows="2" type="textarea" placeholder='请输入评价' />
          </el-form-item>
          <el-form-item class="item" label="描述" required>
            <el-input v-model="state.data.description" :rows="2" type="textarea" placeholder='请输入描述' />
          </el-form-item>

        </el-form>
        <template #footer>
          <span class="dialog-footer" align="center">
            <el-button @click="onCancel()">取 消</el-button>
            <el-button type="primary" @click="onSubmit" :loading="state.loading">保
              存</el-button>
          </span>
        </template>
      </el-card>
    </div>
    <Info v-else-if="currentView === 'Info'"></Info>
  </div>

</template>
<script setup lang="ts">
import { addSite, listSiteGroup, updateSite } from '@/api/dicm/site';
import { listTenant } from '@/api/dicm/tenant';
import { getCurrentInstance, onMounted, reactive, ref } from 'vue';
import Info from './info.vue';
import router from '@/router';
import { ElMessage } from 'element-plus';

let currentView = ref<'Info' | null>(null);
// 接收父组件传参
const props = defineProps({
  item: {},
  title: {
    type: String,
    default: () => "",
  },
})
const { proxy } = getCurrentInstance() as any;
const state = reactive({
  loading: false,
  data: {
    name: undefined,
    _name: undefined,
    slug: undefined,
    status: undefined,
    group_id: undefined,
    latitude: undefined,
    longitude: undefined,
    shipping_address: undefined,
    tenant_id: undefined,
    time_zone: undefined,
    region_id: undefined,
    facility: undefined,
    physical_address: undefined,
    custom_field_data: undefined,
    comments: undefined,
    description: undefined,
  },
  groupOptions: [],
  tenantOptions: [],
  regionOptions: [
    { id: 1, name: "diqu1" }
  ],
  tzOptions: [
    { name: "上海", value: "Shanghai" },
    { name: "伦敦", value: "London" }
  ]
})

const getData = () => {
  if (props.item !== null) {
    state.data = props.item
  }
  state.title = props.title;
  listTenant({}).then((response: any) => {
    if (response.code != 200) {
      state.loading = false;
    }
    state.tenantOptions = response.data.data;

  })
  listSiteGroup({}).then((res: any) => {
    if (res.code != 200) {
      state.loading = false;
    }
    state.groupOptions = res.data.data;
  })

}



// 取消按钮
const onCancel = () => {
  if (props.item !== null) {
    currentView.value = 'Info'
  } else {
    router.push({ path: `/dicm/sites` })
  }
}
// 返回按钮
const handleExit = () => {
  if (props.item !== null) {
    currentView.value = 'Info'
  } else {
    router.push({ path: `/dicm/sites` })
  }
}
// 保存按钮
const onSubmit = () => {
  state.loading = true
  if (props.item !== null) {
    updateSite(state.data).then((res: any) => {
      // console.log("编辑站点数据= ", state.data);
      if (res.code == 200) {
        ElMessage.success("修改【 " + state.data.name + " 】站点信息成功");
      }
      state.loading = false
    })
    currentView.value = 'Info'
  }
  else {
    // console.log("添加站点：", state.data);
    addSite(state.data).then((res: any) => {
      if (res.code === 200) {
        ElMessage({
          type: "success",
          message: "操作成功",
        });
      }
      state.loading = false
    })
    router.push({ path: `/dicm/sites` })
  }
}

// 页面加载时
onMounted(() => {
  getData();
})

</script>
<style lang="css" scoped>
.part {
  margin: 3px;
}

.item {
  margin: 20px;
}

.editform {
  border: 2px;
}
</style>