import { getGitCommitCountRequest, getGitCommitInfoRequest, getSystemRunTimeInfoRequest, IGitCommit, MenUsageMap } from '@/service';
import { useEffect, useState } from 'react';
import { EChartsOption } from 'echarts';
import { AllMemUsageOptions, SystemCPUUsageOptions, SystemMemUsedSituationOptions } from '@/pages/Dashboard/options.ts';

export const useDashBoard = () => {
  const [totalOption, setTotalOption] = useState<EChartsOption>(SystemMemUsedSituationOptions);
  const [cpuUsageOption, setCpuUsageOption] = useState(SystemCPUUsageOptions);
  const [allMenUsageOption, setAllMenUsageOption] = useState(AllMemUsageOptions);
  const [gitCommits, setGitCommits] = useState<IGitCommit[]>([]);
  const [commitCount, setCommitCount] = useState(0);

  const getSystemInfoAction = () => {
    getSystemRunTimeInfoRequest().then((res) => {
      setTotalOption({
        ...totalOption,
        series: [
          {
            data: res.data.map((item: any) => item.memUsage.total),
          },
          {
            data: res.data.map((item: any) => item.memUsage.used),
          },
        ],
      });
      setCpuUsageOption({
        ...cpuUsageOption,
        series: [
          {
            data: res.data.map((item: any) => item.cpuUsage),
          },
        ],
      });
      setAllMenUsageOption({
        ...allMenUsageOption,
        series: {
          // @ts-ignore
          data: MenUsageMap.map(({ label, key }) => ({ name: label, value: res.data[0].memUsage[key] || 0 })),
        },
      });
    });
  };

  const getGitCommitAction = () => {
    getGitCommitInfoRequest().then((res) => {
      setGitCommits(res.data);
    });
    getGitCommitCountRequest().then((res) => {
      setCommitCount(res.data);
    });
  };
  let timer: number;

  useEffect(() => {
    getSystemInfoAction();
    getGitCommitAction();
    timer = setInterval(getSystemInfoAction, 1000);
    return () => {
      timer && clearInterval(timer);
    };
  }, []);

  return {
    commitCount,
    gitCommits,
    totalOption,
    cpuUsageOption,
    allMenUsageOption,
  };
};
