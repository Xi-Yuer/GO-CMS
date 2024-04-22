import { memo, useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '@/store';
import { ConfigProvider, theme } from 'antd';
import { constants } from '@/constant';
import { changeThemeMode } from '@/store/UIStore';
import { Router } from '@/components';

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
            borderRadius: 4,
            colorPrimary: '#00aeff',
          },
        }}>
        <Router />
      </ConfigProvider>
    </>
  );
};

export default memo(APP);
