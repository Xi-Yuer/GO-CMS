import { FC, memo } from 'react';
import { useUserPageHooks } from '@/pages/System/User/hooks.tsx';
import { Button, Col, DatePicker, Form, Input, Modal, Row, Select, Table } from 'antd';
import { IGetUsersParams, IUpdateUserParams } from '@/service';
import { PlusOutlined, RedoOutlined, SearchOutlined } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';

const SystemUser: FC = () => {
  const {
    columns,
    users,
    departments,
    searchFormRef,
    editFormRef,
    editUserModalOpen,
    roles,
    isEdit,
    onFinish,
    onReset,
    editUserConfirm,
    setEditUserModalOpen,
    editUserAction,
  } = useUserPageHooks();
  const { t } = useTranslation();
  return (
    <>
      <div className='bg-white p-4 pt-10 mb-4 rounded dark:bg-[#141414]'>
        <Form form={searchFormRef} onReset={onReset} onFinish={onFinish} autoComplete='off' name='searchFormRef'>
          <Row gutter={[10, 10]}>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <Form.Item<IGetUsersParams> name='account' label={t('account')}>
                <Input allowClear />
              </Form.Item>
            </Col>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <Form.Item<IGetUsersParams> name='nickname' label={t('nickName')}>
                <Input allowClear />
              </Form.Item>
            </Col>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <Form.Item<IGetUsersParams> name='departmentID' label={t('department')}>
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
            </Col>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <Form.Item<IGetUsersParams> name='status' label={t('status')}>
                <Select allowClear>
                  <Select.Option value='0'>{t('on')}</Select.Option>
                  <Select.Option value='1'>{t('off')}</Select.Option>
                </Select>
              </Form.Item>
            </Col>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <Form.Item<IGetUsersParams> name='startTime' label={t('createTime')}>
                <DatePicker allowClear></DatePicker>
              </Form.Item>
            </Col>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <Form.Item<IGetUsersParams> name='endTime' label={t('updateTime')}>
                <DatePicker allowClear></DatePicker>
              </Form.Item>
            </Col>
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
              <div className='flex flex-nowrap'>
                <Button type='primary' className='mx-2' htmlType='reset' icon={<RedoOutlined />}>
                  {t('reset')}
                </Button>
                <Button type='primary' className='mx-2' htmlType='submit' icon={<SearchOutlined />}>
                  {t('search')}
                </Button>
                <Button type='primary' className='mx-2' htmlType='button' icon={<PlusOutlined />} onClick={() => editUserAction()}>
                  {t('add')}
                </Button>
              </div>
            </Col>
          </Row>
        </Form>
      </div>
      <div>
        <Table dataSource={users} columns={columns} bordered={true} rowKey='id'></Table>
      </div>
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
              {roles.map((item) => {
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
