-- iFlow Monitor Platform Database Init
-- Version: v1.0

CREATE DATABASE IF NOT EXISTS iflow_monitor CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE iflow_monitor;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    email VARCHAR(128),
    role ENUM('admin', 'user', 'viewer') DEFAULT 'user',
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 角色表
CREATE TABLE IF NOT EXISTS roles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL UNIQUE,
    description TEXT,
    permissions JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 集群表
CREATE TABLE IF NOT EXISTS clusters (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL,
    kubeconfig TEXT,
    endpoint VARCHAR(256),
    status TINYINT DEFAULT 1,
    created_by BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 数据源表
CREATE TABLE IF NOT EXISTS datasources (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL,
    type VARCHAR(32) NOT NULL,
    address VARCHAR(256),
    port INT DEFAULT 0,
    is_default TINYINT DEFAULT 0,
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 审计日志表
CREATE TABLE IF NOT EXISTS audit_logs (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT,
    action VARCHAR(64),
    resource VARCHAR(128),
    details JSON,
    ip_address VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 集群节点表
CREATE TABLE IF NOT EXISTS cluster_nodes (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    cluster_id BIGINT NOT NULL,
    node_name VARCHAR(128),
    node_ip VARCHAR(64),
    status VARCHAR(32),
    labels JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
