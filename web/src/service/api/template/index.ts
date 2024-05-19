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
  controller: string;
  service: string;
  dto: string;
  repository: string;
}

export const createTemplateRequest = (data: ICreateTemplateParams) => {
  return request.post<AxiosResponse<ICreateTemplateResponse>>({
    url: '/template',
    data: data,
  });
};
