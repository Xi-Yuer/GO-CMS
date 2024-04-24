import { FC, memo } from 'react';
import { Col, Popover, Row, Timeline } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useDashBoard } from '@/pages/Dashboard/hooks.ts';
import { useAppSelector } from '@/store';
import { Card } from '@/components';
import { BugOutlined, SmileOutlined } from '@ant-design/icons';
import classNames from 'classnames';

const Dashboard: FC = () => {
  const { totalOption, cpuUsageOption, allMenUsageOption, gitCommits, commitCount } = useDashBoard();
  const { themeMode } = useAppSelector((state) => state.UIStore);

  const echartsContent = [
    <ReactECharts option={totalOption} theme={themeMode} style={{ height: '100%' }} />,
    <ReactECharts option={cpuUsageOption} theme={themeMode} style={{ height: '100%' }} />,
    <ReactECharts option={allMenUsageOption} theme={themeMode} style={{ height: '100%' }} />,
    <div className='flex h-[100px] justify-center items-center text-2xl dark:bg-[#110f25]'>
      <span className='text-[#ff6787] text-5xl'>{commitCount}</span>
      <span className='ml-4'>Git提交总次</span>
    </div>,
  ];

  const timeLine: any[] = gitCommits.map((item) => {
    return {
      color: '#00CCFF',
      dot: <SmileOutlined />,
      children: item.children.map((item) => {
        return (
          <div key={item.commitID} className='cursor-pointer'>
            <Popover
              placement='bottom'
              content={
                <div>
                  <p>Auth:{item.author}</p>
                  <p>Date:{item.date}</p>
                </div>
              }>
              <p>{item.message}</p>
            </Popover>
          </div>
        );
      }),
    };
  });
  return (
    <div>
      <Row gutter={[10, 10]} className='h-[130px]'>
        {echartsContent.map((item, index) => {
          return (
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }} key={index}>
              <Card>{item}</Card>
            </Col>
          );
        })}
        <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 24 }} lg={{ span: 12 }}>
          <div className='flex flex-col justify-between h-full gap-2'>
            <div
              className={classNames('flex-1 bg-white dark:bg-[#110f25] rounded-md shadow-md p-4 max-h-[520px] overflow-scroll no-scrollbar', {
                physicDarkDashBoard: themeMode === 'dark',
              })}>
              LEFT
            </div>
            <div
              className={classNames('flex-1 bg-white dark:bg-[#110f25] rounded-md shadow-md p-4 max-h-[520px] overflow-scroll no-scrollbar', {
                physicDarkDashBoard: themeMode === 'dark',
              })}>
              LEFT
            </div>
          </div>
        </Col>
        <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 24 }} lg={{ span: 12 }}>
          <div
            className={classNames('bg-white dark:bg-[#110f25] rounded-md shadow-md p-4 max-h-[520px] overflow-scroll no-scrollbar', {
              physicDarkDashBoard: themeMode === 'dark',
            })}>
            <Timeline
              items={timeLine}
              mode='alternate'
              pending={
                <div className='text-[#00c7fc] gap-2 flex'>
                  <BugOutlined />
                  Coding...
                </div>
              }
              reverse={true}></Timeline>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default memo(Dashboard);
