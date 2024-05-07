import { Key, ReactNode, useEffect, useState } from 'react';
import { Button, DatePicker, Input, TableProps, Tag, TreeDataNode, TreeProps } from 'antd';
import { DownloadOutlined } from '@ant-design/icons';
import {
  createRoleRequest,
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

export const useRolePageHooks = () => {
  const [formRef] = useForm();
  const [limit, setLimit] = useState(20);
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
      <Button type='primary' icon={<DownloadOutlined />} onClick={() => exportRolesRequest(selected)}>
        导出
      </Button>
    ),
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
  const editRoleUnderUserAction = async (row: IRoleResponse) => {
    const userResponse = await getUserByRoleIDRequest({ roleID: row.id, limit: 20, offset: 0 });
    setRoleUnderUsersList(userResponse.data);
    setCurrentEditRole(row);
    setEditRoleUnderUserOpen(true);
  };

  const userColumns: TableProps<IUserResponse>['columns'] = [
    {
      title: t('index'),
      width: '80',
      align: 'center',
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
      render: (_, { id }) => (
        <div className='gap-2 flex text-red-500 items-center cursor-pointer justify-center'>
          <span
            onClick={() => {
              console.log(id);
            }}>
            移除
          </span>
        </div>
      ),
    },
  ];

  const roleColumns: TableProps<IRoleResponse>['columns'] = [
    {
      title: t('index'),
      width: '80',
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
      render: (_, row) => {
        return (
          <div className='gap-2 flex text-[#5bb4ef] items-center cursor-pointer justify-center'>
            <span onClick={() => editRolePermissionAction(row)}>{t('permissionEdit')}</span>
            <span onClick={() => editRoleUnderUserAction(row)}>授权用户</span>
            <span onClick={() => editRoleAction(row)}>{t('edit')}</span>
            <span className='text-red-500' onClick={() => deleteRoleAction(row.id)}>
              {t('delete')}
            </span>
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
    limit,
    total,
    roleColumns,
    userColumns,
    roles,
    SearchFormComponent,
    roleMenusSelected,
    roleInterfaceSelected,
    loading,
    editRoleModalOpen,
    isEdit,
    formRef,
    menus,
    editRolePermissionOpen,
    editRoleUnderUserOpen,
    allInterface,
    roleUnderUsersList,
    onPageTreeCheck,
    onInterfaceTreeCheck,
    onFinish,
    setEditRoleModalOpen,
    setSelected,
    setPage,
    setLimit,
    setEditRolePermissionOpen,
    setEditRoleUnderUserOpen,
    editPermissionConfirm,
  };
};
