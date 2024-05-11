import { Key, ReactNode, useEffect, useRef, useState } from 'react';
import { Button, DatePicker, Input, TableProps, Tag, TreeDataNode, TreeProps } from 'antd';
import { DownloadOutlined } from '@ant-design/icons';
import {
  bindUserRequest,
  createRoleRequest,
  deBindUserRequest,
  deleteRolesRequest,
  exportRolesRequest,
  getAllMenusRequest,
  getRolesRequest,
  getUserByRoleIDRequest,
  IQueryRoleParams,
  IRoleResponse,
  IUpdateRoleParams,
  IUserResponse,
  updateRoleRequest,
} from '@/service';
import { useTranslation } from 'react-i18next';
import { useSearchFrom } from '@/hooks/useSearchForm.tsx';
import dayjs from 'dayjs';
import { useForm } from 'antd/es/form/Form';
import { buildInterfaceToAntdTree, buildMenuToAntdTree, getAllChildrenMenusID, getAllInterfaceKeys } from '@/utils';
import { getInterfaceAllListRequest } from '@/service/api/interface';
import { IUserPageRefProps } from '@/pages/System/User/hooks.tsx';
import Auth from '@/components/Auth';
import { constants } from '@/constant';

export const useRolePageHooks = () => {
  const userPageRef = useRef<IUserPageRefProps>(null);
  const [formRef] = useForm();
  const [limit, setLimit] = useState(10);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<Key[]>([]);
  const [roleMenusSelected, setRoleMenusSelected] = useState<string[]>([]);
  const [roleInterfaceSelected, setRoleInterfaceSelected] = useState<string[]>([]);
  const [roles, setRoles] = useState<IRoleResponse[]>([]);
  const [menus, setMenus] = useState<TreeDataNode[]>([]);
  const [allInterface, setAllInterface] = useState<TreeDataNode[]>([]);
  const [currentEditRole, setCurrentEditRole] = useState<IRoleResponse>();
  const [isEdit, setIsEdit] = useState(false);
  const [editRoleModalOpen, setEditRoleModalOpen] = useState(false);
  const [editRolePermissionOpen, setEditRolePermissionOpen] = useState(false);
  const [editRoleUnderUserOpen, setEditRoleUnderUserOpen] = useState(false);
  const { t } = useTranslation();
  const searchConfig: { label: string; name: keyof IQueryRoleParams; component: ReactNode }[] = [
    {
      label: '角色ID',
      name: 'id',
      component: <Input allowClear />,
    },
    {
      label: '角色名称',
      name: 'roleName',
      component: <Input allowClear />,
    },
    {
      label: '角色描述',
      name: 'description',
      component: <Input allowClear />,
    },
    {
      label: '创建时间',
      name: 'startTime',
      component: <DatePicker allowClear style={{ width: '200px' }} />,
    },
    {
      label: '更新时间',
      name: 'endTime',
      component: <DatePicker allowClear style={{ width: '200px' }} />,
    },
  ];
  const { SearchFormComponent } = useSearchFrom({
    getDataRequestFn: (values) => getPageData(values),
    onNewRecordFn: () => {
      setIsEdit(false);
      setEditRoleModalOpen(true);
    },
    formItems: searchConfig,
    operateComponent: !!selected.length && (
      <Auth permission={constants.permissionDicMap.EXPORT_ROLE}>
        <Button type='primary' icon={<DownloadOutlined />} onClick={() => exportRolesRequest(selected)}>
          {t('export')}
        </Button>
      </Auth>
    ),
    formName: 'roleSearchUserForm',
  });

  const getPageData = (values?: IRoleResponse) => {
    setLoading(true);
    getRolesRequest({ limit, offset: (page - 1) * limit, ...values } as IQueryRoleParams)
      .then((res) => {
        setRoles(res.data.list);
        setTotal(res.data.total);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const deleteRoleAction = (id: string) => {
    deleteRolesRequest(id).then(() => getPageData());
  };

  const editRoleAction = async (row: IRoleResponse) => {
    setCurrentEditRole(row);
    setIsEdit(true);
    setEditRoleModalOpen(true);
  };

  const onFinish = () => {
    formRef.validateFields().then((values) => {
      // 编辑
      if (isEdit) {
        updateRoleRequest({ ...currentEditRole, ...values } as IUpdateRoleParams).then(() => {
          setEditRoleModalOpen(false);
          getPageData();
        });
      } else {
        // 新增
        createRoleRequest(values).then(() => {
          setEditRoleModalOpen(false);
          getPageData();
        });
      }
    });
  };

  const [allInterfaceKeys, setAllInterfaceKeys] = useState<string[]>([]);
  const editRolePermissionAction = async (row: IRoleResponse) => {
    const allMenusResult = await getAllMenusRequest();
    const allInterfaceResult = await getInterfaceAllListRequest();
    const menus = buildMenuToAntdTree(allMenusResult.data);
    const interfaceToAntdTree = buildInterfaceToAntdTree(allInterfaceResult.data);
    setAllInterfaceKeys(getAllInterfaceKeys(allInterfaceResult.data));
    setAllInterface(interfaceToAntdTree);
    setMenus(menus);
    const allChildrenMenuID = getAllChildrenMenusID(allMenusResult.data);
    // 过滤父级菜单
    setRoleMenusSelected(row.pageID?.filter((item) => allChildrenMenuID.includes(item)) || []);
    setRoleInterfaceSelected(row.interfaceID || []);
    setCurrentEditRole(row);
    setEditRolePermissionOpen(true);
  };

  const [menuSelectsIsChanged, setMenuSelectsIsChanged] = useState(false);
  const onPageTreeCheck: TreeProps['onCheck'] = (checked, e) => {
    setMenuSelectsIsChanged(true);
    setRoleMenusSelected([...(checked as string[]), ...(e.halfCheckedKeys as string[])] as string[]);
  };

  const onInterfaceTreeCheck: TreeProps['onCheck'] = (checkedKeysValue) => {
    setRoleInterfaceSelected(checkedKeysValue as string[]);
  };

  const editPermissionConfirm = () => {
    updateRoleRequest({
      ...currentEditRole,
      pageID: menuSelectsIsChanged ? roleMenusSelected : currentEditRole?.pageID,
      interfaceID: roleInterfaceSelected.filter((item) => !allInterfaceKeys.includes(item)),
    } as IUpdateRoleParams).then(() => {
      setEditRolePermissionOpen(false);
      getPageData();
    });
  };

  const [roleUnderUsersList, setRoleUnderUsersList] = useState<IUserResponse[]>([]);
  const [roleUnderUserListTotal, setRoleUnderUserListTotal] = useState(0);
  const [roleUnderUserLimit, setRoleUnderUserLimit] = useState(10);
  const [roleUnderUserPage, setRoleUnderUserPage] = useState(1);

  const editRoleUnderUserAction = async (row: IRoleResponse) => {
    setCurrentEditRole(row);
    setRoleUnderUserPage(1);
    setRoleUnderUserLimit(10);
    await getUnderRoleUsersAction(row, 1, 10);
    setEditRoleUnderUserOpen(true);
  };
  const getUnderRoleUsersAction = async (row: IRoleResponse, page: number, pageSize: number) => {
    const userResponse = await getUserByRoleIDRequest({ roleID: row.id, limit: pageSize, offset: (page - 1) * pageSize });
    setRoleUnderUsersList(userResponse.data.list);
    setRoleUnderUserListTotal(userResponse.data.total);
  };
  const deBindRoleUnderUserAction = async (row: IUserResponse) => {
    if (!currentEditRole?.id) return;
    await deBindUserRequest({ userID: row.id, roleID: currentEditRole!.id });
    await getUnderRoleUsersAction(currentEditRole, roleUnderUserPage, roleUnderUserLimit);
    userPageRef?.current?.getPageData();
  };

  const roleUnderUserPageChange = async (page: number, pageSize: number) => {
    if (!currentEditRole?.id) return;
    setRoleUnderUserPage(page);
    setRoleUnderUserLimit(pageSize);
    await getUnderRoleUsersAction(currentEditRole, page, pageSize);
  };

  const addRoleUnderUserAction = async (row: IUserResponse) => {
    await bindUserRequest({ userID: row.id, roleID: currentEditRole!.id });
    await getUnderRoleUsersAction(currentEditRole!, roleUnderUserPage, roleUnderUserLimit);
  };
  const userColumns: TableProps<IUserResponse>['columns'] = [
    {
      title: t('index'),
      align: 'center',
      render: (_, __, index) => (roleUnderUserPage - 1) * roleUnderUserLimit + index + 1,
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
      title: t('operate'),
      key: 'action',
      align: 'center',
      render: (_, row) => (
        <div className='gap-2 flex text-red-500 items-center cursor-pointer justify-center'>
          <Auth permission={constants.permissionDicMap.UNBIND_USER}>
            <span onClick={() => deBindRoleUnderUserAction(row)}>{t('remove')}</span>
          </Auth>
        </div>
      ),
    },
  ];

  const unSelectedUserColumns: TableProps<IUserResponse>['columns'] = [
    {
      title: t('index'),
      align: 'center',
      render: (_, __, index) => (roleUnderUserPage - 1) * roleUnderUserLimit + index + 1,
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
      render: (_, { status }) => {
        return (
          <Tag color={status ? 'green' : 'red'} className='cursor-pointer'>
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
        return <span className='text-[12px]'>{dayjs(createTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('updateTime'),
      dataIndex: 'updateTime',
      key: 'updateTime',
      align: 'center',
      render: (_, { updateTime }) => {
        return <span className='text-[12px]'>{dayjs(updateTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('operate'),
      key: 'action',
      align: 'center',
      render: (_, row) => (
        <div className='gap-2 flex text-[#48ade9] items-center cursor-pointer justify-center'>
          <span onClick={() => deBindRoleUnderUserAction(row)}>{t('add')}</span>
        </div>
      ),
    },
  ];

  const roleColumns: TableProps<IRoleResponse>['columns'] = [
    {
      title: t('index'),
      align: 'center',
      render: (_, __, index) => (page - 1) * limit + index + 1,
    },
    {
      title: t('roleName'),
      dataIndex: 'roleName',
      key: 'roleName',
    },
    {
      title: t('roleDescription'),
      dataIndex: 'description',
      key: 'description',
    },
    {
      title: t('createTime'),
      dataIndex: 'createTime',
      key: 'createTime',
      align: 'center',
      render: (_, { createTime }) => {
        return <span className='text-sm'>{dayjs(createTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('updateTime'),
      dataIndex: 'updateTime',
      key: 'updateTime',
      align: 'center',
      render: (_, { updateTime }) => {
        return <span className='text-sm'>{dayjs(updateTime).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
    {
      title: t('operate'),
      key: 'action',
      align: 'center',
      render: (_, row) => {
        return (
          <div className='gap-2 flex text-[#5bb4ef] items-center cursor-pointer justify-center'>
            <Auth permission={[constants.permissionDicMap.UPDATE_ROLE]}>
              <span onClick={() => editRolePermissionAction(row)}>{t('permissionEdit')}</span>
            </Auth>
            <Auth permission={[constants.permissionDicMap.UPDATE_ROLE]}>
              <span onClick={() => editRoleUnderUserAction(row)}>{t('authUser')}</span>
            </Auth>
            <Auth permission={constants.permissionDicMap.UPDATE_ROLE}>
              <span onClick={() => editRoleAction(row)}>{t('edit')}</span>
            </Auth>
            <Auth permission={constants.permissionDicMap.DELETE_ROLE}>
              <span className='text-red-500' onClick={() => deleteRoleAction(row.id)}>
                {t('delete')}
              </span>
            </Auth>
          </div>
        );
      },
    },
  ];

  useEffect(() => {
    if (editRoleModalOpen && isEdit) {
      formRef.setFieldsValue(currentEditRole);
    }
  }, [editRoleModalOpen]);

  useEffect(() => {
    getPageData();
  }, [limit, page]);

  return {
    userPageRef,
    limit,
    roleUnderUserLimit,
    total,
    roleUnderUserListTotal,
    roleColumns,
    userColumns,
    unSelectedUserColumns,
    roles,
    SearchFormComponent,
    roleMenusSelected,
    roleInterfaceSelected,
    loading,
    editRoleModalOpen,
    isEdit,
    formRef,
    menus,
    currentEditRole,
    editRolePermissionOpen,
    editRoleUnderUserOpen,
    allInterface,
    roleUnderUsersList,
    addRoleUnderUserAction,
    onPageTreeCheck,
    onInterfaceTreeCheck,
    onFinish,
    setEditRoleModalOpen,
    setSelected,
    setPage,
    setLimit,
    roleUnderUserPageChange,
    setEditRolePermissionOpen,
    setEditRoleUnderUserOpen,
    editPermissionConfirm,
  };
};
