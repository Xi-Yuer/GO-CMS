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
      show: false,
    },
  },
  animation: true,
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
};

export const gitCommitsOptions: EChartsOption = {
  tooltip: {
    trigger: 'item',
  },
  calendar: {
    top: 120,
    left: 30,
    right: 30,
    cellSize: ['auto', 13],
    range: '2024',
    splitLine: {
      show: false,
    },
    itemStyle: {
      borderWidth: 0.5,
    },
    dayLabel: {
      nameMap: 'ZH',
    },
    monthLabel: {
      nameMap: 'ZH',
    },
    yearLabel: {
      show: false,
    },
  },
  visualMap: {
    type: 'piecewise',
    pieces: [
      {
        value: 0,
        color: '#e8eaee',
      },
      {
        min: 1,
        max: 3,
        color: '#8ce4a2',
      },
      {
        min: 4,
        max: 6,
        color: '#2dbb60',
      },
      {
        min: 7,
        max: 20,
        color: '#186235',
      },
    ],
    orient: 'horizontal',
    left: 'center',
    top: 65,
    min: 1,
    max: 20,
  },
};
