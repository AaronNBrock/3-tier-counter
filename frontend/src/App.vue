<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Count: {{ this.count }}</h1>
        <div v-if="error" class="alert alert-danger" role="alert">
          <strong>{{ error }}</strong>
        </div>
        <button v-on:click="increment" class="btn btn-primary mr-1">
          increment
        </button>
        <button v-on:click="getCount" class="btn btn-light mr-1">sync</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",

  data() {
    return {
      error: null,
      count: 0,
    };
  },

  methods: {
    getCount() {
      axios
        .get("http://" + config.API_ENDPOINT + "/api/count")
        .then((response) => {
          this.count = response.data.count;
          this.error = null;
        })
        .catch((error) => {
          this.error = error;
        });
    },

    increment() {
      this.count += 1;
      axios
        .post("http://" + config.API_ENDPOINT + "/api/count")
        .then((response) => {
          this.count = response.data.count;
          this.error = null;
        })
        .catch((error) => {
          this.error = error;
        });
    },
  },

  mounted() {
    this.getCount();
    console.log("API_ENDPOINT: " + config.API_ENDPOINT);

    window.setInterval(() => {
      this.getCount();
    }, 5000);
  },
};
</script>

<style>
body {
  background-color: purple;
}
</style>
