import Login from '@/pages/Login';
import Main from '@/pages/Main';
import NotFont from '@/pages/NotFont';
import { Suspense } from 'react';
import { RouteObject } from 'react-router-dom';

export const routes: RouteObject[] = [
  {
    path: '/Login',
    element: <Login />,
  },
  {
    path: '/',
    element: (
      <Suspense fallback={<div>Loading...</div>}>
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
