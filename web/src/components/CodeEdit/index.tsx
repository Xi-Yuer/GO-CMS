import { forwardRef, memo, useEffect, useImperativeHandle, useState } from 'react';
import { Editor, loader, Monaco } from '@monaco-editor/react';
import { defineMonacoTheme } from '@/theme';
import { useTheme } from '@/hooks/useTheme.ts';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker';
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker';
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';
import * as monaco from 'monaco-editor';

loader.config({ monaco });
self.MonacoEnvironment = {
  getWorker(_, label) {
    if (label === 'json') {
      return new jsonWorker();
    }
    if (label === 'css' || label === 'scss' || label === 'less') {
      return new cssWorker();
    }
    if (label === 'html' || label === 'handlebars' || label === 'razor') {
      return new htmlWorker();
    }
    if (label === 'typescript' || label === 'javascript') {
      return new tsWorker();
    }
    return new editorWorker();
  },
};

export interface ICodeEditeInterface {
  lang: string;
  code: string;
}

export interface ICodeEditRefProps {
  getCode: () => string | undefined;
}

let monacoInstance: Monaco | undefined;
await loader.init().then((monaco) => {
  monacoInstance = monaco;
});

const CodeEdit = forwardRef<ICodeEditRefProps, ICodeEditeInterface>(({ lang, code }, ref) => {
  const { themeMode } = useTheme();
  const [eCode, setECode] = useState<string | undefined>('');
  const handleEditorDidMount = (monacoInstance: Monaco) => {
    defineMonacoTheme(monacoInstance, themeMode);
    monacoInstance.editor.setTheme('naruto');
  };

  useEffect(() => {
    if (!monacoInstance) return;
    handleEditorDidMount(monacoInstance);
    setECode(code);
  }, [themeMode, code, monacoInstance]);

  useImperativeHandle(ref, () => ({
    getCode: () => {
      return monacoInstance?.editor.getModels()[0]?.getValue();
    },
  }));

  return (
    <div className='h-full bg-white py-4 rounded dark:bg-[#1e1f32]'>
      <Editor
        beforeMount={handleEditorDidMount}
        options={{
          minimap: { enabled: false },
          contextmenu: false,
        }}
        onChange={(e) => setECode(e)}
        value={eCode}
        theme='naruto'
        defaultLanguage={lang}
        defaultValue={code}
      />
    </div>
  );
});

export default memo(CodeEdit);
