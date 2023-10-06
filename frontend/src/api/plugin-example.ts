import axios from 'axios';

export function getEmployee(id: number) {
  return axios.get<any>(`/plugin-example/employee/info/${id}`);
}

export function getEmployees() {
  return axios.get<any>(`/plugin-example/employees`);
}
