import { RouteObject, useLocation, useNavigate, useRoutes } from 'react-router-dom';
import { useAppSelector } from '@/store';
import { builderMenuRoutes, getFirstMenu } from '@/utils';
import { useEffect, useState } from 'react';
import routes from '@/router';
import NotFont from '@/pages/NotFont';
import { constants } from '@/constant';

export const useAppRouter = () => {
  const navigate = useNavigate();
  const { menus, token } = useAppSelector((state) => state.UserStore);
  const { pathname } = useLocation();
  const [routesWithMenus, setRoutesWithMenus] = useState<RouteObject[]>([]);

  useEffect(() => {
    if (!token) {
      navigate(constants.routePath.login, { replace: true });
    }
  }, [token, navigate]);

  useEffect(() => {
    const routesWithDynamicMenus = [...routes]; // 假设 routes 是一个外部定义的数组
    routesWithDynamicMenus[1].children = builderMenuRoutes(menus);
    routesWithDynamicMenus[1].children?.push({
      path: '*',
      element: <NotFont />,
    });

    setRoutesWithMenus(routesWithDynamicMenus);
    if (pathname === constants.routePath.main) {
      if (!menus.length) return;
      const firstMenuPath = getFirstMenu(menus)?.pagePath || constants.routePath.login;
      navigate(firstMenuPath);
    }
  }, [menus, pathname, navigate]);

  return {
    element: useRoutes(routesWithMenus),
  };
};
