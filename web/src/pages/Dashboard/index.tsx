import { FC, memo } from 'react';
import { Col, Row } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useDashBoard } from '@/pages/Dashboard/hooks.ts';
import { useAppSelector } from '@/store';
import { Card } from '@/components';

const Dashboard: FC = () => {
  const { totalOption } = useDashBoard();
  const { themeMode } = useAppSelector((state) => state.UIStore);

  const echartsContent = [
    <ReactECharts option={totalOption} theme={themeMode} />,
    <ReactECharts option={totalOption} theme={themeMode} />,
    <ReactECharts option={totalOption} theme={themeMode} />,
    <ReactECharts option={totalOption} theme={themeMode} />,
  ];
  return (
    <div>
      <Row gutter={[6, 6]} className='h-[130px]'>
        {echartsContent.map((_, index) => {
          return (
            <Col xs={{ span: 24 }} sm={{ span: 24 }} md={{ span: 12 }} lg={{ span: 6 }} key={index}>
              <Card>
                <ReactECharts option={totalOption} style={{ height: '100%' }} theme={themeMode} />
              </Card>
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default memo(Dashboard);
