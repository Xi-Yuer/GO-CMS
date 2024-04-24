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

export const MenUsageMap = [
  {
    key: 'available',
    label: '可用内存',
  },
  {
    key: 'used',
    label: '已使用内存',
  },
  {
    key: 'free',
    label: '空闲内存',
  },
  {
    key: 'active',
    label: '活跃内存',
  },
  {
    key: 'inactive',
    label: '非活跃内存',
  },
  {
    key: 'wired',
    label: '已分配内存',
  },
  {
    key: 'laundry',
    label: '被回收内存',
  },
];

export interface IGitCommit {
  date: string;
  children: {
    author: string;
    commitID: string;
    date: string;
    email: string;
    message: string;
  }[];
}

export const getSystemRunTimeInfoRequest = () => {
  return request.get<AxiosResponse<SystemInfo[]>>({
    url: '/system',
  });
};

export const getGitCommitInfoRequest = () => {
  return request.get<AxiosResponse<IGitCommit[]>>({
    url: '/log/commits',
  });
};

export const getGitCommitCountRequest = () => {
  return request.get<AxiosResponse<number>>({
    url: '/log/commits/count',
  });
};
