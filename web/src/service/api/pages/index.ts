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

export const deleteMenuRequest = (id: string) => {
  return request.delete<AxiosResponse>({
    url: `/pages/${id}`,
  });
};

export const updateMenuRequest = (id: string, data: menuType) => {
  return request.patch<AxiosResponse>({
    url: `/pages/${id}`,
    data,
  });
};

export const createMenuRequest = (data: menuType) => {
  return request.post<AxiosResponse>({
    url: '/pages',
    data,
  });
};
