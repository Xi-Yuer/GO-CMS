import zhCN from 'antd/locale/zh_CN';
import enUS from 'antd/locale/en_US';
import { Locale } from 'antd/es/locale';

export const constants = {
  localStorage: {
    lang: 'lang',
    token: 'token',
  },
  langMap: {
    zhCN,
    enUS,
  } as Record<string, Locale | undefined>,
  routePath: {
    login: '/login',
    main: '/',
  },
};
