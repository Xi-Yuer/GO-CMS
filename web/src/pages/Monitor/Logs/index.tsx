import { FC, memo, useEffect, useState } from 'react';
import { getSysTemLogsRequest, ISysTemLogsResponse } from '@/service/api/logs';
import { Pagination, Table, TableColumnsType, Tag } from 'antd';
import { useTranslation } from 'react-i18next';
import dayjs from 'dayjs';

const Logs: FC = () => {
  const { t } = useTranslation();
  const [list, setList] = useState<ISysTemLogsResponse[]>([]);
  const [limit, setLimit] = useState(10);
  const [curPage, setCurPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);

  const methodColorMap: Record<string, string> = {
    GET: 'green',
    POST: 'blue',
    PUT: 'orange',
    DELETE: 'red',
    PATCH: 'purple',
    OPTIONS: 'cyan',
  };

  const getPageData = async () => {
    try {
      setLoading(true);
      const result = await getSysTemLogsRequest({ limit, offset: (curPage - 1) * limit });
      setList(result.data.list);
      setTotal(result.data.total);
    } finally {
      setLoading(false);
    }
  };

  const columns: TableColumnsType<ISysTemLogsResponse> = [
    {
      title: t('userName'),
      dataIndex: 'userName',
      key: 'userName',
      align: 'center',
    },
    {
      title: t('requestMethod'),
      dataIndex: 'requestMethod',
      key: 'requestMethod',
      align: 'center',
      render: (text) => {
        return <Tag color={methodColorMap[text]}>{text}</Tag>;
      },
    },
    {
      title: t('requestPath'),
      dataIndex: 'requestPath',
      key: 'requestPath',
      align: 'center',
    },
    {
      title: t('requestDuration'),
      dataIndex: 'requestDuration',
      key: 'requestDuration',
      align: 'center',
      render: (text) => {
        return <span className='text-[#5bb4ef]'>{text}</span>;
      },
    },
    {
      title: t('requestStatus'),
      dataIndex: 'requestStatus',
      key: 'requestStatus',
      align: 'center',
      render: (text) => {
        return <Tag color={text === '200' ? 'green' : 'warning'}>{text}</Tag>;
      },
    },
    {
      title: t('createTime'),
      dataIndex: 'createTime',
      key: 'createTime',
      align: 'center',
      render: (text) => {
        return <span>{dayjs(text).format('YYYY-MM-DD HH:mm:ss')}</span>;
      },
    },
  ];

  useEffect(() => {
    getPageData().then();
  }, [limit, curPage]);
  return (
    <>
      <Table dataSource={list} loading={loading} columns={columns} rowKey='id' pagination={false} bordered></Table>
      <Pagination
        className='flex justify-end mt-2'
        total={total}
        onChange={(page, pageSize) => {
          setCurPage(page);
          setLimit(pageSize);
        }}></Pagination>
    </>
  );
};

export default memo(Logs);
