import { getGitCommitCountRequest, getGitCommitInfoRequest, getSystemRunTimeInfoRequest, IGitCommit, MenUsageMap } from '@/service';
import { useEffect, useRef, useState } from 'react';
import { EChartsOption } from 'echarts';
import { AllMemUsageOptions, gitCommitsOptions, SystemCPUUsageOptions, SystemMemUsedSituationOptions } from '@/pages/Dashboard/options.ts';
import dayjs from 'dayjs';
import { useTranslation } from 'react-i18next';

export const useDashBoard = () => {
  const { t } = useTranslation();
  const [totalOption, setTotalOption] = useState<EChartsOption>({});
  const [cpuUsageOption, setCpuUsageOption] = useState({});
  const [allMenUsageOption, setAllMenUsageOption] = useState({});
  const [gitCommits, setGitCommits] = useState<IGitCommit[]>([]);
  const [commitCount, setCommitCount] = useState(0);
  const [gitCommitFrequency, setGitCommitFrequency] = useState({});
  const gitCommitsRef = useRef<HTMLDivElement>(null);
  const [loading, setLoading] = useState(false);
  const getSystemInfoAction = () => {
    getSystemRunTimeInfoRequest().then((res) => {
      setTotalOption({
        ...SystemMemUsedSituationOptions,
        series: [
          {
            name: t('memUsageTotal'),
            type: 'line',
            showSymbol: false,
            data: res.data?.map((item: any) => item.memUsage.total),
            color: 'gray',
            lineStyle: {
              width: 1,
            },
          },
          {
            name: t('memUsageInUsed'),
            type: 'line',
            showSymbol: false,
            data: res.data?.map((item: any) => item.memUsage.used),
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
                    color: '#04a8fb80', // 0% 处的颜色
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
      });
      setCpuUsageOption({
        ...SystemCPUUsageOptions,
        series: [
          {
            name: 'CPU',
            type: 'line',
            smooth: true,
            showSymbol: false,
            data: res.data?.map((item: any) => item.cpuUsage),
          },
        ],
      });
      setAllMenUsageOption({
        ...AllMemUsageOptions,
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
            // @ts-ignore
            data: MenUsageMap?.map(({ label, key }) => ({ name: label, value: res.data[0].memUsage[key] || 0 })),
          },
        ],
      });
    });
  };

  const getGitCommitAction = () => {
    setLoading(true);
    getGitCommitInfoRequest()
      .then((res) => {
        setGitCommits(res.data);
        setGitCommitFrequency({
          ...gitCommitsOptions,
          series: {
            type: 'heatmap',
            coordinateSystem: 'calendar',
            data: res.data?.map((item) => [dayjs(item.date).format('YYYY-MM-DD'), item.children.length]),
          },
        });
      })
      .finally(() => {
        setLoading(false);
      });
    getGitCommitCountRequest()
      .then((res) => {
        setCommitCount(res.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };
  let timer: number;

  useEffect(() => {
    getSystemInfoAction();
    getGitCommitAction();
    // timer = setInterval(getSystemInfoAction, 1000);
    return () => {
      timer && clearInterval(timer);
    };
  }, []);

  return {
    loading,
    gitCommitFrequency,
    gitCommitsRef,
    commitCount,
    gitCommits,
    totalOption,
    cpuUsageOption,
    allMenUsageOption,
  };
};
