import { FC, memo, ReactNode } from 'react';
import classNames from 'classnames';
import { useAppSelector } from '@/store';

const Card: FC<{ children: ReactNode }> = ({ children }) => {
  const { themeMode } = useAppSelector((state) => state.UIStore);
  return (
    <div
      className={classNames('flex-1 bg-white dark:bg-[#110f25] rounded-md shadow-md p-4 max-h-[520px] min-h-[130px] overflow-scroll no-scrollbar', {
        physicDarkDashBoard: themeMode === 'dark',
      })}>
      {children}
    </div>
  );
};

export default memo(Card);
