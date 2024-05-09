import { FC, memo, useEffect, useState } from 'react';
import { useAppSelector } from '@/store';
import { menuType } from '@/types/menus';
import { getTheCurrentRoutePathAllMenuEntity } from '@/utils';
import { useLocation } from 'react-router-dom';
import { Breadcrumb } from 'antd';

const AppBreadcrumb: FC = () => {
  const { menus } = useAppSelector((state) => state.UserStore);
  const { pathname } = useLocation();
  const [breadCrumb, setBreadCrumb] = useState<menuType[]>([]);

  useEffect(() => {
    setBreadCrumb(getTheCurrentRoutePathAllMenuEntity(pathname, menus));
  }, [menus, pathname]);
  return <Breadcrumb items={breadCrumb.map((item) => ({ title: item.pageName }))} className='font-bold select-none' />;
};

export default memo(AppBreadcrumb);
