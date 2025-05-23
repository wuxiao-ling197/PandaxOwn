<template>
        <div class="container">
                <div v-if="currentView === null">
                        <el-card class="view" shadow="always">
                                <!-- <template #extra> -->
                                <el-button type="info" size="small" @click="handleExit()"
                                        style="margin-bottom: 10px;"><el-icon style="margin-right: 5px;">
                                                <ArrowLeftBold />
                                        </el-icon>
                                        返回</el-button>
                                <!-- </template> -->
                                <template #header>
                                        <div class="card-header"
                                                style="font-weight: 600;margin: 3px 5px;text-justify: auto;">
                                                <span>{{ state.title }}</span>
                                        </div>
                                </template>
                                <el-form class="editform" :model="state.data" label-width="auto">
                                        <h4>基本信息</h4>
                                        <el-form-item class="item" label="名称" prop="state.data.name" required><el-input
                                                        v-model="state.data.name"
                                                        placeholder="请输入姓名"></el-input></el-form-item>
                                        <el-form-item class="item" label="别名" prop="state.data._name"><el-input
                                                        v-model="state.data._name"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="租户" prop="state.data.tenant_id" required>
                                                <template #header>
                                                        <span>租户
                                                                <el-tooltip content="机柜的管理人员、维护人员。" placement="top">
                                                                        <el-icon>
                                                                                <QuestionFilled />
                                                                        </el-icon>
                                                                </el-tooltip>
                                                        </span>
                                                </template>
                                                <el-select v-model="state.data.tenant_id" collapse-tags
                                                        collapse-tags-tooltip placeholder="选择租户">
                                                        <el-option v-for="item in state.tenantOptions"
                                                                :label="item.name" :value="item.id" />
                                                </el-select>
                                        </el-form-item>
                                        <el-form-item class="item" label="资产标签" prop="state.data.asset_tag">
                                                <!-- <el-select
                                                        v-model="state.data.asset_tag" collapse-tags
                                                        collapse-tags-tooltip placeholder="选择标签">
                                                        <el-option v-for="item in state.tagOptions" :label="item.name"
                                                                :value="item.name" />
                                                </el-select> -->
                                                <el-input
                                                        v-model="state.data.asset_tag"
                                                        placeholder="请输入唯一资产标签"></el-input>
                                        </el-form-item>
                                        <el-form-item class="item" label="机架角色" prop="state.data.role_id" required>
                                                <template #header>
                                                        <span>机柜角色
                                                                <el-tooltip content="即该机柜以及其中设备主要做什么用途，比如存储？xxx系统服务器？等等"
                                                                        placement="top">
                                                                        <el-icon>
                                                                                <QuestionFilled />
                                                                        </el-icon>
                                                                </el-tooltip>
                                                        </span>
                                                </template>
                                                <el-select v-model="state.data.role_id" collapse-tags
                                                        collapse-tags-tooltip placeholder="选择角色">
                                                        <el-option v-for="item in state.roleOptions" :label="item.name"
                                                                :value="item.id" />
                                                </el-select></el-form-item>
                                        <el-form-item class="item" label="设备" prop="state.data.facility_id" required>
                                                <template #header>
                                                        <span>安装设备
                                                                <el-tooltip content="该机柜中安装的设备。" placement="top">
                                                                        <el-icon>
                                                                                <QuestionFilled />
                                                                        </el-icon>
                                                                </el-tooltip>
                                                        </span>
                                                </template>
                                                <el-select v-model="state.data.facility_id" multiple collapse-tags
                                                        collapse-tags-tooltip :max-collapse-tags="999"
                                                        placeholder="选择设备">
                                                        <el-option v-for="item in state.facilityOptions"
                                                                :label="item.name" :value="item.id" />
                                                </el-select></el-form-item>
                                        <el-form-item class="item" label="站点" prop="state.data.site_id"
                                                required><el-tree-select v-model="state.data.site_id"
                                                        :data="state.siteOptions"
                                                        :props="{ value: 'id', label: 'name', children: 'children' }"
                                                        check-strictly placeholder="选择站点" /></el-form-item>
                                        <el-form-item class="item" label="物理位置" prop="state.data.location_id" required>
                                                <!-- <el-input
                                                v-model="state.data.location_id"
                                                aria-required="true"></el-input> -->
                                                <el-tree-select v-model="state.data.location_id"
                                                        :data="state.locationOptions"
                                                        :props="{ value: 'id', label: 'name', children: 'children' }"
                                                        check-strictly placeholder="选择物理位置" />
                                        </el-form-item>
                                        <el-form-item class="item" label="状态" required><el-input
                                                        v-model="state.data.status"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="自定义配置数据" required>
                                                <el-input v-model="state.data.custom_field_data" :rows="5"
                                                        type="textarea" placeholder='例如：
{
    "ids":[5,6],
    "group_id":8
}' />
                                        </el-form-item>
                                        <h4>物理数据</h4>
                                        <el-form-item class="item" label="机架类型" prop="state.data.type"
                                                required><el-input v-model="state.data.type"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="序列" prop="state.data.serial"
                                                required><el-input v-model="state.data.serial"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="内部宽度" prop="state.data.width"
                                                required><el-input v-model="state.data.width"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="U高度" required>
                                                <el-input title="U高度" v-model="state.data.u_height" />
                                        </el-form-item>
                                        <el-form-item class="item" label="起始U位" required><el-input
                                                        v-model="state.data.starting_unit"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="安装深度" required><el-input
                                                        v-model="state.data.mounting_depth"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="外部深度" required><el-input
                                                        v-model="state.data.outer_depth"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="外部宽度" required><el-input
                                                        v-model="state.data.outer_width"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="外部尺寸单位" required><el-input
                                                        v-model="state.data.outer_unit"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="描述性单位" required><el-input
                                                        v-model="state.data.desc_units"
                                                        aria-required="true"></el-input></el-form-item>
                                </el-form>
                                <el-form class="editform" :model="state.data" label-width="auto">
                                        <h4>重量数据</h4>
                                        <el-form-item class="item" label="重量单位" required><el-input
                                                        v-model="state.data.weight_unit"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="实际总重量" required><el-input
                                                        v-model="state.data._abs_weight"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="理论最大承重" required><el-input
                                                        v-model="state.data._abs_max_weight"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="设备重量" required><el-input
                                                        v-model="state.data.weight"
                                                        aria-required="true"></el-input></el-form-item>
                                        <el-form-item class="item" label="最大承重限制" required><el-input
                                                        v-model="state.data.max_weight"
                                                        aria-required="true"></el-input></el-form-item>
                                        <h4>其他</h4>
                                        <el-form-item class="item" label="评价" required>
                                                <el-input v-model="state.data.comments" :rows="2" type="textarea"
                                                        placeholder='请输入评价' />
                                        </el-form-item>
                                        <el-form-item class="item" label="描述" required>
                                                <el-input v-model="state.data.description" :rows="2" type="textarea"
                                                        placeholder='请输入描述' />
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
                <List v-else-if="currentView === 'List'"></List>
        </div>
