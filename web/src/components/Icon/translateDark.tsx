import { Image } from 'antd';
import EnZh from '@/assets/svg/en-zh.svg';
import { memo } from 'react';

const TranslateDark = () => {
  return <Image src={EnZh} preview={false} className='cursor-pointer' />;
};

export default memo(TranslateDark);
