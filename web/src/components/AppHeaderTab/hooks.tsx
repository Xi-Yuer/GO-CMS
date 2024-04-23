import { useAppDispatch, useAppSelector } from '@/store';
import { useLocation, useNavigate } from 'react-router-dom';
import { menuType } from '@/types/menus';
import React from 'react';
import { addTabHeader, changeDefaultOpenKeys, changeDefaultSelectedKeys, changeTabHeader, delTabHeader } from '@/store/UIStore';
import { getFirstMenu, getTheCurrentRoutePathAllMenuPath } from '@/utils';
import { CloseCircleOutlined, CloseOutlined, CloseSquareOutlined, SyncOutlined } from '@ant-design/icons';

export const useAppHeaderTab = () => {
  const { TabHeader } = useAppSelector((state) => state.UIStore);
  const { menus } = useAppSelector((state) => state.UserStore);
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const dispatch = useAppDispatch();

  const onclick = (item: menuType) => {
    changeDefaultOpenMenu(item.pagePath, menus);
    navigate(item.pagePath);
  };

  const onclose = (e: React.MouseEvent<HTMLElement>, item: menuType, index: number) => {
    e.preventDefault();
    const firstMenu = getFirstMenu(menus);
    // 当前 tab 只有一个，且点击的 tab 是用户路由的第一项，则不处理（因为tab至少需要一项）
    if (TabHeader.length === 1 && pathname === firstMenu.pagePath) return;

    // 当前 tab 只有一个了
    if (TabHeader.length === 1) {
      dispatch(delTabHeader(index));
      dispatch(addTabHeader(firstMenu));
      changeDefaultOpenMenu(firstMenu.pagePath, menus);
      // 默认跳转到菜单的第一项
      navigate(firstMenu.pagePath);
      return;
    }

    // 当前点击的 tab 是最后一个，且点击的 tab 当前处于选中状态
    if (index === TabHeader.length - 1 && pathname === item.pagePath) {
      dispatch(delTabHeader(index));
      changeDefaultOpenMenu(TabHeader[index - 1].pagePath, menus);
      navigate(TabHeader[index - 1].pagePath);
      return;
    }
    // 当前点击的 tab 是第一个，且点击的 tab 当前处于选中状态
    if (index === 0 && pathname === item.pagePath) {
      changeDefaultOpenMenu(TabHeader[index + 1].pagePath, menus);
      navigate(TabHeader[index + 1].pagePath);
    }
    dispatch(delTabHeader(index));
  };

  const onCloseLeftAll = (item: menuType, index: number) => {
    const newTabHeader = [...TabHeader];
    newTabHeader.splice(0, index);
    dispatch(changeTabHeader(newTabHeader));
    navigate(item.pagePath);
  };

  const onCloseRightAll = (item: menuType, index: number) => {
    const newTabHeader = [...TabHeader];
    newTabHeader.splice(index + 1);
    dispatch(changeTabHeader(newTabHeader));
    navigate(item.pagePath);
  };

  const onCloseOtherAll = (item: menuType, index: number) => {
    const newTabHeader = TabHeader[index];
    dispatch(changeTabHeader([newTabHeader]));
    navigate(item.pagePath);
  };

  const content = (item: menuType, index: number) => {
    return (
      <div className='flex flex-col gap-3 text-gray-600 cursor-pointer dark:text-gray-200 text-sm'>
        {index !== 0 && (
          <div className='flex gap-2' onClick={() => onCloseLeftAll(item, index)}>
            <CloseOutlined />
            关闭左侧标签页
          </div>
        )}
        {index !== TabHeader.length - 1 && (
          <div className='flex gap-2' onClick={() => onCloseRightAll(item, index)}>
            <CloseCircleOutlined />
            关闭右侧标签页
          </div>
        )}
        {TabHeader.length > 1 && (
          <div className='flex gap-2' onClick={() => onCloseOtherAll(item, index)}>
            <CloseSquareOutlined />
            关闭其他标签页
          </div>
        )}
        <div className='flex gap-2' onClick={() => window.location.reload()}>
          <SyncOutlined />
          刷新当前标签页
        </div>
      </div>
    );
  };

  const changeDefaultOpenMenu = (path: string, menus: menuType[]) => {
    dispatch(changeDefaultSelectedKeys([path]));
    dispatch(changeDefaultOpenKeys(getTheCurrentRoutePathAllMenuPath(path, menus)));
  };
  return {
    content,
    onclose,
    onclick,
    TabHeader,
    pathname,
  };
};
