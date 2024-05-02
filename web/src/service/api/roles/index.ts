import request from '@/service/request';
import { AxiosResponse } from 'axios';
import { IHasTotalResponse } from '@/service';
import React from 'react';
import { saveAs } from 'file-saver';

export interface IRoleResponse {
  id: string;
  roleName: string;
  description: string;
  pagesID: null | string;
  interfacesID: null | string;
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

export const exportRolesRequest = (ids: React.Key[]) => {
  return request
    .post({
      url: '/roles/export',
      data: {
        ids,
      },
      responseType: 'blob',
    })
    .then((res: any) => {
      saveAs(new Blob([res]), '角色表.xlsx');
    });
};
