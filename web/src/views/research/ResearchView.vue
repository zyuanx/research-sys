<script setup>
import { ref } from 'vue'
import { v4 } from 'uuid'
import ResearchMaterial from '@/components/research/ResearchMaterial.vue'
import ResearchDesign from '@/components/research/ResearchDesign.vue'
import ResearchAttribute from '@/components/research/ResearchAttribute.vue'

import { research, factorItems, } from '@/utils/factor'


const researchData = ref(research)
const editIndex = ref(0)

function setEditIdx(idx) {
  editIndex.value = idx
}

function itemAdd(item) {
  const _item = factorItems[item]
  _item['fieldID'] = v4()
  researchData.value.items.push(_item)
}

// function updateResearch(data) {
//   console.log('updateResearch', data)
//   researchData.value = data
// }
</script>


<template>
  <div>
    <el-row>
      <el-col :span="6">
        <ResearchMaterial @item-add="itemAdd"></ResearchMaterial>
      </el-col>
      <el-col :span="10">
        <ResearchDesign :research="researchData" @set-edit-idx="setEditIdx" :editIndex="editIndex"></ResearchDesign>
      </el-col>
      <el-col :span="8">
        <ResearchAttribute :researchItem="researchData.items[editIndex]"></ResearchAttribute>
      </el-col>
    </el-row>

  </div>
</template>