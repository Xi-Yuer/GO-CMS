import { FC, memo, useEffect, useState } from 'react';
import { deleteFileRequest, FileResponse, getCookie, getFileList } from '@/service/api/file';
import { Button, Pagination, Table, TableColumnsType, Upload, UploadProps } from 'antd';
import dayjs from 'dayjs';
import { addListenerGetFileList, emitUploadFile, removeListenerGetFileList } from '@/utils/event';
import { formatFileSize } from '@/utils/format';
import { UploadOutlined } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/Auth';
import { constants } from '@/constant';

const File: FC = () => {
  const [limit, setLimit] = useState(10);
  const [total, setTotal] = useState(0);
  const [currentPage, setCurrentPage] = useState(1);
  const [list, setList] = useState<FileResponse[]>([]);
  const [loading, setLoading] = useState(false);
  const { t } = useTranslation();

  const getPageList = async () => {
    setLoading(true);
    const result = await getFileList({ limit, offset: (currentPage - 1) * limit });
    setList(result.data.list);
    setTotal(result.data.total);
    setLoading(false);
  };

  const columns: TableColumnsType<FileResponse> = [
    {
      title: t('fileName'),
      dataIndex: 'fileName',
      key: 'fileName',
      align: 'center',
    },
    {
      title: t('fileSize'),
      dataIndex: 'fileSize',
      key: 'fileName',
      align: 'center',
      render: (text) => {
        return <span>{formatFileSize(text)}</span>;
      },
    },
    {
      title: t('uploadUser'),
      dataIndex: 'uploadUser',
      key: 'uploadUser',
      align: 'center',
    },
    {
      title: t('uploadTime'),
      dataIndex: 'uploadTime',
      key: 'uploadTime',
      align: 'center',
      render: (text) => {
        return dayjs(text).format('YYYY-MM-DD HH:mm:ss');
      },
    },
    {
      title: t('operate'),
      dataIndex: 'operate',
      align: 'center',
      render: (_, { fileID }) => {
        return (
          <div className='flex gap-2 text-[#00a7fb] justify-center items-center cursor-pointer'>
            <Auth permission={constants.permissionDicMap.DOWNLOAD_FILE}>
              <span
                onClick={async () => {
                  await getCookie();
                  const a = document.createElement('a');
                  a.href = `/cms/upload/download/aHref/${fileID}`;
                  document.body.appendChild(a);
                  a.click();
                  document.body.removeChild(a);
                }}>
                {t('download')}
              </span>
            </Auth>
            <Auth permission={constants.permissionDicMap.DELETE_FILE}>
              <span
                className='text-red-500'
                onClick={async () => {
                  await deleteFileRequest(fileID);
                  await getPageList();
                }}>
                {t('delete')}
              </span>
            </Auth>
          </div>
        );
      },
    },
  ];
  useEffect(() => {
    getPageList().then();
  }, [limit, currentPage]);

  useEffect(() => {
    addListenerGetFileList(getPageList);
    return () => {
      removeListenerGetFileList(getPageList);
    };
  }, []);

  const props: UploadProps = {
    beforeUpload: async (file) => {
      emitUploadFile(file);
      return false;
    },
  };
  return (
    <>
      <div className='mb-2 flex justify-between items-center bg-white p-4 rounded dark:bg-[#001620]'>
        <span className='font-bold'>{t('fileList')}</span>
        <Auth permission={constants.permissionDicMap.UPLOAD_FILE}>
          <Upload showUploadList={false} {...props}>
            <Button type='primary' icon={<UploadOutlined />}>
              {t('uploadFile')}
            </Button>
          </Upload>
        </Auth>
      </div>
      <Table dataSource={list} loading={loading} columns={columns} pagination={false} rowKey='fileID' bordered></Table>
      <Pagination
        className='flex justify-end mt-2'
        total={total}
        onChange={(page, pageSize) => {
          setCurrentPage(page);
          setLimit(pageSize);
        }}></Pagination>
    </>
  );
};

export default memo(File);
