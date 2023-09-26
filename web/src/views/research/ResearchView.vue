<script setup>
import { ref } from 'vue'
import { v4 } from 'uuid'
import ResearchMaterial from '@/components/research/ResearchMaterial.vue'
import ResearchDesign from '@/components/research/ResearchDesign.vue'
import ResearchAttribute from '@/components/research/ResearchAttribute.vue'

import { research, factorItems, factorValues } from './factor'


const researchData = ref(research)
const editIndex = ref(0)

function setEditIdx(idx) {
  editIndex.value = idx
}

function itemAdd(item) {
  console.log('itemAdd', item)
  const _item = factorItems[item]
  console.log(_item)
  _item['fieldID'] = v4()
  researchData.value.items.push(_item)
  const _value = factorValues[item]
  researchData.value.values[_item.fieldID] = _value
}

</script>


<template>
  <div>
    <a-row>
      <a-col :span="6">
        <ResearchMaterial @item-add="itemAdd"></ResearchMaterial>
      </a-col>
      <a-col :span="10">
        <ResearchDesign :research="researchData" @set-edit-idx="setEditIdx" :editIndex="editIndex"></ResearchDesign>
      </a-col>
      <a-col :span="8">
        <ResearchAttribute :research="researchData" :editIndex="editIndex"></ResearchAttribute>
      </a-col>
    </a-row>

  </div>
</template>