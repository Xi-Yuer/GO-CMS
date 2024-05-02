import React, { FC, memo } from 'react';
import { useRolePageHooks } from '@/pages/System/Role/hooks.tsx';
import { Pagination, Table } from 'antd';

const SystemRole: FC = () => {
  const { roles, columns, SearchFormComponent, total, limit, loading, setPage, setLimit, setSelected } = useRolePageHooks();
  return (
    <>
      {SearchFormComponent}
      <Table
        dataSource={roles}
        rowSelection={{
          onChange: (selectedRowKeys: React.Key[]) => {
            setSelected(selectedRowKeys);
          },
        }}
        loading={loading}
        columns={columns}
        bordered={true}
        pagination={false}
        rowKey='id'></Table>
      <Pagination
        total={total}
        className='flex justify-end mt-2'
        pageSize={limit}
        onChange={(e) => setPage(e)}
        showSizeChanger
        onShowSizeChange={(_, size) => setLimit(size)}></Pagination>
    </>
  );
};

export default memo(SystemRole);
