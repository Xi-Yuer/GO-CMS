import Login from '@/pages/Login';
import Main from '@/pages/Main';
import NotFont from '@/pages/NotFont';
import { Suspense } from 'react';

export const routes = [
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
