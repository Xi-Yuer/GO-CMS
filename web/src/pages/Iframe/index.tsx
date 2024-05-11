import { FC, memo } from 'react';

export interface IProps {
  src: string;
}

const Iframe: FC<IProps> = ({ src }) => {
  return (
    <iframe
      src={src}
      title='YouTube video player'
      className='w-full h-full'
      allow='accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture'
      allowFullScreen></iframe>
  );
};

export default memo(Iframe);
