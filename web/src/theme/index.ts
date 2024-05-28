import { MenuTheme, theme, ThemeConfig } from 'antd';
import { useAppSelector } from '@/store';
import { Monaco } from '@monaco-editor/react';
import { editor } from 'monaco-editor';
import IStandaloneThemeData = editor.IStandaloneThemeData;

export const useTheme: () => [ThemeConfig] = () => {
  const { themeMode } = useAppSelector((state) => state.UIStore);
  const algorithm = themeMode === 'dark' ? theme.darkAlgorithm : theme.defaultAlgorithm;
  return [
    {
      algorithm,
      token: {
        borderRadius: 4,
        colorPrimary: '#00aeff',
        colorBgContainer: themeMode === 'dark' ? '#001624' : '#fff',
        colorBgElevated: themeMode === 'dark' ? '#0c1d27' : '#fff',
      },
      components: {
        Spin: {
          colorBgContainer: themeMode === 'dark' ? '#001624' : '#fff',
          colorBgMask: themeMode === 'dark' ? '#001624' : '#fff',
        },
      },
    },
  ];
};

export function defineMonacoTheme(monaco: Monaco, themeMode: MenuTheme | undefined) {
  monaco.editor.defineTheme('naruto', themeMode === 'dark' ? monacoDarkTheme : monacoLightTheme);
}

export const monacoLightTheme: IStandaloneThemeData = {
  base: 'vs', // 以哪个默认主题为基础："vs" | "vs-dark" | "hc-black" | "hc-light"
  inherit: true,
  rules: [
    // 高亮规则，即给代码里不同token类型的代码设置不同的显示样式
    { token: 'identifier', foreground: '#d06733' },
    { token: 'number', foreground: '#6bbeeb' },
    { token: 'keyword', foreground: '#05a4d5' },
  ],
  colors: {
    'scrollbarSlider.background': '#edcaa6', // 滚动条背景
    'editor.foreground': '#0d0b09', // 基础字体颜色
    'editor.background': '#ffffff', // 背景颜色
    'editorCursor.foreground': '#d4b886', // 焦点颜色
    'editor.lineHighlightBackground': '#00000000', // 焦点所在的一行的背景颜色
    'editorLineNumber.foreground': '#008800', // 行号字体颜色
  },
};

export const monacoDarkTheme: IStandaloneThemeData = {
  base: 'vs-dark', // 使用vs-dark作为基础主题
  inherit: true,
  rules: [
    { token: 'identifier', foreground: '#819ee9' }, // 修改标识符颜色为黄色
    { token: 'number', foreground: '#6bbeeb' }, // 数字颜色保持不变
    { token: 'keyword', foreground: '#ff79c6' }, // 修改关键字颜色为粉色
  ],
  colors: {
    'scrollbarSlider.background': '#819ee9', // 修改滚动条滑块背景颜色
    'editor.foreground': '#819ee9', // 修改字体颜色为灰白色
    'editor.background': '#1e1f33', // 修改背景颜色为深灰色
    'editorCursor.foreground': '#bbbbbb', // 修改光标颜色为浅灰色
    'editor.lineHighlightBackground': '#00000000', // 修改焦点行背景颜色为深灰色
    'editorLineNumber.foreground': '#819ee9', // 修改行号颜色为灰色
  },
};
