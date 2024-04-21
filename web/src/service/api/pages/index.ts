import request from '@/service/request';
import { menuType } from '@/types/menus';
import { AxiosResponse } from 'axios';

export const getUserMenusRequest = () => {
  return request.get<AxiosResponse<menuType[]>>({
    url: '/pages/menus',
  });
};
