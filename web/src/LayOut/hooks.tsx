import { useEffect } from 'react';
import { Layout, MenuProps } from 'antd';
import { menuType } from '@/types/menus';
import { Icon } from '@/components';
import { useAppSelector } from '@/store';
import { useLocation, useNavigate } from 'react-router-dom';
import { getFirstMenu, getMenuByPath, getTheCurrentRoutePathAllMenuPath } from '@/utils';
import { useDispatch } from 'react-redux';
import { addTabHeader, changeDefaultOpenKeys, changeDefaultSelectedKeys } from '@/store/UIStore';

export const useMainPage = () => {
  const { Header, Sider, Content } = Layout;
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const { menus } = useAppSelector((state) => state.UserStore);

  const { pathname } = useLocation();
  useEffect(() => {
    dispatch(changeDefaultSelectedKeys([pathname]));
    dispatch(changeDefaultOpenKeys(getTheCurrentRoutePathAllMenuPath(pathname, menus)));
  }, [pathname, menus]);

  useEffect(() => {
    // 首页重定向
    if (menus.length && pathname === '/') {
      navigate(getFirstMenu(menus).pagePath);
    }
  }, []);

  // 跳转至首页
  const navigateHome = () => {
    if (menus && menus.length) {
      const firstMenu = getFirstMenu(menus);
      dispatch(changeDefaultSelectedKeys([firstMenu.pagePath] || []));
      dispatch(changeDefaultOpenKeys(getTheCurrentRoutePathAllMenuPath(firstMenu.pagePath, menus)));
      dispatch(addTabHeader(getMenuByPath(firstMenu.pagePath, menus) || ({} as menuType)));
      navigate(firstMenu.pagePath);
    }
  };

  const onSelect: MenuProps['onSelect'] = (e) => {
    dispatch(changeDefaultSelectedKeys(e.selectedKeys));
    dispatch(changeDefaultOpenKeys(e.keyPath));
    dispatch(addTabHeader(getMenuByPath(e.key, menus) || ({} as menuType)));
    navigate(e.key);
  };

  const onOpenChange = (e: string[]) => {
    dispatch(changeDefaultOpenKeys(e));
  };

  return {
    menus: getItem(menus),
    Header,
    Sider,
    Content,
    onSelect,
    onOpenChange,
    navigateHome,
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
