import { forwardRef, memo, Ref, useImperativeHandle, useRef } from 'react';
import { Button, Drawer, message, Progress, Table, TableColumnsType, Upload, UploadProps } from 'antd';
import { DoubleLeftOutlined, UploadOutlined } from '@ant-design/icons';
import { useAppDispatch, useAppSelector } from '@/store';
import { changeUploadBarShowDrawer } from '@/store/UIStore';
import { IUploadFile, updateUploadFileList } from '@/store/UploadStore';
import { RcFile } from 'antd/es/upload';
import { BigFileUpload, calculateFileHash } from '@/pages/File/Upload/hooks/useUpload.ts';
import { checkFileUploadSize } from '@/service/api/file';
import { emitGetFileList, emitUploadFile } from '@/utils/event';
import { formatFileSize } from '@/utils/format';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/Auth';
import { constants } from '@/constant';

export interface AppUploadsRefProps {
  addUploadFile: (file: RcFile) => void;
}

const Index = forwardRef((_, ref: Ref<AppUploadsRefProps>) => {
  const uploadBar = useAppSelector((state) => state.UIStore.uploadBar);
  const uploadList = useAppSelector((state) => state.UploadStore.uploadList);
  const [messageApi, contextHolder] = message.useMessage();
  const dispatch = useAppDispatch();
  const { t } = useTranslation();
  const columns: TableColumnsType<IUploadFile> = [
    {
      title: t('fileName'),
      dataIndex: 'fileName',
      key: 'fileName',
      align: 'center',
    },
    {
      title: t('fileSize'),
      dataIndex: 'fileSize',
      key: 'fileSize',
      align: 'center',
      render: (text) => {
        return <span>{formatFileSize(text)}</span>;
      },
    },
    {
      title: t('progress'),
      dataIndex: 'progress',
      width: 300,
      key: 'progress',
      align: 'center',
      render: (text) => {
        return <Progress percent={text} size='small' />;
      },
    },
  ];
  const uploadFileListRef = useRef<IUploadFile[]>([]);
  useImperativeHandle(ref, () => ({
    addUploadFile: async (file: RcFile) => {
      // 1. 计算文件哈希
      messageApi.loading({ content: t('calculateFileHash'), key: 'calculateFileHash', duration: 0 });
      const fileHash = await calculateFileHash(file);
      messageApi.success({ content: t('startUpload'), key: 'calculateFileHash' });
      dispatch(changeUploadBarShowDrawer(true));
      const uploadSizeResult = await checkFileUploadSize(fileHash);
      let uploadFile = {
        fileName: file.name,
        fileSize: file.size,
        identifier: fileHash,
        progress: Math.ceil((uploadSizeResult.data.hasReadySize / file.size) * 100),
      } as IUploadFile;
      uploadFileListRef.current = [...uploadFileListRef.current, uploadFile];
      // 文件秒传
      dispatch(updateUploadFileList(uploadFileListRef.current));
      if (uploadSizeResult.data.hasReadySize !== file.size) {
        // 无需上传已上传过的部分
        file.slice(uploadSizeResult.data.hasReadySize);
        await BigFileUpload(file, uploadSizeResult.data.hasReadySize, fileHash, (size) => {
          const list = uploadFileListRef.current.map((item) => {
            if (item.identifier === fileHash) {
              return {
                ...item,
                progress: size,
              };
            } else {
              return item;
            }
          });
          // 根据 identify 去重
          const exist: string[] = [];
          const uploadFileList: IUploadFile[] = [];
          list.forEach((item) => {
            if (!exist.includes(item.identifier)) {
              exist.push(item.identifier);
              uploadFileList.push(item);
            }
          });
          uploadFileListRef.current = uploadFileList;
          dispatch(updateUploadFileList(uploadFileList));
        });
      } else {
        // 根据 identify 去重
        const list = uploadFileListRef.current.map((item) => {
            if (item.identifier === fileHash) {
              return {
                ...item,
                progress: 100,
              };
            } else {
              return item;
            }
          }),
          exist: string[] = [],
          uploadFileList: IUploadFile[] = [];
        list.forEach((item) => {
          if (!exist.includes(item.identifier)) {
            exist.push(item.identifier);
            uploadFileList.push(item);
          }
        });
        uploadFileListRef.current = uploadFileList;
        dispatch(updateUploadFileList(uploadFileList));
      }
      messageApi.success({ content: t('uploadFinish'), key: 'calculateFileHash' });
      emitGetFileList();
    },
  }));

  const props: UploadProps = {
    beforeUpload: async (file) => {
      emitUploadFile(file);
      return false;
    },
  };

  return (
    <div>
      {contextHolder}
      {uploadBar.showSider && (
        <div className='flex justify-center items-center absolute z-10 right-0 w-5 cursor-pointer h-screen top-0 bg-[#ffffff50] dark:bg-[#001620] hover:bg-white dark:hover:bg-[#001620] shadow'>
          <div
            onClick={() => dispatch(changeUploadBarShowDrawer(true))}
            className='w-full h-full flex justify-center items-center dark:text-white animate-pulse'>
            <DoubleLeftOutlined />
          </div>
          <Drawer
            title={t('fileList')}
            width='60%'
            mask={false}
            onClose={() => dispatch(changeUploadBarShowDrawer(false))}
            open={uploadBar.showDrawer}
            extra={
              <Auth permission={constants.permissionDicMap.UPLOAD_FILE}>
                <Upload showUploadList={false} {...props}>
                  <Button type='primary' icon={<UploadOutlined />}>
                    {t('continueUploadFile')}
                  </Button>
                </Upload>
              </Auth>
            }>
            <Table dataSource={uploadList} columns={columns} bordered rowKey='identifier'></Table>
          </Drawer>
        </div>
      )}
    </div>
  );
});

export default memo(Index);
