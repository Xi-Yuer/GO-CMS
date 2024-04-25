import { FC, memo, useEffect, useState } from 'react';
import { getUsers, IUserResponse } from '@/service';
import { useUserPageHooks } from '@/pages/System/User/hooks.tsx';
import { Table } from 'antd';

const SystemUser: FC = () => {
  const [users, setUsers] = useState<IUserResponse[]>([]);
  const { columns } = useUserPageHooks();
  useEffect(() => {
    getPageData();
  }, []);
  const getPageData = () => {
    getUsers({ limit: 20, offset: 0 }).then((res) => {
      setUsers(res.data);
    });
  };

  return (
    <>
      <div>
        <Table dataSource={users} columns={columns} bordered={true} rowKey='id'></Table>
      </div>
    </>
  );
};

export default memo(SystemUser);
