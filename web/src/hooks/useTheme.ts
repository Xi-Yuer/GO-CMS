import { useAppDispatch, useAppSelector } from '@/store';
import { changeThemeMode } from '@/store/UIStore';

export const useTheme = () => {
  const themeMode = useAppSelector((state) => state.UIStore.themeMode);
  const dispatch = useAppDispatch();
  const toggleTheme = () => {
    dispatch(changeThemeMode(themeMode === 'dark' ? 'light' : 'dark'));
  };

  return {
    themeMode,
    toggleTheme,
  };
};
