import { FC, memo, useEffect, useState } from 'react';
import { Button, Drawer, Form, Input, InputNumber, Modal, Select, Space, Switch, Table, TableColumnsType, Tag } from 'antd';
import { createMenuRequest, deleteMenuRequest, getAllMenusRequest, updateMenuRequest } from '@/service';
import { useTranslation } from 'react-i18next';
import dayjs from 'dayjs';
import { useForm } from 'antd/es/form/Form';
import { CheckOutlined, CloseOutlined, PlusOutlined } from '@ant-design/icons';
import { menuType } from '@/types';
import { Icon, IconSelect } from '@/components';
import {
  addInterfaceRequest,
  deleteInterfaceRequest,
  getInterfaceListByPageIDRequest,
  IInterfaceResponse,
  updateInterfaceRequest,
} from '@/service/api/interface';
import Auth from '@/components/Auth';
import { constants } from '@/constant';

const SystemMenu: FC = () => {
  const { t } = useTranslation();
  const [form] = useForm();
  const [resourceFormRef] = useForm();
  const [menuList, setMenuList] = useState<menuType[]>([]);
  const [modalOpen, setModalOpen] = useState(false);
  const [resourceModalOpen, setResourceModalOpen] = useState(false);
  const [editResourceOpen, setEditResourceOpen] = useState(false);
  const [isEdit, setIsEdit] = useState(false);
  const [editCurrentMenu, setEditCurrentMenu] = useState<menuType>();
  const [editCurrentResource, setEditCurrentResource] = useState<IInterfaceResponse>();
  const [isEditResource, setIsEditResource] = useState(false);
  const [resourceList, setResourceList] = useState<IInterfaceResponse[]>([]);

  const getPageData = async () => {
    const result = await getAllMenusRequest();
    setMenuList(result.data);
  };

  const getPageInterfaceAction = async (id: string) => {
    if (!id) return;
    const result = await getInterfaceListByPageIDRequest(id);
    setResourceList(result.data);
  };

  const methodColorMap: Record<string, string> = {
    GET: 'green',
    POST: 'blue',
    PUT: 'orange',
    DELETE: 'red',
    PATCH: 'purple',
    OPTIONS: 'cyan',
  };
  const columns: TableColumnsType<menuType> = [
    {
      title: t('menuName'),
      dataIndex: 'pageName',
      align: 'center',
      key: 'pageName',
    },
    {
      title: t('menuIcon'),
      dataIndex: 'pageIcon',
      align: 'center',
      key: 'pageIcon',
      render: (text) => {
        return <Icon name={text}></Icon>;
      },
    },
    {
      title: t('menuPath'),
      dataIndex: 'pagePath',
      align: 'center',
      key: 'pagePath',
    },
    {
      title: t('menuComponent'),
      dataIndex: 'pageComponent',
      align: 'center',
      key: 'pageComponent',
    },
    {
      title: t('order'),
      dataIndex: 'pageOrder',
      align: 'center',
      key: 'pageOrder',
    },
    {
      title: t('menuType'),
      dataIndex: 'isOutSite',
      align: 'center',
      key: 'isOutSite',
      render: (text) => {
        return text ? <Tag color='warning'>{t('routeExternalType')}</Tag> : <Tag color='green'>{t('routeInnerType')}</Tag>;
      },
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
            <Auth permission={constants.permissionDicMap.UPDATE_MENU}>
              <span
                onClick={() => {
                  setEditCurrentMenu(row);
                  setIsEdit(true);
                  setModalOpen(true);
                  form.setFieldsValue(row);
                }}>
                {t('edit')}
              </span>
            </Auth>
            <Auth permission={constants.permissionDicMap.ADD_MENU}>
              <span
                onClick={() => {
                  setEditCurrentMenu(row);
                  setIsEdit(false);
                  setModalOpen(true);
                  form.resetFields();
                }}>
                {t('addMenu')}
              </span>
            </Auth>
            {!row?.children?.length && (
              <Auth permission={constants.permissionDicMap.GET_PAGE_INTERFACE}>
                <span
                  onClick={async () => {
                    setEditCurrentMenu(row);
                    await getPageInterfaceAction(row.pageID);
                    setEditResourceOpen(true);
                  }}>
                  {t('resource')}
                </span>
              </Auth>
            )}
            <Auth permission={constants.permissionDicMap.UPDATE_MENU}>
              <span
                className='text-red-500'
                onClick={async () => {
                  await deleteMenuRequest(row.pageID);
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
  const resourceColumns: TableColumnsType<IInterfaceResponse> = [
    {
      title: t('interfaceName'),
      dataIndex: 'interfaceName',
      align: 'center',
      key: 'interfaceName',
    },
    {
      title: t('interfacePath'),
      dataIndex: 'interfacePath',
      align: 'center',
      key: 'interfacePath',
    },
    {
      title: t('interfaceDic'),
      dataIndex: 'interfaceDic',
      align: 'center',
      key: 'interfaceDic',
    },
    {
      title: t('interfaceMethod'),
      dataIndex: 'interfaceMethod',
      align: 'center',
      key: 'interfaceMethod',
      render: (text) => {
        return <Tag color={methodColorMap[text]}>{text}</Tag>;
      },
    },
    {
      title: t('interfaceDesc'),
      dataIndex: 'interfaceDesc',
      align: 'center',
      key: 'interfaceDesc',
    },
    {
      title: t('createTime'),
      dataIndex: 'createTime',
      align: 'center',
      key: 'createTime',
      render: (text) => {
        return <span>{dayjs(text).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('updateTime'),
      dataIndex: 'updateTime',
      align: 'center',
      key: 'updateTime',
      render: (text) => {
        return <span>{dayjs(text).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('operate'),
      dataIndex: 'action',
      align: 'center',
      key: 'action',
      render: (_, row) => {
        return (
          <div className='text-[#00b0f0] flex gap-2 items-center justify-center cursor-pointer'>
            <Auth permission={constants.permissionDicMap.UPDATE_PAGE_INTERFACE}>
              <span
                onClick={() => {
                  setIsEditResource(true);
                  setEditCurrentResource(row);
                  setResourceModalOpen(true);
                  resourceFormRef.setFieldsValue(row);
                }}>
                {t('edit')}
              </span>
            </Auth>
            <Auth permission={constants.permissionDicMap.DELETE_PAGE_INTERFACE}>
              <span
                className='text-red-500'
                onClick={async () => {
                  await deleteInterfaceRequest(row.id);
                  await getPageInterfaceAction(row.interfacePageID);
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
        editCurrentMenu?.pageID && (await updateMenuRequest(editCurrentMenu.pageID, values));
      } else {
        await createMenuRequest({ ...values, ParentPage: editCurrentMenu?.pageID });
      }
      setModalOpen(false);
      getPageData().then();
    });
  };

  const addResource = () => {
    setIsEditResource(false);
    setResourceModalOpen(true);
    resourceFormRef.resetFields();
  };

  const resourceEditOrAdd = () => {
    resourceFormRef.validateFields().then(async (values) => {
      if (isEditResource) {
        editCurrentResource?.id && (await updateInterfaceRequest({ ...editCurrentResource, ...values }));
      } else {
        await addInterfaceRequest({ ...values, interfacePageID: editCurrentMenu?.pageID });
      }
      setResourceModalOpen(false);
      await getPageInterfaceAction(editCurrentMenu!.pageID);
    });
  };

  useEffect(() => {
    getPageData().then();
  }, []);
  return (
    <>
      <div className='mb-2 flex justify-between items-center bg-white p-4 rounded dark:bg-[#001620]'>
        <span className='font-bold'>{t('menuList')}</span>
        <Auth permission={constants.permissionDicMap.ADD_MENU}>
          <Button
            type='primary'
            icon={<PlusOutlined />}
            onClick={() => {
              setEditCurrentMenu(undefined);
              form.resetFields();
              setIsEdit(false);
              setModalOpen(true);
            }}>
            {t('add')}
          </Button>
        </Auth>
      </div>
      <Table columns={columns} dataSource={menuList} bordered key='menuTable' rowKey='pageID' />
      <Modal open={modalOpen} onCancel={() => setModalOpen(false)} title={isEdit ? t('edit') : t('add')} onOk={onOk}>
        <Form form={form} labelAlign='left' labelCol={{ span: 6 }} autoComplete='off'>
          <Form.Item label={t('menuName')} name='pageName' rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item rules={[{ required: true }]} noStyle shouldUpdate>
            {({ getFieldValue }) => {
              return (
                (!getFieldValue('isOutSite') && (
                  <Form.Item label={t('menuPath')} name='pagePath' rules={[{ required: true }]}>
                    <Input allowClear placeholder={t('pleaseEnter')}></Input>
                  </Form.Item>
                )) ||
                null
              );
            }}
          </Form.Item>
          <Form.Item label={t('menuIcon')} name='pageIcon' rules={[{ required: true }]}>
            <IconSelect></IconSelect>
          </Form.Item>
          <Form.Item noStyle shouldUpdate rules={[{ required: true }]}>
            {({ getFieldValue }) => {
              return (
                (!getFieldValue('isOutSite') && (
                  <Form.Item label={t('menuComponent')} name='pageComponent' rules={[{ required: true }]}>
                    <Input allowClear placeholder={t('pleaseEnter')}></Input>
                  </Form.Item>
                )) ||
                null
              );
            }}
          </Form.Item>
          <Form.Item label={t('order')} name='pageOrder' rules={[{ required: true }]}>
            <InputNumber type='number' style={{ width: '100%' }} placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item label={t('isOutSite')} name='isOutSite' rules={[{ required: true }]} initialValue={false}>
            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
          </Form.Item>
          <Form.Item rules={[{ required: true }]} noStyle shouldUpdate>
            {({ getFieldValue }) => {
              return (
                (getFieldValue('isOutSite') && (
                  <Form.Item label={t('outSiteLink')} name='outSiteLink' rules={[{ required: true, type: 'url' }]}>
                    <Input allowClear placeholder={t('pleaseEnter')}></Input>
                  </Form.Item>
                )) ||
                null
              );
            }}
          </Form.Item>
        </Form>
      </Modal>
      <Drawer
        open={editResourceOpen}
        title={t('resource')}
        width='70%'
        onClose={() => setEditResourceOpen(false)}
        extra={
          <Space>
            <Button onClick={() => setEditResourceOpen(false)}>{t('cancel')}</Button>
            <Auth permission={constants.permissionDicMap.ADD_PAGE_INTERFACE}>
              <Button type='primary' onClick={addResource} icon={<PlusOutlined />}>
                {t('add')}
              </Button>
            </Auth>
          </Space>
        }>
        <Table dataSource={resourceList} bordered columns={resourceColumns} rowKey='id'></Table>
      </Drawer>
      <Modal
        open={resourceModalOpen}
        onCancel={() => setResourceModalOpen(false)}
        onOk={resourceEditOrAdd}
        title={isEditResource ? t('edit') : t('add')}
        mask={false}>
        <Form form={resourceFormRef} labelCol={{ span: 6 }} autoComplete='off'>
          <Form.Item label={t('interfaceName')} name='interfaceName' rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item label={t('interfacePath')} name='interfacePath' rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item label={t('interfaceMethod')} name='interfaceMethod' rules={[{ required: true }]}>
            <Select placeholder={t('pleaseEnter')}>
              <Select.Option value='GET'>GET</Select.Option>
              <Select.Option value='POST'>POST</Select.Option>
              <Select.Option value='PATCH'>PATCH</Select.Option>
              <Select.Option value='DELETE'>DELETE</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label={t('interfaceDic')} name='interfaceDic' rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
          <Form.Item label={t('interfaceDesc')} name='interfaceDesc' rules={[{ required: true }]}>
            <Input placeholder={t('pleaseEnter')} />
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default memo(SystemMenu);
