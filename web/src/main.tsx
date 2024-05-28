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
import { App as AntdApp, Image, Spin } from 'antd';
import App from '@/App.tsx';
import ErrorBoundary from 'antd/es/alert/ErrorBoundary';
import LoadingGIF from '@/assets/image/loading.gif';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <ErrorBoundary>
      <PersistGate
        loading={<Spin spinning={true} fullscreen={true} indicator={<Image src={LoadingGIF} />} style={{ width: '100px', height: '100px' }}></Spin>}
        persistor={persistStore(store)}>
        <HashRouter>
          <AntdApp>
            <App />
          </AntdApp>
        </HashRouter>
      </PersistGate>
    </ErrorBoundary>
  </Provider>,
);
