import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import LanguageDetector from 'i18next-browser-languagedetector';
import { constants } from '@/constant';
import { cache } from '@/utils';
import en from './en';
import zh from './zh';

const i18next = i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    debug: false,
    fallbackLng: cache.get(constants.localStorage.lang) === 'enUS' ? 'en' : 'zh' || 'zh',
    interpolation: {
      escapeValue: false,
    },
    resources: {
      en,
      zh,
    },
  });

export default i18next;
