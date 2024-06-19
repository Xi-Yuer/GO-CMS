import { FC, memo } from 'react';
import { Popover, Tag } from 'antd';
import { Icon } from '@/components';
import '@/styles/common/index.css';
import { useAppHeaderTab } from '@/components/AppHeaderTab/hooks.tsx';
import { CSSTransition, TransitionGroup } from 'react-transition-group';

const AppHeaderTab: FC = () => {
  const { content, onclick, onclose, TabHeader, pathname } = useAppHeaderTab();
  return (
    <div className='flex py-4 min-h-[55px] relative bg-white border-t px-8 dark:bg-[#001624] dark:border-[#a6aeb3] overflow-y-scroll no-scrollbar'>
      <TransitionGroup>
        {!!TabHeader.length &&
          TabHeader?.map((item, index) => {
            return (
              <CSSTransition
                appear={true}
                timeout={1000}
                key={item.pageID}
                classNames={{
                  enter: 'enter',
                  enterActive: 'enter-active',
                  exit: 'exit',
                  exitActive: 'exit-active',
                }}>
                <Popover placement='bottomLeft' key={item.pageID} trigger={'contextMenu'} content={() => content(item, index)}>
                  <Tag
                    color={pathname === item.pagePath ? '#2db7f5' : ''}
                    bordered={false}
                    icon={<Icon name={item.pageIcon as any} />}
                    key={item.pageID}
                    closable={true}
                    className='cursor-pointer select-none'
                    onClose={(e) => onclose(e, item, index)}
                    onClick={() => onclick(item)}>
                    {item.pageName}
                  </Tag>
                </Popover>
              </CSSTransition>
            );
          })}
      </TransitionGroup>
    </div>
  );
};

export default memo(AppHeaderTab);
