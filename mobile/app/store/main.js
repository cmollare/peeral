import Vue from 'nativescript-vue';
import Vuex from 'vuex';
Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    currentMessage: ""
  },
  mutations: {
  },
  actions: {
    parseCurrentMessage({ commit }) {
      
    }
  }
});

export default store;
