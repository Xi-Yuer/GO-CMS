import { RouteObject } from 'react-router-dom';
import { constants } from '@/constant';

import Login from '@/pages/Login';
import Main from '@/LayOut';
import NotFond from '@/pages/NotFond';

export const routes: RouteObject[] = [
  {
    path: constants.routePath.login,
    element: <Login />,
  },
  {
    path: constants.routePath.main,
    element: <Main></Main>,
    children: [],
  },
  {
    path: '*',
    element: <NotFond />,
  },
];

export default routes;
