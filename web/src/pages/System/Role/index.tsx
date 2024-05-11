import React, { FC, memo } from 'react';
import { useRolePageHooks } from '@/pages/System/Role/hooks.tsx';
import { Button, Card, Drawer, Form, Input, Modal, Pagination, Space, Table, Tabs, Tree } from 'antd';
import { useTranslation } from 'react-i18next';
import { IUpdateRoleParams } from '@/service';
import { SwapOutlined } from '@ant-design/icons';
import User from '@/pages/System/User';
import { constants } from '@/constant';

const SystemRole: FC = () => {
  const {
    userPageRef,
    roles,
    roleColumns,
    userColumns,
    menus,
    SearchFormComponent,
    total,
    roleUnderUserListTotal,
    limit,
    roleUnderUserLimit,
    loading,
    isEdit,
    formRef,
    roleInterfaceSelected,
    currentEditRole,
    editRoleModalOpen,
    editRolePermissionOpen,
    editRoleUnderUserOpen,
    onPageTreeCheck,
    allInterface,
    roleMenusSelected,
    onInterfaceTreeCheck,
    roleUnderUsersList,
    addRoleUnderUserAction,
    setEditRolePermissionOpen,
    setEditRoleUnderUserOpen,
    setPage,
    setLimit,
    roleUnderUserPageChange,
    setSelected,
    setEditRoleModalOpen,
    onFinish,
    editPermissionConfirm,
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
        columns={roleColumns}
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
        <Form form={formRef} autoComplete='off' labelAlign='left' labelCol={{ span: 6 }} id='editFormRef'>
          <Form.Item<IUpdateRoleParams> name='roleName' label={t('roleName')} rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item<IUpdateRoleParams> name='description' label={t('roleDescription')} rules={[{ required: true }]}>
            <Input />
          </Form.Item>
        </Form>
      </Modal>
      <Drawer
        title={t('permissionEdit')}
        width='40%'
        destroyOnClose
        onClose={() => setEditRolePermissionOpen(false)}
        open={editRolePermissionOpen}
        extra={
          <Space>
            <Button onClick={() => setEditRolePermissionOpen(false)}>{t('cancel')}</Button>
            <Button type='primary' onClick={editPermissionConfirm}>
              {t('confirm')}
            </Button>
          </Space>
        }>
        <Tabs
          size='small'
          items={[
            {
              label: t('menusPermission'),
              key: '1',
              children: (
                <>
                  <Tree
                    treeData={menus}
                    checkable
                    defaultExpandedKeys={roleMenusSelected}
                    defaultSelectedKeys={roleMenusSelected}
                    defaultCheckedKeys={roleMenusSelected}
                    multiple
                    selectable={false}
                    onCheck={onPageTreeCheck}></Tree>
                </>
              ),
            },
            {
              label: t('resourcePermission'),
              key: '2',
              children: (
                <>
                  <Tree
                    treeData={allInterface}
                    defaultExpandedKeys={roleInterfaceSelected}
                    defaultSelectedKeys={roleInterfaceSelected}
                    defaultCheckedKeys={roleInterfaceSelected}
                    multiple
                    checkable
                    selectable={false}
                    onCheck={onInterfaceTreeCheck}></Tree>
                </>
              ),
            },
          ]}></Tabs>
      </Drawer>
      <Drawer open={editRoleUnderUserOpen} title={t('authUser')} width='80%' destroyOnClose onClose={() => setEditRoleUnderUserOpen(false)}>
        <div className='flex gap-2 justify-between'>
          <Card title={t('readyAuthUser')} style={{ width: '50%' }}>
            <Table dataSource={roleUnderUsersList} columns={userColumns} pagination={false} rowKey='id' bordered></Table>
            <Pagination
              total={roleUnderUserListTotal}
              className='flex justify-end mt-2'
              pageSize={roleUnderUserLimit}
              onChange={(page, pageSize) => roleUnderUserPageChange(page, pageSize)}
              showSizeChanger></Pagination>
          </Card>
          <div className='my-auto'>
            <SwapOutlined />
          </div>
          <Card title={t('notAuthUser')}>
            <User ref={userPageRef} module={constants.module.ROLE} context={currentEditRole?.id} operation={addRoleUnderUserAction}></User>
          </Card>
        </div>
      </Drawer>
    </>
  );
};

export default memo(SystemRole);
