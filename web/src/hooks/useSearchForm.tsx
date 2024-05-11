import { Button, Col, Form, Row } from 'antd';
import { ReactNode } from 'react';
import { PlusOutlined, RedoOutlined, SearchOutlined } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/Auth';
import { constants } from '@/constant';

interface SearchFormProps<T extends Record<string, any>> {
  getDataRequestFn: (values: any) => void;
  onNewRecordFn: () => void;
  formItems: { label: string; name: keyof T; component: ReactNode }[];
  operateComponent?: ReactNode;
  formName: string;
  showAddBtn?: boolean;
}

export const useSearchFrom = <T extends Record<string, any>>(props: SearchFormProps<T>) => {
  const { getDataRequestFn, onNewRecordFn, formItems, operateComponent, formName, showAddBtn = true } = props;
  const [searchFormRef] = Form.useForm();
  const { t } = useTranslation();
  const onFinish = () => getDataRequestFn(searchFormRef.getFieldsValue() as T);
  const onReset = () => {
    searchFormRef?.resetFields();
    getDataRequestFn({});
  };

  const SearchFormComponent = (
    <div className='bg-white p-4 pt-10 mb-4 rounded dark:bg-[#001620]'>
      <Form form={searchFormRef} onReset={onReset} onFinish={onFinish} autoComplete='off' name={formName}>
        <Row gutter={[10, 10]}>
          {formItems.map((item) => {
            return (
              <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 24 }} lg={{ span: 6 }} key={item.name as string}>
                <Form.Item name={item.name as any} label={item.label}>
                  {item.component}
                </Form.Item>
              </Col>
            );
          })}
          <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 24 }} lg={{ span: 12 }}>
            <div className='flex flex-nowrap'>
              <Button type='primary' className='mx-2' htmlType='reset' icon={<RedoOutlined />}>
                {t('reset')}
              </Button>
              <Button type='primary' className='mx-2' htmlType='submit' icon={<SearchOutlined />}>
                {t('search')}
              </Button>
              {showAddBtn && (
                <Auth permission={[constants.permissionDicMap.ADD_USER, constants.permissionDicMap.ADD_ROLE]}>
                  <Button type='primary' className='mx-2' htmlType='button' icon={<PlusOutlined />} onClick={() => onNewRecordFn()}>
                    {t('add')}
                  </Button>
                </Auth>
              )}
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
