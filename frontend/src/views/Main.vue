<template>
<div>
  <div>
    <button class="btn btn-primary" @click="start_requests()">Start requests</button>
    <button class="btn btn-primary" @click="stop_requests()">Stop requests</button>
  </div>
</div>
</template>

<script>
import axios from "axios";

export default {
  methods: {
    start_requests() {
      axios.post("http://localhost:3000/?request_state=true");
    },
    stop_requests() {
      axios.post("http://localhost:3000/?request_state=false");
    },
  },
};
</script>