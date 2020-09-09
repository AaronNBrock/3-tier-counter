<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Count: {{ this.count }}</h1>
          <button v-on:click="increment" class="btn btn-primary">increment</button>
          <button v-on:click="getCount" class="btn btn-light">sync</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',

  data() { return {
    count: 0,
  } },

  methods: {
    getCount() {
      axios.get("http://" + config.API_ENDPOINT + "/api/count")
        .then((response) => {
          this.count = response.data.count;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        })
    },

    increment() {
      axios.post("http://" + config.API_ENDPOINT + "/api/count")
        .then((response) => {
          this.count = response.data.count;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        })
    }
  },

  mounted() {
    this.getCount()
    console.log("API_ENDPOINT: " + config.API_ENDPOINT)
  }
}
</script>
