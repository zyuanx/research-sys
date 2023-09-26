<script setup>
import { ref } from 'vue'
const props = defineProps(['research', 'editIndex'])
const emit = defineEmits(['setEditIdx'])
const research = props.research
const rules = ref([])
function conLabel(idx, label) {
  return (idx + 1) + '. ' + label
}
// const dragIndex = ref(0)

function setEditIdx(idx) {
  emit('setEditIdx', idx)
}
</script>


<template>
  <div class="container" style="margin: 10px 5px 10px 5px;">
    <h1>{{ research.title }}</h1>
    <p>{{ research.description }}</p>
    <el-form ref="formRef" :model="research" :rules="rules" layout="vertical">
      <transition-group name="drag">
        <div v-for="(item, index) in research.items" :key="item.fieldID" draggable :class="[editIndex === index ? 'bg-select' : '', 'form-item']"
          @click="setEditIdx(index)">
          <el-form-item ref="name" :label="conLabel(index, item.label)" name="name" :required="item.required">
            <div v-if="item.factor === 'input'">
              <el-input :value="item.value" :placeholder="item.placeholder" />
            </div>
            <div v-else-if="item.factor === 'textarea'">
              <el-input type="textarea" :value="item.value" :placeholder="item.placeholder" :rows="4" />
            </div>
            <div v-else-if="item.factor === 'radio'">
              <el-radio-group :value="item.value">
                <el-radio v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value">{{ op.label }}</el-radio>
              </el-radio-group>
            </div>
            <div v-else-if="item.factor === 'checkbox'">
              <el-checkbox-group :value="item.value">
                <el-checkbox v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value">{{ op.label }}</el-checkbox>
              </el-checkbox-group>
            </div>
          </el-form-item>
        </div>

      </transition-group>
    </el-form>

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