import Request from '@/service/request/lib';
import { IResponse } from '@/service/request/lib/type.ts';
import { cache, message } from '@/utils';
import { constants } from '@/constant';

const request = new Request<IResponse>(import.meta.env.VITE_APP_BASE_URL, 1000 * 60, {
  requestInterceptor: {
    onFulfilled(config) {
      const token = cache.get(constants.localStorage.token);
      if (token) {
        config.headers.Authorization = token;
      }
      return config;
    },
    onRejected(error) {
      return error;
    },
  },
  responseInterceptor: {
    onFulfilled: (value) => {
      if (value.data.code > 201) {
        message.error(value.data.msg);
      }
      if (value.data.code == 401) {
        cache.clear();
        window.location.replace(constants.routePath.login);
        message.error('登录过期，请重新登录');
        return;
      }
      return value.data;
    },
    onRejected(error) {
      const { msg } = error.response.data;
      message.error(msg);
      if (error.response.data.code == 401) {
        cache.clear();
        window.location.replace(constants.routePath.login);
        message.error('登录过期，请重新登录');
        return;
      }
      return error;
    },
  },
});

export default request;
