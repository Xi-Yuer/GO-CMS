import { menuType } from '@/types/menus';

import DashBoard from '@/pages/Dashboard/';
import SystemUser from '@/pages/System/User';
import SystemRole from '@/pages/System/Role';
import SystemDepartment from '@/pages/System/Department';
import SystemMenu from '@/pages/System/Menu';
import Test from '@/pages/Test';
import * as React from 'react';
import { RouteObject } from 'react-router-dom';
import NotFont from '@/pages/NotFont';

const pagesMap: Record<string, React.ReactNode | null> = {
  '/dashboard': <DashBoard />,
  '/system/user': <SystemUser />,
  '/system/role': <SystemRole />,
  '/system/department': <SystemDepartment />,
  '/system/menu': <SystemMenu />,
  '/test/test2': <Test />,
};

// 返回所有 Main 下的路由
export const builderMenuRoutes = (menus: menuType[]) => {
  const mainChildrenRoutes: RouteObject[] = [];

  function recursionMenu(menus: menuType[]) {
    menus.forEach((item) => {
      if (item.children?.length) {
        recursionMenu(item.children);
      } else {
        mainChildrenRoutes.push({
          path: item.pagePath,
          element: pagesMap[item.pagePath] ? pagesMap[item.pagePath] : <NotFont />,
        });
      }
    });
  }

  recursionMenu(menus);
  return mainChildrenRoutes;
};

// 根据当前路径，返回当前路径的父级菜单
export const getTheCurrentRoutePathAllMenuPath: (path: string, menus: menuType[]) => string[] = (path: string, menus: menuType[]) => {
  const paths: string[] = [];

  function traverse(node: menuType): any {
    paths.push(node.pagePath);
    if (node.pagePath === path) {
      return paths;
    }
    if (node.children) {
      for (const child of node.children) {
        const result = traverse(child);
        if (result) {
          return result;
        }
      }
    }
    paths.pop();
    return null;
  }

  for (const page of menus) {
    const result = traverse(page);
    if (result) {
      return result;
    }
  }
  return [];
};

export const getTheCurrentRoutePathAllMenuEntity: (path: string, menus: menuType[]) => menuType[] = (path: string, menus: menuType[]) => {
  const entity: menuType[] = [];

  function traverse(node: menuType): any {
    entity.push(node);
    if (node.pagePath === path) {
      return entity;
    }
    if (node.children) {
      for (const child of node.children) {
        const result = traverse(child);
        if (result) {
          return result;
        }
      }
    }
    entity.pop();
    return null;
  }

  for (const page of menus) {
    const result = traverse(page);
    if (result) {
      return result;
    }
  }
  return [];
};

export const getMenuByPath: (path: string, menus: menuType[]) => menuType | undefined = (path: string, menus: menuType[]) => {
  for (const item of menus) {
    if (item.pagePath === path) {
      return item;
    }
    if (item.children?.length) {
      const menu = getMenuByPath(path, item.children);
      if (menu) {
        return menu;
      }
    }
  }
  return undefined;
};

// 获取第一个菜单
export const getFirstMenu: (menus: menuType[]) => menuType = (menus: menuType[]) => {
  if (menus && menus.length) {
    if (!menus[0].children?.length) {
      return menus[0];
    }
    if (menus[0].children?.length) {
      return getFirstMenu(menus[0].children);
    }
  }
  return {} as menuType;
};

// 获取用户所有菜单的第一项子菜单
export const getFirstMenuChildren: (menus: menuType[]) => menuType[] = (menus: menuType[]) => {
  const firstMenuChildren: menuType[] = [];

  function recursionMenu(menus: menuType[]) {
    menus.forEach((item) => {
      if (item.children?.length) {
        recursionMenu(item.children);
      } else {
        // 排除仪表盘页面
        if (item.pagePath !== '/dashboard') {
          firstMenuChildren.push(item);
        }
      }
    });
  }

  recursionMenu(menus);
  return firstMenuChildren;
};
