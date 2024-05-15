import { combineReducers, configureStore } from '@reduxjs/toolkit';
import { FLUSH, PAUSE, PERSIST, persistReducer, PURGE, REGISTER, REHYDRATE } from 'redux-persist';
import UseUIStoreSlice from '@/store/UIStore';
import UserStoreSlice from '@/store/UserStore';
import useUploadStore from '@/store/UploadStore';
import storage from 'redux-persist/es/storage';
import { useDispatch, useSelector } from 'react-redux';

const reducers = combineReducers({
  UIStore: UseUIStoreSlice,
  UserStore: UserStoreSlice,
  UploadStore: useUploadStore,
});

const persistedReducer = persistReducer(
  {
    key: 'cms',
    version: 1,
    storage,
    whitelist: ['UIStore', 'UserStore'],
    blacklist: ['UploadStore'],
  },
  reducers,
);

const store = configureStore({
  reducer: persistedReducer,
  devTools: import.meta.env.NODE_ENV !== 'production',
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

export type RootState = ReturnType<typeof store.getState>;
export const useAppSelector = <T>(selector: (state: RootState) => T) => useSelector(selector);
export const useAppDispatch = useDispatch;
export default store;
