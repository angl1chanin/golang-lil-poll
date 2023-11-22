<template>
  <div class="poll__body">
    <div class="poll__header">
      <p class="poll__option-name">{{ name }}</p>
      <p class="poll__score">{{ percentage }}%</p>
    </div>
    <div class="poll__scorebar">
      <span :style="percentageWidth"></span>
    </div>
  </div>
</template>

<script setup>
import {computed} from "vue";

const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  percentage: {
    type: Number,
    required: true,
  }
})

const percentageWidth = computed(() => {
  return `width: ${props.percentage}%`
})
</script>

<style lang="scss" scoped>
.poll {
  &__body {
    width: 100%;
  }

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 20px;
  }

  &__scorebar {
    width: 100%;
    height: 10px;
    border-radius: 5px;
    background-color: #EDEEF0;
    position: relative;
    overflow: hidden;

    span {
      position: absolute;
      top: 0;
      bottom: 0;
      left: 0;
      background-color: #1FCC92;
    }
  }
}
</style>