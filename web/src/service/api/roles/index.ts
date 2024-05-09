import request from '@/service/request';
import { AxiosResponse } from 'axios';
import { IHasTotalResponse } from '@/service';
import React from 'react';
import { saveAs } from 'file-saver';

export interface IRoleResponse {
  id: string;
  roleName: string;
  description: string;
  pageID: null | string[];
  interfaceID: null | string[];
  createTime: string;
  updateTime: string;
}

export interface Page {
  limit: number;
  offset: number;
}

export type IQueryRoleParams = Partial<Omit<IRoleResponse, 'createTime' | 'updateTime'>> & {
  limit: number;
  offset: number;
  startTime: string;
  endTime: string;
};

export interface IUpdateRoleParams {
  id: string;
  roleName: string;
  description: string;
  pageID: string[];
  interfaceID: string[];
}

export const getRolesRequest = (params?: IQueryRoleParams & Page) => {
  return request.post<AxiosResponse<IHasTotalResponse<IRoleResponse[]>>>({
    url: '/roles/query',
    data: params,
  });
};

export const deleteRolesRequest = (id: string) => {
  return request.delete({
    url: `/roles/${id}`,
  });
};

export const exportRolesRequest = async (ids: React.Key[]) => {
  const res = await request.post({
    url: '/roles/export',
    data: {
      ids,
    },
    responseType: 'blob',
  });
  saveAs(new Blob([res]), '角色表.xlsx');
};

export const updateRoleRequest = (data: IUpdateRoleParams) => {
  return request.patch({
    url: `/roles/${data.id}`,
    data,
  });
};

export const createRoleRequest = (data: IUpdateRoleParams) => {
  return request.post({
    url: '/roles',
    data,
  });
};

export interface IDeBindUserParams {
  userID: string;
  roleID: string;
}

export const deBindUserRequest = (data: IDeBindUserParams) => {
  return request.post({
    url: '/roles/deBindUser',
    data,
  });
};

export const bindUserRequest = (data: IDeBindUserParams) => {
  return request.post({
    url: '/roles/bindUser',
    data,
  });
};
