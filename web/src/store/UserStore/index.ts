import { createSlice } from '@reduxjs/toolkit';
import { cache } from '@/utils';
import { constants } from '@/constant';

const UserStoreSlice = createSlice({
  name: 'UserStore',
  initialState: {
    token: null,
    userInfo: null,
    menus: [],
  },
  reducers: {
    changeToken(state, action) {
      cache.set(constants.localStorage.token, action.payload);
      state.token = action.payload;
    },
    changeUserInfo(state, action) {
      state.userInfo = action.payload;
    },
    changeMenus(state, action) {
      state.menus = action.payload;
    },
  },
});

export default UserStoreSlice.reducer;
export const { changeToken, changeUserInfo, changeMenus } = UserStoreSlice.actions;
