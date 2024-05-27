import { lazy, Suspense } from 'react';
import { RouteObject } from 'react-router-dom';
import { Spin } from 'antd';
import { constants } from '@/constant';

const Login = lazy(() => import('@/pages/Login'));
const Main = lazy(() => import('@/LayOut'));
const NotFond = lazy(() => import('@/pages/NotFond'));

export const routes: RouteObject[] = [
  {
    path: constants.routePath.login,
    element: <Login />,
  },
  {
    path: constants.routePath.main,
    element: (
      <Suspense fallback={<Spin spinning={true} fullscreen={true}></Spin>}>
        <Main></Main>
      </Suspense>
    ),
    children: [],
  },
  {
    path: '*',
    element: <NotFond />,
  },
];

export default routes;
