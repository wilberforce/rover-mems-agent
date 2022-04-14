import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "popper.js";

import "@fortawesome/fontawesome-free/css/fontawesome.css";
import "@fortawesome/fontawesome-free/css/solid.css";

import jQuery from "jquery";
window.$ = window.jQuery = jQuery;

/*
import "./globals";
import "./helpers";
import "./output";
import "./script";
import "./canvas";
*/
import "./agent";
import "./websocket";
import { ecuTypeSelectChanged, runUserAction } from "./agent.js";
//import { downloadData } from "./output.js";
//window.downloadData=downloadData;
window.runUserAction=runUserAction;
window.ecuTypeSelectChanged=ecuTypeSelectChanged;

const ecuConnectedElement = document.getElementById('ecuConnected');


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
