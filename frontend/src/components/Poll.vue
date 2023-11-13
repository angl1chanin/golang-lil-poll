<template>
  <div class="poll app-window">
    <div class="poll__wrapper">
      <p class="poll__title title">Poll</p>
      <Preloader v-if="isLoading" />
      <div v-else-if="poll != null">
        <p class="poll__name">{{ poll.name }}</p>
        <ul class="poll__options">
          <li v-for="option in poll.options" :key="option.id">
            <Option
                :option="option"
                :selected="selected"
                @vote="vote"
                :total-votes="poll.total_votes"
            />
          </li>
        </ul>
      </div>
      <p v-else class="empty">No polls yet</p>
    </div>
  </div>
</template>

<script>
// components
import Option from "@/components/Option";
import Preloader from "@/components/Preloader";
import notices from "@/mixins/notices";

export default {
  name: "Poll",
  components: {
    Option,
    Preloader
  },
  mixins: [
    notices
  ],
  data() {
    return {
      api: "http://127.0.0.1:3301/v1/poll",
      isLoading: false,
      selected: false,
      poll: null,
    }
  },
  methods: {
    loadPoll() {
      this.isLoading = true
      fetch(this.api+"/1")
        .then(response => {
          if (response.ok) {
            this.isLoading = false
            return response.json()
          } else {
            this.danger("Error while loading poll")
          }
        })
        .then(json => {
          this.poll = json
        })
        .catch(error => {
          console.log("Poll server isn't available")
        })
    },
    updatePoll() {
      fetch(this.api+"/1")
          .then(response => {
            if (response.ok) {
              return response.json()
            } else {
              this.danger("Error while loading poll")
            }
          })
          .then(json => {
            this.poll = json
          })
          .catch(error => {
            console.log("Poll server isn't available")
          })
    },
    updateSelectState() {
      this.selected = true
    },
    vote(id) {
      const config = {
        method: "POST",
      }

      fetch(this.api+"/option/vote/"+id, config)
          .then(response => {
            if (response.ok) {
              this.success("You have successfully voted")
            } else {
              this.danger("Error while loading poll")
            }
          })
          .catch(error => {
            console.log("Poll not available")
          })
      this.updateSelectState()
      this.updatePoll()
    }
  },
  mounted() {
    this.loadPoll()
  }
}
</script>

<style lang="scss" scoped>
.app-window {
  background-color: #fff;
  padding: 30px;
  border-radius: 32px;
  box-shadow: 0px 2px 8px 0px rgba(34, 60, 80, 0.2);
}

.button {
  background: rgba(118, 118, 128, 0.12);
  color: var(--text-color);
  border: none;
  font-family: 'Poppins', sans-serif;
  font-size: 16px;
  font-weight: 400;
  text-align: center;
  border-radius: 12px;
  line-height: 22px;
  cursor: pointer;
  padding: 10px 0;
}

.title {
  color: var(--text-color);
  font-size: 18px;
  font-weight: 500;
  text-transform: capitalize;
}

.poll {
  &__title {
    margin-bottom: 35px;
  }

  &__name {
    margin-bottom: 15px;
    font-weight: 600;
  }

  &__options {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
}
</style>