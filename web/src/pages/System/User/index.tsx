import React, { FC, memo } from 'react';
import { useUserPageHooks } from '@/pages/System/User/hooks.tsx';
import { Button, Form, Input, Modal, Pagination, Select, Table } from 'antd';
import { IUpdateUserParams } from '@/service';
import { useTranslation } from 'react-i18next';
import { useSearchFrom } from '@/hooks/useSearchForm.tsx';
import { DownloadOutlined } from '@ant-design/icons';

const SystemUser: FC = () => {
  const {
    total,
    columns,
    users,
    departments,
    editFormRef,
    editUserModalOpen,
    roles,
    isEdit,
    limit,
    searchConfig,
    selected,
    loading,
    exportUsersAction,
    setSelected,
    getPageData,
    setPage,
    setLimit,
    editUserConfirm,
    setEditUserModalOpen,
    editUserAction,
  } = useUserPageHooks();
  const { t } = useTranslation();
  const { SearchFormComponent } = useSearchFrom({
    getDataRequestFn: getPageData,
    onNewRecordFn: editUserAction,
    formItems: searchConfig,
    operateComponent: !!selected.length && (
      <Button type='primary' icon={<DownloadOutlined />} onClick={exportUsersAction}>
        导出
      </Button>
    ),
  });
  return (
    <>
      {SearchFormComponent}
      <Table
        dataSource={users}
        loading={loading}
        columns={columns}
        rowSelection={{
          onChange: (selectedRowKeys: React.Key[]) => {
            setSelected(selectedRowKeys);
          },
        }}
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
      <Modal destroyOnClose open={editUserModalOpen} title={isEdit ? t('edit') : t('add')} onOk={editUserConfirm} onCancel={() => setEditUserModalOpen(false)}>
        <Form form={editFormRef} autoComplete='off' labelAlign='right' id='editFormRef'>
          {!isEdit && (
            <Form.Item<IUpdateUserParams> name='account' label={t('account')} rules={[{ required: !isEdit }]}>
              <Input />
            </Form.Item>
          )}
          <Form.Item<IUpdateUserParams> name='nickname' label={t('nickName')} rules={[{ required: !isEdit }]}>
            <Input />
          </Form.Item>
          <Form.Item<IUpdateUserParams> name='password' label={t('password')} rules={[{ required: !isEdit }]}>
            <Input.Password />
          </Form.Item>
          <Form.Item<IUpdateUserParams> name='departmentID' label={t('department')} rules={[{ required: !isEdit }]}>
            <Select allowClear>
              {departments?.map((item) => {
                return (
                  <Select.Option key={item.id} value={item.id}>
                    {item.departmentName}
                  </Select.Option>
                );
              })}
            </Select>
          </Form.Item>
          <Form.Item<IUpdateUserParams> name='rolesID' label={t('role')} rules={[{ required: !isEdit }]}>
            <Select allowClear mode='multiple'>
              {roles?.map((item) => {
                return (
                  <Select.Option key={item.id} value={item.id}>
                    {item.roleName}
                  </Select.Option>
                );
              })}
            </Select>
          </Form.Item>
          <Form.Item<IUpdateUserParams> name='status' label={t('status')} rules={[{ required: !isEdit }]}>
            <Select>
              <Select.Option value='0'>{t('off')}</Select.Option>
              <Select.Option value='1'>{t('on')}</Select.Option>
            </Select>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default memo(SystemUser);
