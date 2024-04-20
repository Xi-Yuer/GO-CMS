import { memo, useEffect } from 'react';
import { RouterProvider } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from '@/store';
import { ConfigProvider, theme } from 'antd';
import router from '@/router';
import { constants } from '@/constant';
import { changeThemeMode } from '@/store/UIStore';

const APP = () => {
  const { langMode, themeMode } = useAppSelector((state) => state.UIStore);
  const dispatch = useAppDispatch();
  const algorithm = themeMode === 'dark' ? theme.darkAlgorithm : theme.defaultAlgorithm;
  useEffect(() => {
    dispatch(changeThemeMode(themeMode));
  }, [dispatch, themeMode]);
  return (
    <>
      <ConfigProvider
        locale={constants.langMap[langMode]}
        theme={{
          algorithm,
          token: {
            colorPrimary: '#00aeff',
          },
        }}>
        <RouterProvider router={router}></RouterProvider>
      </ConfigProvider>
    </>
  );
};

export default memo(APP);
