import React from 'react';
import { Link } from 'react-router-dom';

function Sidebar() {
  return (
    <nav style={{ width: '200px', background: '#f5222d' }}>
      <ul>
        <li><Link to="/">仪表盘</Link></li>
        <li><Link to="/clusters">集群</Link></li>
        <li><Link to="/monitor">监控</Link></li>
        <li><Link to="/alerts">告警</Link></li>
        <li><Link to="/logs">日志</Link></li>
        <li><Link to="/settings">设置</Link></li>
      </ul>
    </nav>
  );
}

export default Sidebar;
