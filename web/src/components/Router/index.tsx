import { useLocation, useNavigate, useRoutes } from 'react-router-dom';
import { useAppSelector } from '@/store';
import { builderMenuRoutes, getFirstMenu } from '@/utils';
import { useEffect } from 'react';
import routes from '@/router';
import NotFont from '@/pages/NotFont';
import { constants } from '@/constant';

export const Router = () => {
  const navigate = useNavigate();
  const { menus, token } = useAppSelector((state) => state.UserStore);
  const { pathname } = useLocation();
  routes[1].children = builderMenuRoutes(menus) as any;
  routes[1].children?.push({
    path: '*',
    element: <NotFont />,
  });
  useEffect(() => {
    if (!token) {
      navigate(constants.routePath.login, { replace: true });
    }
    if (pathname === constants.routePath.main) {
      navigate(getFirstMenu(menus).pagePath || constants.routePath.login);
    }
  }, []);
  return useRoutes(routes);
};
