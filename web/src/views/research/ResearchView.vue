<script setup>
import { ref } from 'vue'
import ResearchMaterial from '@/components/research/ResearchMaterial.vue'
import ResearchDesign from '@/components/research/ResearchDesign.vue'
import ResearchAttribute from '@/components/research/ResearchAttribute.vue'

import { research } from './factor'


const researchData = ref(research)
</script>


<template>
  <div>
    <a-row>
      <a-col :span="6">
        <ResearchMaterial></ResearchMaterial>
      </a-col>
      <a-col :span="10">
        <ResearchDesign :research="researchData"></ResearchDesign>
      </a-col>
      <a-col :span="8">
        <ResearchAttribute></ResearchAttribute>
      </a-col>
    </a-row>

  </div>
</template>