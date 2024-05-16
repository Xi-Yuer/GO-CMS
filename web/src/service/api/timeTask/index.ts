import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface ITimeTaskResponse {
  cron: string;
  lastRunTime: string;
  runTimes: number;
  status: boolean;
  taskName: string;
  timeTaskID: string;
}

export const getTimeTaskListRequest = () => {
  return request.get<AxiosResponse<ITimeTaskResponse[]>>({
    url: '/timeTask',
  });
};

export const startTimeTaskRequest = (id: string) => {
  return request.post({
    url: `/timeTask/start/${id}`,
  });
};

export const stopTimeTaskRequest = (id: string) => {
  return request.post({
    url: `/timeTask/stop/${id}`,
  });
};

export interface IUpdateTimeTaskRequest {
  timeTaskName: string;
  cron: string;
  status: boolean;
}

export const updateTimeTaskRequest = (id: string, data: IUpdateTimeTaskRequest) => {
  return request.patch({
    url: `/timeTask/update/${id}`,
    data,
  });
};
