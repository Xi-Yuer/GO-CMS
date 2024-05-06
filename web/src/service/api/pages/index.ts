import request from '@/service/request';
import { menuType } from '@/types/menus';
import { AxiosResponse } from 'axios';

export const getUserMenusRequest = () => {
  return request.get<AxiosResponse<menuType[]>>({
    url: '/pages/user',
  });
};

export const getRoleMenusRequest = (id: string) => {
  return request.get<AxiosResponse<menuType[]>>({
    url: `/pages/role/${id}`,
  });
};

export const getAllMenusRequest = () => {
  return request.get<AxiosResponse<menuType[]>>({
    url: '/pages',
  });
};
