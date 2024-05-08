import request from '@/service/request';
import { Md5 } from 'ts-md5';
import { AxiosResponse } from 'axios';

export type LoginParamsType = {
  account: string;
  password: string;
  captcha: string;
};

export interface LoginUserResponseType {
  id: string;
  account: string;
  nickname: string;
  isAdmin: number;
  avatar: null;
  rolesID: string[];
  interfaceDic: string[];
  departmentID: string;
  createTime: Date;
  updateTime: Date;
  status: string;
}

export const loginRequest = (data: LoginParamsType) => {
  return request.post<AxiosResponse<{ user: LoginUserResponseType } & { token: string }>>({
    url: '/auth/login',
    data: {
      ...data,
      password: Md5.hashStr(data.password),
    },
  });
};

export const getCaptchaRequest = () => {
  return request.get({
    url: '/auth/captcha',
    responseType: 'blob',
  });
};
