import { FC, memo } from 'react';

const Footer: FC = () => {
  return (
    <>
      <div className='flex text-[#6d727e] py-2 flex-col w-full justify-center items-center bg-[#ededed] dark:bg-[#2e2e2e] transition-all duration-150 dark:border-0 border-t border-2 border-gray-300'>
        <span>&copy; Design Xi-Yuer</span>
        <span>备案号：20233321323</span>
      </div>
    </>
  );
};

export default memo(Footer);
