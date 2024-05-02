import request from '@/service/request';

// 获取页面接口
export const getInterfaceListQuest = (id: string) => {
  return request.get({
    url: `/interface/page/${id}`,
  });
};
