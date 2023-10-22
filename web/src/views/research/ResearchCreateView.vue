<script setup>
import { ref } from 'vue'
import { v4 as uuidv4 } from 'uuid'
import ResearchMaterial from '@/components/research/ResearchMaterial.vue'
import ResearchDesign from '@/components/research/ResearchDesign.vue'
import ResearchAttribute from '@/components/research/ResearchAttribute.vue'
import ResearchSetting from '@/components/research/ResearchSetting.vue'
import { Setting, Operation } from '@element-plus/icons-vue'

import { research, factorItems } from '@/utils/factor'

import { createResearch } from '@/api/research'
import ResearchPreview from '@/components/research/ResearchPreview.vue'
const researchData = ref(research)
const editIndex = ref(0)


function itemAdd(item) {
  const _item = { ...factorItems[item] }
  _item['fieldID'] = uuidv4()
  researchData.value.items.push(_item)
}
const researchSetting = ref(false)

const drawer = ref(false)

const researchRules = ref({})
const researchValues = ref({})


function previewResearch() {
  for (const item of researchData.value.items) {
    researchRules.value[item.fieldID] = []
    if (item.required) {
      if (['input', 'textarea'].includes(item.factor)) {
        researchRules.value[item.fieldID].push({ required: true, message: '请输入' + item.label, trigger: 'blur' })
      } else if (['radio', 'select'].includes(item.factor)) {
        researchRules.value[item.fieldID].push({ required: true, message: '请选择' + item.label, trigger: 'change' })
      } else if (['checkbox'].includes(item.factor)) {
        researchRules.value[item.fieldID].push({ type: 'array', required: true, message: '请选择' + item.label, trigger: 'change' })
      } else if (['timePicker', 'datePicker'].includes(item.factor)) {
        researchRules.value[item.fieldID].push({ type: 'date', required: true, message: '请选择' + item.label, trigger: 'change' })
      } else {
        researchRules.value[item.fieldID].push({ required: true, message: '请输入' + item.label, trigger: 'blur' })
      }
    }
    if (item.minLength) {
      researchRules.value[item.fieldID].push({ min: item.minLength, message: '最少输入' + item.minLength + '个字符', trigger: 'blur' })
    }
    if (item.maxLength) {
      researchRules.value[item.fieldID].push({ max: item.maxLength, message: '最多输入' + item.maxLength + '个字符', trigger: 'blur' })
    }
    researchValues.value[item.fieldID] = item.value
  }
  drawer.value = true
}

async function CreateResearch() {
  const payload = {
    title: researchData.value.config.title,
    description: researchData.value.config.description,
    pattern: researchData.value.pattern,
    config: researchData.value.config,
    items: researchData.value.items,
    once: researchData.value.config.once,
    open: researchData.value.config.open,
  }
  const start = new Date(researchData.value.config.startAt)
  const end = new Date(researchData.value.config.endAt)
  if (start > end) {
    return
  }
  payload.startAt = start.toISOString()
  payload.endAt = end.toISOString()
  const res = await createResearch(payload)
  console.log(res)
}
</script>


<template>
  <div>
    <el-row>
      <el-col :span="6">
        <ResearchMaterial class="container" @item-add="itemAdd"></ResearchMaterial>
      </el-col>
      <el-col :span="10">
        <div class="container" style="margin: 10px 5px 10px 5px;">
          <div class="tools" style="display: flex;justify-content: end;align-items: center;">
            <el-button type="primary" text @click="previewResearch">预览</el-button>
            <el-button type="warning" text @click="CreateResearch">发布</el-button>
            <el-button :type="researchSetting ? 'primary' : 'success'" @click="researchSetting = !researchSetting"
              :icon="researchSetting ? Setting : Operation">
              {{ researchSetting ? '设置' : '属性' }}
            </el-button>
          </div>

          <ResearchDesign :research="researchData" v-model:editIndex="editIndex" style="overflow-y: auto;height: 80vh;"></ResearchDesign>
        </div>

      </el-col>
      <el-col :span="8">
        <ResearchSetting v-if="researchSetting" class="container" v-model:pattern="researchData.pattern" v-model:config="researchData.config">
        </ResearchSetting>
        <ResearchAttribute v-else class="container" v-model:researchItem="researchData.items[editIndex]" style="overflow-y: auto;height: 80vh;">
        </ResearchAttribute>
      </el-col>
    </el-row>
    <el-drawer v-model="drawer" title="Preview" :with-header="false">
      <ResearchPreview :research="researchData" :rules="researchRules" :values="researchValues"></ResearchPreview>
    </el-drawer>
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