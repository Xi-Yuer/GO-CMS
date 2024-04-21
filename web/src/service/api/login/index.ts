import request from '@/service/request';

export type LoginParamsType = {
  account: string;
  password: string;
  captcha: string;
};

export const loginRequest = (data: LoginParamsType) => {
  return request.post({
    url: '/auth/login',
    data,
  });
};

export const getCaptchaRequest = () => {
  return request.get({
    url: '/auth/captcha',
    responseType: 'blob',
  });
};
