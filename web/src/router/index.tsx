import Login from '@/pages/login';
import Main from '@/pages/main';

export const routes = [
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/',
    element: <Main />,
    children: [],
  },
];

export default routes;
