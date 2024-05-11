import Login from '@/pages/Login';
import Main from '@/pages/Main';
import NotFont from '@/pages/NotFont';
import { Suspense } from 'react';
import { RouteObject } from 'react-router-dom';
import { Spin } from 'antd';

export const routes: RouteObject[] = [
  {
    path: '/Login',
    element: <Login />,
  },
  {
    path: '/',
    element: (
      <Suspense fallback={<Spin spinning={true} fullscreen={true}></Spin>}>
        <Main></Main>
      </Suspense>
    ),
    children: [],
  },
  {
    path: '*',
    element: <NotFont />,
  },
];

export default routes;
