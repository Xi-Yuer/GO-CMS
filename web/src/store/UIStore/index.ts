import { createSlice } from '@reduxjs/toolkit';
import { constants } from '@/constant';
import { cache } from '@/utils';
import { MenuTheme } from 'antd';
import { menuType } from '@/types/menus';

interface IUIStore {
  isFold: boolean;
  langMode: 'enUS' | 'zhCN';
  themeMode: MenuTheme | undefined;
  TabHeader: menuType[];
  defaultSelectedKeys: string[];
  defaultOpenKeys: string[];
  uploadBar: {
    showSider: boolean;
    showDrawer: boolean;
  };
}

const useUIStoreSlice = createSlice({
  name: 'UIStore',
  initialState: {
    isFold: false,
    langMode: 'zhCN',
    themeMode: 'light',
    TabHeader: [],
    defaultSelectedKeys: [],
    defaultOpenKeys: [],
    uploadBar: {
      showSider: true,
      showDrawer: false,
    },
  } as IUIStore,
  reducers: {
    changeFold(state, action) {
      state.isFold = action.payload;
    },

    changeLang(state, action: { payload: 'enUS' | 'zhCN'; type: string }) {
      state.langMode = action.payload;
    },

    changeThemeMode(state, action: { payload: 'dark' | 'light' | undefined; type: string }) {
      state.themeMode = action.payload;
      if (action.payload === 'dark') {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
      cache.set(constants.localStorage.lang, action.payload);
    },

    addTabHeader(state, action: { payload: menuType; type: string }) {
      const tabHeader = action.payload;
      state.TabHeader.forEach((item, index) => {
        if (item.pageID === tabHeader.pageID) {
          state.TabHeader.splice(index, 1);
        }
      });
      state.TabHeader.push(tabHeader);
    },

    delTabHeader(state, action: { payload: number; type: string }) {
      const result = state.TabHeader;
      result.splice(action.payload, 1);
      state.TabHeader = result;
    },

    changeTabHeader(state, action: { payload: menuType[]; type: string }) {
      state.TabHeader = action.payload;
    },

    changeDefaultOpenKeys(state, action: { payload: string[]; type: string }) {
      state.defaultOpenKeys = action.payload;
    },

    changeDefaultSelectedKeys(state, action: { payload: string[]; type: string }) {
      state.defaultSelectedKeys = action.payload;
    },
    changeUploadBarShowSider(state, action: { payload: boolean; type: string }) {
      state.uploadBar.showSider = action.payload;
    },
    changeUploadBarShowDrawer(state, action: { payload: boolean; type: string }) {
      state.uploadBar.showDrawer = action.payload;
    },
  },
});

export const {
  changeFold,
  changeLang,
  changeThemeMode,
  addTabHeader,
  delTabHeader,
  changeTabHeader,
  changeDefaultOpenKeys,
  changeDefaultSelectedKeys,
  changeUploadBarShowSider,
  changeUploadBarShowDrawer,
} = useUIStoreSlice.actions;
export default useUIStoreSlice.reducer;
