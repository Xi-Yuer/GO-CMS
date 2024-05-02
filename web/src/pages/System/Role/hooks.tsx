import React, { ReactNode, useEffect, useState } from 'react';
import { Button, DatePicker, Input, TableProps } from 'antd';
import { DownloadOutlined } from '@ant-design/icons';
import {
  deleteRolesRequest,
  exportRolesRequest,
  getAllMenusRequest,
  getRolesRequest,
  IQueryRoleParams,
  IRoleResponse,
  IUpdateRoleParams,
  updateRoleRequest,
} from '@/service';
import { useTranslation } from 'react-i18next';
import { useSearchFrom } from '@/hooks/useSearchForm.tsx';
import dayjs from 'dayjs';
import { useForm } from 'antd/es/form/Form';
import { menuType } from '@/types/menus';

export const useRolePageHooks = () => {
  const [formRef] = useForm();
  const [limit, setLimit] = useState(20);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<React.Key[]>([]);
  const [roles, setRoles] = useState<IRoleResponse[]>([]);
  const [menus, setMenus] = useState<menuType[]>([]);
  const [currentEditRole, setCurrentEditRole] = useState<IRoleResponse>();
  const [isEdit, setIsEdit] = useState(false);
  const [editRoleModalOpen, setEditRoleModalOpen] = useState(false);
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
    const menusResult = await getAllMenusRequest();
    setMenus(menusResult.data);
    setCurrentEditRole(row);
    setIsEdit(true);
    setEditRoleModalOpen(true);
  };

  const onFinish = () => {
    formRef.validateFields().then((values) => {
      updateRoleRequest({ ...currentEditRole, ...values } as IUpdateRoleParams).then(() => {
        setEditRoleModalOpen(false);
        getPageData();
      });
    });
  };

  const columns: TableProps<IRoleResponse>['columns'] = [
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
    columns,
    roles,
    searchConfig,
    SearchFormComponent,
    selected,
    loading,
    editRoleModalOpen,
    isEdit,
    formRef,
    menus,
    currentEditRole,
    onFinish,
    setEditRoleModalOpen,
    setSelected,
    getPageData,
    setPage,
    setLimit,
  };
};
