import { FC, memo } from 'react';
import { Form, Select } from 'antd';
import * as AllIcons from '@ant-design/icons';
import { useTranslation } from 'react-i18next';
import { Icon } from '@/components';

const IconSelect: FC = () => {
  const { t } = useTranslation();
  const options = Object.keys(AllIcons)
    .slice(0, 300)
    .map((key: any) => {
      return {
        value: key,
        label: (
          <>
            <Icon name={key} props={{ className: 'mr-2' }}></Icon>
            {key}
          </>
        ),
      };
    });
  return (
    <Form.Item name='pageIcon' shouldUpdate noStyle>
      <Select placeholder={t('pleaseEnter')} showSearch options={options}></Select>
    </Form.Item>
  );
};

export default memo(IconSelect);
