import axios from 'axios';
import { LoginData, LoginRes } from '../types/auth';

export function login(data: LoginData) {
  return axios.post<LoginRes>('api/auth/login', data)
}