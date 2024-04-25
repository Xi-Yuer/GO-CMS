import { memo, useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '@/store';
import { ConfigProvider, theme } from 'antd';
import { constants } from '@/constant';
import { changeThemeMode } from '@/store/UIStore';
import { Router } from '@/components';
import 'dayjs/locale/zh-cn';
import dayjs from 'dayjs';

const APP = () => {
  dayjs.locale('zh-cn');
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
