import { FC, memo } from 'react';
import { Image, Layout, Menu } from 'antd';
import { AppBreadcrumb, AppHeaderTab, ThemeBar, Translate } from '@/components';
import { useMainPage } from '@/pages/Main/hooks.tsx';
import { useTheme } from '@/hooks/useTheme';
import { Outlet, useNavigate } from 'react-router-dom';
import { useAppSelector } from '@/store';
import Logo from '@/assets/svg/logo.svg';
import classNames from 'classnames';

const Main: FC = () => {
  const { Sider, Header, Content, menus, onSelect, onOpenChange } = useMainPage();
  const { defaultSelectedKeys, defaultOpenKeys } = useAppSelector((state) => state.UIStore);
  const navigate = useNavigate();
  const { themeMode } = useTheme();

  return (
    <>
      <Layout className='h-screen overflow-hidden select-none'>
        <Sider width='250px' theme={themeMode}>
          <div
            className='h-16 bg-white dark:bg-[#001624] animate__animated animate__backInDown dark:text-white flex items-center justify-center  text-xl font-bold cursor-pointer'
            onClick={() => menus[0] && navigate(menus[0].key)}>
            <Image src={Logo} width={30} />
            <span>Go-React-Admin</span>
          </div>
          <Menu
            onSelect={onSelect}
            style={{ width: 260 }}
            mode='inline'
            theme={themeMode}
            items={menus}
            className='h-screen select-none'
            onOpenChange={onOpenChange}
            selectedKeys={defaultSelectedKeys}
            openKeys={defaultOpenKeys}
            defaultSelectedKeys={defaultSelectedKeys}
            defaultOpenKeys={defaultOpenKeys}
          />
        </Sider>
        <Layout>
          <Header className='flex items-center justify-between bg-white dark:bg-[#001624] dark:text-white px-6'>
            <div className='flex-1'>
              <AppBreadcrumb />
            </div>
            <div className='flex items-center'>
              <ThemeBar />
              <div className='mt-3'>
                <Translate />
              </div>
            </div>
          </Header>
          <AppHeaderTab />
          <Content className='px-8 py-4'>
            <div
              className={classNames('w-full h-full p-4', {
                physicLightCard: themeMode === 'light',
                physicDarkCard: themeMode === 'dark',
              })}>
              <Outlet />
            </div>
          </Content>
        </Layout>
      </Layout>
    </>
  );
};

export default memo(Main);
