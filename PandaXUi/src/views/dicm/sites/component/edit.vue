<template>
    <div class="app-container">
        <h1>站点数据编辑视图</h1>
    </div>

</template>
<script setup lang="ts">
import { getCurrentInstance } from 'vue';


const props = defineProps({
  title: {
    type: String,
    default: () => "",
  },
})

const { proxy } = getCurrentInstance() as any;
</script>
<style lang="css" scoped>

</style>