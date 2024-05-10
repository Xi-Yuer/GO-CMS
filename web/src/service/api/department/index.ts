import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface IDepartmentResponse {
  id: string;
  departmentName: string;
  parentDepartment: string;
  departmentOrder: number;
  departmentDescription: string;
  children: IDepartmentResponse[];
  createTime: string;
  updateTime: string;
}

export interface IAddDepartmentRequest {
  departmentName: string;
  parentDepartment: string;
  departmentOrder: number;
  departmentDescription: string;
}

export const getDepartmentRequest = () => {
  return request.get<AxiosResponse<IDepartmentResponse[]>>({
    url: '/department',
  });
};

export const addDepartmentRequest = (data: IAddDepartmentRequest) => {
  return request.post<AxiosResponse<IDepartmentResponse>>({
    url: '/department',
    data,
  });
};

export const deleteDepartmentRequest = (id: string) => {
  return request.delete<AxiosResponse<IDepartmentResponse>>({
    url: `/department/${id}`,
  });
};

export const updateDepartmentRequest = (id: string, data: IAddDepartmentRequest) => {
  return request.patch<AxiosResponse<IDepartmentResponse>>({
    url: `/department/${id}`,
    data,
  });
};
