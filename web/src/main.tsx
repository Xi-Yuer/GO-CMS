import ReactDOM from 'react-dom/client';
import '@/utils/echarts';
import 'animate.css';
import '@/styles/index.css';
import '@/local/index';
import store from '@/store';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import { persistStore } from 'redux-persist';
import { HashRouter } from 'react-router-dom';
import { App as AntdApp, Spin } from 'antd';
import App from '@/App.tsx';
import ErrorBoundary from 'antd/es/alert/ErrorBoundary';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <ErrorBoundary>
      <PersistGate loading={<Spin spinning={true} fullscreen={true}></Spin>} persistor={persistStore(store)}>
        <HashRouter>
          <AntdApp>
            <App />
          </AntdApp>
        </HashRouter>
      </PersistGate>
    </ErrorBoundary>
  </Provider>,
);
