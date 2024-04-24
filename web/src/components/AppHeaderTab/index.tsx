import { FC, memo } from 'react';
import { Popover, Tag } from 'antd';
import { Icon } from '@/components';
import '@/styles/common/index.css';
import { useAppHeaderTab } from '@/components/AppHeaderTab/hooks.tsx';

const AppHeaderTab: FC = () => {
  const { content, onclick, onclose, TabHeader, pathname } = useAppHeaderTab();
  return (
    <div className='flex py-4 bg-white border-t px-8 dark:bg-[#001624] dark:border-[#a6aeb3] overflow-y-scroll no-scrollbar'>
      {!!TabHeader.length &&
        TabHeader.map((item, index) => {
          return (
            <Popover placement='bottomLeft' key={item.pageID} trigger={'contextMenu'} content={() => content(item, index)}>
              <Tag
                color={pathname === item.pagePath ? '#2db7f5' : ''}
                bordered={false}
                icon={<Icon name={item.pageIcon as any} />}
                key={item.pageID}
                closable={true}
                className='cursor-pointer animate__animated animate__slideInRight select-none'
                onClose={(e) => onclose(e, item, index)}
                onClick={() => onclick(item)}>
                {item.pageName}
              </Tag>
            </Popover>
          );
        })}
    </div>
  );
};

export default memo(AppHeaderTab);
