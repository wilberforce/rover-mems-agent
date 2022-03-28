import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "popper.js";

import "@fortawesome/fontawesome-free/css/fontawesome.css";
import "@fortawesome/fontawesome-free/css/solid.css";

import jQuery from "jquery";
window.$ = window.jQuery = jQuery;

import "./globals";
import "./helpers";
import "./output";
import "./script";
import "./canvas";
import "./agent";
import "./websocket";
import { ecuTypeSelectChanged, runUserAction } from "./agent.js";
import { downloadData } from "./output.js";

jQuery(function () {
  jQuery('[data-toggle="popover"]').popover();
})
window.downloadData=downloadData;
window.runUserAction=runUserAction;
window.ecuTypeSelectChanged=ecuTypeSelectChanged;

import { createApp } from "vue";
import App from "./App.vue";

const app = createApp(App);

app.mount("#app");

export { app, jQuery };
