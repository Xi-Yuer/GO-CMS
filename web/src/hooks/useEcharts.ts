import { useAppSelector } from '@/store';
import * as echarts from 'echarts';
import { EChartsOption } from 'echarts';
import { RefObject, useEffect } from 'react';
import { EChartsInstance } from 'echarts-for-react';

export const useEcharts = (el: RefObject<HTMLDivElement>) => {
  const { themeMode } = useAppSelector((state) => state.UIStore);
  let Echarts: EChartsInstance;
  // 初始化echarts实例
  const setOption = (option: EChartsOption) => {
    Echarts?.setOption(option);
  };

  useEffect(() => {
    Echarts = echarts.init(el.current, themeMode, { renderer: 'svg' });
    window.addEventListener('resize', () => Echarts.resize());
    return () => {
      window.removeEventListener('resize', () => Echarts.resize());
    };
  }, []);
  return {
    setOption,
  };
};
