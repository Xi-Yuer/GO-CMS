import { FC, memo, ReactNode } from 'react';

const Card: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className='bg-white dark:bg-[#141414] rounded-md shadow-md h-full p-2'>{children}</div>;
};

export default memo(Card);
