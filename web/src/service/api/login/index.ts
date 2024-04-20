import request from '@/service/request';

export const login = () => {};
export const getCaptchaRequest = () => {
  return request.get({
    url: '/auth/captcha',
    responseType: 'blob',
  });
};
