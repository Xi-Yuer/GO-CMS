import { FC, memo } from 'react';
import { Image, Layout, Menu } from 'antd';
import { ThemeBar, Translate } from '@/components';
import Logo from '@/assets/svg/logo.svg';
import { useMainPage } from '@/pages/main/static.tsx';
import { Outlet } from 'react-router-dom';

const Main: FC = () => {
  const { Sider, Header, Content, menus, onClick, defaultOpenKeys } = useMainPage();

  return (
    <>
      <Layout className='bg-white h-screen overflow-hidden'>
        <Sider width='15%'>
          <div className='h-16 dark:bg-[#141414] dark:text-white flex items-center justify-center text-xl font-bold'>
            <Image src={Logo} width={30} />
            <span>Go-React-Admin</span>
          </div>
          <Menu onClick={onClick} style={{ width: 260 }} defaultOpenKeys={defaultOpenKeys} mode='inline' items={menus} className='h-screen' />
        </Sider>
        <Layout>
          <Header className='flex items-center justify-end dark:bg-[#141414] dark:text-white'>
            <ThemeBar />
            <div className='mt-3'>
              <Translate />
            </div>
          </Header>
          <Content className='p-[30px]'>
            <Outlet />
          </Content>
        </Layout>
      </Layout>
    </>
  );
};

export default memo(Main);
