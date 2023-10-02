<script setup>
import { ref, computed } from 'vue'
import { VueDraggable } from 'vue-draggable-plus'
import { Edit, Delete } from '@element-plus/icons-vue'
const props = defineProps(['research', 'editIndex'])
const emit = defineEmits(['update:research', 'update:editIndex'])
const rules = ref([])

const _research = computed({
  get: () => props.research,
  set: (val) => {
    console.log("update:research", val)
    emit('update:research', val)
  }
})
const _editIndex = computed({
  get: () => props.editIndex,
  set: (val) => {
    console.log("update:editIndex", val)
    emit('update:editIndex', val)
  }
})

function remove(index) {
  if (_research.value.items.length <= 1) {
    return
  }
  if (_research.value.items.length - 1 === index) {
    _editIndex.value = index - 1
  } else {
    _editIndex.value = index
  }
  _research.value.items.splice(index, 1)
}

</script>


<template>
  <div>
    <h1>{{ _research.config.title }}</h1>
    <p>{{ _research.config.description }}</p>
    <el-form ref="formRef" id="editForm" :model="_research" :rules="rules" layout="vertical" :label-position="_research.pattern.labelPosition"
      :size="_research.pattern.size" :label-width="_research.pattern.labelWidth">
      <VueDraggable v-model="_research.items" animation="150">
        <div v-for="(item, index) in _research.items" :key="item.fieldID"
          :class="[_editIndex === index ? 'bg-select' : '', 'form-item', 'center-form']">
          <div style="width: 100%;text-align: right;">
            <el-button type="primary" :icon="Edit" size="small" circle @click.stop="_editIndex = index"></el-button>
            <el-button type="danger" :icon="Delete" size="small" circle @click.stop="remove(index)"></el-button>
          </div>
          <el-form-item ref="name" :label="(index + 1) + '. ' + (item.label)" name="name" :required="item.required">

            <div v-if="item.factor === 'input'" class="form-item-div">
              <el-input :value="item.value" :placeholder="item.placeholder" :show-word-limit="item.showWordLimit" :minlength="item.minLength"
                :maxlength="item.maxLength" />
            </div>
            <div v-else-if="item.factor === 'textarea'" class="form-item-div">
              <el-input type="textarea" :value="item.value" :placeholder="item.placeholder" :rows="item.rows" :show-word-limit="item.showWordLimit"
                :minlength="item.minLength" :maxlength="item.maxLength" />
            </div>
            <div v-else-if="item.factor === 'radio'" class="form-item-div">
              <el-radio-group :value="item.value">
                <el-radio v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value">{{ op.label }}</el-radio>
              </el-radio-group>
            </div>
            <div v-else-if="item.factor === 'checkbox'" class="form-item-div">
              <el-checkbox-group :value="item.value" :min="item.min" :max="item.max">
                <el-checkbox v-for="(op, idx) in item.options" :key="idx" :id="op.value" :label="op.label" :value="op.value">{{ op.label
                }}</el-checkbox>
              </el-checkbox-group>
            </div>
            <div v-else-if="item.factor === 'select'" class="form-item-div">
              <el-select :value="item.value" placeholder="请选择">
                <el-option v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value"></el-option>
              </el-select>
            </div>
            <div v-else-if="item.factor === 'timePicker'" class="form-item-div">
              <el-time-picker :value="item.value" arrow-control :placeholder="item.placeholder" />
            </div>
            <div v-else-if="item.factor === 'datePicker'" class="form-item-div">
              <el-date-picker :value="item.value" :type="item.type" :placeholder="item.placeholder" />
            </div>
          </el-form-item>
        </div>

      </VueDraggable>
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

.center-form {
  cursor: move;
  padding: 5px;
  margin-top: 2px;
  border-radius: 5px;

  &:hover {
    background-color: #ebfaeb;
  }
}

.form-item-div {
  width: 100%;
}
</style>