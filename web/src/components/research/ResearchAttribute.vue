<script setup>
import { computed } from 'vue';
import {
  Delete,
  Plus
} from '@element-plus/icons-vue'

import { factorRules } from '@/utils/factor'

const props = defineProps(['researchItem'])
const emit = defineEmits(['updateResearch'])
const _researchItem = computed({
  get: () => props.researchItem,
  set: (val) => {
    emit('updateResearch', val)
  }
})

const checkFactor = computed(() => {
  return ['radio', 'checkbox', 'select'].includes(_researchItem.value.factor)
})

function deleteOption(index) {
  _researchItem.value.options.splice(index, 1)
}
function addOption() {
  _researchItem.value.options.push({
    label: '选项xx',
    value: '值xx'
  })
}

function addRuleOption() {
  _researchItem.value.rules.push({
    label: '',
    reg: '',
    tip: ''
  })
}


</script>


<template>
  <div class="container" style="margin: 10px 10px 10px 5px;">
    <h1>组件属性</h1>
    <h2>基本</h2>
    <el-form-item label="Id">
      <el-input :value="_researchItem.fieldID" disabled :bordered="false" />
    </el-form-item>
    <el-form-item label="Title">
      <el-input v-model="_researchItem.label" />
    </el-form-item>
    <el-form-item label="Placeholder">
      <el-input v-model="_researchItem.placeholder" />
    </el-form-item>
    <h2 v-if="checkFactor">选项</h2>
    <div v-if="checkFactor">
      <div v-for="(option, index) in _researchItem.options" :key="index" class="options">
        <el-input v-model="option.label" placeholder="label" />
        <span style="margin: 2.5px;"></span>
        <el-input v-model="option.value" placeholder="value" />
        <span style="margin: 2.5px;"></span>
        <el-button type="danger" :icon="Delete" circle @click="deleteOption(index)" />
      </div>
      <div>
        <el-button link type="primary" :icon="Plus" @click="addOption">
          Add Option
        </el-button>
      </div>
    </div>
    <h2>验证规则</h2>
    <el-form-item label="Required">
      <el-switch v-model="_researchItem.required" />
    </el-form-item>
    <div>
      <div v-for="(rule, index) in _researchItem.rules" :key="index">
        <el-select v-model="rule.label" style="width: 120px">
          <el-option v-for="(options, idx) in factorRules" :key="idx" :label="options.label" :value="options.label">
          </el-option>
        </el-select>
        <el-form-item label="Reg">
          <el-input v-model:value="rule.reg" />
        </el-form-item>
        <el-form-item label="Tip">
          <el-input v-model:value="rule.tip" />
        </el-form-item>
      </div>
      <div>
        <el-button link type="primary" :icon="Plus" @click="addRuleOption">
          Add Rule Option
        </el-button>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.container {
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  padding: 10px;
  height: 100%;

  .form-item {
    border-radius: 2px;
    padding: 10px;
  }

  .bg-select {
    border-top: 2px solid #3498db;
    background-color: #f2f6fc;
  }
}

.options {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 5px;
}
</style>
