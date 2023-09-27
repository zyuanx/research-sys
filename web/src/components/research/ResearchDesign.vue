<script setup>
import { ref, computed } from 'vue'
const props = defineProps(['research', 'editIndex'])
const emit = defineEmits(['update:editIndex'])
const research = props.research
const rules = ref([])
function conLabel(idx, label) {
  return (idx + 1) + '. ' + label
}

const _editIndex = computed({
  get: () => props.editIndex,
  set: (val) => {
    emit('update:editIndex', val)
  }
})
// const dragIndex = ref(0)


</script>


<template>
  <div style="margin: 10px 5px 10px 5px;">
    <h1>{{ research.config.title }}</h1>
    <p>{{ research.config.description }}</p>
    <el-form ref="formRef" :model="research" :rules="rules" layout="vertical" :label-position="research.pattern.labelPosition"
      :size="research.pattern.size" :label-width="research.pattern.labelWidth">
      <transition-group name="drag" class="tg">
        <div v-for="(item, index) in research.items" :key="item.fieldID" draggable :class="[_editIndex === index ? 'bg-select' : '', 'form-item']"
          @click="_editIndex = index">
          <el-form-item ref="name" :label="conLabel(index, item.label)" name="name" :required="item.required">
            <div v-if="item.factor === 'input'" class="form-item-div">
              <el-input :value="item.value" :placeholder="item.placeholder" />
            </div>
            <div v-else-if="item.factor === 'textarea'" class="form-item-div">
              <el-input type="textarea" :value="item.value" :placeholder="item.placeholder" :rows="4" />
            </div>
            <div v-else-if="item.factor === 'radio'" class="form-item-div">
              <el-radio-group :value="item.value">
                <el-radio v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value">{{ op.label }}</el-radio>
              </el-radio-group>
            </div>
            <div v-else-if="item.factor === 'checkbox'" class="form-item-div">
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
.form-item {
  border-radius: 2px;
  padding: 10px;
}

.bg-select {
  border-top: 2px solid #3498db;
  background-color: #f2f6fc;
}

// .drag-move {
//   transition: transform 0.3s;
// }

// .el-input {
//   width: initial !important;
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

.form-item-div {
  width: 100%;
}
</style>