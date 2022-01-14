import Vue from "nativescript-vue";
import { localize, androidLaunchEventLocalizationHandler } from '@nativescript/localize';
import { on, launchEvent } from '@nativescript/core/application';
import Vuex from 'vuex';

Vue.use(Vuex);

import Login from "./pages/Login";

import store from './store/main';

Vue.config.silent = false
Vue.filter("L", localize);

on(launchEvent, (args) => {
  if (args.android) {
    androidLaunchEventLocalizationHandler();
  }
});

new Vue({
    render: h => h('frame', [h(Login)]),
    store: store
}).$start();
