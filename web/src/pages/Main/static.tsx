import { useState } from 'react';
import { useAppSelector } from '@/store';
import { Layout, MenuProps } from 'antd';
import { menuType } from '@/types/menus';
import { Icon } from '@/components';
import { useLocation, useNavigate } from 'react-router-dom';
import { getTheCurrentRoutePathAllMenuPath } from '@/utils/builder';

export const useMainPage = () => {
  const { Header, Sider, Content } = Layout;
  const navigate = useNavigate();
  const { menus } = useAppSelector((state) => state.UserStore);
  const { pathname } = useLocation();
  const [defaultSelectedKeys, setDefaultSelectedKeys] = useState<string[]>([pathname]);
  const [defaultOpenKeys, setDefaultOpenKeys] = useState<string[]>(getTheCurrentRoutePathAllMenuPath(pathname, menus));

  const onSelect: MenuProps['onSelect'] = (e) => {
    setDefaultOpenKeys(e.keyPath);
    setDefaultSelectedKeys(e.selectedKeys);
    navigate(e.key);
  };
  return {
    menus: getItem(menus),
    Header,
    Sider,
    Content,
    defaultOpenKeys,
    defaultSelectedKeys,
    onSelect,
  };
};

type returnType = {
  key: string;
  icon: any;
  label: string;
  children?: returnType[];
};

function getItem(menu: menuType[]): returnType[] {
  return menu?.map((item) => {
    return {
      key: item.pagePath,
      icon: <Icon name={item.pageIcon as any} />,
      label: item.pageName,
      children: item.children && getItem(item.children),
    };
  });
}
