import { createSlice } from '@reduxjs/toolkit';
import { constants } from '@/constant';
import { cache } from '@/utils';
import { MenuTheme } from 'antd';
import { menuType } from '@/types/menus';

interface IUIStore {
  isFold: boolean;
  langMode: string;
  themeMode: MenuTheme | undefined;
  TabHeader: menuType[];
  defaultSelectedKeys: string[];
  defaultOpenKeys: string[];
}

const useUIStoreSlice = createSlice({
  name: 'UIStore',
  initialState: {
    isFold: false,
    langMode: 'zh',
    themeMode: 'light',
    TabHeader: [],
    defaultSelectedKeys: [],
    defaultOpenKeys: [],
  } as IUIStore,
  reducers: {
    changeFold(state, action) {
      state.isFold = action.payload;
    },

    changeLang(state, action) {
      state.langMode = action.payload;
    },

    changeThemeMode(state, action) {
      state.themeMode = action.payload;
      if (action.payload === 'dark') {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
      cache.set(constants.localStorage.lang, action.payload);
    },

    addTabHeader(state, action) {
      const tabHeader = action.payload;
      state.TabHeader.forEach((item, index) => {
        if (item.pageID === tabHeader.pageID) {
          state.TabHeader.splice(index, 1);
        }
      });
      state.TabHeader.push(tabHeader);
    },

    delTabHeader(state, action) {
      const result = state.TabHeader;
      result.splice(action.payload, 1);
      state.TabHeader = result;
    },

    changeTabHeader(state, action) {
      state.TabHeader = action.payload;
    },

    changeDefaultOpenKeys(state, action) {
      state.defaultOpenKeys = action.payload;
    },

    changeDefaultSelectedKeys(state, action) {
      state.defaultSelectedKeys = action.payload;
    },
  },
});

export const { changeFold, changeLang, changeThemeMode, addTabHeader, delTabHeader, changeTabHeader, changeDefaultOpenKeys, changeDefaultSelectedKeys } =
  useUIStoreSlice.actions;
export default useUIStoreSlice.reducer;
