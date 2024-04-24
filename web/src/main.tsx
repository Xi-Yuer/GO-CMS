import ReactDOM from 'react-dom/client';
import '@/utils/echarts';
import 'animate.css';
import 'dayjs/locale/zh-cn';
import '@/styles/index.css';
import '@/local/index';
import store from '@/store';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import { persistStore } from 'redux-persist';
import { BrowserRouter } from 'react-router-dom';
import { App as AntdApp } from 'antd';
import App from '@/App.tsx';
import ErrorBoundary from 'antd/es/alert/ErrorBoundary';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <ErrorBoundary>
      <PersistGate loading={null} persistor={persistStore(store)}>
        <BrowserRouter>
          <AntdApp>
            <App />
          </AntdApp>
        </BrowserRouter>
      </PersistGate>
    </ErrorBoundary>
  </Provider>,
);
