import { forwardRef, memo, Ref, useImperativeHandle, useRef } from 'react';
import { Drawer, message, Progress, Table, TableColumnsType } from 'antd';
import { DoubleLeftOutlined } from '@ant-design/icons';
import { useAppDispatch, useAppSelector } from '@/store';
import { changeUploadBarShowDrawer } from '@/store/UIStore';
import { IUploadFile, updateUploadFileList } from '@/store/UploadStore';
import { RcFile } from 'antd/es/upload';
import { BigFileUpload, calculateFileHash } from '@/pages/File/Upload/hooks/useUpload.ts';
import { checkFileUploadSize } from '@/service/api/file';
import { emitGetFileList } from '@/utils/event';
import { formatFileSize } from '@/utils/format';

export interface AppUploadsRefProps {
  addUploadFile: (file: RcFile) => void;
}

const Index = forwardRef((_, ref: Ref<AppUploadsRefProps>) => {
  const uploadBar = useAppSelector((state) => state.UIStore.uploadBar);
  const uploadList = useAppSelector((state) => state.UploadStore.uploadList);
  const dispatch = useAppDispatch();
  const [messageApi, contextHolder] = message.useMessage();
  const columns: TableColumnsType<IUploadFile> = [
    {
      title: '文件名',
      dataIndex: 'fileName',
      key: 'fileName',
      align: 'center',
    },
    {
      title: '大小',
      dataIndex: 'fileSize',
      key: 'fileSize',
      align: 'center',
      render: (text) => {
        return <span>{formatFileSize(text)}</span>;
      },
    },
    {
      title: '进度',
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
      messageApi.loading({ content: '计算文件哈希,请稍后...', key: 'calculateFileHash', duration: 0 });
      const fileHash = await calculateFileHash(file);
      messageApi.success({ content: '开始上传文件', key: 'calculateFileHash' });
      dispatch(changeUploadBarShowDrawer(true));
      let uploadFile = {
        fileName: file.name,
        fileSize: file.size,
        identifier: fileHash,
        progress: 0,
      } as IUploadFile;
      uploadFileListRef.current = [...uploadFileListRef.current, uploadFile];
      const uploadSizeResult = await checkFileUploadSize(fileHash);
      // 文件秒传
      if (uploadSizeResult.data.hasReadySize !== file.size) {
        // 无需上传已上传过的部分
        file.slice(uploadSizeResult.data.hasReadySize);
        await BigFileUpload(file, fileHash, (size) => {
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
      messageApi.success({ content: '文件已上传完成', key: 'calculateFileHash' });
      emitGetFileList();
    },
  }));
  return (
    <div>
      {contextHolder}
      {uploadBar.showSider && (
        <div className='flex justify-center items-center absolute z-10 right-0 w-5 cursor-pointer h-screen top-0 bg-[#ffffff50] hover:bg-white shadow'>
          <div onClick={() => dispatch(changeUploadBarShowDrawer(true))} className='w-full h-full flex justify-center items-center'>
            <DoubleLeftOutlined />
          </div>
          <Drawer title='文件上传列表' width='60%' mask={false} onClose={() => dispatch(changeUploadBarShowDrawer(false))} open={uploadBar.showDrawer}>
            <Table dataSource={uploadList} columns={columns} bordered rowKey='identifier'></Table>
          </Drawer>
        </div>
      )}
    </div>
  );
});

export default memo(Index);
