import React, { FC, memo } from 'react';
import { useRolePageHooks } from '@/pages/System/Role/hooks.tsx';
import { Form, Input, Modal, Pagination, Table } from 'antd';
import { useTranslation } from 'react-i18next';
import { IUpdateRoleParams } from '@/service';

const SystemRole: FC = () => {
  const {
    roles,
    columns,
    SearchFormComponent,
    total,
    limit,
    loading,
    editRoleModalOpen,
    isEdit,
    formRef,
    setPage,
    setLimit,
    setSelected,
    setEditRoleModalOpen,
    onFinish,
  } = useRolePageHooks();
  const { t } = useTranslation();
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
      <Modal open={editRoleModalOpen} title={isEdit ? t('edit') : t('add')} onOk={onFinish} onCancel={() => setEditRoleModalOpen(false)}>
        <Form form={formRef} autoComplete='off' labelAlign='right' id='editFormRef'>
          <Form.Item<IUpdateRoleParams> name='roleName' label={t('roleName')} rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item<IUpdateRoleParams> name='description' label={t('roleDescription')} rules={[{ required: true }]}>
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default memo(SystemRole);
