<template>
  <div class="poll__option">
    <div class="poll__vote">
      <label :for="option.name">
        <input
            :disabled="selected"
            type="radio"
            name=""
            :id="option.name"
            @click="$emit('vote', option.id)"
        >
      </label>
    </div>
    <Chart :percentage="percentage" :name="option.name" />
  </div>
</template>

<script>
import Chart from "@/components/Chart";

export default {
  name: "PollOption",
  components: {
    Chart,
  },
  data() {
    return {
    }
  },
  props: {
    option: {
      type: Object,
      required: true,
    },
    selected: {
      type: Boolean,
      required: true,
    },
    totalVotes: {
      type: Number,
      required: true,
    },
  },
  computed: {
    percentage() {
      if (this.totalVotes > 0) {
        return Math.floor((this.option.votes/this.totalVotes)*100)
      } else {
        return 0
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.poll {
  &__option {
    display: flex;
    gap: 15px;
  }
}
</style>