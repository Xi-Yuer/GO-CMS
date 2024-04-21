import { menuType } from '@/types/menus';
import { RouteObject } from 'react-router-dom';

export const builderMenuRoutes = (menus: menuType[]) => {
  // 递归处理菜单
  const recursionMenus: (menus: menuType[]) => RouteObject[] = (menus: menuType[]) => {
    return menus?.map((item) => {
      return {
        path: item.pagePath,
        element: <div>APP</div>,
        children: item.children ? recursionMenus(item.children) : [],
      };
    });
  };
  return recursionMenus(menus);
};
