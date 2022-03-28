<template>
  <br />
  ECU/Module type:
  <select class="select">
    <option value="">--Select one--</option>
    <option v-for="(ecuType, index) in ecuTypes" v-bind:key="index">{{ecuType}}</option>
  </select>
  <hr />
</template>

<script lang="ts">
export default {
  name: "app",
  data() {
    return {
      ecuTypes: {
        "1.x": "MEMS 1.2, 1.3 or 1.6",
        "1.9": "MEMS 1.9",
        "2J": "MEMS 2J",
        "3": "MEMS 3",
        "rc5": "Airbag RC5"
      },
    };
  },
};
</script>

<style lang="scss"></style>
