import { lazy, Suspense } from 'react';
import { RouteObject } from 'react-router-dom';
import { Image, Spin } from 'antd';
import { constants } from '@/constant';
import LoadingGIF from '@/assets/image/loading.gif';

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
      <Suspense fallback={<Spin spinning={true} fullscreen={true} indicator={<Image src={LoadingGIF} style={{ width: '150px', height: '150px' }} />}></Spin>}>
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
