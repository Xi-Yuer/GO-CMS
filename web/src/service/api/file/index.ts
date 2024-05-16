import request from '@/service/request';
import { IHasTotalResponse, Page } from '@/service';
import { AxiosResponse } from 'axios';

export interface FileResponse {
  fileID: string;
  fileName: string;
  fileSize: number;
  uploadUser: string;
  uploadTime: string;
}

export const getFileList = (params: Page) => {
  return request.get<AxiosResponse<IHasTotalResponse<FileResponse[]>>>({
    url: '/upload',
    params,
  });
};

export interface IUploadFileChunk {
  identifier: string;
  file: Blob;
}

export const uploadFileChunk = (params: IUploadFileChunk) => {
  const formData = new FormData();
  formData.append('file', params.file);
  formData.append('identifier', params.identifier);
  return request.post({
    url: '/upload ',
    data: formData,
  });
};

export const checkFileUploadSize = (identifier: string) => {
  return request.post<AxiosResponse<{ hasReadySize: number }>>({
    url: '/upload/check',
    data: {
      identifier,
    },
  });
};

export interface IMergeFileChunk {
  identifier: string;
  fileName: string;
  fileSize: number;
}

export const mergeFileChunk = (data: IMergeFileChunk) => {
  return request.post({
    url: '/upload/finish',
    data,
  });
};

export const deleteFileRequest = (fileID: string) => {
  return request.delete({
    url: `/upload/del/${fileID}`,
  });
};

export const getCookie = () => {
  return request.get({
    url: '/auth/cookie',
  });
};
