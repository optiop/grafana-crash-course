<script setup lang="ts">
import { computed } from 'vue'
import { onMounted } from 'vue'
import { handleBackground } from './layoutHelper'

const props = defineProps({
  background: {
    default: '',
  },
})

const style = computed(() => handleBackground(props.background, true))
</script>

<template>
  <div class="slidev-layout cover" :style="style">
    <div class="my-auto w-full">
      <div class="w-10/12 mx-auto bg-black bg-opacity-70 rounded-lg p-10">
        <slot />
      </div>
    </div>
  </div>
</template>
