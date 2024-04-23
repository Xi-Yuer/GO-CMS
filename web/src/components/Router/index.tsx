import { useNavigate, useRoutes } from 'react-router-dom';
import { useAppSelector } from '@/store';
import { builderMenuRoutes } from '@/utils';
import { useEffect } from 'react';
import routes from '@/router';

export const Router = () => {
  const navigate = useNavigate();
  const { menus, token } = useAppSelector((state) => state.UserStore);
  routes[1].children = builderMenuRoutes(menus) as any;
  useEffect(() => {
    if (!token) {
      navigate('/login');
    }
  }, []);
  return useRoutes(routes);
};
