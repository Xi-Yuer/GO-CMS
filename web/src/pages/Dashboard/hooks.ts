import { getSystemRunTimeInfoRequest } from '@/service';
import { useEffect, useState } from 'react';
import { EChartsOption } from 'echarts';
import { useTotalOption } from '@/pages/Dashboard/options.ts';

export const useDashBoard = () => {
  const options = useTotalOption();
  const [totalOption, setTotalOption] = useState<EChartsOption>(options);

  const getSystemInfoAction = () => {
    getSystemRunTimeInfoRequest().then((res) => {
      setTotalOption({
        ...options,
        series: [
          {
            data: res.data.map((item: any) => item.memUsage.total),
          },
          {
            data: res.data.map((item: any) => item.memUsage.used),
          },
        ],
      });
    });
  };

  useEffect(() => {
    getSystemInfoAction();
    setInterval(getSystemInfoAction, 1000);
  }, []);
  return {
    totalOption,
  };
};
