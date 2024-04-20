import Request from '@/service/request/lib';
import { IResponse } from '@/service/request/lib/type.ts';

const request = new Request<IResponse>(import.meta.env.VITE_APP_BASE_URL, 1000 * 60, {
  requestInterceptor: {
    onFulfilled(config) {
      return config;
    },
    onRejected(error) {
      return error;
    },
  },
  responseInterceptor: {
    onFulfilled: (value) => {
      return value.data;
    },
    onRejected(error) {
      return error;
    },
  },
});

export default request;
