import React from 'react';

function Dashboard() {
  return (
    <div>
      <h1>监控仪表盘</h1>
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: '20px' }}>
        <div style={{ padding: '20px', border: '1px solid #ddd' }}>
          <h3>集群数量</h3>
          <p>0</p>
        </div>
        <div style={{ padding: '20px', border: '1px solid #ddd' }}>
          <h3>告警数量</h3>
          <p>0</p>
        </div>
        <div style={{ padding: '20px', border: '1px solid #ddd' }}>
          <h3>监控指标</h3>
          <p>0</p>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
