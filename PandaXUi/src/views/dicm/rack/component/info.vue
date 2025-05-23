<template>
    <div class="app-container">
        <div v-if="currentView === 'Infoo'">
            <el-card>
                <template #header>
                    <div>
                        <el-descriptions class="margin-top" :title="state.title" :column="1" :size="size"
                            :style="blockMargin">
                            <el-descriptions-item><span>创建于 {{ dayjs(state.data.created).format('YYYY-MM-DD HH:mm:ss')
                                    }}
                                    更新于 {{ dayjs(state.data.last_updated).format('YYYY-MM-DD HH:mm:ss')
                                    }}</span></el-descriptions-item>

                            <template #extra>
                                <el-button type="primary" @click="handleAddDevice()"><el-icon style="margin-right: 5px;">
                                        <CirclePlusFilled />
                                    </el-icon>添加设备</el-button>
                                <el-button type="warning"><el-icon style="margin-right: 5px;">
                                        <MagicStick />
                                    </el-icon>收藏</el-button>
                                <el-button type="primary" @click="handleUpdate()"><el-icon style="margin-right: 5px;">
                                        <EditPen />
                                    </el-icon>编辑</el-button>
                                <el-button type="info" @click="handleExit()"><el-icon style="margin-right: 5px;">
                                        <ArrowLeftBold />
                                    </el-icon>
                                    返回</el-button>
                            </template>
                        </el-descriptions>
                        <el-tabs v-model="activeName" class="demo-tabs" type="border-card" @tab-click="handleClick">
                            <el-tab-pane label="机柜" name="rack">
                                <el-row>
                                    <el-col :span="12" style="width: 100%;justify-content: center;font-weight: 500;">
                                        <el-descriptions border class="margin-info" title="基础属性" :column="1"
                                           >
                                            <el-descriptions-item label="编码">{{ state.data.id }}</el-descriptions-item>
                                            <el-descriptions-item label="名称">{{ state.data.name
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="别名">{{ state.data._name
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="标签"><el-tag>
                                                    {{ state.data.asset_tag == "" ? "暂未分配标签" : state.data.asset_tag
                                                    }}
                                                </el-tag></el-descriptions-item>
                                            <el-descriptions-item label="机柜角色">
                                                <el-link type="primary" :underline="false" @click="handle2Role()"
                                                    style="font-weight: 600;">
                                                    {{ state.data.role_id == 0 ? "-----" : state.data.role_id }}
                                                </el-link></el-descriptions-item>
                                            <el-descriptions-item label="站点">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Site(state.data.site_id)" style="font-weight: 600;">
                                                    {{ state.data.site_id == 0 ? "-----" : state.data.site_id }}
                                                </el-link>
                                            </el-descriptions-item>
                                            <el-descriptions-item label="设备">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Facility(state.data.facility_id)"
                                                    style="font-weight: 600;">
                                                    {{ state.data.facility_id == 0 ? "-----" : state.data.facility_id }}
                                                </el-link>
                                            </el-descriptions-item>
                                            <el-descriptions-item label="租户">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Tenant(state.data.tenant_id)"
                                                    style="font-weight: 600;">
                                                    {{ state.data.tenant_id == 0 ? "-----" : state.data.tenant_id }}
                                                </el-link></el-descriptions-item>
                                            <el-descriptions-item label="物理位置">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Location(state.data.location_id)"
                                                    style="font-weight: 600;">
                                                    {{ state.data.location_id == 0 ? "-----" : state.data.location_id }}
                                                </el-link></el-descriptions-item>
                                            <!-- :type="getTagType(status)" -->
                                            <el-descriptions-item label="状态"><el-tag :type="state.data.status === 'abandon' ? 'danger' : 'success'"
                                                    disable-transitions>
                                                    {{ state.data.status
                                                    }}
                                                </el-tag></el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="物理数据" :column="1"
                                           >
                                            <el-descriptions-item label="机架类型">{{ state.data.type == "" ? "-----" :
                                                state.data.type
                                                }}
                                            </el-descriptions-item>
                                            <el-descriptions-item label="序列">{{ state.data.serial == "" ? "-----" :
                                                state.data.serial
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="内部宽度">{{ state.data.width == 0 ? "-----" :
                                                state.data.width
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="U高度">{{ state.data.u_height == 0 ? "-----" :
                                                state.data.u_height
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="起始U位">{{ state.data.starting_unit == 0 ?
                                                "-----"
                                                :
                                                state.data.starting_unit
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="安装深度">{{ state.data.mounting_depth == 0 ?
                                                "-----" :
                                                state.data.mounting_depth
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="外部深度">{{ state.data.outer_depth == 0 ? "-----"
                                                :
                                                state.data.outer_depth
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="外部宽度">{{ state.data.outer_width == 0 ? "-----"
                                                :
                                                state.data.outer_width
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="外部尺寸单位">{{ state.data.outer_unit == "" ?
                                                "-----"
                                                : state.data.outer_unit
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="描述性单位">{{ state.data.desc_units == true ? "是" :
                                                "否"
                                                }}</el-descriptions-item>

                                        </el-descriptions>
                                    </el-col>
                                    <el-divider />
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="重量数据" :column="1"
                                           >
                                            <el-descriptions-item label="重量单位">{{ state.data.weight_unit == "" ? "-----"
                                                : state.data.weight_unit
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="实际总重量">{{ state.data._abs_weight == 0 ? "-----"
                                                : state.data._abs_weight
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="理论最大承重">{{ state.data._abs_max_weight == 0 ?
                                                "-----" :
                                                state.data._abs_max_weight
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="设备重量">{{ state.data.weight == 0 ? "-----" :
                                                state.data.weight
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="最大承重限制">{{ state.data.max_weight == 0 ? "-----"
                                                : state.data.max_weight
                                                }}</el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="相关对象" :column="1"
                                           >
                                            <el-descriptions-item label="设备">设备{{ state.data.comments == "" ? "暂无评价" :
                                                state.data.comments
                                                }}</el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-divider />
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="其他数据" :column="1"
                                           >
                                            <el-descriptions-item label="评价">{{ state.data.comments == "" ? "暂无评价" :
                                                state.data.comments
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="描述">{{ state.data.description == "" ? "暂无说明" :
                                                state.data.description
                                                }}</el-descriptions-item>
                                            <!-- <el-descriptions-item label="添加时间">{{ state.data.created
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="更新时间">{{ state.data.last_updated
                                                }}</el-descriptions-item> -->
                                        </el-descriptions>
                                    </el-col>
                                </el-row>
                            </el-tab-pane>
                            <el-tab-pane label="自定义配置数据" name="custom_field_data">
                                <el-descriptions>
                                    <el-descriptions-item label="" size="large" style="height: 10%;">{{
                                        state.data.custom_field_data == "" ? "暂无配置信息" : state.data.custom_field_data
                                        }}</el-descriptions-item>
                                </el-descriptions>
                            </el-tab-pane>
                            <el-tab-pane label="机柜预留" name="reserve">
                                <el-table :data="state.reserve" row-key="id" border default-expand-all>
                                    <!-- <el-table-column prop="id" label="编码" width="100" fixed /> -->
                                    <el-table-column prop="units" label="单位" width="100" />
                                    <el-table-column prop="rack_id" label="相关机柜" width="100" />
                                    <el-table-column prop="tenant_id" label="租户" width="100" />
                                    <el-table-column prop="user_id" label="用户" width="100" />
                                    <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" />
                                    <el-table-column prop="description" label="描述" width="150" />
                                    <el-table-column prop="comments" label="评价" width="150" />
                                    <!-- <el-table-column prop="created" label="添加" width="180" />
                                    <el-table-column prop="last_updated" label="更新" width="180" /> -->
                                </el-table>
                            </el-tab-pane>
                            <el-tab-pane label="待上架设备" name="waiting_device">
                                <el-table :data="state.device" row-key="id" border default-expand-all>
                                    <el-table-column prop="id" label="编码" width="100" fixed />
                                    <el-table-column prop="name" label="名称" width="100" />
                                    <!-- <el-table-column prop="rack_id" label="相关机柜" width="100" />
                                    <el-table-column prop="tenant_id" label="租户" width="100" />
                                    <el-table-column prop="user_id" label="用户" width="100" />
                                    <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" />
                                    <el-table-column prop="description" label="描述" width="150" />
                                    <el-table-column prop="comments" label="评价" width="150" />
                                    <el-table-column prop="created" label="添加" width="180" />
                                    <el-table-column prop="last_updated" label="更新" width="180" /> -->
                                </el-table>
                            </el-tab-pane>
                            <el-tab-pane label="操作历史" name="rack_history">
                                <el-table :data="state.history" row-key="id" border default-expand-all>
                                    <el-table-column prop="user_id" label="操作人" width="100" fixed />
                                    <el-table-column prop="name" label="名称" width="100" />
                                    <!-- <el-table-column prop="rack_id" label="相关机柜" width="100" />
                                    <el-table-column prop="tenant_id" label="租户" width="100" />
                                    <el-table-column prop="user_id" label="用户" width="100" />
                                    <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" />
                                    <el-table-column prop="description" label="描述" width="150" />
                                    <el-table-column prop="comments" label="评价" width="150" />

                                    <el-table-column prop="last_updated" label="更新" width="180" /> -->
                                    <el-table-column prop="type" label="操作类型" width="180" />
                                    <el-table-column prop="created" label="操作时间" width="180" />
                                </el-table>
                            </el-tab-pane>
                        </el-tabs>
                    </div>
                </template>
            </el-card>
        </div>
        <Editt v-else-if="currentView === 'Editt'" :item="editData" :title="state.title"></Editt>
    </div>
    <Role v-if="isRole" :title="state.title" />
</template>
<script setup lang="ts">
import { getCurrentInstance, onMounted, reactive, computed, ref } from 'vue';
import { ElMessage, type ComponentSize, dayjs } from 'element-plus'
import { listRackInfo, listRackReserve } from '@/api/dicm/rack';
import { useRoute } from 'vue-router';
import type { TabsPaneContext } from 'element-plus'
import Editt from './edit.vue'
import router from '@/router';
import Role from './role.vue'
import List from '../index.vue'
const activeName = ref('rack')
// tab数据分页
const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}
const isRole = ref(false)

let currentView = ref<'Infoo' | 'Editt' | 'List'>('Infoo');

const props = defineProps({
    title: {
        type: String,
        default: () => "",
    },
})

const route = useRoute()
const rackName = route.params.name
const editData=ref()
const state = reactive({
    loading: true,
    data: {}, //接收机柜数据
    reserve: [], //机柜预留信息
    device: [],//机柜中设备信息
    history: [],//机柜管理操作记录
    title: "",
})

const { proxy } = getCurrentInstance() as any;

const getDetail = () => {
    state.title = "机柜/" + rackName
    listRackInfo(rackName).then((res: any) => {
        // console.log('获取详情返回数据=', res.data);
        state.data = res.data
        editData.value=res.data
        rackreserveList()
    })
}

// 机柜预留列表
const rackreserveList = () => {
    listRackReserve({}).then((response: any) => {
        if (response.code == 200) {
            state.reserve = response.data.data;
        }
    })
}

//跳转站点 
const handle2Site = (siteId: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/dicm/sites', query: { id: siteId } })
}
const handle2Facility = (siteId: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/dicm/sites', query: { id: siteId } })
}
const handle2Tenant = (siteId: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/tenant', query: { id: siteId } })
}
const handle2Location = (siteId: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/dicm/sites', query: { id: siteId } })
}
const handle2Role = () => {
    isRole.value = true
}

/** 新增设备按钮操作 todo*/
const handleAddDevice = () => {
        ElMessage.success("在机柜中部署安装设备");
};
/** 修改按钮操作 */
const handleUpdate = () => {
    currentView.value = "Editt"
    state.title = "编辑机柜信息";
};
// 返回到列表 
const handleExit=()=>{
    router.push({ path: `/dicm/racks`})
}

onMounted(() => {
    getDetail()
})

</script>
<style lang="css" scoped>
.el-descriptions {
    margin-top: 8px;
}

.cell-item {
    display: flex;
    align-items: center;
}

.margin-top {
    box-shadow: inset;
}

.margin-info {
    width: 95%;
    box-shadow: inset;
    margin-top: 10px;
}
</style>