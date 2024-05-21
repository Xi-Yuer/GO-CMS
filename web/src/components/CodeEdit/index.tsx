import { Dispatch, forwardRef, memo, SetStateAction, useEffect, useImperativeHandle, useState } from 'react';
import { Editor, loader, Monaco } from '@monaco-editor/react';
import { defineMonacoTheme } from '@/theme';
import { useTheme } from '@/hooks/useTheme.ts';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker';
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker';
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';
import * as monaco from 'monaco-editor';
import { code } from '@/service/api/template';

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
  editCode: Dispatch<SetStateAction<code | undefined>> | undefined;
  defaultLang: string | undefined;
  defaultCode: string | undefined;
}

export interface ICodeEditRefProps {
  getCode: () => string | undefined;
}

let monacoInstance: Monaco | undefined;
await loader.init().then((monaco) => {
  monacoInstance = monaco;
});

const CodeEdit = forwardRef<ICodeEditRefProps, ICodeEditeInterface>(({ defaultCode, editCode, defaultLang }, ref) => {
  const { themeMode } = useTheme();
  const [eCode, setECode] = useState<string | undefined>('');
  const handleEditorDidMount = (monacoInstance: Monaco) => {
    defineMonacoTheme(monacoInstance, themeMode);
    monacoInstance.editor.setTheme('naruto');
  };

  useEffect(() => {
    if (!monacoInstance) return;
    handleEditorDidMount(monacoInstance);
  }, [themeMode, monacoInstance]);

  useEffect(() => {
    setECode(defaultCode);
  }, [defaultCode, editCode]);

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
        onChange={(e: any) => {
          editCode &&
            editCode({
              code: e,
              lang: defaultLang || 'text',
            });
        }}
        value={eCode}
        theme='naruto'
        defaultLanguage={defaultLang || 'text'}
        defaultValue={eCode}
      />
    </div>
  );
});

export default memo(CodeEdit);
