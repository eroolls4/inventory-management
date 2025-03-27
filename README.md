# Inventory Management System

##  Environment Configuration

### Backend `.env` Template
```env
PORT=3000 # Application port, default 3000
DB_HOST=db # Database host
DB_USER=postgres # Database username
DB_PASSWORD= # Populate with your database password
DB_NAME=inventoryDB # Database name
DB_PORT=5432 # Database port
REDIS_HOST=redis # Redis host
REDIS_PORT=6379 # Redis port
REDIS_PASSWORD= # Populate with your Redis password (if any)
SECRET_KEY= # Generate a long, random string for JWT/session security

FRONTEND
VITE_APP_API_URL=http://localhost:3000 # Backend API URL
VITE_APP_FE_PORT=5000 # Frontend application port