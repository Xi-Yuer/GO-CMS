import { Image } from 'antd';
import ZhEn from '@/assets/svg/zh-en.svg';
import { memo } from 'react';

const TranslateLight = () => {
  return <Image src={ZhEn} preview={false} className='cursor-pointer' />;
};

export default memo(TranslateLight);
