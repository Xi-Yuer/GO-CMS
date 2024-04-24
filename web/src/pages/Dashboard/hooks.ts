import { getSystemRunTimeInfoRequest } from '@/service';
import { useEffect, useState } from 'react';
import { EChartsOption } from 'echarts';
import { SystemMemUsedSituationOptions } from '@/pages/Dashboard/options.ts';

export const useDashBoard = () => {
  const [totalOption, setTotalOption] = useState<EChartsOption>(SystemMemUsedSituationOptions);
  const getSystemInfoAction = () => {
    getSystemRunTimeInfoRequest().then((res) => {
      setTotalOption({
        ...SystemMemUsedSituationOptions,
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
    // setInterval(getSystemInfoAction, 1000);
  }, []);

  return {
    totalOption,
  };
};
