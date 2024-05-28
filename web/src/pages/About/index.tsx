import { FC, memo } from 'react';
import { Descriptions, DescriptionsProps } from 'antd';

const AboutPage: FC = () => {
  const fontEndItems: DescriptionsProps['items'] = [
    {
      key: '1',
      label: 'UI框架',
      children: 'React: ^18.2.0',
    },
    {
      key: '2',
      label: '组件库',
      children: 'Antd: ^5.16.2,',
    },
    {
      key: '3',
      label: '可视化',
      children: (
        <div className='flex flex-col'>
          <span>echarts: ^5.5.0</span>
          <span>echarts-for-react: ^3.0.2</span>
        </div>
      ),
    },
    {
      key: '4',
      label: '国际化',
      children: (
        <div className='flex flex-col'>
          <span>i18next: ^23.11.2 </span> <span>react-i18next: ^14.1.0</span>
        </div>
      ),
    },
    {
      key: '5',
      label: '状态管理',
      children: (
        <div className='flex flex-col'>
          <span>react-redux: ^9.1.1</span> <span>redux-persist: ^6.0.0</span> <span>@reduxjs/toolkit: ^2.2.3</span>
        </div>
      ),
    },
    {
      key: '6',
      label: 'CSS解决方案',
      children: 'TailwindCSS: ^3.4.3',
    },
  ];
  const backEndItems: DescriptionsProps['items'] = [
    {
      key: '1',
      label: '开发语言',
      children: 'Golang: 1.21.1',
    },
    {
      key: '2',
      label: 'Web框架',
      children: 'github.com/gin-gonic/gin v1.9.1',
    },
    {
      key: '3',
      label: '数据库',
      children: 'MySQL: ^8.3.0',
    },
    {
      key: '4',
      label: 'JWT',
      children: 'github.com/golang-jwt/jwt/v5: v5.2.1',
    },
    {
      key: '5',
      label: '主键算法',
      children: 'github.com/bwmarrin/snowflake: v0.3.0',
    },
    {
      key: '6',
      label: 'DevOps',
      children: 'Docker',
    },
  ];
  return (
    <div className='w-full h-full bg-white p-4 dark:bg-[#001620] flex flex-col gap-10'>
      <Descriptions bordered title='前端' items={fontEndItems} />
      <Descriptions bordered title='后端' items={backEndItems} />
    </div>
  );
};

export default memo(AboutPage);
