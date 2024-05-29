import { FC, memo } from 'react';

const Footer: FC = () => {
  return (
    <>
      <div className='flex text-[#6d727e] py-2 flex-col w-full justify-center items-center bg-[#ededed] dark:bg-[#2e2e2e] transition-all duration-150 dark:border-0 border-t border-2 border-gray-300'>
        <span>&copy; Design Xi-Yuer</span>
        <span>
          <a href='https://beian.miit.gov.cn/#/Integrated/index' target='_blank'>
            备案号： 蜀ICP备2022015920号-2
          </a>
        </span>
      </div>
    </>
  );
};

export default memo(Footer);
