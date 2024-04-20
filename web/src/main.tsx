import ReactDOM from 'react-dom/client';
import 'dayjs/locale/zh-cn';
import '@/styles/index.css';
import '@/local/index';
import store from '@/store';
import App from './App.tsx';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import { persistStore } from 'redux-persist';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <PersistGate loading={null} persistor={persistStore(store)}>
      <App />
    </PersistGate>
  </Provider>,
);
