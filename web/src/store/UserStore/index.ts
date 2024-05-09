import { createSlice } from '@reduxjs/toolkit';
import { cache } from '@/utils';
import { constants } from '@/constant';
import { menuType } from '@/types/menus';
import { LoginUserResponseType } from '@/service';

interface IUserStore {
  token: string | null;
  userInfo: LoginUserResponseType | undefined;
  menus: menuType[];
}

const UserStoreSlice = createSlice({
  name: 'UserStore',
  initialState: {
    token: null,
    userInfo: undefined,
    menus: [],
  } as IUserStore,
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
