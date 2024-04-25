import { FC, memo } from 'react';
import TranslateDark from '@/components/Icon/translateDark.tsx';
import { useTheme } from '@/hooks/useTheme.ts';
import TranslateLight from '@/components/Icon/translateLight.tsx';
import { Dropdown, MenuProps } from 'antd';
import { useTranslation } from 'react-i18next';
import { changeLang } from '@/store/UIStore';
import { useDispatch } from 'react-redux';

const Translate: FC = () => {
  const { themeMode } = useTheme();
  const { i18n } = useTranslation();
  const dispatch = useDispatch();
  const items: MenuProps['items'] = [
    {
      key: 'zhCN',
      label: '简体中文',
    },
    {
      key: 'enUS',
      label: 'English',
    },
  ];
  const onClick: MenuProps['onClick'] = async ({ key }) => {
    dispatch(changeLang(key as 'zhCN' | 'enUS'));
    await i18n.changeLanguage(key === 'zhCN' ? 'zh' : 'en');
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
