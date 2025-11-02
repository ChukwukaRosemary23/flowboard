-- Migration: Add access control tables
-- File: migrations/000004_add_access_control.up.sql

BEGIN;

-- Create roles table
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add BTREE index on name for faster lookups
CREATE INDEX IF NOT EXISTS idx_roles_name ON roles(name);

-- Seed roles data
INSERT INTO roles (name, description) VALUES
    ('owner', 'Board creator with full control'),
    ('admin', 'Can manage members and all board content'),
    ('member', 'Can create and edit cards and lists'),
    ('viewer', 'Read-only access to board')
ON CONFLICT (name) DO NOTHING;

-- Create permissions table
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    resource VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add BTREE index on name for faster permission lookups
CREATE INDEX IF NOT EXISTS idx_permissions_name ON permissions(name);

-- Seed permissions data
INSERT INTO permissions (name, resource, action, description) VALUES
    ('view_board', 'board', 'view', 'Can view board and its contents'),
    ('edit_board', 'board', 'edit', 'Can edit board settings'),
    ('delete_board', 'board', 'delete', 'Can delete the board'),
    ('manage_members', 'board', 'manage', 'Can add/remove members'),
    ('create_list', 'list', 'create', 'Can create new lists'),
    ('edit_list', 'list', 'edit', 'Can edit list properties'),
    ('delete_list', 'list', 'delete', 'Can delete lists'),
    ('create_card', 'card', 'create', 'Can create new cards'),
    ('edit_card', 'card', 'edit', 'Can edit card properties'),
    ('delete_card', 'card', 'delete', 'Can delete cards'),
    ('move_card', 'card', 'move', 'Can move cards between lists'),
    ('comment_card', 'card', 'comment', 'Can add comments to cards'),
    ('upload_attachment', 'card', 'upload', 'Can upload attachments')
ON CONFLICT (name) DO NOTHING;

-- Create role_permissions table (maps which permissions each role has)
CREATE TABLE IF NOT EXISTS role_permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(role_id, permission_id)
);

-- Add compound BTREE index for faster searches on both columns
CREATE INDEX IF NOT EXISTS idx_role_permissions_compound ON role_permissions(role_id, permission_id);

-- Seed role_permissions data
-- Owner has ALL permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p WHERE r.name = 'owner'
ON CONFLICT DO NOTHING;

-- Admin has all except delete_board
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p 
WHERE r.name = 'admin' AND p.name != 'delete_board'
ON CONFLICT DO NOTHING;

-- Member can create/edit/delete lists and cards
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p 
WHERE r.name = 'member' AND p.name IN (
    'view_board', 'create_list', 'edit_list', 'delete_list',
    'create_card', 'edit_card', 'delete_card', 'move_card',
    'comment_card', 'upload_attachment'
)
ON CONFLICT DO NOTHING;

-- Viewer can only view
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p 
WHERE r.name = 'viewer' AND p.name = 'view_board'
ON CONFLICT DO NOTHING;

-- Create board_members table (tracks which users have access to which boards)
CREATE TABLE IF NOT EXISTS board_members (
    id SERIAL PRIMARY KEY,
    board_id INTEGER NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER NOT NULL REFERENCES roles(id),
    invited_by INTEGER REFERENCES users(id),
    invited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    accepted_at TIMESTAMP,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(board_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_board_members_board ON board_members(board_id);
CREATE INDEX IF NOT EXISTS idx_board_members_user ON board_members(user_id);
CREATE INDEX IF NOT EXISTS idx_board_members_role ON board_members(role_id);
CREATE INDEX IF NOT EXISTS idx_board_members_status ON board_members(status);

-- Add compound BTREE index based on CheckPermission query
CREATE INDEX IF NOT EXISTS idx_board_members_compound ON board_members(board_id, user_id, status);

-- Update existing boards table to add owner_id
ALTER TABLE boards ADD COLUMN IF NOT EXISTS owner_id INTEGER REFERENCES users(id);
CREATE INDEX IF NOT EXISTS idx_boards_owner ON boards(owner_id);

-- Migrate existing data: Set current board creators as owners
INSERT INTO board_members (board_id, user_id, role_id, status)
SELECT 
    b.id,
    b.owner_id,
    (SELECT id FROM roles WHERE name = 'owner'),
    'active'
FROM boards b
WHERE b.owner_id IS NOT NULL
ON CONFLICT (board_id, user_id) DO NOTHING;

COMMIT;