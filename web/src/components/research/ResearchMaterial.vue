<script setup>
const emit = defineEmits(['itemAdd'])
function itemAdd(item) {
  console.log('itemAdd', item)
  emit('itemAdd', item)
}
</script>


<template>
  <div style="margin: 10px 5px 10px 10px;;">
    <el-space wrap>
      <el-button type="primary" @click="itemAdd('input')">短文本</el-button>
      <el-button type="primary" @click="itemAdd('textarea')">长文本</el-button>
      <el-button type="primary" @click="itemAdd('radio')">单选框</el-button>
      <el-button type="primary" @click="itemAdd('checkbox')">复选框</el-button>
      <el-button type="primary" @click="itemAdd('select')">下拉选择</el-button>
      <el-button type="primary" @click="itemAdd('time')">时间选择</el-button>
      <el-button type="primary" @click="itemAdd('date')">日期选择</el-button>
    </el-space>
  </div>
</template>

<style lang="scss" scoped>
.container {
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  padding: 10px;
  height: 100%;
}
</style>