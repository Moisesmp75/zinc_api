<template>
  <div class="flex justify-center items-center mt-8 gap-8">
    <button
      @click="prevPage"
      :disabled="!meta_data.HasPrevPage"
      class="py-2 px-4 bg-green-500 text-white rounded disabled:opacity-50"
    >Prev</button>
    <span class="text-sm">
      Página
      <input type="number" v-model="meta_data.CurrentPage" min="1" :max="meta_data.TotalPages" class="text-center" @keyup.enter="inputPage">
      de {{ meta_data.TotalPages }}
    </span> 
    <button
      @click="nextPage"
      :disabled="!meta_data.HasNextPage"
      class="py-2 px-4 bg-green-500 text-white rounded disabled:opacity-50"
      >Next</button>
  </div>
</template>
<script>
export default {
  props: {
    meta_data: {
      type: Object,
      required: true
    },
    prevPage: {
      type: Function,
      required: true
    },
    nextPage: {
      type: Function,
      required: true
    },
    inputPage: {
      type: Function,
      required: true
    }
  },
}
</script>
<style lang="">
  
</style>