import axios from 'axios';

export function getEmployeeInfo(id: string) {
  return axios.get(`/example/employee/${id}`).then((res) => res.data);
}

export function hello() {
  return 'hi';
}
