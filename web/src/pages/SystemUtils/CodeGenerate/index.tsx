import { FC, memo, useEffect, useRef, useState } from 'react';
import CodeEdit, { ICodeEditRefProps } from '@/components/CodeEdit';
import { code, createTemplateRequest, server, web } from '@/service/api/template';
import { ConfigProvider, Layout, Menu, MenuProps } from 'antd';
import Sider from 'antd/es/layout/Sider';
import { Content } from 'antd/es/layout/layout';
import { GenerateFolderMenu } from '@/utils';
import { useAppSelector } from '@/store';

const CodeGenerate: FC = () => {
  const CodeEditRef = useRef<ICodeEditRefProps>(null);
  const { themeMode } = useAppSelector((state) => state.UIStore);
  const [currentSelected, setCurrentSelected] = useState<keyof server | keyof web>('controllerFile');
  const [showEditor, setShowEditor] = useState(false);
  const [controllerCode, setControllerCode] = useState<code>();
  const [serviceCode, setServiceCode] = useState<code>();
  const [repositoryCode, setRepositoryCode] = useState<code>();
  const [dtoCode, setDtoCode] = useState<code>();
  const [codeModules, setCodeModules] = useState({
    controllerFile: {
      code: controllerCode,
      setState: setControllerCode,
    },
    serviceFile: {
      code: serviceCode,
      setState: setServiceCode,
    },
    repositoryFile: {
      code: repositoryCode,
      setState: setRepositoryCode,
    },
    dtoFile: {
      code: dtoCode,
      setState: setDtoCode,
    },
  });

  const createTemplateAction = async () => {
    const result = await createTemplateRequest({
      tableName: 'Users',
      fields: [
        {
          name: 'Name',
          type: 'string',
          default: '',
        },
        {
          name: 'Age',
          type: 'int',
          default: '0',
        },
      ],
    });
    setControllerCode(result.data.server.controllerFile);
    setServiceCode(result.data.server.serviceFile);
    setRepositoryCode(result.data.server.repositoryFile);
    setDtoCode(result.data.server.dtoFile);
    setCodeModules({
      controllerFile: {
        code: result.data.server.controllerFile,
        setState: setControllerCode,
      },
      serviceFile: {
        code: result.data.server.serviceFile,
        setState: setServiceCode,
      },
      repositoryFile: {
        code: result.data.server.repositoryFile,
        setState: setRepositoryCode,
      },
      dtoFile: {
        code: result.data.server.dtoFile,
        setState: setDtoCode,
      },
    });
    setShowEditor(true);
  };
  useEffect(() => {
    createTemplateAction().then();
  }, []);
  const [defaultSelectedKeys, setDefaultSelectedKeys] = useState<string[]>(['controllerFile']);
  const [defaultOpenKeys, setDefaultOpenKeys] = useState<string[]>(['Server', 'Controller', 'controllerFile']);

  const onSelect: MenuProps['onSelect'] = (e) => {
    setDefaultSelectedKeys(e.selectedKeys);
    setDefaultOpenKeys(e.keyPath);
    setCurrentSelected(e.selectedKeys[0] as keyof server | keyof web);
  };

  useEffect(() => {
    setCodeModules({
      controllerFile: {
        code: controllerCode,
        setState: setControllerCode,
      },
      serviceFile: {
        code: serviceCode,
        setState: setServiceCode,
      },
      repositoryFile: {
        code: repositoryCode,
        setState: setRepositoryCode,
      },
      dtoFile: {
        code: dtoCode,
        setState: setDtoCode,
      },
    });
  }, [controllerCode, serviceCode, repositoryCode, dtoCode]);

  type MenuItem = Required<MenuProps>['items'][number];
  const menus: MenuItem[] = GenerateFolderMenu('UserTable');

  return (
    <div className='h-full flex flex-col'>
      <div className='h-[200px] mb-4 rounded bg-white dark:bg-[#1c1d2c]'></div>
      {showEditor ? (
        <Layout className='flex-1 overflow-hidden rounded-md'>
          <Sider width={280} collapsible={false}>
            <div className='h-full bg-white dark:bg-[#1c1d2c]'>
              <ConfigProvider
                theme={{
                  components: {
                    Menu: {
                      subMenuItemBg: themeMode === 'dark' ? '#1e1f32' : '#fff',
                      itemHoverBg: themeMode === 'dark' ? '#1e1f32' : '#f6f6f6',
                      itemSelectedBg: themeMode === 'dark' ? '#2f3245' : '#f6f6f6',
                      fontSize: 12,
                      iconMarginInlineEnd: 5,
                    },
                  },
                }}>
                <Menu
                  mode='inline'
                  items={menus}
                  onSelect={onSelect}
                  onOpenChange={(openKeys) => setDefaultOpenKeys(openKeys)}
                  selectedKeys={defaultSelectedKeys}
                  openKeys={defaultOpenKeys}
                  defaultSelectedKeys={defaultSelectedKeys}
                  defaultOpenKeys={defaultOpenKeys}
                  className='h-full rounded dark:bg-[#1e1f32]'></Menu>
              </ConfigProvider>
            </div>
          </Sider>
          <Content className='flex-1'>
            {codeModules[currentSelected]?.code?.code && codeModules[currentSelected]?.code?.lang && (
              <CodeEdit
                defaultCode={codeModules[currentSelected]?.code?.code}
                defaultLang={codeModules[currentSelected]?.code?.lang}
                editCode={codeModules[currentSelected].setState}
                ref={CodeEditRef}
              />
            )}
          </Content>
        </Layout>
      ) : (
        <div className='w-full h-full flex-1 overflow-hidden rounded-md bg-[#ffffff] dark:bg-[#1a1b26]'></div>
      )}
    </div>
  );
};

export default memo(CodeGenerate);
