import { FC, memo, useEffect, useState } from 'react';
import { useTheme } from '@/hooks/useTheme.ts';
import { ThemeBar, Translate } from '@/components/index';
import { Button, Form, Image, Input, message } from 'antd';
import { EyeInvisibleOutlined, EyeTwoTone, UserOutlined } from '@ant-design/icons';
import { getCaptchaRequest, getUserMenusRequest, LoginParamsType, loginRequest } from '@/service';
import { useTranslation } from 'react-i18next';
import { FieldType } from '@/pages/Login/type.ts';
import { changeMenus, changeToken, changeUserInfo } from '@/store/UserStore';
import { useAppDispatch } from '@/store';
import { useNavigate } from 'react-router-dom';
import { getFirstMenu, sleep } from '@/utils';
import { changeTabHeader } from '@/store/UIStore';
import Logo from '@/assets/svg/logo.svg';
import classNames from 'classnames';

const Login: FC = () => {
  const navigate = useNavigate();
  const [captcha, setCaptcha] = useState<string>();
  const [loading, setLoading] = useState(false);
  const dispatch = useAppDispatch();
  const [formRef] = Form.useForm();
  const { themeMode } = useTheme();
  const { t } = useTranslation();

  useEffect(() => getCaptcha(), []);

  const getCaptcha = () => {
    getCaptchaRequest().then((res) => {
      const imageUrl = URL.createObjectURL(new Blob([res]));
      setCaptcha(imageUrl);
    });
  };

  const form: FieldType = {
    account: 'admin',
    password: '123',
    captcha: '',
  };

  const onFinish = (values: FieldType) => {
    setLoading(true);
    loginRequest(values as LoginParamsType)
      .then(async (r) => {
        const { token, user } = r.data;
        dispatch(changeUserInfo(user));
        dispatch(changeToken(token));
        await sleep(1000);
        const result = await getUserMenusRequest();
        dispatch(changeMenus(result.data));
        if (result.data?.length) {
          navigate(getFirstMenu(result.data).pagePath || '/');
          dispatch(changeTabHeader([getFirstMenu(result.data)]));
        } else {
          message.error('暂无任何菜单，请联系管理员');
        }
      })
      .catch(() => {
        getCaptcha();
        formRef.resetFields(['captcha']);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <>
      <div
        className={classNames('w-screen h-screen flex justify-center items-center tran dark:text-[#fff]', {
          'bg-[#ededed]': themeMode === 'light',
          'bg-[#2e2e2e]': themeMode === 'dark',
        })}>
        <div className='absolute top-0 left-0 right-0 flex justify-end p-4'>
          <ThemeBar />
          <Translate />
        </div>
        <div
          className={classNames('w-3/5 h-[500px] rounded-xl tran loginBg flex justify-between overflow-hidden', {
            physicLight: themeMode === 'light',
            physicDark: themeMode === 'dark',
          })}>
          <div className='flex flex-col flex-1 p-10 pt-20'>
            <p className='text-2xl font-bold'>{t('welcome')}</p>
            <div className='mt-10 ml-[-40px] hidden lg:flex'>
              <Image src={Logo} preview={false} className='animate__animated animate__backInUp' />
              <div className='w-full mt-10'>
                <p className='text-xl font-bold hidden lg:block'>{t('slogan')}</p>
                <p className='mt-2 text-sm text-[#6c727f] dark:text-[#fff] tran'>{t('description')}</p>
              </div>
            </div>
          </div>
          <div className='w-[500px] flex bg-gray-100 dark:bg-[#262626] tran'>
            <div className='flex flex-col items-center gap-4 flex-grow px-20 mt-20'>
              <p className='text-2xl font-bold'>{t('title')}</p>
              <p className='text-sm text-[#6c727f] dark:text-[#fff] tran'>A management platform using Golang and React</p>
              <Form onFinish={onFinish} initialValues={form} form={formRef} autoComplete='off'>
                <Form.Item<FieldType> name='account' rules={[{ required: true, message: t('accountRequired') }]}>
                  <Input suffix={<UserOutlined />} size='large' placeholder={t('completeAccount')} />
                </Form.Item>
                <Form.Item<FieldType> name='password' rules={[{ required: true, message: t('passwordRequired') }]}>
                  <Input.Password
                    size='large'
                    iconRender={(visible) => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
                    placeholder={t('completePassword')}
                  />
                </Form.Item>
                <Form.Item<FieldType> name='captcha' rules={[{ required: true, message: t('captchaRequired') }]}>
                  <div className='flex items-center w-full gap-1'>
                    <Input size='large' className='flex-1' placeholder={t('completeCaptcha')} />
                    <Image preview={false} src={captcha} height={36} width={150} onClick={getCaptcha} className='bg-white cursor-pointer tran rounded-s' />
                  </div>
                </Form.Item>
                <div className='w-full'>
                  <Button block htmlType='submit' type='primary' loading={loading}>
                    {t('login')}
                  </Button>
                </div>
              </Form>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default memo(Login);
