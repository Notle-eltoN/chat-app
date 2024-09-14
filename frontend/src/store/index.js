import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import io from 'socket.io-client';

Vue.use(Vuex);

const socket = io('ws://localhost:8080');

export default new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || '',
    user: null,
    messages: [],
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
      localStorage.setItem('token', token);
    },
    setUser(state, user) {
      state.user = user;
    },
    addMessage(state, message) {
      state.messages.push(message);
    },
    clearMessages(state) {
      state.messages = [];
    },
  },
  actions: {
    login({ commit }, credentials) {
      return axios.post('http://localhost:8080/login', credentials)
        .then(response => {
          commit('setToken', response.data.token);
          commit('setUser', credentials.username);
        });
    },
    signup({ commit }, credentials) {
      return axios.post('http://localhost:8080/signup', credentials)
        .then(() => {
          commit('setUser', credentials.username);
        });
    },
    sendMessage({ state }, message) {
      socket.emit('chatMessage', {
        content: message,
        username: state.user,
      });
    },
    listenForMessages({ commit }) {
      socket.on('message', (message) => {
        commit('addMessage', message);
      });
    },
  },
});
