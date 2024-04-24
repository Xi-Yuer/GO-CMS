import { EChartsOption } from 'echarts';

export const SystemMemUsedSituationOptions: EChartsOption = {
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

export const SystemCPUUsageOptions: EChartsOption = {
  animation: true, //控制动画示否开启
  animationDuration: 3000,
  animationEasing: 'bounceOut', //缓动动画
  animationThreshold: 8, //动画元素的阈值
  title: {
    text: 'CPU使用率',
    textStyle: {
      fontSize: 13,
    },
    left: '50%',
    textAlign: 'center',
  },
  tooltip: {
    trigger: 'axis',
    formatter(params: any[]) {
      return `${params[0].marker}CPU使用率  ${params[0].data.toFixed(2)}%`;
    },
  } as any,
  grid: {
    bottom: '20%',
    top: '10%',
    left: '3%',
    right: '3%',
  },
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
    axisTick: {
      show: false,
    },
  },
  series: [
    {
      name: 'CPU',
      type: 'line',
      smooth: true,
      showSymbol: false,
      data: [],
    },
  ],
};

export const AllMemUsageOptions: EChartsOption = {
  tooltip: {
    trigger: 'item',
    formatter(params: any) {
      return `${params.marker} ${params.name} ${params.percent}%`;
    },
  },
  grid: {
    bottom: '20%',
    top: '10%',
    left: '3%',
    right: '3%',
  },
  series: [
    {
      name: '系统运行总情况',
      type: 'pie',
      radius: ['50%', '70%'],
      avoidLabelOverlap: false,
      padAngle: 5,
      itemStyle: {
        borderRadius: 10,
      },
      label: {
        show: true,
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 10,
          fontWeight: 'bold',
        },
      },
      labelLine: {
        show: false,
      },
      data: [],
    },
  ],
};
