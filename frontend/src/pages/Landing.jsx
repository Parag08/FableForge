import { GoogleLogin } from '@react-oauth/google';
import { useNavigate } from 'react-router-dom';

export function LandingPage() {
  const navigate = useNavigate();

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <GoogleLogin
        onSuccess={credentialResponse => {
          navigate('/home');
        }}
        onError={() => {
          console.log('Login Failed');
          // Handle login failure here
        }}
      />
    </div>
  );
}