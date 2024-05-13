import request from '@/service/request';
import { IHasTotalResponse, Page } from '@/service';
import { AxiosResponse } from 'axios';

export interface ISysTemLogsResponse {
  id: string;
  requestDuration: string;
  requestMethod: string;
  requestPath: string;
  requestStatus: string;
  userID: string;
  userIP: string;
  userName: string;
  createTime: string;
}

export const getSysTemLogsRequest = (params: Page) => {
  return request.get<AxiosResponse<IHasTotalResponse<ISysTemLogsResponse[]>>>({
    url: '/log/system',
    params,
  });
};
