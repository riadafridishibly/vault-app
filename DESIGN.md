# Design Document for Vault App

## Goals

- [ ] Password Manager (no browser integration)
  - Browser integration:
    - `stdio` rpc
    - extension (not goal)
  - Features
    - Copy Paste Management
- [x] Password Generator
  - Generates strong password
  - [ ] Save to password manager
- [ ] Encrypts File/Directory
  - Local files encrypt
  - X25519
  - ED25519
- [ ] Secure Notes (Encrypted Notes)

## Implementaion

### UI

- Password Manager
  - Table
    - Row
      - Website
      - Email/Username
      - Password
      - Tags
      - Edit
  - New
    - [x] Modal (May change)
    - [] backend implementation
  - Edit
    - Modal

### Backend

- Password Manager
  - Models
    - ## PasswordItem
