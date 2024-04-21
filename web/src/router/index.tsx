import Login from '@/pages/login';
import Main from '@/pages/main';
import WithAuth from '@/components/auth';
import { createBrowserRouter } from 'react-router-dom';

const routes = createBrowserRouter([
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/',
    element: (
      <WithAuth>
        <Main />
      </WithAuth>
    ),
    children: [],
  },
]);

export default routes;
