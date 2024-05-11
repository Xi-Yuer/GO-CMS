import { PropsWithChildren, ReactNode } from 'react';
import { useAppSelector } from '@/store';

interface Props {
  permission: string | string[];
}

const Auth: (props: PropsWithChildren<Props>) => ReactNode = ({ children, permission }) => {
  const allInterfaceDic = useAppSelector((state) => state.UserStore.userInterfaceDic);
  if (Array.isArray(permission)) {
    return permission.some((item) => allInterfaceDic.includes(item)) ? children : null;
  }
  if (allInterfaceDic.includes(permission)) {
    return children;
  }
  return null;
};

export default Auth;
