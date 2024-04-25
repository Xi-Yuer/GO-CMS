import { Form, TableProps, Tag } from 'antd';
import {
  createUsersRequest,
  deleteUsersRequest,
  getDepartmentRequest,
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
import { useEffect, useState } from 'react';
import dayjs from 'dayjs';
import { Md5 } from 'ts-md5';
import { useTranslation } from 'react-i18next';
import { AxiosResponse } from 'axios';

export const useUserPageHooks = () => {
  const [searchFormRef] = Form.useForm();
  const [editFormRef] = Form.useForm();
  const { t } = useTranslation();
  const columns: TableProps<IUserResponse>['columns'] = [
    {
      title: t('id'),
      dataIndex: 'id',
      key: 'id',
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
      render: (_, { createTime }) => {
        return <span>{dayjs(createTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('updateTime'),
      dataIndex: 'updateTime',
      key: 'updateTime',
      render: (_, { updateTime }) => {
        return <span>{dayjs(updateTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('operate'),
      key: 'action',
      align: 'center',
      render: (_, { id }) => (
        <div className='gap-2 flex text-[#5bb4ef] items-center cursor-pointer justify-center'>
          <span onClick={() => editUserAction(id)}>{t('edit')}</span>
          <span className='text-red-500' onClick={() => deleteUsersAction(id)}>
            {t('delete')}
          </span>
        </div>
      ),
    },
  ];

  const [departments, setDepartments] = useState<IDepartmentResponse[]>([]);
  const [users, setUsers] = useState<IUserResponse[]>([]);
  const [roles, setRoles] = useState<IRoleResponse[]>([]);
  const [editUserModalOpen, setEditUserModalOpen] = useState(false);
  const [currentEditUser, setCurrentEditUser] = useState<IUserResponse | undefined>();
  const [isEdit, setIsEdit] = useState(true);

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
    getUsersRequest({ limit: 20, offset: 0, ...value }).then((res) => {
      setUsers(res.data);
    });
  };

  const getRoleAction = () => {
    getDepartmentRequest().then((res) => {
      setDepartments(res.data);
    });
  };

  const updateUserAction = (params: IUpdateUserParams) => {
    updateUsersRequest(params).then(() => getPageData());
  };

  const deleteUsersAction = (id: string) => {
    deleteUsersRequest(id).then(() => getPageData());
  };

  const editUserAction = async (id?: string) => {
    const { data: user } = id ? await getUserRequest(id) : ({ data: newUserParams as any } as AxiosResponse<IUserResponse>);
    const { data: roles } = await getRolesRequest();
    setIsEdit(!!id);
    setRoles(roles);
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
    getRoleAction();
  }, []);

  return {
    users,
    roles,
    departments,
    columns,
    searchFormRef,
    editFormRef,
    editUserModalOpen,
    isEdit,
    onFinish,
    onReset,
    editUserConfirm,
    setEditUserModalOpen,
    editUserAction,
  };
};
