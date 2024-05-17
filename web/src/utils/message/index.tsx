import { JointContent } from 'antd/es/message/interface';

export const MESSAGE_EVENT_NAME = 'show_message';

export enum MESSAGE_TYPES {
  SUCCESS = 'success',
  ERROR = 'error',
  INFO = 'info',
  WARNING = 'warning',
  LOADING = 'loading',
}

const dispatch = (type: MESSAGE_TYPES, content: JointContent, duration?: number | VoidFunction, onClose?: VoidFunction) => {
  window.dispatchEvent(
    new CustomEvent(MESSAGE_EVENT_NAME, {
      detail: {
        params: {
          content,
          duration,
          onClose,
        },
        type: type,
      },
    }),
  );
};

export const message = {
  success(content: JointContent, duration?: number | VoidFunction, onClose?: VoidFunction) {
    dispatch(MESSAGE_TYPES.SUCCESS, content, duration, onClose);
  },
  error(content: JointContent, duration?: number | VoidFunction, onClose?: VoidFunction) {
    dispatch(MESSAGE_TYPES.ERROR, content, duration, onClose);
  },
  info(content: JointContent, duration?: number | VoidFunction, onClose?: VoidFunction) {
    dispatch(MESSAGE_TYPES.INFO, content, duration, onClose);
  },
  warning(content: JointContent, duration?: number | VoidFunction, onClose?: VoidFunction) {
    dispatch(MESSAGE_TYPES.WARNING, content, duration, onClose);
  },
  loading(content: JointContent, duration?: number | VoidFunction, onClose?: VoidFunction) {
    dispatch(MESSAGE_TYPES.LOADING, content, duration, onClose);
  },
};
