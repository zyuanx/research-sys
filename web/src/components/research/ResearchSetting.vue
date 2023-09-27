<script setup>
import { computed } from 'vue';
const props = defineProps(['pattern', 'config'])
const emit = defineEmits(['update:pattern', 'update:config'])
const _pattern = computed({
  get: () => props.pattern,
  set: (val) => {
    emit('update:pattern', val)
  }
})
const _config = computed({
  get: () => props.config,
  set: (val) => {
    emit('update:config', val)
  }
})
</script>
<template>
  <div style="margin: 10px 10px 10px 5px;">
    <h1>问卷设置</h1>
    <el-form-item label="Title">
      <el-input v-model="_config.title" />
    </el-form-item>
    <el-form-item label="Description">
      <el-input type="textarea" v-model="_config.description" :rows="5" />
    </el-form-item>
    <el-form-item label="表单大小">
      <el-radio-group v-model="_pattern.size">
        <el-radio-button label="large">大</el-radio-button>
        <el-radio-button label="default">中</el-radio-button>
        <el-radio-button label="small">小</el-radio-button>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="标签对齐">
      <el-radio-group v-model="_pattern.labelPosition">
        <el-radio-button label="left">左对齐</el-radio-button>
        <el-radio-button label="right">右对齐</el-radio-button>
        <el-radio-button label="top">顶部对齐</el-radio-button>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="标签宽度">
      <el-input type="text" v-model="_pattern.labelWidth"></el-input>
    </el-form-item>
    <el-form-item label="开始时间" prop="startAt">
      <el-date-picker v-model="_config.startAt" type="datetime" placeholder="选择开始时间"></el-date-picker>
    </el-form-item>
    <el-form-item label="结束时间" prop="endAt">
      <el-date-picker v-model="_config.endAt" type="datetime" placeholder="选择结束时间"></el-date-picker>
    </el-form-item>
    <el-form-item label="重复填写">
      <el-radio-group v-model="_config.once">
        <el-radio :label="1">是</el-radio>
        <el-radio :label="0">否</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="开放填写">
      <el-radio-group v-model="_config.open">
        <el-radio :label="1">是</el-radio>
        <el-radio :label="0">否</el-radio>
      </el-radio-group>
      <el-tag type="warning" size="small" style="margin-left: 10px;">发布后无法修改</el-tag>
    </el-form-item>
  </div>
</template>

<style lang="scss" scoped></style>