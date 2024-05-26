import { FC, memo, useEffect, useRef, useState } from 'react';
import { code, createTemplateRequest, downloadTemplateRequest, ICreateTemplateParams, reactType, server } from '@/service/api/template';
import { EyeOutlined, FileZipOutlined, MinusCircleOutlined, PlusCircleOutlined } from '@ant-design/icons';
import { Button, ConfigProvider, Drawer, Empty, Form, Input, Layout, Menu, MenuProps, Select } from 'antd';
import { GenerateFolderMenu } from '@/utils';
import { useAppSelector } from '@/store';
import { useForm } from 'antd/es/form/Form';
import { useTranslation } from 'react-i18next';
import CodeEdit, { ICodeEditRefProps } from '@/components/CodeEdit';
import { saveAs } from 'file-saver';
import Sider from 'antd/es/layout/Sider';
import { Content } from 'antd/es/layout/layout';

const CodeGenerate: FC = () => {
  const [form] = useForm();
  const { t } = useTranslation();
  const CodeEditRef = useRef<ICodeEditRefProps>(null);
  const [canDownLoad, setCanDownLoad] = useState(false);
  const { themeMode } = useAppSelector((state) => state.UIStore);
  const [currentSelected, setCurrentSelected] = useState<keyof server | keyof reactType>('controllerFile');
  const [showEditor, setShowEditor] = useState(false);
  const [controllerCode, setControllerCode] = useState<code>();
  const [serviceCode, setServiceCode] = useState<code>();
  const [repositoryCode, setRepositoryCode] = useState<code>();
  const [routeFileCode, setRouteFileCode] = useState<code>();
  const [dtoCode, setDtoCode] = useState<code>();

  const [reactSearchFormCode, setReactSearchFormCode] = useState<code>();
  const [reactTableCode, setReactTableCode] = useState<code>();
  const [reactTableHookCode, setReactTableHookCode] = useState<code>();
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
    routeFile: {
      code: routeFileCode,
      setState: setRouteFileCode,
    },
    dtoFile: {
      code: dtoCode,
      setState: setDtoCode,
    },
    searchForm: {
      code: reactSearchFormCode,
      setState: setReactSearchFormCode,
    },
    table: {
      code: reactTableCode,
      setState: setReactTableCode,
    },
    tableHook: {
      code: reactTableHookCode,
      setState: setReactTableHookCode,
    },
  });

  const createTemplateAction = async (value: ICreateTemplateParams) => {
    const result = await createTemplateRequest(value);
    setControllerCode(result.data.server.controllerFile);
    setServiceCode(result.data.server.serviceFile);
    setRepositoryCode(result.data.server.repositoryFile);
    setRouteFileCode(result.data.server.routeFile);
    setDtoCode(result.data.server.dtoFile);
    setReactTableCode(result.data.web.react.table);
    setReactTableHookCode(result.data.web.react.tableHook);
    setReactSearchFormCode(result.data.web.react.searchForm);
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
      routeFile: {
        code: result.data.server.routeFile,
        setState: setRouteFileCode,
      },
      dtoFile: {
        code: result.data.server.dtoFile,
        setState: setDtoCode,
      },
      searchForm: {
        code: result.data.web.react.searchForm,
        setState: setReactSearchFormCode,
      },
      tableHook: {
        code: result.data.web.react.tableHook,
        setState: setReactTableHookCode,
      },
      table: {
        code: result.data.web.react.table,
        setState: setReactTableCode,
      },
    });
    setShowEditor(true);
  };
  const [defaultSelectedKeys, setDefaultSelectedKeys] = useState<string[]>(['controllerFile']);
  const [defaultOpenKeys, setDefaultOpenKeys] = useState<string[]>(['Server', 'Controller', 'controllerFile']);

  const onSelect: MenuProps['onSelect'] = (e) => {
    setDefaultSelectedKeys(e.selectedKeys);
    setDefaultOpenKeys(e.keyPath);
    setCurrentSelected(e.selectedKeys[0] as keyof server | keyof reactType);
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
      routeFile: {
        code: routeFileCode,
        setState: setRouteFileCode,
      },
      dtoFile: {
        code: dtoCode,
        setState: setDtoCode,
      },
      searchForm: {
        code: reactSearchFormCode,
        setState: setReactSearchFormCode,
      },
      tableHook: {
        code: reactTableHookCode,
        setState: setReactTableHookCode,
      },
      table: {
        code: reactTableCode,
        setState: setReactTableCode,
      },
    });
  }, [controllerCode, serviceCode, repositoryCode, routeFileCode, dtoCode, reactSearchFormCode, reactTableHookCode, reactTableCode]);

  type MenuItem = Required<MenuProps>['items'][number];
  const menus: MenuItem[] = GenerateFolderMenu('UserTable');

  const [open, setOpen] = useState(false);
  const onFinish = (values: any) => {
    createTemplateAction(values).then(() => {
      setOpen(true);
      setCanDownLoad(true);
    });
  };
  const downLoadAction = () => {
    downloadTemplateRequest({
      tableName: 'UserTable',
      controller: controllerCode?.code,
      service: serviceCode?.code,
      repository: repositoryCode?.code,
      route: routeFileCode?.code,
      dto: dtoCode?.code,
      searchForm: reactSearchFormCode?.code,
      tableHook: reactTableHookCode?.code,
      table: reactTableCode?.code,
    }).then((res) => {
      saveAs(res, 'UserTableTemplate.zip');
    });
  };

  return (
    <div className='h-full flex flex-col relative bg-[#ffffff] dark:bg-[#041527] rounded-md'>
      <div className='mb-4 rounded p-4 overflow-auto relative no-scrollbar'>
        <Form name='form' form={form} onFinish={onFinish} style={{ maxWidth: 600 }} autoComplete='off'>
          <Form.Item label='包名（package）' name='package' rules={[{ required: true }]} initialValue='cms'>
            <Input></Input>
          </Form.Item>
          <Form.Item label='表名称' name='tableName' rules={[{ required: true }]} initialValue='Users'>
            <Input></Input>
          </Form.Item>{' '}
          <Form.List
            name='fields'
            initialValue={[
              { name: 'Name', type: 'string', default: '' },
              { name: 'Age', type: 'int', default: '0' },
            ]}>
            {(fields, { add, remove }) => (
              <>
                {fields.map(({ key, name, ...restField }) => (
                  <div key={key} className='w-[1000px]'>
                    <div className='flex items-center justify-between gap-4'>
                      <Form.Item {...restField} name={[name, 'name']} className='flex-1' label='字段名称' rules={[{ required: true, type: 'string' }]}>
                        <Input placeholder={t('pleaseEnter')} />
                      </Form.Item>
                      <Form.Item {...restField} name={[name, 'type']} label='字段类型' rules={[{ required: true }]}>
                        <Select
                          style={{ width: '200px' }}
                          placeholder={t('pleaseSelect')}
                          options={[
                            { label: 'int', value: 'int' },
                            { label: 'string', value: 'string' },
                            { label: 'float', value: 'float' },
                            { label: 'boolean', value: 'boolean' },
                            { label: 'date', value: 'date' },
                            { label: 'datetime', value: 'datetime' },
                            { label: 'time', value: 'time' },
                          ]}
                        />
                      </Form.Item>
                      <Form.Item {...restField} name={[name, 'default']} label='默认值'>
                        <Input placeholder={t('pleaseEnter')} />
                      </Form.Item>
                      <Form.Item>
                        <PlusCircleOutlined onClick={() => add()} />
                      </Form.Item>
                      <Form.Item>
                        <MinusCircleOutlined onClick={() => remove(name)} />
                      </Form.Item>
                    </div>
                  </div>
                ))}
              </>
            )}
          </Form.List>
          <Button type='primary' htmlType='submit' icon={<EyeOutlined />} className='absolute top-0 left-[650px] mt-4 mr-4'>
            预览
          </Button>
        </Form>
      </div>
      <Drawer
        title=''
        width='100%'
        height='100%'
        placement='left'
        closable={true}
        onClose={() => setOpen(false)}
        open={open}
        getContainer={false}
        extra={
          <Button type='primary' icon={<FileZipOutlined />} onClick={downLoadAction} disabled={!canDownLoad}>
            下载（zip）
          </Button>
        }>
        {showEditor ? (
          <Layout className='flex-1 overflow-hidden rounded-md h-full'>
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
          <div className='w-full h-full justify-center items-center flex-1 overflow-hidden rounded-md bg-[#ffffff] dark:bg-[#1a1b26]'>
            <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} className='mt-40' />
          </div>
        )}
      </Drawer>
    </div>
  );
};

export default memo(CodeGenerate);
