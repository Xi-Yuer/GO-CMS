import { RouterProvider } from 'react-router-dom';
import { memo } from 'react';
import router from '@/router';

const Router = () => {
  return <RouterProvider router={router}></RouterProvider>;
};

export default memo(Router);
