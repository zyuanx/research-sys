<script setup>
import { ref, defineEmits } from 'vue'
const props = defineProps(['research', 'editIndex'])
const emit = defineEmits(['setEditIdx'])
console.log('ResearchDesign', props.research)
const research = props.research
const values = props.research.values
const rules = ref([])
function conLabel(idx, label) {
  return (idx + 1) + '. ' + label
}
const dragIndex = ref(0)

function setEditIdx(idx) {
  emit('setEditIdx', idx)
}
</script>


<template>
  <div class="container" style="margin: 10px 5px 10px 5px;">
    <h1>{{ research.title }}</h1>
    <p>{{ research.description }}</p>
    <a-form ref="formRef" :model="research" :rules="rules" layout="vertical">
      <transition-group name="drag">
        <div v-for="(item, index) in research.items" :key="item.fieldID" draggable :class="[editIndex === index ? 'bg-select' : '', 'form-item']"
          @click="setEditIdx(index)">
          <a-form-item ref="name" :label="conLabel(index, item.label)" name="name">
            <div v-if="item.factor === 'input'">
              <a-input :value="values[item.fieldID]" :placeholder="item.placeholder" />
            </div>
            <div v-else-if="item.factor === 'textarea'">
              <a-textarea :value="values[item.fieldID]" :placeholder="item.placeholder" :rows="4" />
            </div>
            <div v-else-if="item.factor === 'radio'">
              <a-radio-group :value="values[item.fieldID]">
                <a-radio v-for="(op, idx) in item.options" :key="idx" :value="op.value">{{ op.label }}</a-radio> </a-radio-group>
            </div>
            <div v-else-if="item.factor === 'checkbox'">
              <a-checkbox-group :value="values[item.fieldID]">
                <a-checkbox v-for="(op, idx) in item.options" :key="idx" :value="op.value">{{ op.label }}</a-checkbox>
              </a-checkbox-group>
            </div>
          </a-form-item>
        </div>

      </transition-group>
    </a-form>

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

// .drag-move {
//   transition: transform 0.3s;
// }


.bg-drag {
  border: 1.5px dashed #909399;
}

.center-form {
  cursor: move;
  padding: 5px;
  margin-top: 2px;
  border-radius: 5px;

  &:hover {
    background-color: #f2f6fc;
  }
}
</style>