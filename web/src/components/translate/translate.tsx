import { memo } from 'react';
import TranslateDark from '@/components/icon/translateDark.tsx';
import { useTheme } from '@/hooks/useTheme.ts';
import TranslateLight from '@/components/icon/translateLight.tsx';
import { Dropdown, MenuProps } from 'antd';
import { useTranslation } from 'react-i18next';

const Translate = () => {
  const { themeMode } = useTheme();
  const { i18n } = useTranslation();
  const items: MenuProps['items'] = [
    {
      key: 'zh',
      label: '简体中文',
    },
    {
      key: 'en',
      label: 'English',
    },
  ];
  const onClick: MenuProps['onClick'] = ({ key }) => {
    i18n.changeLanguage(key).then((r) => r);
  };

  return (
    <div className='mx-1'>
      <Dropdown menu={{ items, onClick }} placement='bottomLeft' trigger={['click']}>
        <div>{themeMode === 'light' ? <TranslateDark /> : <TranslateLight />}</div>
      </Dropdown>
    </div>
  );
};

export default memo(Translate);
