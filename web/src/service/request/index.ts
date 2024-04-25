import Request from '@/service/request/lib';
import { IResponse } from '@/service/request/lib/type.ts';
import { message } from 'antd';
import { cache } from '@/utils';
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
        message.error(value.data.msg).then((r) => r);
      }
      return value.data;
    },
    onRejected(error) {
      const { msg } = error.response.data;
      message.error(msg).then((r) => r);
      return error;
    },
  },
});

export default request;
