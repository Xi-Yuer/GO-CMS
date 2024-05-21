import request from '@/service/request';
import { AxiosResponse } from 'axios';

interface Field {
  name: string;
  type: string;
  default: string;
}

export interface ICreateTemplateParams {
  package: string;
  tableName: string;
  fields: Field[];
}

export interface ICreateTemplateResponse {
  server: server;
  web: web;
}

export interface server {
  controllerFile: code;
  serviceFile: code;
  repositoryFile: code;
  routeFile: code;
  dtoFile: code;
}

export interface web {
  react: reactType;
}

export interface reactType {
  searchForm: code;
  table: code;
  tableHook: code;
}

export interface code {
  code: string;
  lang: string;
}

export interface IDownloadTemplateParams {
  tableName: string;
  controller: string | undefined;
  service: string | undefined;
  repository: string | undefined;
  route: string | undefined;
  dto: string | undefined;
  searchForm: string | undefined;
  table: string | undefined;
  tableHook: string | undefined;
}

export const createTemplateRequest = (data: ICreateTemplateParams) => {
  return request.post<AxiosResponse<ICreateTemplateResponse>>({
    url: '/template',
    data: data,
  });
};

export const downloadTemplateRequest = (data: IDownloadTemplateParams) => {
  return request.post<Blob>({
    url: '/template/download',
    data: data,
    responseType: 'blob',
  });
};
