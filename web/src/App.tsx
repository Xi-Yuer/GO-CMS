import { memo, useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '@/store';
import { ConfigProvider } from 'antd';
import { constants } from '@/constant';
import { changeThemeMode } from '@/store/UIStore';
import { Router } from '@/components';
import 'dayjs/locale/zh-cn';
import dayjs from 'dayjs';
import { useTheme } from '@/theme';

dayjs.locale('zh-cn');

const APP = () => {
  const [theme] = useTheme();
  const { langMode, themeMode } = useAppSelector((state) => state.UIStore);
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(changeThemeMode(themeMode));
  }, [dispatch, themeMode]);

  return (
    <>
      <ConfigProvider locale={constants.langMap[langMode]} theme={theme}>
        <Router />
      </ConfigProvider>
    </>
  );
};

export default memo(APP);
