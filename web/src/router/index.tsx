import Login from '@/pages/Login';
import Main from '@/LayOut';
import NotFont from '@/pages/NotFont';
import { Suspense } from 'react';
import { RouteObject } from 'react-router-dom';
import { Spin } from 'antd';
import { constants } from '@/constant';

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
    element: <NotFont />,
  },
];

export default routes;
