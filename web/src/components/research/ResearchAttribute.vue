<script setup>
import { computed } from 'vue';
import {
  Delete,
  Plus, Minus
} from '@element-plus/icons-vue'

import { factorRules, factorRulesMap } from '@/utils/factor'

const props = defineProps(['researchItem'])
const emit = defineEmits(['update:researchItem'])
const _researchItem = computed({
  get: () => props.researchItem,
  set: (val) => {
    emit('update:researchItem', val)
  }
})

const checkFactor = computed(() => {
  return ['radio', 'checkbox', 'select'].includes(_researchItem.value.factor)
})

const hasPlaceholder = computed(() => {
  return ['input', 'textarea'].includes(_researchItem.value.factor)
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
    ...factorRules[0]
  })
}

function ruleChange(index, val) {
  // console.log('ruleChange')
  _researchItem.value.rules[index] = {
    ...factorRulesMap[val]
  }
}

</script>


<template>
  <div style="margin: 10px 10px 10px 5px;">
    <h1>组件属性</h1>
    <span>{{ _researchItem }}</span>
    <h2>基本</h2>
    <el-form-item label="Id">
      <el-input :value="_researchItem.fieldID" disabled :bordered="false" />
    </el-form-item>
    <el-form-item label="Label">
      <el-input v-model="_researchItem.label" />
    </el-form-item>
    <el-form-item label="Placeholder" v-if="hasPlaceholder">
      <el-input v-model="_researchItem.placeholder" />
    </el-form-item>
    <el-form-item label="Min" v-if="_researchItem.factor === 'checkbox'">
      <el-input-number v-model="_researchItem.min" :min="0" :max="10" />
    </el-form-item>
    <el-form-item label="Max" v-if="_researchItem.factor === 'checkbox'">
      <el-input-number v-model="_researchItem.max" :min="1" />
    </el-form-item>
    <el-form-item label="MinLength" v-if="hasPlaceholder">
      <el-input-number v-model="_researchItem.minLength" :min="0" :max="10" />
    </el-form-item>
    <el-form-item label="MaxLength" v-if="hasPlaceholder">
      <el-input-number v-model="_researchItem.maxLength" :min="1" />
    </el-form-item>
    <el-form-item label="showWordLimit" v-if="hasPlaceholder">
      <el-switch v-model="_researchItem.showWordLimit" />
    </el-form-item>
    <el-form-item label="type" v-if="_researchItem.factor === 'datePicker'">
      <el-radio-group v-model="_researchItem.type">
        <el-radio-button label="date">date</el-radio-button>
        <el-radio-button label="month">month</el-radio-button>
        <el-radio-button label="year">year</el-radio-button>
      </el-radio-group>
    </el-form-item>
    <h2 v-if="checkFactor">选项</h2>
    <div v-if="checkFactor">
      <div v-for="(option, index) in _researchItem.options" :key="index" class="options">
        <el-input v-model="option.label" placeholder="label" />
        <span style="margin: 2.5px;"></span>
        <el-input v-model="option.value" placeholder="value" />
        <span style="margin: 2.5px;"></span>
        <el-button type="danger" :icon="Minus" circle size="small" @click="deleteOption(index)" />
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
        <el-divider />
        <div style="display: flex;justify-content: space-between;">
          <el-form-item label="Label">
            <el-select v-model="rule.id" style="width: 120px" @change="ruleChange(index, $event)">
              <el-option v-for="(options, idx) in factorRules" :key="idx" :label="options.label" :value="options.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-button type="danger" :icon="Delete" @click="_researchItem.rules.splice(index, 1)" />
        </div>
        <el-form-item label="Reg">
          <el-input v-model="rule.reg" :disabled="rule.id !== 'custom'" />
        </el-form-item>
        <el-form-item label="Tip">
          <el-input v-model="rule.tip" :disabled="rule.id !== 'custom'" />
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
.options {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 5px;
}
</style>
