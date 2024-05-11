import React, { ReactNode, useEffect, useImperativeHandle, useState } from 'react';
import { DatePicker, Form, Input, Select, TableProps, Tag, TreeSelect } from 'antd';
import {
  createUsersRequest,
  deleteUsersRequest,
  exportUsersRequest,
  getDepartmentRequest,
  getOutRoleUsersRequest,
  getRolesRequest,
  getUserRequest,
  getUsersRequest,
  IDepartmentResponse,
  IGetUsersParams,
  IRoleResponse,
  IUpdateUserParams,
  IUserResponse,
  updateUsersRequest,
} from '@/service';
import { useTranslation } from 'react-i18next';
import { AxiosResponse } from 'axios';
import { constants } from '@/constant';
import { Md5 } from 'ts-md5';
import dayjs from 'dayjs';

export interface IUserPageHooks {
  module?: string;
  context?: string;
  operation?: (val: IUserResponse) => Promise<void>;
}

export interface IUserPageRefProps {
  getPageData: (params?: IGetUsersParams) => Promise<AxiosResponse<any>>;
}

export const useUserPageHooks = (userPageRef: any, props?: IUserPageHooks) => {
  const [searchFormRef] = Form.useForm();
  const [editFormRef] = Form.useForm();
  const { t } = useTranslation();
  const columns: TableProps<IUserResponse>['columns'] = [
    {
      title: t('index'),
      align: 'center',
      hidden: props?.module === constants.module.ROLE,
      render: (_, __, index) => (page - 1) * limit + index + 1,
    },
    {
      title: t('account'),
      dataIndex: 'account',
      key: 'account',
    },
    {
      title: t('nickName'),
      dataIndex: 'nickname',
      key: 'nickname',
    },
    {
      title: t('role'),
      dataIndex: 'isAdmin',
      key: 'isAdmin',
      align: 'center',
      render: (_, { isAdmin }) => {
        return <Tag color={isAdmin ? 'gold' : 'green'}>{isAdmin ? t('superMan') : t('normalMan')}</Tag>;
      },
    },
    {
      title: t('status'),
      dataIndex: 'status',
      hidden: props?.module === constants.module.ROLE,
      key: 'status',
      align: 'center',
      render: (_, { status, id }) => {
        return (
          <Tag color={status === '1' ? 'green' : 'red'} className='cursor-pointer' onClick={() => updateUserAction({ id, status: status === '1' ? '0' : '1' })}>
            {status === '1' ? t('on') : t('off')}
          </Tag>
        );
      },
    },
    {
      title: t('createTime'),
      dataIndex: 'createTime',
      key: 'createTime',
      align: 'center',
      render: (_, { createTime }) => {
        return <span>{dayjs(createTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('updateTime'),
      dataIndex: 'updateTime',
      key: 'updateTime',
      align: 'center',
      render: (_, { updateTime }) => {
        return <span>{dayjs(updateTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('operate'),
      key: 'action',
      align: 'center',
      render: (_, row) =>
        props?.module !== constants.module.ROLE ? (
          <div className='gap-2 flex text-[#5bb4ef] items-center cursor-pointer justify-center'>
            <span onClick={() => editUserAction(row.id)}>{t('edit')}</span>
            <span className='text-red-500' onClick={() => deleteUsersAction(row.id)}>
              {t('delete')}
            </span>
          </div>
        ) : (
          <div className='gap-2 flex text-[#5bb4ef] items-center cursor-pointer justify-center'>
            <span
              onClick={() =>
                props?.operation &&
                props?.operation(row).then(() => {
                  getPageData();
                })
              }>
              {t('auth')}
            </span>
          </div>
        ),
    },
  ];
  const [departments, setDepartments] = useState<any[]>([]);
  const [limit, setLimit] = useState(10);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<React.Key[]>([]);
  const [users, setUsers] = useState<IUserResponse[]>([]);
  const [total, setTotal] = useState(0);
  const [roles, setRoles] = useState<IRoleResponse[]>([]);
  const [editUserModalOpen, setEditUserModalOpen] = useState(false);
  const [currentEditUser, setCurrentEditUser] = useState<IUserResponse | undefined>();
  const [isEdit, setIsEdit] = useState(true);
  const searchConfig: { label: string; name: keyof IGetUsersParams; component: ReactNode }[] = [
    {
      label: t('account'),
      name: 'account',
      component: <Input allowClear />,
    },
    {
      label: t('nickName'),
      name: 'nickname',
      component: <Input allowClear />,
    },
    {
      label: t('department'),
      name: 'departmentID',
      component: <TreeSelect allowClear treeData={departments}></TreeSelect>,
    },
    {
      label: t('status'),
      name: 'status',
      component: (
        <Select allowClear>
          <Select.Option value='1'>{t('on')}</Select.Option>
          <Select.Option value='0'>{t('off')}</Select.Option>
        </Select>
      ),
    },
    {
      label: t('createTime'),
      name: 'startTime',
      component: <DatePicker allowClear style={{ width: '100%' }} />,
    },
    {
      label: t('updateTime'),
      name: 'endTime',
      component: <DatePicker allowClear style={{ width: '100%' }} />,
    },
  ];

  const onFinish = (values: any) => getPageData(values);
  const newUserParams = {
    id: '',
    account: '',
    nickname: '',
    password: '',
    departmentID: '',
    rolesID: [],
    status: '1',
  };

  const onReset = () => {
    searchFormRef?.resetFields();
    getPageData();
  };

  const getPageData = (value?: IGetUsersParams) => {
    setLoading(true);
    const getUsersRequestFn =
      props?.module === constants.module.ROLE
        ? () =>
            getOutRoleUsersRequest(
              {
                limit: limit,
                offset: (page - 1) * limit,
                ...value,
              },
              props?.context,
            )
        : () => getUsersRequest({ limit: limit, offset: (page - 1) * limit, ...value });
    getUsersRequestFn()
      .then((res) => {
        setUsers(res.data.list);
        setTotal(res.data.total);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const getRoleAction = () => {
    getDepartmentRequest().then((res) => {
      setDepartments(mapInterface(res.data));
    });
  };

  function mapInterface(department: IDepartmentResponse[]): any[] {
    return (
      department?.map((item) => {
        return {
          title: item.departmentName,
          key: item.id,
          value: item.id,
          children: (item.children && mapInterface(item.children)) || [],
        };
      }) || []
    );
  }

  const updateUserAction = (params: IUpdateUserParams) => {
    updateUsersRequest(params).then(() => getPageData());
  };

  const deleteUsersAction = (id: string) => {
    deleteUsersRequest(id).then(() => getPageData());
  };

  const exportUsersAction = async () => {
    await exportUsersRequest(selected);
  };

  const editUserAction = async (id?: string) => {
    const { data: user } = id ? await getUserRequest(id) : ({ data: newUserParams as any } as AxiosResponse<IUserResponse>);
    const { data: roles } = await getRolesRequest();
    setIsEdit(!!id);
    setRoles(roles.list);
    setCurrentEditUser(user);
    setEditUserModalOpen(true);
  };

  useEffect(() => {
    if (editUserModalOpen) {
      editFormRef.setFieldsValue(currentEditUser);
    }
  }, [editUserModalOpen]);

  const editUserConfirm = async () => {
    const isValid = await editFormRef.validateFields();
    if (!isValid) return;
    const params: IUpdateUserParams = {
      ...editFormRef.getFieldsValue(),
      id: currentEditUser?.id,
    };
    if (editFormRef.getFieldsValue().password) {
      params.password = Md5.hashStr(editFormRef.getFieldsValue().password);
    }
    isEdit ? await updateUsersRequest(params) : await createUsersRequest(params);
    setEditUserModalOpen(false);
    getPageData();
  };

  useEffect(() => {
    getPageData();
  }, [page, limit]);

  useEffect(() => {
    getRoleAction();
  }, []);

  useImperativeHandle(userPageRef, () => ({
    getPageData,
  }));

  return {
    userPageRef,
    total,
    users,
    roles,
    departments,
    columns,
    editFormRef,
    editUserModalOpen,
    isEdit,
    limit,
    page,
    selected,
    searchConfig,
    loading,
    exportUsersAction,
    setSelected,
    getPageData,
    setLimit,
    setPage,
    onFinish,
    onReset,
    editUserConfirm,
    setEditUserModalOpen,
    editUserAction,
  };
};
