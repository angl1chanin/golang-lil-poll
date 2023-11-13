<template>
  <div class="main">
    <div class="main__container container">
      <h1 class="main__title">Lil todo</h1>
      <div class="main__chat">
        <Poll/>
      </div>
    </div>
  </div>
</template>

<script>
// styles
import "mosha-vue-toastify/dist/style.css";
import "@/assets/style/reset.css";

// components
import Poll from "@/components/Poll";

export default {
  name: 'App',
  components: {
    Poll,
  }
}
</script>

<style lang="scss">
html {
  font-size: 10px;
}

body {
  background: linear-gradient(207deg, #F4FBF1 0%, #FFF4F1 100%) center / cover;
  font-family: 'Inter', sans-serif;
  font-size: 14px;
  height: 100vh;
}

.container {
  width: min(100% - 3rem, 120rem);
  margin-inline: auto;
}

.main {
  &__container {
    height: 100vh;

    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  &__title {
    font-size: 36px;
    font-weight: 500;
    text-transform: capitalize;
    text-align: center;
    margin-bottom: 100px;
  }

  &__chat {
    width: 386px;
  }
}
</style>
