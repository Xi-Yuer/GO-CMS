import React, { ReactNode, useEffect, useState } from 'react';
import { Button, DatePicker, Input, TableProps } from 'antd';
import { DownloadOutlined } from '@ant-design/icons';
import { deleteRolesRequest, exportRolesRequest, getRolesRequest, IQueryRoleParams, IRoleResponse } from '@/service';
import { useTranslation } from 'react-i18next';
import { useSearchFrom } from '@/hooks/useSearchForm.tsx';
import dayjs from 'dayjs';

export const useRolePageHooks = () => {
  const [limit, setLimit] = useState(20);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<React.Key[]>([]);
  const [roles, setRoles] = useState<IRoleResponse[]>([]);
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
    onNewRecordFn: () => {},
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
      render: (_, { id }) => {
        return (
          <div className='gap-2 flex text-[#5bb4ef] items-center cursor-pointer justify-center'>
            <span onClick={() => {}}>{t('edit')}</span>
            <span className='text-red-500' onClick={() => deleteRoleAction(id)}>
              {t('delete')}
            </span>
          </div>
        );
      },
    },
  ];

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
    setSelected,
    getPageData,
    setPage,
    setLimit,
  };
};
