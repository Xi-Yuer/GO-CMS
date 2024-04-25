import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface IGetUsersParams {
  offset: number;
  limit: number;
  nickname?: string;
  status?: number;
  account?: number;
  id?: string;
  startTime?: string;
  endTime?: string;
  departmentID?: string;
}

export interface IUserResponse {
  id: string;
  account: string;
  nickname: string;
  avatar: string | null;
  rolesID: string[] | null;
  interfaceDic: string[] | null;
  departmentID: string;
  createTime: string;
  updateTime: string;
  status: number;
}

export const getUsers = (params: IGetUsersParams) => {
  return request.get<AxiosResponse<IUserResponse[]>>({
    url: '/users',
    params,
  });
};
