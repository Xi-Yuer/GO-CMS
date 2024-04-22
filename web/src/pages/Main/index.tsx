import { FC, memo } from 'react';
import { Image, Layout, Menu } from 'antd';
import { AppBreadcrumb, ThemeBar, Translate } from '@/components';
import Logo from '@/assets/svg/logo.svg';
import { useMainPage } from '@/pages/Main/static.tsx';
import { Outlet, useNavigate } from 'react-router-dom';
import classNames from 'classnames';
import { useTheme } from '@/hooks/useTheme.ts';

const Main: FC = () => {
  const { Sider, Header, Content, menus, onSelect, defaultSelectedKeys, defaultOpenKeys } = useMainPage();
  const navigate = useNavigate();
  const { themeMode } = useTheme();

  return (
    <>
      <Layout className='h-screen overflow-hidden'>
        <Sider width='250px' theme={themeMode}>
          <div
            className='h-16 bg-white dark:bg-[#001624] animate__animated animate__backInDown dark:text-white flex items-center justify-center  text-xl font-bold cursor-pointer'
            onClick={() => navigate(menus[0].key)}>
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
          <Content className='p-[30px]'>
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
