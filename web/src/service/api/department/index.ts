import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface IDepartmentResponse {
  id: string;
  departmentName: string;
  parentDepartment: string;
  createTime: string;
  updateTime: string;
}

export const getDepartmentRequest = () => {
  return request.get<AxiosResponse<IDepartmentResponse[]>>({
    url: '/department',
  });
};
