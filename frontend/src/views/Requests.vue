<template>
  <div>
      <table class="table table-striped table-sm">
    <thead>
      <tr>
        <th>Timestamp</th>
        <th>Request time(ns)</th>
        <th>Url</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="request of requests" :key="request">
        <td>{{request.Timestamp}}</td>
        <td>{{request.Req_time}}</td>
        <td>{{request.Url}}</td>
      </tr>
    </tbody>
  </table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      requests: [],
    };
  },

  // Fetches posts when the component is created.
  created() {
    axios
      .get('http://localhost:3000/requests')
      .then((response) => {
        // JSON responses are automatically parsed
        console.log(response.data)
        this.requests = response.data.data
      })
  },
};
</script>