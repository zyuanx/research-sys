<script setup>
import { ref } from 'vue'
const props = defineProps(['research'])
console.log('ResearchDesign', props.research)
const research = props.research
const values = props.research.values
const rules = ref([])
function conLabel(idx, label) {
  return (idx + 1) + '. ' + label
}
</script>


<template>
  <div style="padding: 20px;">
    <h1>{{ research.title }}</h1>
    <p>{{ research.description }}</p>
    <a-form ref="formRef" :model="research" :rules="rules" layout="vertical">
      <a-form-item v-for="(item, index) in research.items" :key="item.fieldID" ref="name" :label="conLabel(index, item.label)" name="name">
        <div v-if="item.factor === 'input'">
          <a-input :value="values[item.fieldID]" :placeholder="item.placeholder" />
        </div>
        <div v-else-if="item.factor === 'textarea'">
          <a-textarea :value="values[item.fieldID]" :placeholder="item.placeholder" :rows="4" />
        </div>
        <div v-else-if="item.factor === 'radio'">
          <a-radio-group :value="values[item.fieldID]">
            <a-radio v-for="(op, idx) in item.options" :key="idx" :value="op.value">{{ op.label }}</a-radio>
          </a-radio-group>
        </div>
        <div v-else-if="item.factor === 'checkbox'">
          <a-checkbox-group :value="values[item.fieldID]">
            <a-checkbox v-for="(op, idx) in item.options" :key="idx" :value="op.value">{{ op.label }}</a-checkbox>
          </a-checkbox-group>
        </div>

      </a-form-item>
    </a-form>
  </div>
</template>