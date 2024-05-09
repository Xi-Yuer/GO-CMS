import request from '@/service/request';
import { AxiosResponse } from 'axios';
import { saveAs } from 'file-saver';
import React from 'react';

export interface IGetUsersParams {
  id?: string;
  offset: number;
  limit: number;
  nickname?: string;
  status?: number;
  account?: number;
  startTime?: string;
  endTime?: string;
  departmentID?: string;
}

export interface IUserResponse {
  id: string;
  account: string;
  nickname: string;
  isAdmin: 0 | 1;
  avatar?: string | null;
  rolesID: string[] | null;
  interfaceDic: string[] | null;
  departmentID: string;
  createTime: string;
  updateTime: string;
  status: '0' | '1';
}

export interface IHasTotalResponse<T> {
  total: number;
  list: T;
}

export interface IUpdateUserParams {
  id: string;
  account?: string;
  nickname?: string;
  password?: string;
  rolesID?: string[];
  status?: '0' | '1';
  departmentID?: string;
}

export interface IPage {
  limit: number;
  offset: number;
}

export const getUsersRequest = (params: IGetUsersParams) => {
  return request.post<AxiosResponse<IHasTotalResponse<IUserResponse[]>>>({
    url: '/users/query',
    data: params,
  });
};

export const updateUsersRequest = (params: IUpdateUserParams) => {
  return request.patch<AxiosResponse<IUserResponse[]>>({
    url: `/users/${params.id}`,
    data: params,
  });
};

export const deleteUsersRequest = (id: string) => {
  return request.delete({
    url: `/users/${id}`,
  });
};

export const getUserRequest = (id: string) => {
  return request.get<AxiosResponse<IUserResponse>>({
    url: `/users/${id}`,
  });
};

export const createUsersRequest = (params: IUpdateUserParams) => {
  return request.post<AxiosResponse<IUserResponse>>({
    url: '/users',
    data: params,
  });
};

export const exportUsersRequest = async (ids: React.Key[]) => {
  const res = await request.post({
    url: '/users/export',
    data: {
      ids,
    },
    responseType: 'blob',
  });
  saveAs(new Blob([res]), '用户表.xlsx');
};

export const getUserByRoleIDRequest = (params: { roleID?: string } & IPage) => {
  return request.get<AxiosResponse<IHasTotalResponse<IUserResponse[]>>>({
    url: `/users/role/${params.roleID}`,
    params: {
      limit: params.limit,
      offset: params.offset,
    },
  });
};

export const getOutRoleUsersRequest = (data: IGetUsersParams, id?: string) => {
  return request.post<AxiosResponse<IHasTotalResponse<IUserResponse[]>>>({
    url: `/users/query/role/${id}`,
    data,
  });
};
