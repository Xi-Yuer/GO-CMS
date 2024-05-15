import { EventEmitter } from 'events';
import { RcFile } from 'antd/es/upload';

const eventEmitter = new EventEmitter();

// 组件间通讯的事件
export const EVENT_UPLOAD_FILE = 'uploadFile';
export const GET_FILE_LIST = 'getFileList';

// 注册事件监听
export function addListenerUploadFile(handler: (file: RcFile) => void) {
  eventEmitter.on(EVENT_UPLOAD_FILE, handler);
}

export function addListenerGetFileList(handler: () => void) {
  eventEmitter.on(GET_FILE_LIST, handler);
}

export function removeListenerGetFileList(handler: () => void) {
  eventEmitter.removeListener(GET_FILE_LIST, handler);
}

export function removeListenerUploadFile(handler: (file: RcFile) => void) {
  eventEmitter.removeListener(EVENT_UPLOAD_FILE, handler);
}

// 触发事件
export function emitUploadFile(file: RcFile) {
  eventEmitter.emit(EVENT_UPLOAD_FILE, file);
}

export function emitGetFileList() {
  eventEmitter.emit(GET_FILE_LIST);
}
