import { Button, Col, Form, Row } from 'antd';
import { ReactNode } from 'react';
import { PlusOutlined, RedoOutlined, SearchOutlined } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';

interface SearchFormProps<T extends Record<string, any>> {
  getDataRequestFn: (values: any) => void;
  onNewRecordFn: () => void;
  formItems: { label: string; name: keyof T; component: ReactNode }[];
  operateComponent?: ReactNode;
}

export const useSearchFrom = <T extends Record<string, any>>(props: SearchFormProps<T>) => {
  const { getDataRequestFn, onNewRecordFn, formItems, operateComponent } = props;
  const [searchFormRef] = Form.useForm();
  const { t } = useTranslation();
  const onFinish = () => getDataRequestFn(searchFormRef.getFieldsValue() as T);
  const onReset = () => {
    searchFormRef?.resetFields();
    getDataRequestFn({});
  };

  const SearchFormComponent = (
    <div className='bg-white p-4 pt-10 mb-4 rounded dark:bg-[#141414]'>
      <Form form={searchFormRef} onReset={onReset} onFinish={onFinish} autoComplete='off' name='searchFormRef'>
        <Row gutter={[10, 10]}>
          {formItems.map((item) => {
            return (
              <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }} key={item.name as string}>
                <Form.Item name={item.name as any} label={item.label}>
                  {item.component}
                </Form.Item>
              </Col>
            );
          })}
          <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }}>
            <div className='flex flex-nowrap'>
              <Button type='primary' className='mx-2' htmlType='reset' icon={<RedoOutlined />}>
                {t('reset')}
              </Button>
              <Button type='primary' className='mx-2' htmlType='submit' icon={<SearchOutlined />}>
                {t('search')}
              </Button>
              <Button type='primary' className='mx-2' htmlType='button' icon={<PlusOutlined />} onClick={() => onNewRecordFn()}>
                {t('add')}
              </Button>
              {operateComponent}
            </div>
          </Col>
        </Row>
      </Form>
    </div>
  );

  return {
    SearchFormComponent,
  };
};
