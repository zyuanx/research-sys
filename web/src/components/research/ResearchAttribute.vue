<script setup>
import { ref, onMounted } from 'vue';

const props = defineProps(['research', 'editIndex'])
const research = props.research
console.log('ResearchAttribute', research.items)
const item = research.items[props.editIndex]
console.log('ResearchAttributeItem', item)


</script>


<template>
  <div class="container" style="margin: 10px 10px 10px 5px;">
    <h1>组件属性</h1>
    <a-form-item label="Id">
      <a-input :value="item.fieldID" disabled :bordered="false" />
    </a-form-item>
  </div>
</template>