import zhCN from 'antd/locale/zh_CN';
import enUS from 'antd/locale/en_US';
import { Locale } from 'antd/es/locale';

export const constants = {
  localStorage: {
    lang: 'lang',
    token: 'token',
  },
  langMap: {
    zhCN: zhCN,
    enUS: enUS,
  } as Record<string, Locale | undefined>,
  routePath: {
    login: '/login',
    main: '/',
  },
  module: {
    ROLE: 'role',
  },
  permissionDicMap: {
    EXPORT_USER: 'POST:/users/export',
    ADD_USER: 'POST:/users',
    DELETE_USER: 'DELETE:/users/:id',
    UPDATE_USER: 'PATCH:/users/:id',
    EXPORT_ROLE: 'POST:/roles/export',
    ADD_ROLE: 'POST:/roles',
    DELETE_ROLE: 'DELETE:/roles/:id',
    UPDATE_ROLE: 'PATCH:/roles/:id',
    BIND_USER: 'POST:/roles/bindUser',
    UNBIND_USER: 'POST:/roles/deBindUser',
    GET_ALL_MENU: 'GET:/pages',
    GET_MENU: 'GET:/pages/user',
    ADD_MENU: 'POST:/pages',
    UPDATE_MENU: 'PATCH:/pages/:id',
    GET_ROLE_MENU: 'GET:/pages/role/:id',
    ADD_DEPARTMENT: 'POST:/department',
    DELETE_DEPARTMENT: 'DELETE:/department/:id',
    UPDATE_DEPARTMENT: 'PATCH:/department/:id',
    GET_PAGE_INTERFACE: 'GET:/interface/page/:id',
    GET_ALL_INTERFACE: 'GET:/interface',
    GET_ALL_USER_IS_OUT_ROLE_ID: 'POST:/users/query/role/:id',
    ADD_PAGE_INTERFACE: 'POST:/interface',
    DELETE_PAGE_INTERFACE: 'DELETE:/interface/:id',
    UPDATE_PAGE_INTERFACE: 'PATCH:/interface/:id',
    CREATE_TEMPLATE: 'POST:/template',
    DOWNLOAD_TEMPLATE: 'POST:/template/download',
  },
};
