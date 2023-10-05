<script setup>
import { computed, ref } from 'vue';
import { cloneDeep } from 'lodash-unified'
const props = defineProps(['research', 'rules', 'values'])
const _research = computed(() => cloneDeep(props.research))
const _rules = computed(() => props.rules)
const _values = computed(() => props.values)

const formRef = ref(null)


async function submitForm() {
  if (!formRef.value) { return }
  await formRef.value.validate((valid) => {
    if (valid) {
      console.log('submit!');
    } else {
      console.log('error submit!!');
      return false;
    }
  });
}

function resetForm() {
  if (!formRef.value) { return }
  formRef.value.resetFields();
}
</script>

<template>
  <div>
    <h1>{{ _research.config.title }}</h1>
    <p>{{ _research.config.description }}</p>
    <el-form ref="formRef" :model="_values" :rules="_rules" layout="vertical" :label-position="_research.pattern.labelPosition"
      :size="_research.pattern.size" :label-width="_research.pattern.labelWidth">
      <div v-for="(item, index) in _research.items" :key="item.fieldID" :class="['form-item', 'center-form']">
        <el-form-item ref="name" :label="(index + 1) + '. ' + (item.label)" name="name" :required="item.required" :prop="item.fieldID">
          <div v-if="item.factor === 'input'" class="form-item-div">
            <el-input v-model="_values[item.fieldID]" :placeholder="item.placeholder" :show-word-limit="item.showWordLimit" :minlength="item.minLength"
              :maxlength="item.maxLength" />
          </div>
          <div v-else-if="item.factor === 'textarea'" class="form-item-div">
            <el-input type="textarea" v-model="_values[item.fieldID]" :placeholder="item.placeholder" :rows="item.rows"
              :show-word-limit="item.showWordLimit" :minlength="item.minLength" :maxlength="item.maxLength" />
          </div>
          <div v-else-if="item.factor === 'radio'" class="form-item-div">
            <el-radio-group v-model="_values[item.fieldID]">
              <el-radio v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value">{{ op.label }}</el-radio>
            </el-radio-group>
          </div>
          <div v-else-if="item.factor === 'checkbox'" class="form-item-div">
            <el-checkbox-group v-model="_values[item.fieldID]" :min="item.min" :max="item.max">
              <el-checkbox v-for="(op, idx) in item.options" :key="idx" :id="op.value" :label="op.label" :value="op.value">
                {{ op.label }}
              </el-checkbox>
            </el-checkbox-group>
          </div>
          <div v-else-if="item.factor === 'select'" class="form-item-div">
            <el-select v-model="_values[item.fieldID]" placeholder="请选择">
              <el-option v-for="(op, idx) in item.options" :key="idx" :label="op.label" :value="op.value"></el-option>
            </el-select>
          </div>
          <div v-else-if="item.factor === 'timePicker'" class="form-item-div">
            <el-time-picker v-model="_values[item.fieldID]" arrow-control :placeholder="item.placeholder" />
          </div>
          <div v-else-if="item.factor === 'datePicker'" class="form-item-div">
            <el-date-picker v-model="_values[item.fieldID]" :type="item.type" :placeholder="item.placeholder" />
          </div>
        </el-form-item>
      </div>
      <el-form-item>
        <el-button type="primary" @click="submitForm()">Submit</el-button>
        <el-button @click="resetForm()">Reset</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style lang="scss" scoped>
.form-item {
  border-radius: 2px;
  padding: 10px;
}

.center-form {
  cursor: move;
  padding: 5px;
  margin-top: 2px;
  border-radius: 5px;

}

.form-item-div {
  width: 100%;
}
</style>