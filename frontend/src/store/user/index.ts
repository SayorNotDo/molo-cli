import { defineStore } from 'pinia';
import { LoginData } from '../../types/auth';
import { login as userLogin } from '../../api/auth';

const useUserStore = defineStore('user', {
  state: () => ({}),
  getters: {},
  actions: {
    async login(loginForm: LoginData) {
      try {
        const res = await userLogin(loginForm);
      } catch (err) {
        throw err;
      }
    }
  }
});

export default useUserStore;