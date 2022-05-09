import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "popper.js";

import "@fortawesome/fontawesome-free/css/fontawesome.css";
import "@fortawesome/fontawesome-free/css/solid.css";

import jQuery from "jquery";
window.$ = window.jQuery = jQuery;

jQuery(function () {
  jQuery('[data-toggle="popover"]').popover();
})

import { createApp } from "vue";
import App from "./App.vue";

import chart from "vue3-charts";

const app = createApp(App);
app.use(chart);
app.mount("#app");

export { app, jQuery };