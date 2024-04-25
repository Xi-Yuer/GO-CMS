import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface IRoleResponse {
  id: string;
  roleName: string;
  description: string;
  pagesID: null | string;
  interfacesID: null | string;
  createTime: string;
  updateTime: string;
}

export const getRolesRequest = (params?: any) => {
  return request.get<AxiosResponse<IRoleResponse[]>>({
    url: '/roles',
    params,
  });
};
