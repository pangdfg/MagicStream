import axios from 'axios';

const API_BASE_URL = import.meta.env.API_BASE_URL; 

const axiosPrivate = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',

  },
  withCredentials: true, 
});

export default axiosPrivate;