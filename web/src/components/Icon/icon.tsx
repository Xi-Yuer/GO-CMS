import { FC } from 'react';
import * as AllIcons from '@ant-design/icons';
import { AntdIconProps } from '@ant-design/icons/es/components/AntdIcon';

const Icon: FC<{ name: keyof typeof AllIcons; props?: AntdIconProps }> = ({ name, props }) => {
  const Comp = AllIcons[name] as any;
  return Comp ? <Comp {...props} /> : name;
};

export default Icon;
