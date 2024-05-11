import { createElement } from 'react';
import { createRoot } from 'react-dom/client';

export const message = {
  error(message: string, duration: number = 4000) {
    const el = createElement(
      'div',
      {
        className:
          'animate__animated animate__fadeInDown text-[#2a2a2a] text-sm absolute top-10 left-1/2  px-4 py-2 justify-center items-center text-center bg-white rounded shadow',
      },
      [
        <div className='flex gap-2 items-center justify-center'>
          <svg className='icon' viewBox='0 0 1024 1024' version='1.1' xmlns='http://www.w3.org/2000/svg' width='15' height='15'>
            <path
              d='M512 0C229.376 0 0 229.376 0 512s229.376 512 512 512 512-229.376 512-512S794.624 0 512 0z m218.624 672.256c15.872 15.872 15.872 41.984 0 57.856-8.192 8.192-18.432 11.776-29.184 11.776s-20.992-4.096-29.184-11.776L512 569.856l-160.256 160.256c-8.192 8.192-18.432 11.776-29.184 11.776s-20.992-4.096-29.184-11.776c-15.872-15.872-15.872-41.984 0-57.856L454.144 512 293.376 351.744c-15.872-15.872-15.872-41.984 0-57.856 15.872-15.872 41.984-15.872 57.856 0L512 454.144l160.256-160.256c15.872-15.872 41.984-15.872 57.856 0 15.872 15.872 15.872 41.984 0 57.856L569.856 512l160.768 160.256z'
              fill='#ff4246'></path>
          </svg>
          {message}
        </div>,
      ],
    );
    createRoot(document.getElementById('#message')!).render(el);
    let timer: number;
    timer = setTimeout(() => {
      document.getElementById('#message')!.removeChild(document.getElementById('#message')!.children[0]);
      timer && clearTimeout(timer);
    }, duration);
  },
};
