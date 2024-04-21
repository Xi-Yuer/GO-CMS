import { useRoutes } from 'react-router-dom';
import { useAppSelector } from '@/store';
import { builderMenuRoutes } from '@/utils/builder';
import routes from '@/router';

export const DynamicComponentLoader = () => {
  const { menus } = useAppSelector((state) => state.UserStore);
  routes[1].children = builderMenuRoutes(menus) as any;
  return useRoutes(routes);
};
