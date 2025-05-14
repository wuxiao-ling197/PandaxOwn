<template>
    编辑信息页面
    <div class="container">
        <el-card class="view" :title="state.title">
            <el-form :model="state.data">
                <h4>基本信息</h4>
                <el-form-item class="item" label="名称" prop="state.data.name" required><el-input
                        v-model="state.data.name" placeholder="请输入姓名"></el-input></el-form-item>
                <el-form-item class="item" label="别名" prop="state.data._name" ><el-input aria-required="true"></el-input></el-form-item>

                <el-form-item class="item" label="类型" prop="state.data.type" required><el-input aria-required="true"></el-input></el-form-item>

                <el-form-item class="item" label="资产标签" prop="state.data.asset_tag"  required><el-input aria-required="true"></el-input></el-form-item>

                <el-form-item class="item" label="机架功能" prop="state.data.role_id" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="序列号" prop="state.data.serial" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="站点" prop="state.data.site_id" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="租户" prop="state.data.tenant_id" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="理论最大承重" prop="state.data.weight" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="重量单位" prop="state.data.weight_unit" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="内部宽度" prop="state.data.width" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="最大承重限制" prop="state.data._abs_max_weight" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="实际总重量" prop="state.data._abs_weight" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="名称" prop="state.data.name" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="名称" prop="state.data.name" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="描述" prop="state.data.name" required><el-input aria-required="true"></el-input></el-form-item>
                <el-form-item class="item" label="评价" prop="state.data.name" required><el-input aria-required="true"></el-input></el-form-item>

                <el-form-item class="item" label="自定义配置数据">
                    <el-input title="自定义配置数据" v-model="custom_filed_data" style="width: 240px" :rows="2" type="textarea"
                        placeholder="Please input" />
                </el-form-item>

            </el-form>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
import { reactive, ref } from 'vue';
const custom_filed_data = ref('')
const state = reactive({
    data: [],
    title: "",
})

state.title = "编辑中————"

</script>
<style lang="css" scoped>
.part {
    margin: 3px;
}

.item {
    margin: 5px;
}
</style>