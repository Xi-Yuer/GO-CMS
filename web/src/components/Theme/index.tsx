import { LightSvg } from '@/components/Icon/light.tsx';
import { DarkSvg } from '@/components/Icon/dark.tsx';
import { Switch } from 'antd';
import { useTheme } from '@/hooks/useTheme.ts';
import { memo } from 'react';

const ThemeBar = () => {
  const { toggleTheme, themeMode } = useTheme();
  return (
    <>
      <Switch checkedChildren={<LightSvg />} unCheckedChildren={<DarkSvg />} onChange={toggleTheme} checked={themeMode === 'dark'} />
    </>
  );
};

export default memo(ThemeBar);
