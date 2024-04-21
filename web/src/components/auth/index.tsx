import { FC, memo, ReactElement, ReactNode, useEffect } from 'react';
import { useAppSelector } from '@/store';
import { useNavigate, useNavigation } from 'react-router-dom';
import { constants } from '@/constant';

const withAuth: FC<{ children: ReactNode }> = ({ children }): ReactElement | null => {
  const { token } = useAppSelector((state) => state.UserStore);
  const navigate = useNavigate();
  const navigation = useNavigation();
  useEffect(() => {
    if (!token && navigation.location?.pathname !== constants.routePath.login) {
      navigate(constants.routePath.login);
    }
  }, []);
  return <>{children}</>;
};

export default memo(withAuth);
