import { combineReducers, configureStore } from '@reduxjs/toolkit';
import { FLUSH, PAUSE, PERSIST, persistReducer, PURGE, REGISTER, REHYDRATE } from 'redux-persist';
import useUIStoreSlice from '@/store/UIStore';
import storage from 'redux-persist/es/storage';
import { useDispatch, useSelector } from 'react-redux';

const reducers = combineReducers({
  UIStore: useUIStoreSlice,
});

const persistedReducer = persistReducer(
  {
    key: 'cms',
    version: 1,
    storage,
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
