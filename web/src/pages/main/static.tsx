import { useState } from 'react';
import { useAppSelector } from '@/store';
import { Layout, MenuProps } from 'antd';
import { menuType } from '@/types/menus';
import { Icon } from '@/components';
import { useNavigate } from 'react-router-dom';

export const useMainPage = () => {
  const { Header, Sider, Content } = Layout;
  const navigate = useNavigate();
  const { menus } = useAppSelector((state) => state.UserStore);
  const [defaultOpenKeys] = useState([]);

  const onClick: MenuProps['onClick'] = (e) => {
    navigate(e.key);
  };

  return {
    menus1: menus,
    menus: getItem(menus),
    Header,
    Sider,
    Content,
    defaultOpenKeys,
    onClick,
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