</template>
<script lang="ts" setup>
import { addRacks, listRackRole, updateRacks } from '@/api/dicm/rack';
import { listLocations, listSites } from '@/api/dicm/site';
import { listTenant } from '@/api/dicm/tenant';
import router from '@/router';
import { ElMessage } from 'element-plus';
import { onMounted, reactive, ref } from 'vue';
import Info from './info.vue';
import List from '../index.vue';

let currentView = ref<'Info' | 'List' | null>(null);
// 接收父组件传参
const props = defineProps<{
        item: {} | null;
        title: {
                type: String,
                default: () => "",
        }
}>();
console.log("接收父组件传参：", props.item);


const state = reactive({
        loading: false,
        data: {
                name: undefined,
                _name: undefined,
                type: undefined,
                serial: undefined,
                width: undefined,
                u_height: undefined,
                starting_unit: undefined,
                mounting_depth: undefined,
                outer_depth: undefined,
                outer_width: undefined,
                weight_unit: undefined,
                outer_unit: undefined,
                desc_units: undefined,
                _abs_weight: undefined,
                _abs_max_weight: undefined,
                status: undefined,
                description: undefined,
                site_id: undefined,
                facility_id: undefined,
                tenant_id: undefined,
                weight: undefined,
                location_id: undefined,
                max_weight: undefined,
                asset_tag: undefined,
                role_id: undefined,
                custom_field_data: undefined,
                comments: undefined,

        },
        title: "",
        // 下拉选择
        tenantOptions: [],
        tagOptions: [{ id: 1, name: "test" }, { id: 2, name: "add" }, { id: 3, name: "真就唯一使用了？使用后就违反唯一性了？" }],
        roleOptions: [],
        facilityOptions: [{ id: 1, name: "test" }],
        siteOptions: [],
        locationOptions: [],
})


// 取消按钮
const onCancel = () => {
        if (props.item !== null) {
                currentView.value = 'Info'
        } else {
                // router.push({ path: `/dicm/racks` })
                currentView.value = 'List'
        }
}
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
        listRackRole({}).then((response: any) => {
                if (response.code != 200) {
                        state.loading = false;
                }
                state.roleOptions = response.data.data;
        })
        listSites({}).then((response: any) => {
                if (response.code != 200) {
                        state.loading = false;
                }
                state.siteOptions = response.data.data;
        })
        listLocations({}).then((response: any) => {
                if (response.code != 200) {
                        state.loading = false;
                }
                state.locationOptions = response.data.data;
        })
}
const handleExit = () => {
        if (props.item !== null) {
                currentView.value = 'Info'
        } else {
                // router.push({ path: `/dicm/racks` })
                currentView.value = 'List'
        }
}
// 保存按钮
const onSubmit = () => {
        state.loading = true
        if (props.item !== null) {
                updateRacks(state.data).then((res: any) => {
                        // console.log("编辑机柜数据= ", state.data);
                        if (res.code == 200) {
                                ElMessage.success("修改机柜实例成功");
                        }
                        state.loading = false
                })
                currentView.value = 'Info'
        }
        else {
                console.log("添加机柜：", state.data);
                addRacks(state.data).then((res: any) => {
                        if (res.code === 200) {
                                ElMessage({
                                        type: "success",
                                        message: "操作成功",
                                });
                        }
                        state.loading = false
                })
                router.push({ path: `/dicm/racks` })
                currentView.value = 'List'
        }
}

// 页面加载时
onMounted(() => {
        // 查询列表
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