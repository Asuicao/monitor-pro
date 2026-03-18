import React from 'react';
import { useNavigate } from 'react-router-dom';
import { setToken } from '../utils/auth';
import api from '../api';

function Login() {
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setToken('demo-token');
    navigate('/');
  };

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <form onSubmit={handleLogin}>
        <h2>iFlow Monitor Pro</h2>
        <input type="text" placeholder="用户名" defaultValue="admin" />
        <input type="password" placeholder="密码" defaultValue="password" />
        <button type="submit">登录</button>
      </form>
    </div>
  );
}

export default Login;
