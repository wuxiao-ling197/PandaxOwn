<template>
    <div class="app-container">
        <div v-if="currentView === 'Infoo'">
            <el-card>
                <template #header>
                    <div>
                        <el-descriptions class="margin-top" :title="state.title" :column="1" :size="size"
                            :style="blockMargin">
                            <el-descriptions-item><span>创建于 {{ dayjs(state.data.created).format('YYYY-MM-DD HH:mm:ss') }}
                                    更新于 {{ dayjs(state.data.last_updated).format('YYYY-MM-DD HH:mm:ss')
                                    }}</span></el-descriptions-item>

                            <template #extra>
                                <el-button><el-icon style="margin-right: 5px;">
                                        <MagicStick />
                                    </el-icon>收藏</el-button>
                                <el-button @click="handleUpdate()"><el-icon style="margin-right: 5px;">
                                        <EditPen />
                                    </el-icon>编辑</el-button>
                            </template>
                        </el-descriptions>
                        <el-tabs v-model="activeName" class="demo-tabs" type="border-card" @tab-click="handleClick">
                            <el-tab-pane label="机柜" name="rack">
                                <el-row>
                                    <el-col :span="12" style="width: 100%;justify-content: center;font-weight: 500;">
                                        <el-descriptions border class="margin-info" title="基础属性" :column="1"
                                            :size="size" :style="blockMargin">
                                            <el-descriptions-item label="编码">{{ state.data.id }}</el-descriptions-item>
                                            <el-descriptions-item label="名称">{{ state.data.name
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="别名">{{ state.data._name
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="标签"><el-tag>
                                                    {{ state.data.asset_tag == "" ? "暂未分配标签" : state.data.asset_tag
                                                    }}
                                                </el-tag></el-descriptions-item>
                                            <el-descriptions-item label="机柜用途">{{ state.data.role_id == 0 ? "-----" :
                                                state.data.role_id
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="站点">{{ state.data.site_id == 0 ? "-----" :
                                                state.data.site_id
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="设备">{{ state.data.facility_id == 0 ? "-----" :
                                                state.data.facility_id
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="租户">{{ state.data.tenant_id == 0 ? "-----" :
                                                state.data.tenant_id
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="物理位置">{{ state.data.location_id == 0 ? "-----" :
                                                state.data.location_id
                                                }}</el-descriptions-item>
                                            <!-- :type="getTagType(status)" -->
                                            <el-descriptions-item label="状态"><el-tag>
                                                    {{ state.data.status
                                                    }}
                                                </el-tag></el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="物理数据" :column="1"
                                            :size="size" :style="blockMargin">
                                            <el-descriptions-item label="机架类型">{{ state.data.type == "" ? "-----" :
                                                state.data.type
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="序列">{{ state.data.serial == "" ? "-----" :
                                                state.data.serial
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="内部宽度">{{ state.data.width == 0 ? "-----" :
                                                state.data.width
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="U高度">{{ state.data.u_height == 0 ? "-----" :
                                                state.data.u_height
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="起始U位">{{ state.data.starting_unit == 0 ? "-----"
                                                :
                                                state.data.starting_unit
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="安装深度">{{ state.data.mounting_depth == 0 ?
                                                "-----" :
                                                state.data.mounting_depth
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="外部深度">{{ state.data.outer_depth == 0 ? "-----" :
                                                state.data.outer_depth
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="外部宽度">{{ state.data.outer_width == 0 ? "-----" :
                                                state.data.outer_width
                                                }}</el-descriptions-item>
                                            <el-descriptions-item label="外部尺寸单位">{{ state.data.outer_unit == "" ? "-----"
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
                                            :size="size" :style="blockMargin">
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
                                            :size="size" :style="blockMargin">
                                            <el-descriptions-item label="设备">设备{{ state.data.comments == "" ? "暂无评价" :
                                                state.data.comments
                                                }}</el-descriptions-item>
                                        </el-descriptions>
                                    </el-col>
                                    <el-divider />
                                    <el-col :span="12">
                                        <el-descriptions border class="margin-info" title="其他数据" :column="1"
                                            :size="size" :style="blockMargin">
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
                            <el-tab-pane label="机柜预留" name="reserve">机柜预留列表</el-tab-pane>
                            <el-tab-pane label="待上架设备" name="waiting_device">待上架设备列表</el-tab-pane>
                        </el-tabs>
                    </div>
                </template>
            </el-card>
        </div>
        <Editt v-else-if="currentView === 'Editt'"></Editt>
    </div>

</template>
<script setup lang="ts">
import { getCurrentInstance, onMounted, reactive, computed, ref } from 'vue';
import { ElMessage, type ComponentSize, dayjs } from 'element-plus'
import { addRacks, listRackInfo, updateRacks } from '@/api/dicm/rack';
import { useRoute } from 'vue-router';
import type { TabsPaneContext } from 'element-plus'
import Editt from './edit.vue'

const activeName = ref('rack')
// tab数据分页
const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}

let currentView = ref<'Infoo' | 'Editt'>('Infoo');

const props = defineProps({
    title: {
        type: String,
        default: () => "",
    },
})

const route = useRoute()
const rackName = route.params.name
const state = reactive({
    data: [],
    title: "",
})

const { proxy } = getCurrentInstance() as any;

const getDetail = () => {
    console.log("详情页获取动态参数=", rackName);
    state.title = "机柜/" + rackName
    listRackInfo(rackName).then((res: any) => {
        console.log('获取详情返回数据=', res.data);
        state.data = res.data
    })
}

const loadData = (row: any) => {
    if (row) {
        // 编辑数据
        console.log("编辑机柜数据= ", row);
        updateRacks(state.data).then((res: any) => {
            console.log("编辑机柜数据= ", row);
            if (res.code == 200) {
                ElMessage.success("修改机柜实例成功");
            }
        })
    } else {
        // 添加数据
        addRacks(state.data).then((res: any) => {
            console.log("新建机柜= ", state.data);
            if (res.code == 200) {
                ElMessage.success("新增成功");
            }
        })
    }
}
/** 新增按钮操作 */
const handleAdd = () => {
    state.title = "添加机柜实例";
    // 添加数据
    addRacks(state.data).then((res: any) => {
        console.log("新建机柜= ", state.data);
        if (res.code == 200) {
            ElMessage.success("新增成功");
        }
    })
};
/** 修改按钮操作 */
const handleUpdate = () => {
    console.log("切换到编辑视图");
    currentView.value="Editt"
    state.title = "编辑中";
    // 编辑数据
    // console.log("编辑机柜数据= ", state.data);
    // updateRacks(state.data).then((res: any) => {
    //     console.log("编辑机柜数据= ", state.data);
    //     if (res.code == 200) {
    //         ElMessage.success("修改机柜实例成功");
    //     }
    // })
};

// 添加标签
// const getTagType = (status:any) {
//       if (status.includes('在线')) return 'success'
//       if (status.includes('离线')) return 'danger'
//       if (status.includes('警告')) return 'warning'
//       return 'info' // 默认类型
//     }


const size = ref<ComponentSize>('default')

const blockMargin = computed(() => {
    const marginMap = {
        large: '32px',
        default: '28px',
        small: '24px',
    }
})

onMounted(() => {
    getDetail()
})

defineExpose({
    loadData,
});
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