import axios from "axios";

const api = axios.create({
  baseURL: process.env.REACT_APP_BASE_API_URL || "/api",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export { api };
