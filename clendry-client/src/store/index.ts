import { Commit, createStore } from 'vuex'

export default createStore({
   state: {
    searchValue: '',
    authenticated: false,
    appName: 'Clendry ',
    logo: require('../assets/images/logo.png'),
    darklogo:require('../assets/images/logo-white.png'),
    dark: false,
    namespaced: true,
    user:{
      name:'Bill Yerds',
      image:require('../assets/images/user/1.jpg'),
    }
  },
  mutations: {
    SET_AUTH: (state: { authenticated: boolean }, auth: boolean) => state.authenticated = auth,
    SET_SEARCH: (state: { searchValue: string }, value: string) => state.searchValue = value,
    layoutModeCommit (state, payload) {
      state.dark = payload
      if (!payload) {
        state.logo = require('../assets/images/logo.png')
      } else {
        state.logo = require('../assets/images/logo-white.png')
      }
    }
  },
  actions: {
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('SET_AUTH', auth),
    layoutModeAction (context, payload) {
      context.commit('layoutModeCommit', payload.dark)
    },
    setSearch: ({ commit }: { commit: Commit }, value: string) => commit('SET_SEARCH', value),
  },
  getters: {
    appName: state => { return state.appName },
    logo: state => { return state.logo },
    darklogo: state => { return state.darklogo },
    image1: state => { return state.user.image },
    name: state => { return state.user.name },
    dark: state => { return state.dark },
    searchVal: state => { return state.searchValue},
  },
  modules: {
  }
})
