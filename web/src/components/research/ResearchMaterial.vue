<script setup>
import { defineEmits } from 'vue'
const emit = defineEmits(['itemAdd'])
function itemAdd(item) {
  console.log('itemAdd', item)
  emit('itemAdd', item)
}
</script>


<template>
  <div class="container" style="margin: 10px 5px 10px 10px;;">
    <a-space wrap>
      <a-button type="primary" @click="itemAdd('input')">短文本</a-button>
      <a-button type="primary" @click="itemAdd('textarea')">长文本</a-button>
      <a-button type="primary" @click="itemAdd('radio')">单选框</a-button>
      <a-button type="primary" @click="itemAdd('checkbox')">复选框</a-button>
      <a-button type="primary" @click="itemAdd('select')">下拉选择</a-button>
      <a-button type="primary" @click="itemAdd('time')">时间选择</a-button>
      <a-button type="primary" @click="itemAdd('date')">日期选择</a-button>
    </a-space>
    <!-- <div class="btns-flex">
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('radio')">单选框</el-button>
        </div>
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('checkbox')">复选框</el-button>
        </div>
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('text')">短文本</el-button>
        </div>
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('textarea')">长文本</el-button>
        </div>
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('select')">下拉选择</el-button>
        </div>
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('timePicker')">时间选择</el-button>
        </div>
        <div class="btns-item">
          <el-button type="primary" size="small" @click="factorAdd('datePicker')">日期选择</el-button>
        </div>
      </div> -->
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