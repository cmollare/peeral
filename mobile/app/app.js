import Vue from "nativescript-vue";
import { localize, androidLaunchEventLocalizationHandler } from "nativescript-localize";
import { on, launchEvent } from '@nativescript/core/application';

import Login from "./pages/Login";

Vue.config.silent = false
Vue.filter("L", localize);

on(launchEvent, (args) => {
  if (args.android) {
    androidLaunchEventLocalizationHandler();
  }
});

new Vue({
    render: h => h('frame', [h(Login)]),
}).$start();
