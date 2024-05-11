import { FC, memo, useEffect, useState } from 'react';
import { Button, Form, Input, InputNumber, Modal, Table, TableColumnsType } from 'antd';
import { addDepartmentRequest, deleteDepartmentRequest, getDepartmentRequest, IDepartmentResponse, updateDepartmentRequest } from '@/service';
import { useTranslation } from 'react-i18next';
import dayjs from 'dayjs';
import { useForm } from 'antd/es/form/Form';
import { PlusOutlined } from '@ant-design/icons';
import Auth from '@/components/Auth';
import { constants } from '@/constant';

const SystemDepartment: FC = () => {
  const { t } = useTranslation();
  const [form] = useForm();
  const [departmentList, setDepartmentList] = useState<IDepartmentResponse[]>([]);
  const [modalOpen, setModalOpen] = useState(false);
  const [isEdit, setIsEdit] = useState(false);
  const [editCurrentDepartment, setEditCurrentDepartment] = useState<IDepartmentResponse>();
  const getPageData = async () => {
    const result = await getDepartmentRequest();
    setDepartmentList(result.data);
  };

  const columns: TableColumnsType<IDepartmentResponse> = [
    {
      title: t('departmentName'),
      dataIndex: 'departmentName',
      key: 'departmentName',
    },
    {
      title: t('order'),
      align: 'center',
      dataIndex: 'departmentOrder',
      key: 'departmentOrder',
    },
    {
      title: t('departmentDescription'),
      align: 'center',
      dataIndex: 'departmentDescription',
      key: 'departmentDescription',
    },
    {
      title: t('createTime'),
      align: 'center',
      dataIndex: 'createTime',
      key: 'createTime',
      render: (text) => {
        return dayjs(text).format('YYYY-MM-DD HH:mm:ss');
      },
    },
    {
      title: t('updateTime'),
      align: 'center',
      dataIndex: 'updateTime',
      key: 'updateTime',
      render: (text) => {
        return dayjs(text).format('YYYY-MM-DD HH:mm:ss');
      },
    },
    {
      title: t('operate'),
      dataIndex: 'operation',
      align: 'center',
      key: 'operation',
      render: (_, row) => {
        return (
          <div className='text-[#00b0f0] flex gap-2 items-center justify-center cursor-pointer'>
            <Auth permission={constants.permissionDicMap.UPDATE_DEPARTMENT}>
              <span
                onClick={() => {
                  setEditCurrentDepartment(row);
                  setIsEdit(true);
                  setModalOpen(true);
                  form.setFieldsValue(row);
                }}>
                {t('edit')}
              </span>
            </Auth>
            <Auth permission={constants.permissionDicMap.ADD_DEPARTMENT}>
              <span
                onClick={() => {
                  setEditCurrentDepartment(row);
                  setIsEdit(false);
                  setModalOpen(true);
                  form.resetFields();
                }}>
                {t('addChildDepartment')}
              </span>
            </Auth>
            <Auth permission={constants.permissionDicMap.DELETE_DEPARTMENT}>
              <span
                className='text-red-500'
                onClick={async () => {
                  await deleteDepartmentRequest(row.id);
                  await getPageData();
                }}>
                {t('delete')}
              </span>
            </Auth>
          </div>
        );
      },
    },
  ];

  const onOk = () => {
    form.validateFields().then(async (values) => {
      if (isEdit) {
        if (!editCurrentDepartment?.id) return;
        await updateDepartmentRequest(editCurrentDepartment?.id, values);
      } else {
        await addDepartmentRequest({ ...values, parentDepartment: editCurrentDepartment?.id });
      }
      setModalOpen(false);
      getPageData().then();
    });
  };

  useEffect(() => {
    getPageData().then();
  }, []);
  return (
    <>
      <div className='mb-2 flex justify-between items-center bg-white p-4 rounded dark:bg-[#001620]'>
        <span className='font-bold'>{t('departmentList')}</span>
        <Auth permission={constants.permissionDicMap.ADD_DEPARTMENT}>
          <Button
            type='primary'
            icon={<PlusOutlined />}
            onClick={() => {
              setEditCurrentDepartment(undefined);
              form.resetFields();
              setIsEdit(false);
              setModalOpen(true);
            }}>
            {t('add')}
          </Button>
        </Auth>
      </div>
      <Table columns={columns} dataSource={departmentList} bordered key='departmentTable' rowKey='id' />
      <Modal open={modalOpen} onCancel={() => setModalOpen(false)} title={isEdit ? t('edit') : t('add')} onOk={onOk}>
        <Form form={form} labelAlign='left' labelCol={{ span: 6 }} autoComplete='off'>
          <Form.Item name='departmentName' label={t('departmentName')} rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item name='departmentDescription' label={t('departmentDescription')} rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item name='departmentOrder' label={t('order')} rules={[{ required: true }]}>
            <InputNumber style={{ width: '100%' }} type='number' placeholder={t('pleaseEnter')} />
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default memo(SystemDepartment);
