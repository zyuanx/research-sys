<script setup>
import { ref } from 'vue'
import { v4 } from 'uuid'
import ResearchMaterial from '@/components/research/ResearchMaterial.vue'
import ResearchDesign from '@/components/research/ResearchDesign.vue'
import ResearchAttribute from '@/components/research/ResearchAttribute.vue'
import ResearchSetting from '@/components/research/ResearchSetting.vue'

import { research, factorItems, } from '@/utils/factor'


const researchData = ref(research)
const editIndex = ref(0)


function itemAdd(item) {
  const _item = factorItems[item]
  _item['fieldID'] = v4()
  researchData.value.items.push(_item)
}
const researchSetting = ref(false)
</script>


<template>
  <div>
    <el-row>
      <el-col :span="6">
        <ResearchMaterial class="container" @item-add="itemAdd"></ResearchMaterial>
      </el-col>
      <el-col :span="10">
        <div style="display: flex;justify-content: end;align-items: center;">
          <el-button type="primary" text>primary</el-button>
          <el-button type="primary" text @click="researchSetting = !researchSetting">{{ researchSetting ? '设置' : '属性' }}</el-button>
        </div>
        <ResearchDesign :research="researchData" v-model:editIndex="editIndex" style="overflow-y: auto;height: 80vh;"></ResearchDesign>
      </el-col>
      <el-col :span="8">
        <ResearchSetting v-if="researchSetting" class="container" v-model:pattern="researchData.pattern" v-model:config="researchData.config">
        </ResearchSetting>
        <ResearchAttribute v-else class="container" v-model:researchItem="researchData.items[editIndex]" style="overflow-y: auto;height: 80vh;">
        </ResearchAttribute>
      </el-col>
    </el-row>

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