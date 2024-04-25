import { TableProps, Tag } from 'antd';
import { IUserResponse } from '@/service';

export const useUserPageHooks = () => {
  const columns: TableProps<IUserResponse>['columns'] = [
    {
      title: '用户ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '账号',
      dataIndex: 'account',
      key: 'account',
    },
    {
      title: '昵称',
      dataIndex: 'nickname',
      key: 'nickname',
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      render: (_, { status }) => {
        return <Tag color={status ? '#00b0f0' : '#f86e71'}>{status ? '启用' : '禁用'}</Tag>;
      },
    },
    {
      title: '创建时间',
      dataIndex: 'createTime',
      key: 'createTime',
    },
    {
      title: '更新时间',
      dataIndex: 'updateTime',
      key: 'updateTime',
    },
  ];

  return {
    columns,
  };
};
