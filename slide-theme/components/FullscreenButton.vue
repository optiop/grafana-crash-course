<template>
  <button
    @click="handleClick"
    className="p-2 bg-blue-500 text-white rounded w-10
      flex items-center justify-center"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      className="h-6 w-6"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={2}
        d="M8 3H5a2 2 0 00-2 2v3m0 8v3a2 2 0 002 2h3m8-18h3a2 2 0 012 2v3m0 8v3a2 2 0 01-2 2h-3"
      />
    </svg>
  </button>
</template>

<script>
export default {
  methods: {
    handleClick() {
      const iframe = document.querySelector("iframe");
      if (iframe && iframe.requestFullscreen) {
        iframe.requestFullscreen();
      }
    }
  }
}
</script>