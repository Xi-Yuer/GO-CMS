import type { AxiosInterceptorOptions, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from 'axios'; // 扩展拦截器

export interface RequestConfig extends AxiosRequestConfig {
  requestInterceptor?: {
    onFulfilled?: (value: InternalAxiosRequestConfig) => InternalAxiosRequestConfig | Promise<InternalAxiosRequestConfig>;
    onRejected?: ((error: any) => any) | null;
    options?: AxiosInterceptorOptions;
  };

  responseInterceptor?: {
    onFulfilled?: <T>(value: AxiosResponse) => AxiosResponse<T, any>;
    onRejected?: ((error: any) => any) | null;
    options?: AxiosInterceptorOptions;
  };
}

export type InterceptorType = Partial<Pick<RequestConfig, 'requestInterceptor' | 'responseInterceptor'>>;

export interface IResponse<T = any> {
  data: T;
  message: string;
  code: number;
}
