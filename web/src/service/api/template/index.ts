import request from '@/service/request';
import { AxiosResponse } from 'axios';

interface Field {
  name: string;
  type: string;
  default: string;
}

export interface ICreateTemplateParams {
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
  dtoFile: code;
}

export interface web {}

export interface code {
  code: string;
  lang: string;
}

export const createTemplateRequest = (data: ICreateTemplateParams) => {
  return request.post<AxiosResponse<ICreateTemplateResponse>>({
    url: '/template',
    data: data,
  });
};
