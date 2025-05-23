<template>
    <div class="app-container">
        <div v-if="currentView === null">
            <el-card>
                <template #header>
                    <div>
                        <el-descriptions class="margin-top" :title="state.title" :column="1">
                            <el-descriptions-item><span>创建于 {{ dayjs(state.data.created).format('YYYY-MM-DD HH:mm:ss')
                                    }}
                                    更新于 {{ dayjs(state.data.last_updated).format('YYYY-MM-DD HH:mm:ss')
                                    }}</span></el-descriptions-item>

                            <template #extra>
                                <!-- <el-button type="primary" @click="handleAddDevice()"><el-icon style="margin-right: 5px;">
                                        <EditPen />
                                    </el-icon></el-button> -->
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
                            <el-tab-pane label="站点" name="rack">
                                <el-row>
                                    <el-col :span="12" style="width: 100%;justify-content: center;font-weight: 500;">
                                        <el-descriptions border class="margin-info" title="基础属性" :column="1">
                                            <el-descriptions-item label="编码">{{ state.data.id }}</el-descriptions-item>
                                            <el-descriptions-item label="名称">{{ state.data.name }}</el-descriptions-item>
                                            <el-descriptions-item label="别名">{{
                                                state.data._name}}</el-descriptions-item>
                                            <el-descriptions-item label="标识符">{{
                                                state.data.slug}}</el-descriptions-item>
                                            <el-descriptions-item label="地区">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Region(state.data.region_id)"
                                                    style="font-weight: 600;">
                                                    {{ state.data.region_id == 0 ? "-----" : state.data.region_id }}
                                                </el-link>
                                            </el-descriptions-item>
                                            <el-descriptions-item label="站点组">
                                                <el-link type="primary" :underline="false" @click="handle2Group()"
                                                    style="font-weight: 600;">
                                                    {{ state.data.group_id == 0 ? "-----" : state.data.group_id }}
                                                </el-link></el-descriptions-item>
                                            <el-descriptions-item label="设备">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Facility(state.data.facility)"
                                                    style="font-weight: 600;">
                                                    {{ state.data.facility == "" ? "-----" : state.data.facility }}
                                                </el-link>
                                            </el-descriptions-item>
                                            <el-descriptions-item label="租户">
                                                <el-link type="primary" :underline="false"
                                                    @click="handle2Tenant(state.data.tenant_id)"
                                                    style="font-weight: 600;">
                                                    {{ state.data.tenant_id == 0 ? "-----" : state.data.tenant_id }}
                                                </el-link></el-descriptions-item>
                                            <el-descriptions-item label="物理地址">{{
                                                state.data.physical_address}}</el-descriptions-item>
                                            <el-descriptions-item label="物流地址">{{
                                                state.data.shipping_address}}</el-descriptions-item>
                                            <el-descriptions-item label="状态"><el-tag
                                                    :type="state.data.status === 'abandon' ? 'danger' : 'success'"
                                                    disable-transitions>
                                                    {{ state.data.status }}</el-tag></el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="物理数据" :column="1">
                                            <el-descriptions-item label="GPS">{{ [Number(state.data.latitude),
                                                Number(state.data.longitude)] }}
                                            </el-descriptions-item>
                                            <el-descriptions-item label="时区"><el-tag> {{
                                                    state.data.time_zone}}</el-tag></el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-divider />
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="相关对象" :column="1">
                                            <el-descriptions-item label="设备">设备</el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="其他数据" :column="1">
                                            <el-descriptions-item label="评价">{{ state.data.comments == "" ? "暂无评价" :
                                                state.data.comments
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="描述">{{ state.data.description == "" ? "暂无说明" :
                                                state.data.description
                                                }}</el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                </el-row>
                            </el-tab-pane>
                            <el-tab-pane label="自定义配置数据" name="custom_field_data">
                                <el-descriptions size="default">
                                    <el-descriptions-item label="" size="large" style="height: 10%;">{{
                                        state.data.custom_field_data == "" ? "暂无配置信息" : state.data.custom_field_data
                                        }}</el-descriptions-item>
                                </el-descriptions>
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
                            <el-tab-pane label="操作历史" name="site_history">
                                <el-table :data="state.history" row-key="id" border default-expand-all>
                                    <el-table-column prop="user_id" label="操作人" width="100" fixed />
                                    <el-table-column prop="name" label="名称" width="100" />
                                    <el-table-column prop="rack_id" label="相关机柜" width="100" />
                                    <el-table-column prop="tenant_id" label="租户" width="100" />
                                    <el-table-column prop="user_id" label="用户" width="100" />
                                    <el-table-column prop="custom_field_data" label="自定义配置数据" width="150" />
                                    <el-table-column prop="description" label="描述" width="150" />
                                    <el-table-column prop="comments" label="评价" width="150" />

                                    <el-table-column prop="last_updated" label="更新" width="180" />
                                    <el-table-column prop="type" label="操作类型" width="180" />
                                    <el-table-column prop="created" label="操作时间" width="180" />
                                </el-table>
                            </el-tab-pane>
                        </el-tabs>
                    </div>
                </template>
            </el-card>
        </div>
        <Group v-if="currentView === 'Group'"></Group>
        <Editt v-else-if="currentView === 'Editt'" :item="editData" :title="state.title"></Editt>
    </div>
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router';
import Editt from './edit.vue';
import Group from './group.vue';
import { onMounted, reactive, ref } from 'vue';
import { listSiteInfo } from '@/api/dicm/site';
import { dayjs } from 'element-plus';
import type { TabsPaneContext } from 'element-plus'

import router from '@/router';

let currentView = ref<null | 'Editt' | 'Group'>(null);


const route = useRoute()
const siteName = route.params.name
const activeName = ref('rack')
// tab数据分页
const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}
const editData = ref()
const state = reactive({
    title: "",
    data: {},
    device: [],
    history: [],
})
const isGroup = ref(false)

const getSite = () => {
    listSiteInfo(siteName).then((response: any) => {
        console.log("fanhui站点数据：", response.data);
        if (response.code == 200) {
            state.data = response.data;
            editData.value = response.data
        }
    })
}
console.log("站点数据：", state.data);


//跳转站点 
const handle2Region = (Id: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    // router.push({ name: '/dicm/sites', query: { id: Id } })
}
const handle2Facility = (id: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/dicm/sites', query: { id: id } })
}
const handle2Tenant = (id: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/tenant', query: { id: id } })
}
const handle2Location = (id: any) => {
    // router.push(`/dicm/sites/list?id=${siteId}}`)
    router.push({ name: '/dicm/sites', query: { id: id } })
}
const handle2Group = () => {
    currentView.value = 'Group'
    // isGroup.value = true
}

/** 编辑按钮 */
const handleUpdate = () => {
    currentView.value = "Editt";
    state.title = "编辑站点信息";

}
/** 返回按钮 */
const handleExit = () => {
    router.push({ path: `/dicm/sites` })
}

onMounted(() => {
    state.title = "站点/ " + siteName
    getSite()
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