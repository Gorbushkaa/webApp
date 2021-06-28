<template>
  <section class="container">
    <h1>Comparison of request times</h1>
    <div class="columns">
      <div class="column">
        <bar-chart :chart-data="datacollection"></bar-chart>
      </div>
    </div>
  </section>
</template>

<script>
import BarChart from "../components/BarChart.vue";
import axios from "axios";

export default {
  components: {
    BarChart,
  },
  data() {
    return {
      datacollection: null,
    };
  },
  name: "VueChartJS",
  mounted() {
    this.fillData();
  },

  methods: {
    async fillData() {
      await axios
        .get(`http://${process.env.VUE_APP_BACKEND_URL}/get_last_requests`)
        .then((response) => {
          this.datacollection = {
            datasets: [
              {
                data: [response.data.delay_request],
                label: "Delay request",
                borderColor: "#3cba9f",
                backgroundColor: "#3cba9f",
                borderWidth: 2,
              },
              {
                data: [response.data.status_request],
                label: "Request 200",
                borderColor: "#ffa500",
                backgroundColor: "#ffa500",
                borderWidth: 2,
              },
            ],
          };
        });
    },
  },
};
</script>