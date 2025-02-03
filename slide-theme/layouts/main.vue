<script>
  export default {
    methods: {
      closeSuccessMsg() {
        document.getElementById('successMsg').classList.add('hidden');
      }
    }
  }
</script>
<template>
  <div class="slidev-layout main">
    <div class="my-auto w-full">
      <slot />
    </div>
  </div>

  <div class="abs-br m-6">
    <div id="successMsg" class="w-96 bg-green-100 color-green-800 p-4 rounded-lg hidden">
      <p id="successMsgTxt"> </p>
      <button @click="closeSuccessMsg" class="absolute top-0 right-0 m-2 pr-2 text-xl font-bold">&times;</button>
    </div>
    <div id="failureMsg" class="w-96 bg-red-100 color-red-800 p-4 rounded-lg hidden">
    </div>
  </div>
</template>
