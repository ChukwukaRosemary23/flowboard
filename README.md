FlowBoard - Task Management System
A modern, real-time task management application inspired by Trello. Built with Go, PostgreSQL, and React.
![FlowBoard Logo](docs/images/flowboard-logo.png)
🚀 Features
- **User Authentication** - Secure JWT-based authentication
- **Board Management** - Create and organize multiple project boards
- **Drag & Drop** - Intuitive card movement between lists
- **Real-time Updates** - WebSocket-based live collaboration
- **File Attachments** - Upload and attach files to cards
- **User Assignments** - Assign team members to tasks
- **Labels & Tags** - Categorize tasks with colored labels
- **Comments** - Collaborate with team discussions on cards
🛠️ Tech Stack
**Backend:**
- Go 1.21+
- Gin Web Framework
- GORM (PostgreSQL)
- JWT Authentication
- WebSockets
- bcrypt
**Frontend:**
- React 18
- React Router
- Tailwind CSS
- Axios
- React Beautiful DnD
**Database:**
- PostgreSQL 14+
📋 Prerequisites
- Go 1.21 or higher
- PostgreSQL 14 or higher
- Node.js 18+ (for frontend)
⚙️ Installation
Backend Setup
1. Clone the repository:
```bash
git clone https://github.com/ChukwukaRosemary23/flowboard.git
cd flowboard/backend
```
2. Install dependencies:
```bash
go mod download
```
3. Configure environment variables:
```bash
cp .env.example .env
Edit .env with your database credentials
```
4. Create database:
```bash
createdb flowboard_db
```
5. Run the server:
```bash
go run cmd/api/main.go
```
The backend will start on `http://localhost:8080`
🗄️ Database Schema
- **Users** - User accounts and authentication
- **Boards** - Project boards
- **Lists** - Columns within boards
- **Cards** - Individual tasks
- **Comments** - Card discussions
- **Labels** - Task categorization
- **Attachments** - File uploads
- **Card Members** - Task assignments
📚 API Documentation
Coming soon...
🚧 Current Status
🔨 Work in Progress - Active Development

👤 Author
📧 Contact:
Rosemary Chukwuka
Full Stack Developer | MERN Stack | Go + PostgreSQL

LinkedIn: https://www.linkedin.com/in/chukwuka-rosemary-0944b9244
Email: chukwukarosemary2020@gmail.com
GitHub: https://github.com/ChukwukaRosemary23


📄 License:
This project is open source and available under the MIT License.

Built with ❤️ by Rosemary Chukwuka

🙏 Acknowledgments
Built as a portfolio project to demonstrate full-stack development skills.
⭐ If you find this project impressive, please give it a star!
