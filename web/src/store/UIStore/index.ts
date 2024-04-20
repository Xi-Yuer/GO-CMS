import { createSlice } from '@reduxjs/toolkit';
import { constants } from '@/constant';
import { cache } from '@/utils';

const useUIStoreSlice = createSlice({
  name: 'UIStore',
  initialState: {
    isFold: false,
    langMode: 'zh',
    themeMode: 'light',
  },
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
  },
});

export const { changeFold, changeLang, changeThemeMode } = useUIStoreSlice.actions;
export default useUIStoreSlice.reducer;
