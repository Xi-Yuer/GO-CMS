import { memo, useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '@/store';
import { ConfigProvider, message } from 'antd';
import { constants } from '@/constant';
import { changeThemeMode } from '@/store/UIStore';
import { useAppRouter } from '@/hooks/useAppRoutes.tsx';
import 'dayjs/locale/zh-cn';
import dayjs from 'dayjs';
import { useTheme } from '@/theme';
import { MESSAGE_EVENT_NAME } from '@/utils';
import { MessageInstance } from 'antd/es/message/interface';

dayjs.locale('zh-cn');

const APP = () => {
  const [theme] = useTheme();
  const [api, contextHolder] = message.useMessage();
  const { langMode, themeMode } = useAppSelector((state) => state.UIStore);
  const dispatch = useAppDispatch();
  const { element } = useAppRouter();
  useEffect(() => {
    dispatch(changeThemeMode(themeMode));
  }, [dispatch, themeMode]);

  useEffect(() => {
    const bindEvent = (e: CustomEvent | any) => {
      const func: keyof MessageInstance = e?.detail?.type || 'info';
      const { content, duration, onClose } = e.detail?.params;
      api[func](content, duration, onClose);
    };

    window.addEventListener(MESSAGE_EVENT_NAME, bindEvent);

    return () => {
      window.removeEventListener(MESSAGE_EVENT_NAME, bindEvent);
    };
  }, [api]);

  return (
    <>
      {contextHolder}
      <ConfigProvider locale={constants.langMap[langMode]} theme={theme}>
        {element}
      </ConfigProvider>
    </>
  );
};

export default memo(APP);
