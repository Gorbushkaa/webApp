<template>
  <div>
    <form v-on:submit.prevent="submitForm">
      <div class="form-group">
        <label for="timeout">Set timeout (sec)</label>
        <input type="text" class="form-control" id="timeout" placeholder="Request timeout" v-model="form.timeout">
      </div>
      <div class="form-group">
        <button class="btn btn-primary">Submit</button>
      </div>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "PostFormAxios",
  data() {
    return {
      form: {
        timeout: '',
      }
    }
  },
  methods: {
    submitForm() {
      axios
        .post(`http://${process.env.VUE_APP_BACKEND_URL}/settings`, this.form)
        .then((response) => {
            console.log(response)
        })
    }
  }
};
</script>