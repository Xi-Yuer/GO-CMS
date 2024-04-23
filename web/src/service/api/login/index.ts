import request from '@/service/request';
import { Md5 } from 'ts-md5';

export type LoginParamsType = {
  account: string;
  password: string;
  captcha: string;
};

export const loginRequest = (data: LoginParamsType) => {
  return request.post({
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
