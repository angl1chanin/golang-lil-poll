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
import Option from "@/components/Poll/Option";
import Preloader from "@/components/Preloader";
import notifications from "@/mixins/notifications";

export default {
  name: "Poll",
  components: {
    Option,
    Preloader
  },
  mixins: [
      notifications
  ],
  data() {
    return {
      api: "https://d952-2a00-1fa0-c207-3373-9919-d58b-f508-f162.ngrok-free.app",
      isLoading: false,
      selected: false,
      poll: null,
    }
  },
  methods: {
    loadPoll() {
      const myHeaders = new Headers();
      myHeaders.append("Ngrok-Skip-Browser-Warning", "123");

      const config = {
        headers: myHeaders,
      }

      this.isLoading = true
      fetch(this.api+"/poll/2", config)
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
      const myHeaders = new Headers();
      myHeaders.append("Ngrok-Skip-Browser-Warning", "123");

      const config = {
        headers: myHeaders,
      }

      fetch(this.api+"/poll/2", config)
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
      const myHeaders = new Headers();
      myHeaders.append("Ngrok-Skip-Browser-Warning", "123");

      const config = {
        method: "POST",
        headers: myHeaders,
      }

      fetch(this.api+"/poll/vote/"+id, config)
          .then(response => {
            if (response.ok) {
              this.success("Vote sent")
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
  }

  &__options {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
}
</style>