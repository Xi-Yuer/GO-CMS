import request from '@/service/request';
import { AxiosResponse } from 'axios';

export interface SystemInfo {
  cpuUsage: number;
  memUsage: {
    total: number; // 总内存 (bytes)
    available: number; // 可用内存 (bytes)
    used: number; // 已使用内存 (bytes)
    usedPercent: number; // 内存使用率 (%)
    free: number; // 空闲内存 (bytes)
    active: number; // 活跃内存 (bytes)
    inactive: number; // 非活跃内存 (bytes)
    wired: number; // 已分配内存 (bytes)
    laundry: number; // 被回收内存 (bytes)
  };
}

export const getSystemRunTimeInfoRequest = () => {
  return request.get<AxiosResponse<SystemInfo[]>>({
    url: '/system',
  });
};
