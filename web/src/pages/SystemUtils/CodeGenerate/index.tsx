import { FC, memo, useEffect, useRef, useState } from 'react';
import CodeEdit, { ICodeEditRefProps } from '@/components/CodeEdit';
import { createTemplateRequest } from '@/service/api/template';

const CodeGenerate: FC = () => {
  const [code, setCode] = useState('');
  const CodeEditRef = useRef<ICodeEditRefProps>(null);
  const createTemplateAction = async () => {
    const result = await createTemplateRequest({
      tableName: 'Users',
      fields: [
        {
          name: 'Name',
          type: 'string',
          default: '',
        },
        {
          name: 'Age',
          type: 'int',
          default: '0',
        },
      ],
    });
    setCode(result.data.repository);
  };

  useEffect(() => {
    createTemplateAction().then();
  }, []);
  return (
    <div className='h-full'>
      <CodeEdit lang='go' code={code} ref={CodeEditRef} />
    </div>
  );
};

export default memo(CodeGenerate);
