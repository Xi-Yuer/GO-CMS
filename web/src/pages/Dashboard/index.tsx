import { FC, memo } from 'react';
import ReactECharts from 'echarts-for-react';
import { useDashBoard } from '@/pages/Dashboard/hooks.ts';
import { useAppSelector } from '@/store';
import { useTranslation } from 'react-i18next';
import { getFirstMenuChildren } from '@/utils';
import { Icon } from '@/components';
import { Col, Popover, Row, Spin, Timeline } from 'antd';
import { BugOutlined, SmileOutlined } from '@ant-design/icons';
import classNames from 'classnames';
import dayjs from 'dayjs';

const Dashboard: FC = () => {
  const { t } = useTranslation();
  const { menus } = useAppSelector((state) => state.UserStore);
  const { themeMode } = useAppSelector((state) => state.UIStore);
  const { loading, totalOption, cpuUsageOption, allMenUsageOption, gitCommits, commitCount, gitCommitFrequency, navigateToPage } = useDashBoard();

  const commonStyle = 'bg-white dark:bg-[#110f25] rounded-md shadow-md p-4 no-scrollbar w-full';
  const backgroundArray = ['#ff9066', '#56bafc', '#ab77e3', '#ff7bb6', '#46d6cd'];
  const echartsContent = [
    <ReactECharts option={totalOption} theme={themeMode} style={{ height: '120px', width: '100%' }} />,
    <ReactECharts option={cpuUsageOption} theme={themeMode} style={{ height: '120px', width: '100%' }} />,
    <ReactECharts option={allMenUsageOption} theme={themeMode} style={{ height: '120px', width: '100%' }} />,
    <div className='flex h-[100px] justify-center items-center text-2xl dark:bg-[#110f25]'>
      <span className='text-[#ff6787] text-5xl'>{commitCount}</span>
      <span className='ml-4'>{t('commitCount')}</span>
    </div>,
  ];

  const TimeLine: any[] = gitCommits.map((item) => {
    return {
      color: '#00CCFF',
      dot: <SmileOutlined />,
      label: <p className='text-[#5bb4ef]'>{item.date}</p>,
      children: item.children.map((item) => {
        return (
          <div key={item.commitID} className='cursor-pointer'>
            <Popover
              placement='bottom'
              content={
                <div>
                  <p>Auth:{item.author}</p>
                  <p>Date:{dayjs(item.date).format('YYYY-MM-DD HH:mm:ss')}</p>
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
    <Spin spinning={loading}>
      <Row gutter={[10, 10]}>
        {echartsContent.map((item, index) => {
          return (
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 12 }} xl={6} key={index}>
              <div
                className={classNames(commonStyle, 'h-full animate__animated animate__fadeInRight', {
                  physicDarkDashBoard: themeMode === 'dark',
                })}>
                {item}
              </div>
            </Col>
          );
        })}
        <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 24 }} lg={{ span: 12 }}>
          <div className='flex flex-col justify-between gap-2 animate__animated animate__fadeInRight'>
            <div
              className={classNames(commonStyle, {
                physicDarkDashBoard: themeMode === 'dark',
              })}>
              <span className='text-md font-bold mb-2'>{t('quickStart')}</span>
              <Row className='flex items-center min-h-[150px] justify-between gap-4'>
                {getFirstMenuChildren(menus)
                  ?.slice(0, 4)
                  ?.map((item, index) => {
                    return (
                      <Col
                        xs={{ span: 24 }}
                        sm={{ span: 4 }}
                        md={{ span: 4 }}
                        lg={{ span: 24 }}
                        xl={{ span: 4 }}
                        key={item.pageID}
                        onClick={() => navigateToPage(item)}
                        className={classNames('flex flex-col justify-evenly items-center h-20 rounded cursor-pointer shadow')}
                        style={{ background: backgroundArray[index] + '30' }}>
                        <Icon
                          name={item.pageIcon as any}
                          props={{ className: `text-xl text-[${backgroundArray[index]}]`, style: { color: backgroundArray[index] } }}></Icon>
                        {item.pageName}
                      </Col>
                    );
                  })}
              </Row>
            </div>
            <div
              className={classNames(commonStyle, {
                physicDarkDashBoard: themeMode === 'dark',
              })}>
              <span className='text-md font-bold'>
                {commitCount} {t('commitTitle')}
              </span>
              <ReactECharts option={gitCommitFrequency} theme={themeMode} style={{ width: '100%', height: '250px' }} />
            </div>
          </div>
        </Col>
        <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 24 }} lg={{ span: 12 }}>
          <div
            className={classNames(commonStyle, 'overflow-scroll no-scrollbar h-[515px] animate__animated animate__fadeInRight', {
              physicDarkDashBoard: themeMode === 'dark',
            })}>
            <span className='text-md font-bold'>{t('commitBlameTitle')}</span>
            <Timeline
              items={TimeLine}
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
    </Spin>
  );
};

export default memo(Dashboard);
