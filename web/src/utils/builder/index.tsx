import { menuType } from '@/types/menus';

import DashBoard from '@/pages/Dashboard/';
import SystemUser from '@/pages/System/User';
import SystemRole from '@/pages/System/Role';
import SystemDepartment from '@/pages/System/Department';
import SystemMenu from '@/pages/System/Menu';
import NotFont from '@/pages/NotFont';
import Iframe from '@/components/Iframe';
import Logs from '@/pages/Monitor/Logs';
import File from '@/pages/File/Upload';
import * as React from 'react';
import { RouteObject } from 'react-router-dom';
import { TreeDataNode } from 'antd';
import { IAllPageInterfaceListResponse, IInterfaceResponse } from '@/service/api/interface';
import TimeTask from '@/pages/Monitor/TimeTask';
import CodeGenerator from '@/pages/SystemUtils/CodeGenerate';

const pagesMap: Record<string, React.ReactNode | null> = {
  '/dashboard': <DashBoard />,
  '/system/user': <SystemUser />,
  '/system/role': <SystemRole />,
  '/system/department': <SystemDepartment />,
  '/system/menu': <SystemMenu />,
  '/monitor/logs': <Logs />,
  '/monitor/timeTask': <TimeTask />,
  '/file/upload': <File />,
  '/utils/codeGenerator': <CodeGenerator />,
};

// 返回所有 LayOut 下的路由
export const builderMenuRoutes = (menus: menuType[]) => {
  const mainChildrenRoutes: RouteObject[] = [];

  function recursionMenu(menus: menuType[]) {
    menus?.forEach((item) => {
      if (item.children?.length) {
        recursionMenu(item.children);
      } else {
        if (item.isOutSite) {
          mainChildrenRoutes.push({
            path: item.pagePath,
            element: <Iframe src={item.outSiteLink}></Iframe>,
          });
        } else {
          mainChildrenRoutes.push({
            path: item.pagePath,
            element: pagesMap[item.pagePath] ? pagesMap[item.pagePath] : <NotFont />,
          });
        }
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
    menus?.forEach((item) => {
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

export const buildMenuToAntdTree: (menus: menuType[]) => TreeDataNode[] = (menus: menuType[]) => {
  function recursionMenu(menus: menuType[]): TreeDataNode[] {
    const treeData: TreeDataNode[] = [];
    menus?.forEach((item) => {
      treeData.push({
        title: item.pageName,
        key: item.pageID,
        children: recursionMenu(item.children),
      });
    });
    return treeData;
  }

  return recursionMenu(menus);
};

export const getAllChildrenMenusID: (menus: menuType[]) => string[] = (menus: menuType[]) => {
  const treeData: string[] = [];

  function recursionMenu(menus: menuType[]) {
    menus?.forEach((item) => {
      if (item.children && item.children.length > 0) {
        recursionMenu(item.children);
      } else {
        treeData.push(item.pageID);
      }
    });
  }

  recursionMenu(menus);

  return treeData;
};

export const buildInterfaceToAntdTree: (interfaces: IAllPageInterfaceListResponse[]) => TreeDataNode[] = (interfaces: IAllPageInterfaceListResponse[]) => {
  const treeData: TreeDataNode[] = [];
  interfaces.forEach((item) => {
    treeData.push({
      title: item.key,
      key: item.key,
      selectable: false,
      children: mapInterface(item.children),
    });
  });

  return treeData;
};

export function mapInterface(interfaces: IInterfaceResponse[]): TreeDataNode[] {
  return interfaces.map((item) => {
    return {
      title: item.interfaceName,
      key: item.id,
      children: [],
    };
  });
}

export const getAllInterfaceKeys = (interfaces: IAllPageInterfaceListResponse[]): string[] => {
  const treeData: string[] = [];
  interfaces.forEach((item) => {
    treeData.push(item.key);
  });
  return treeData;
};

export const getAllInterfaceDic = (inter: IAllPageInterfaceListResponse[]): string[] => {
  const treeData: string[] = [];

  function _recursionInterface(interfaces: IAllPageInterfaceListResponse[]) {
    interfaces.forEach((item) => {
      item.children.forEach((child) => {
        treeData.push(child.interfaceDic);
      });
    });
  }

  _recursionInterface(inter);

  return treeData;
};
