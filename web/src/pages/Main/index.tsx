import { FC, memo } from 'react';
import { Image, Layout, Menu } from 'antd';
import { AppBreadcrumb, AppHeaderTab, ThemeBar, Translate } from '@/components';
import { useMainPage } from '@/pages/Main/hooks.tsx';
import { useTheme } from '@/hooks/useTheme';
import { Outlet } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from '@/store';
import Logo from '@/assets/svg/logo.svg';
import classNames from 'classnames';
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import { changeFold } from '@/store/UIStore';

const Main: FC = () => {
  const { Sider, Header, Content, menus, onSelect, onOpenChange, navigateHome } = useMainPage();
  const { defaultSelectedKeys, defaultOpenKeys, isFold } = useAppSelector((state) => state.UIStore);
  const { themeMode } = useTheme();
  const dispatch = useAppDispatch();

  const changeFoldAction = () => {
    dispatch(changeFold(!isFold));
  };
  return (
    <>
      <Layout className='h-screen overflow-hidden select-none'>
        <Sider width='250px' theme={themeMode} className='hidden md:block' trigger={null} collapsible collapsed={isFold}>
          <div
            className='h-16 bg-white truncate overflow-hidden dark:bg-[#001624] animate__animated animate__backInDown dark:text-white flex items-center justify-center text-xl font-bold cursor-pointer'
            onClick={navigateHome}>
            <Image src={Logo} width={30} preview={false} />
            {!isFold && <span>Go-React-Admin</span>}
          </div>
          <Menu
            onSelect={onSelect}
            style={{ width: isFold ? 80 : 260 }}
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
            {isFold ? (
              <MenuUnfoldOutlined className='text-xl mr-2' onClick={changeFoldAction} />
            ) : (
              <MenuFoldOutlined className='text-xl mr-2' onClick={changeFoldAction} />
            )}
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
              className={classNames('w-full h-full min-h-[600px] overflow-y-scroll no-scrollbar p-4', {
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
