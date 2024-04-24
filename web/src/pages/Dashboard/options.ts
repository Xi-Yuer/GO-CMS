import { EChartsOption } from 'echarts';

export const useTotalOption: () => EChartsOption = () => {
  return {
    title: {
      text: '内存占用',
      left: '50%',
      top: '0',
      textAlign: 'center',
      textStyle: {
        fontSize: '12',
      },
    },
    grid: {
      bottom: '20%',
      top: '10%',
      left: '3%',
      right: '3%',
    },
    tooltip: {
      trigger: 'axis',
      formatter(params: any[]) {
        return `
          ${params[0].marker} 总内存：${(params[0].data / (1024 * 1024 * 1024)).toFixed(2)}GB<br/>
          ${params[0].marker} 已使用内存：${(params[1].data / (1024 * 1024 * 1024)).toFixed(2)}GB<br/>
        `;
      },
    } as any,
    xAxis: {
      type: 'category',
      axisLabel: {
        show: false,
      },
      axisTick: {
        show: false,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: false,
      },
      axisLine: {
        show: true,
      },
      axisTick: {
        show: true,
      },
    },
    animation: true,
    series: [
      {
        name: '总内存',
        type: 'line',
        showSymbol: false,
        data: [],
        color: 'gray',
        lineStyle: {
          width: 1,
        },
      },
      {
        name: '已使用内存',
        type: 'line',
        showSymbol: false,
        data: [],
        lineStyle: {
          color: '#04a8fb80',
        },
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0,
                color: '#04a8fb', // 0% 处的颜色
              },
              {
                offset: 1,
                color: '#04a8fb00', // 100% 处的颜色
              },
            ],
          },
        },
      },
    ],
  };
};
