import { FC, memo, useEffect, useRef } from 'react';
import { useFullscreen } from 'ahooks';
import { useMainPage } from '@/LayOut/hooks.tsx';
import { useTheme } from '@/hooks/useTheme';
import { Outlet, useNavigate } from 'react-router-dom';
import { Image, Layout, Menu, Popover } from 'antd';
import { AppBreadcrumb, AppHeaderTab, AppUploads, ThemeBar, Translate } from '@/components';
import { useAppDispatch, useAppSelector } from '@/store';
import { DownOutlined, ExpandOutlined, MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import { changeFold } from '@/store/UIStore';
import { cache } from '@/utils';
import Logo from '@/assets/svg/logo.svg';
import classNames from 'classnames';
import { addListenerUploadFile, removeListenerUploadFile } from '@/utils/event';
import { RcFile } from 'antd/es/upload';
import { AppUploadsRefProps } from '@/components/AppUploads';
import { constants } from '@/constant';
import { useTranslation } from 'react-i18next';

const Main: FC = () => {
  const { t } = useTranslation();
  const appUploadsRef = useRef<AppUploadsRefProps>(null);
  const dispatch = useAppDispatch();
  const fullscreenRef = useRef();
  const { Sider, Header, Content, menus, onSelect, onOpenChange, navigateHome } = useMainPage();
  const { defaultSelectedKeys, defaultOpenKeys, isFold } = useAppSelector((state) => state.UIStore);
  const { themeMode } = useTheme();
  const navigate = useNavigate();
  const { userInfo } = useAppSelector((state) => state.UserStore);
  const [_, { toggleFullscreen }] = useFullscreen(fullscreenRef);

  const changeFoldAction = () => {
    dispatch(changeFold(!isFold));
  };

  const logOutAction = () => {
    cache.clear();
    navigate(constants.routePath.login);
  };

  const uploadFileHandler = (file: RcFile) => {
    appUploadsRef.current?.addUploadFile(file);
  };

  useEffect(() => {
    addListenerUploadFile(uploadFileHandler);
    return () => {
      removeListenerUploadFile(uploadFileHandler);
    };
  }, []);
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
              <div className='flex items-center justify-center text-xl mr-4 cursor-pointer' onClick={toggleFullscreen}>
                <ExpandOutlined />
              </div>
              <ThemeBar />
              <div className='mt-3 mx-2'>
                <Translate />
              </div>
              <div className='mx-2 cursor-pointer'>
                <Popover
                  content={
                    <div className='cursor-pointer hover:text-[#00b0f0]' onClick={logOutAction}>
                      {t('logout')}
                    </div>
                  }
                  trigger='click'>
                  <span>{userInfo?.nickname}</span>
                  <DownOutlined className='mx-2 text-gray-500 dark:text-gray-50' />
                </Popover>
              </div>
            </div>
          </Header>
          <AppHeaderTab />
          <Content className='px-8 py-4'>
            <div
              ref={fullscreenRef as any}
              className={classNames('w-full h-full min-h-[600px] overflow-y-scroll no-scrollbar p-4', {
                physicLightCard: themeMode === 'light',
                physicDarkCard: themeMode === 'dark',
              })}>
              <Outlet />
            </div>
          </Content>
          <AppUploads ref={appUploadsRef} />
        </Layout>
      </Layout>
    </>
  );
};

export default memo(Main);
