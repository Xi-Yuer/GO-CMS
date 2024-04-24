import { useAppDispatch, useAppSelector } from '@/store';
import { changeThemeMode } from '@/store/UIStore';
import { useLocation } from 'react-router-dom';

export const useTheme = () => {
  const themeMode = useAppSelector((state) => state.UIStore.themeMode);
  const dispatch = useAppDispatch();
  const { pathname } = useLocation();
  const toggleTheme = () => {
    dispatch(changeThemeMode(themeMode === 'dark' ? 'light' : 'dark'));
    if (pathname === '/dashboard') {
      window.location.reload();
    }
  };

  return {
    themeMode,
    toggleTheme,
  };
};
