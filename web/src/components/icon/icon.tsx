import { FC } from 'react';
import * as AllIcons from '@ant-design/icons';

const Icon: FC<{ name: keyof typeof AllIcons }> = ({ name }) => {
  const Comp = AllIcons[name] as any;
  return <Comp />;
};

export default Icon;
