import { FC, memo } from 'react';
import { Image } from 'antd';
import NotFontImage from '@/assets/image/404.png';

const NotFont: FC = () => {
  return (
    <div className='flex items-center bg-white h-full dark:bg-[#121121]'>
      <Image src={NotFontImage} preview={false}></Image>
      <div className='text-6xl font-bold flex flex-col text-gray-700 dark:text-[#a6aeb3]'>
        <span>OOPS!</span>
        <span>Page Not Font!</span>
      </div>
    </div>
  );
};

export default memo(NotFont);
