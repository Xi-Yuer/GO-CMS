import { theme, ThemeConfig } from 'antd';
import { useAppSelector } from '@/store';

export const useTheme: () => [ThemeConfig] = () => {
  const { themeMode } = useAppSelector((state) => state.UIStore);
  const algorithm = themeMode === 'dark' ? theme.darkAlgorithm : theme.defaultAlgorithm;
  return [
    {
      algorithm,
      token: {
        borderRadius: 4,
        colorPrimary: '#00aeff',
        colorBgContainer: themeMode === 'dark' ? '#001624' : '#fff',
        colorBgElevated: themeMode === 'dark' ? '#0c1d27' : '#fff',
      },
      components: {
        Table: {
          borderColor: themeMode === 'dark' ? '#004259' : '#eee',
        },
      },
    },
  ];
};
