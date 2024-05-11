import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface IInterfaceResponse {
  id: string;
  interfaceName: string;
  interfaceMethod: string;
  interfacePageID: string;
  interfacePath: string;
  interfaceDesc: string;
  interfaceDic: string;
  createTime: string;
  updateTime: string;
}

export interface IAllPageInterfaceListResponse {
  key: string;
  children: IInterfaceResponse[];
}

// 获取页面接口
export const getInterfaceListByPageIDRequest = (id: string) => {
  return request.get<AxiosResponse<IInterfaceResponse[]>>({
    url: `/interface/page/${id}`,
  });
};

// 按照页面分类获取所有接口
export const getInterfaceAllListRequest = () => {
  return request.get<AxiosResponse<IAllPageInterfaceListResponse[]>>({
    url: `/interface`,
  });
};

export const deleteInterfaceRequest = (id: string) => {
  return request.delete<AxiosResponse<string>>({
    url: `/interface/${id}`,
  });
};

export const addInterfaceRequest = (data: IInterfaceResponse) => {
  return request.post<AxiosResponse<string>>({
    url: `/interface`,
    data,
  });
};

export const updateInterfaceRequest = (data: IInterfaceResponse) => {
  return request.patch<AxiosResponse<string>>({
    url: `/interface/${data.id}`,
    data,
  });
};
