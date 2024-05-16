import { createSlice } from '@reduxjs/toolkit';

export interface IUploadFile {
  fileName: string;
  fileSize: number;
  identifier: string;
  progress: number;
}

interface IUploadStore {
  uploadList: IUploadFile[];
}

const useUploadStore = createSlice({
  name: 'uploadStore',
  initialState: {
    uploadList: [],
  } as IUploadStore,
  reducers: {
    addUploadFile: (state, action) => {
      state.uploadList.push(action.payload);
    },
    updateUploadFileList: (state, action) => {
      state.uploadList = action.payload;
    },
  },
});

export default useUploadStore.reducer;
export const { addUploadFile, updateUploadFileList } = useUploadStore.actions;
